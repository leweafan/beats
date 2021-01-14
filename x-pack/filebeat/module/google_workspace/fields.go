// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package google_workspace

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "google_workspace", asset.ModuleFieldsPri, AssetGoogleWorkspace); err != nil {
		panic(err)
	}
}

// AssetGoogleWorkspace returns asset data.
// This is the base64 encoded gzipped contents of module/google_workspace.
func AssetGoogleWorkspace() string {
	return "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"
}
