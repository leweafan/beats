// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
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
