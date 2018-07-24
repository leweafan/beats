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

package prometheus

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "prometheus", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJy0ksGumzAQRfd8xVX2jw9g0U33T5W6rCrkwAWsGA/1DE/K31cuCY9QomzaWQ7jew5jv+HCa4UpyUgbOGsBmLfACqdva/NUAC21SX4yL7HClwIAvpszRSMhsDG26JKM+DxVFoAOkqxuJHa+r9C5oCyAxECnrNC7PEMzH3ut8OOkGk4/C6DzDK1WfzBviG7kTjKXXaeckWSebp0DycesJe+Rf+8uFM0/tXaPIE9BSy1bcWeZDTZwsxAo0wdTuRlfTc40t+nvnbeGUcx3vnEZrg8Tz2xfGOd634QuOyh3M0dKW61fM2fWgbG34a+hu1mQ2B98fCGX6+ucEqMtGCyYveKnTJtkmtj+B4/3eTwzQbo742bED8b90jYvt6EqtZSJse7a42s7kHp1a6tMDkbnA9cjkp7YqElyPctmmONFa5N6YlKv9s+lRo6SrlhAsMEZXGJ+wLjScMOyhQlar5ey+B0AAP//Qe07WA=="
}
