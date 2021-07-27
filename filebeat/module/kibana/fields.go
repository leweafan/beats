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

package kibana

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "kibana", asset.ModuleFieldsPri, AssetKibana); err != nil {
		panic(err)
	}
}

// AssetKibana returns asset data.
// This is the base64 encoded zlib format compressed contents of module/kibana.
func AssetKibana() string {
	return "eJzMlk1v4zYQhu/5FQNfcqkE2+uv9WGBYruHYtFbb0UhjMWRxJoiFc4waf59QclJFVmxHdctqiPJ951nRsOPBPb0vIW93qHFOwDRYmgLk25gcgfgyRAybaGM84o497oR7ewWvtwBwEELvzgVDN0BFJqM4m07l4DFmnr+8ZPnJtp5F5rDyIhr/DpLKJyHBj1rW8L3LphxJaeHdf14/ZhMzNrZTKvXqUGkXyuCn38CV4BUBIHJv4gAmV2uUUjBk5YKpNIM9EhWUviGeRUJtAUUoboR8MTBCEMcgmD1Q6BXK63SHgD9iXXTlng2/0SL5Wqd0ObzLpnN1acEF8tVspivVrPFbL2YTqeTnrIr256en5xXx9k2mNOJXCcxWa1ekm2Xn8pyMsqsqMBg5ANY+Egqc7s/KJc0Lj3JFxe8EkYldMorQJGrnUOvrkS9vJD/CHM1/7zcKTVNpjilZDajdbKZL4tktV5scoVrtVluLs8AlcrEZe2/5ZP8TBIT6LpAKwapUADfZvOEDFyhJwXi3kngt/tDS9z/APc1+j2JtuX975dDKzIklBXe1Tcl91S7OBSN/zV4DFKRFZ1jRMwa7x61In8yg7caeNEcNQ8ezphT/bND1vnsat6zG/I91najXg18Na8nNPVJ4G8GWXTOhD6vhvitvHWGp0rnFRTBFNoYUhexWxT9SJfDG+f2ofkwdCc7wP43nOXR2v7tfAT9pTcB8B2LfYsFRlvi/mU3vJr7UQVLfjPxHuZZgBbi9V1QxmdCNE9Hw7Kg0O3ifg3ek5XONh5MHch47JoER0N3x9ZgqhvMxuFG/btXVmpcmcZQqaeHtCJU5Dn1VJB/czL9HR6NxuG/aFCqLVQiTbQJxNJZHHvUuvTY1Ud8oIvJ/n9EL7WKL8EES7LDX3IOLgqzVpg6r0tt0dysWLUT+lEpTzy+bd6nYhd8TimOiq8FCn6Y2tnieHPLqsQdF/irU+O7+WwHceMs08Emy499roGqSSo3PEI+1NKjDkOUvwIAAP//j5XtEA=="
}
