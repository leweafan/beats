// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/elastic-agent/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/apm-server.yml
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/fleet-server.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
	// spec/osquerybeat.yml
	// spec/packetbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}
