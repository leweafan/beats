// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gcp

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "gcp", asset.ModuleFieldsPri, AssetGcp); err != nil {
		panic(err)
	}
}

// AssetGcp returns asset data.
// This is the base64 encoded gzipped contents of module/gcp.
func AssetGcp() string {
	return "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"
}
