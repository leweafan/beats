// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package application

import (
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/common/backoff"
	"github.com/elastic/beats/x-pack/agent/pkg/agent/errors"
	"github.com/elastic/beats/x-pack/agent/pkg/core/logger"
	"github.com/elastic/beats/x-pack/agent/pkg/fleetapi"
	"github.com/elastic/beats/x-pack/agent/pkg/scheduler"
)

type dispatcher interface {
	Dispatch(acker fleetAcker, actions ...action) error
}

type agentInfo interface {
	AgentID() string
}

type fleetReporter interface {
	Events() ([]fleetapi.SerializableEvent, func())
}

type fleetAcker interface {
	Ack(action fleetapi.Action) error
	Commit() error
}

// fleetGateway is a gateway between the Agent and the Fleet API, it's take cares of all the
// bidirectional communication requirements. The gateway aggregates events and will periodically
// call the API to send the events and will receive actions to be executed locally.
// The only supported action for now is a "ActionPolicyChange".
type fleetGateway struct {
	log        *logger.Logger
	dispatcher dispatcher
	client     clienter
	scheduler  scheduler.Scheduler
	backoff    backoff.Backoff
	settings   *fleetGatewaySettings
	agentInfo  agentInfo
	reporter   fleetReporter
	done       chan struct{}
	wg         sync.WaitGroup
	acker      fleetAcker
}

type fleetGatewaySettings struct {
	Duration time.Duration
	Jitter   time.Duration
	Backoff  backoffSettings
}

type backoffSettings struct {
	Init time.Duration
	Max  time.Duration
}

func newFleetGateway(
	log *logger.Logger,
	settings *fleetGatewaySettings,
	agentInfo agentInfo,
	client clienter,
	d dispatcher,
	r fleetReporter,
	acker fleetAcker,
) (*fleetGateway, error) {
	scheduler := scheduler.NewPeriodicJitter(settings.Duration, settings.Jitter)
	return newFleetGatewayWithScheduler(
		log,
		settings,
		agentInfo,
		client,
		d,
		scheduler,
		r,
		acker,
	)
}

func newFleetGatewayWithScheduler(
	log *logger.Logger,
	settings *fleetGatewaySettings,
	agentInfo agentInfo,
	client clienter,
	d dispatcher,
	scheduler scheduler.Scheduler,
	r fleetReporter,
	acker fleetAcker,
) (*fleetGateway, error) {
	done := make(chan struct{})

	return &fleetGateway{
		log:        log,
		dispatcher: d,
		client:     client,
		settings:   settings,
		agentInfo:  agentInfo,
		scheduler:  scheduler,
		backoff: backoff.NewEqualJitterBackoff(
			done,
			settings.Backoff.Init,
			settings.Backoff.Max,
		),
		done:     done,
		reporter: r,
		acker:    acker,
	}, nil
}

func (f *fleetGateway) worker() {
	for {
		select {
		case <-f.scheduler.WaitTick():
			f.log.Debug("FleetGateway calling Checkin API")

			// Execute the checkin call and for any errors returned by the fleet API
			// the function will retry to communicate with fleet with an exponential delay and some
			// jitter to help better distribute the load from a fleet of agents.
			resp, err := f.doExecute()
			if err != nil {
				f.log.Error(err)
				continue
			}

			actions := make([]action, len(resp.Actions))
			for idx, a := range resp.Actions {
				actions[idx] = a
			}

			if err := f.dispatcher.Dispatch(f.acker, actions...); err != nil {
				f.log.Errorf("failed to dispatch actions, error: %s", err)
			}

			f.log.Debugf("FleetGateway is sleeping, next update in %s", f.settings.Duration)
		case <-f.done:
			return
		}
	}
}

func (f *fleetGateway) doExecute() (*fleetapi.CheckinResponse, error) {
	f.backoff.Reset()
	for {
		resp, err := f.execute()
		if err != nil {
			f.log.Errorf("Could not communicate with Checking API will retry, error: %s", err)
			if !f.backoff.Wait() {
				return nil, errors.New(
					"execute retry loop was stopped",
					errors.TypeNetwork,
					errors.M(errors.MetaKeyURI, f.client.URI()),
				)
			}
			continue
		}
		return resp, nil
	}
}

func (f *fleetGateway) execute() (*fleetapi.CheckinResponse, error) {
	// get events
	ee, ack := f.reporter.Events()

	// checkin
	cmd := fleetapi.NewCheckinCmd(f.agentInfo, f.client)
	req := &fleetapi.CheckinRequest{
		Events: ee,
	}

	resp, err := cmd.Execute(req)
	if err != nil {
		return nil, err
	}

	// ack events so they are dropped from queue
	ack()
	return resp, nil
}

func (f *fleetGateway) Start() {
	f.wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		f.worker()
	}(&f.wg)
}

func (f *fleetGateway) Stop() {
	close(f.done)
	f.scheduler.Stop()
	f.wg.Wait()
}
