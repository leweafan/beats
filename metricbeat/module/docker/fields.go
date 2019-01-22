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

package docker

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzsm92P4roVwN/nrzi6fahU7YBaXfWBh0rt3VblYe+Odnf6ioxzAi6JnfoDhv3rK38EQuIkEAIzs9o8Ejjn5+PzYfuYR9jgfgaJoBuUDwCa6Qxn8MtH98EvDwAJKipZoZngM/jbAwCAfwlKE62AiixDqjGBVIo8vJs8AEjMkCicwYo8AKi1kHpBBU/ZagYpyRQ+AKQMs0TNnNRH4CTHCot99L6wEqQwRfgkwmOfOU+FzIn9GAhPHBxTmlEFZCmMDmL/qEAazhlfARVcE8ZRqkmQUqWpEh2+eXgTA+uAqxjtIAty1JLRg3L7nJqsfOpYp2h5Tnhy8q6E2+B+J2T9XQeifX7zAkGviYYdUYAvSI2dXsZBr7ExjkmcSyLRGOdKiMbLoD4SjbBboyc4mtDyBU1xDOsFRo1pnVK1lxzXyooFSRKJSuFNdM+f4CC/Zdzse93EcYe9bMzsO8bcFlqctEokhdCLtG6OI1gm+CrysofNPd+EJpmHEymQLHNekrIMVem0Ld56Ari7BdvXQHUkcoG1JluEJSIv3ReEBLomfIUJKMYp+hdM8PgEa7KKuxaRkuwvm+B5TlboJE6aqa8w1yS9L4ZrliP89vQ8Tr7boOSYTQqqo6NXlGSYLNJMkPoXfHmYQYGSItdkdWEOejr8zs2nHRXjgQdUQSjGZyoQa0Y38RmLeFdfRD49g5N3HoHaK435G7CZi1MP761noyLQdZGPbTsv1pswrtgolG/AYMFMlqZrgh3tzRysT7ub1tcw1reDPxlFVi10VEic/KkVTyz/i41X/sPFnWf7NDCY8uhdg2qf8v5hXe4Uv5t8ifJIGtwjgnpYyTO1YeKqRTNTG5hPP49TPCSS+Ip0wKro75Sa3GSudlu5ChIj7bbC5rSMpYeqH9s8tIFWYUVxi8XScRIHQR/xlnvdWNz2ApYB0/bjMwbwD/tTBz+cXTY3IL3oF9mWGimR62DjwuZPpKK2Tat6ZTyKO1JPgoVEar1vBn+d/Do0ki8C3UnWsNso8eMEv7sAGkb9ViLI0mvk7yCIgp3Pcc7XDqPzUJXJcyLrm7QRKxHh7zWmbKkXBUq38323seWqUzkJ7yPIKkbv8V633H+lOGu4d4S15MQtcn3V2tMfdNblXL7sHPsQ8p+WKCb1cP444nmwV8YSf/xKtoRlZJlhVG8qRT76MIWRNK7Oyh5P3bc1up9bP/M7JcCcaV0Gbt0PjhyEWpm3IenUKuqJ44oa0hTWVwsaXtY37DM4yuHPP5Y58rypODGM1pItTVcBiG6Pob5FvmoU/yGSCaOskOmWZAYrXKdj+2CzI/LEjk5wYFqdenY5rjWSTK/pGulmhLQ2ch8qJSxjfKW0RLKJeiXjGleNUtdjSSo4Na5AWQWYQH34t8uy/z6aO8gGKpJ4Joo55OBYdMKCsS/tsiBPFpH2GnT13s5AqtsDeRKXVJkMqe9B4hR1swijCxML9hGivIrSoucwNS9MLxoeVCWJR8gwozTc9ZC189NjyqH5w8kZJ3e0LFcGhI4tHR4sI3uUCliCXLOUNRt/fZEUFsa3cRt45ux/pmQ9QsKKbZGDKUIdiPcAq5gFuSHl/AgWKpYD/gAsBaatRyutPvh7H7s1o2u/qwlbCj+4hEmkOts7hcjrKe2G9wXcLQaWVy4OeKL+SwMjNs99b9X1pr1LXuqHWya1aWy4YOzetDNNR3McVyYjsdQ0cvfespAsA0roGhOPpYAoJShzJxxaNJ3syB2dz4wsMRvaJxm0dpiHHGT19sDdq5HPeHpVJ2bOU1Gme1gShYldrK61LtRsOk0EVRN/w2pCRT5FvmIcpxJTlMgpTknBpv79QmIuNC5IwRbbP0/+8uv0D9OEqSIj+0ffmH3csQQf2fFC17VXpMp7XmMF9ectSuekJ7eBLg7tghjVSHkwQkz5kOKHYxOvKHLhrckULsfdAar9Gl6TSmlRFHcxVdB0FlXsJOwWTK7OdpnqFuc+YYUSVntroXR0MRXncFk7ynJ527nVGl5LM9PlmIuT0/WLc90nJ2H44rZ1hzyhwrRsFDsPTzsN9C/CbCoyXLddxsxYzuJaI9PRdQreQxLs5tTFSaQaLQl/+fo1TPWw7DsoekdoEXh/DuQSlYsqUKjdGqhjpR89sOp1HjjnnsqZ6J9aoCuC2y4oNXUOnvhn5Xeew6c+Jy+vMPGfyEtJHbmvBG9yqh1o6/TCWwunml1LNo56J+RVp6a/exEj1wRmc3ZK6IhtjBL0INqpalkvxJsW8Zi8ooE351TktmSHiQj/HDk27y4N4Ndq79bXIqwcmJPZHhyJ7F6/9gT1ALKg8UhYELrBZqqsnE9KKRpbJBhtOevFuwbH2UjhC3dYYnczVU6S7xMwn41eiR8xYEQ5sDcbMAfCtxMw5yPdL2C6mY4FZilMy7/T7lVl/L9pTv835rpG9ePf9xNFY5Wd8dzhZ7m5Tbm5Y/i01JwfMHzGKkLjh8/P4jOw+Pw/AAD//8MBpec="
}
