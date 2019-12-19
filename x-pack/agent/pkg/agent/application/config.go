// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package application

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/elastic/beats/agent/kibana"
	"github.com/elastic/beats/x-pack/agent/pkg/config"
	fleetreporter "github.com/elastic/beats/x-pack/agent/pkg/reporter/fleet"
	logreporter "github.com/elastic/beats/x-pack/agent/pkg/reporter/log"
)

// TODO(ph) correctly setup global path.
func fleetAgentConfigPath() string {
	return "fleet.yml"
}

// Config define the configuration of the Agent.
type Config struct {
	Management *config.Config `config:"management"`
}

func localDefaultConfig() *Config {
	return &Config{}
}

type managementMode int

// Define the supported mode of management.
const (
	localMode managementMode = iota + 1
	fleetMode
)

var managementModeMap = map[string]managementMode{
	"local": localMode,
	"fleet": fleetMode,
}

func (m *managementMode) Unpack(v string) error {
	mgt, ok := managementModeMap[v]
	if !ok {
		return fmt.Errorf(
			"unknown management mode, received '%s' and valid values are local or fleet",
			v,
		)
	}
	*m = mgt
	return nil
}

// ManagementConfig defines the options for the running of the beats.
type ManagementConfig struct {
	Mode      managementMode      `config:"mode"`
	Reporting *logreporter.Config `config:"reporting.log"`
}

func defaultManagementConfig() *ManagementConfig {
	return &ManagementConfig{
		Mode: localMode,
	}
}

type localConfig struct {
	Management *localManagementConfig `config:"management" yaml:"management"`
}

type localManagementConfig struct {
	Reload    *reloadConfig       `config:"reload" yaml:"reload"`
	Path      string              `config:"path" yaml:"path"`
	Reporting *logreporter.Config `config:"reporting" yaml:"reporting"`
}

type reloadConfig struct {
	Enabled bool          `config:"enabled" yaml:"enabled"`
	Period  time.Duration `config:"period" yaml:"period"`
}

func (r *reloadConfig) Validate() error {
	if r.Enabled {
		if r.Period <= 0 {
			return ErrInvalidPeriod
		}
	}
	return nil
}

func localConfigDefault() *localConfig {
	return &localConfig{
		Management: &localManagementConfig{
			Reload: &reloadConfig{
				Enabled: true,
				Period:  10 * time.Second,
			},
			Reporting: logreporter.DefaultLogConfig(),
		},
	}
}

// FleetAgentConfig is the internal configuration of the agent after the enrollment is done,
// this configuration is not exposed in anyway in the agent.yml and is only internal configuration.
type FleetAgentConfig struct {
	API       *APIAccess    `config:"api" yaml:"api"`
	Reporting *LogReporting `config:"reporting" yaml:"reporting"`
}

// APIAccess contains the required details to connect to the Kibana endpoint.
type APIAccess struct {
	AccessAPIKey string         `config:"access_api_key" yaml:"access_api_key"`
	Kibana       *kibana.Config `config:"kibana" yaml:"kibana"`
}

// LogReporting define the fleet options for log reporting.
type LogReporting struct {
	Log   *logreporter.Config             `config:"log" yaml:"log"`
	Fleet *fleetreporter.ManagementConfig `config:"fleet" yaml:"fleet"`
}

// Validate validates the required fields for accessing the API.
func (e *APIAccess) Validate() error {
	if len(e.AccessAPIKey) == 0 {
		return errors.New("empty access token")
	}

	if e.Kibana == nil || len(e.Kibana.Host) == 0 {
		return errors.New("missing Kibana host configuration")
	}

	return nil
}

func defaultFleetAgentConfig() *FleetAgentConfig {
	return &FleetAgentConfig{
		Reporting: &LogReporting{
			Log:   logreporter.DefaultLogConfig(),
			Fleet: fleetreporter.DefaultFleetManagementConfig(),
		},
	}
}

func createFleetConfigFromEnroll(access *APIAccess) (*FleetAgentConfig, error) {
	if err := access.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid enrollment options")
	}

	cfg := defaultFleetAgentConfig()
	cfg.API = access
	return cfg, nil
}
