// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "elasticsearch", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
