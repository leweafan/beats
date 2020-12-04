// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gcp

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "gcp", asset.ModuleFieldsPri, AssetGcp); err != nil {
		panic(err)
	}
}

// AssetGcp returns asset data.
// This is the base64 encoded gzipped contents of module/gcp.
func AssetGcp() string {
	return "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"
}
