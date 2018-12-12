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

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzsfduT2zbW57v/ClT2wU5VtxK3Y++UH77aHnvypWuSiWsc787W1pYaIo8opEmABkCplb9+CwcXArxIhLoz+/LlwWlJxO/gcnDuAK/JAxzfk0I0jeAvCNFM1/CefPCfS1CFZK1mgr8n//GCEEI+CK4p48o1IlsGdakI3VNW000NhHFC65rAHrgm+tiCWr0g7rH3Lwi5Jpw28N4+sGpE2dWAyBPECPltB/g8EVuid0Ds80TvqCYVcJBUQ4m/WLwXIwq0MIgIZ3rz3gz5IGSJ38AjbVoz5FpUFZTXjM935RaB3G8bUEi12FFeuf5oyaoK5Kg/5r8fhcRvtwynSEMlmT72wwHSCqWYmcA9rTtQhEp47xpTrSXbdBrUuhEl2zIor0ghwYz9ipRQA/7RtaX9phF78z/KS1IIvmXV2nYzmR3Tk2hWKim6djz4H01/e/orfKJfzB6upXrnumsBNTxq90WCaFbUPEy0CDOycg82Xa3ZOobvCUh6cN9MLeTsqnkmGpEkv+2YIkwRSrjg15TT+vgHlHZwdkGZIp2CbVdHWFshCa0qCRU1hJRbYd9LTWUFej2ajbSvowmxzWwnDQV1bGrGH0boxxZyYJHbzJPklfnzipRMXpEe/9sBfgl7VmRRsC0GMIyLMgsFGxAJrQQFXDNexZsl/K2OSkMzINaxcimpnicM0U6BJHcfyasvdx+/xUmBosNdyUrTiS0DSV59xh+3fX/EgYMc9AG/y14YbPUSeUwanAFotXxgyOCSNVQe7UbGcf3noOtD/LDjL6IQS+UJ9OY0CwTB+/27H76fI2gwkrlnnIhC07pnFdyEA9IK9JApNkLUQPmQuJYdTBH/DJqwiO6OWml/b6HvyYZpQ2ZFfm2YNhpI6B3IA1PDaVCgh+v4jH2p8vrC/kiXpBa8OsmhpoWZ9M1RgzIyxIhGpojg9ZHQsoSSHHbAyb3Buze/3Jt290O50mjWpKSNppojXVOliVd0xDRNmOAVfoN0C8G1MTMOVIUGQ9pFPm2v1E9QbkDTkmqKpO3zI8pCsorxRUox1la3nFAp6dGQVloyXilvdBjBSHtLAR61kRy1o2Q0R4SjjX6zqs5YH47RrnA4X/75M2F23kpx4LWgZi23UjQr8iuvjxGM6tpWSMNXjJOGFr9+viJ7RhHm4ZePdxqa/7UDCT9K0ajeVFhFEJ4x2db3lPGtkA1uXsM1XOjegvzTjYHFan+k7IlAZRShufEYVhjt/Jrx7jFZ/ljkjljv899+Ng2c9tHHVLjaRtOzYRRIznQYaqh8PAmx+R0KvRpOsqghGxaRXqowGgMyBC5FQ8PGuBzawhjNbRCGRGrYQ32GRhC+6vsMuoi8evHCeU8boLr3nf5qPy3wnEy7XPcp4TADsDJ/nrd1YmVtekgU8NKbWbWoSANK0QrUitxFT2EzpgKUMrrI2mPWseik3cPWutRGMuIWQh/GbaQSMZl2Gz0Gs8JoJ5R2lNzzvwkklfTjyvxmNZ/5eB9wRGvFyFy/VuNJ8xTPT1zoG1VEgu4kh5JsjnbXtMb9NLNoDVMjHQ47Vuz6jkdzJzvOGa8memO0yh+CL+iNf/LP7M0epPIO88nOuAc9WyE7j/xyphJH2FP65n+YoShNm/abOd0s4WvHJJSJaWTVRvJc2MS3XdUpTW7e6R25+f71uyvy+ub9m7fv375ZvXlzs2x2sUtWxQcnHjeIhELIEjVmGN/I96vUaSq3csO0NDa0edbOVkGNKEB+b0HahTKeu/mgJeUqCmD4eRoQttIhmUcrtNxX9sM6wz8KsgqdpLCnjICyxAY9ACmFXKbqeiJ/M428BCwsRbRwypKZZ2mNdoLZ2QVVKL+QjprWhpG5aYVZbAsNra8T3eq75nBWIwJF79yQKWt6EboBGUNHLj6Z1VsL0B2bOB1Fa0ZVr6Ru3ccJFPzJL8oWu9m0VLMNq41NcmB6R/776jGN6RHy31zQEPlXxRzpB2Yk6VDw2tH5zpn/Wqp3U0I6jekNJOZJmPCs72ctupK0UhSgVOC/1GUwj6xaKfasHPj1YyLGE1hNtEjBGFea8gJWA2fwJJ5vtHaNZiAXTOgU6Ghu7e8NLXaMw2oUazqJ6lqtQ6sU1Nk3yEPrBSsXIU83HS0VGrA5k+vaTM+thKqXuQvA3POOxT6K4gHkGR6z8g7k+U6XCLcatRhDLeCEEdiYDXo6jZF+l4BiSy98cIqi7ILfgLhXgic9LY5+8X628U0tUmiqXAjCGUC0LNf4wDo452EFZm3oud0bmRVQnLEd/hEZ12kPV+TTOKpvAK9IVQBGYktWMU1rUQAdxrFOSILZvty5B8ndR98lI0eJ39XnKZy3iwON2KtYRmUkJuIA2M2qgZJ1zWnqv1gIG9bOIj4nhEIPOnUNVOnr18UZMy4CImiPs97WZsp2h6neyD7BcrEMirrifrl+XM56ronpy38KUdVgd9o89UTIzRD4Jz5zbnxuo1sx0O/0j/7zBLiTkUobc6EQdQ2FdiEo95vZs2onpF5b+/M92dJamUWjvNgJ6eldh13+Yjo1FbpFJq3TOStyQkKTyyyyL5x97aAHJKycsilT4fkkijFfIJz3jV0HjBuz6VitieCnuhIJgwt78iHQtCmOeVo13UCtRtQST4ac9mbO9OUOZ8LSCUxrmLln2Z/spwmQO+OKRIzqErqp6Ol503x/ljMd7Ty+fPqa/OQs6/FqPBOnWwExweRUFjumodCdfIYxJHDkFayqFXn8y7v1ux+uCJXNFWnb4oo0rFXfjrsi1Kqtqd4K2TytJ79+Jh7I9aEAroW6It2m47q7IgfGS3GY6UQab7m8Dw5nksaWNiwJ6l9GwsK4QUood1RfkRI2jPIrspUAG1WeGi1rR11IvjpB/WemtBFod5+uaVlKUArUmEBDi6cN0pPZUVkeqISe2BXpVEfr+kh+uf0Q98HLkYduA5KDhsjP/nv83QTZ/vc+oZTYtD0oiWXJabXYNzorgJJOkywx1IryGdRDNAOtKK1smyTVPVU0RZQ+iZJ8ufs4JmT+VS0tnm9QPeKYmCjheWcQqyimp3Cpcl1GyKKRhrZjSpRzYdPyz0Yugpym+ZwGS0S3SGyXU2SfwWSbpJv40bQrmY4c6Vv/eZiyUoB+pk0eYiCvD8l7ZxnbuvqzEwV6L8iEUJiSJAXVUAl5fDE/AcG5QeLXsqsXpDywFy9VwLepW5tfKkGyvXcaMDNko/RoG96nuY3c0inso4v6v1R91H2Uc10yOTSjSKkWFeO+OClByQCZag4Z7WG7hUKzPUwibVUGlC0ksSmpKbAcLAVcT4Jk1Er5+qh0bjIA+rmZhFIZUAr0NMg2ByWe4Uk08+/aCM4YcEGC5m5rdpQS9R7WrFRY4IN5WYEpuemUKyaB7z3JPk17YHXtJRuhRo63jFfG2OqYL79AbemrNB3lMnV4rPFpNhS5/g8ihdDfni6OoLHlsKDUINp/EeHxHlwENgcDmTiD/TiBGO/JRZCjfTkBmovZ788JsCoPq68yHM5dJtBwv06NMxOy37eTK5GJNt6/o5j8sxUT0ULbaPj8hsmtJqJFITquieo2rs4Ly6Y6Iwg0K2iUvb60ssh0aVBTNEC8rKQIZ+PfX1E0JOsKilJqA4Pq4jEFw8nncqVugGuVWDPOT5w0aBIKn+yDy44ELFdhHvbuY2q1tTkYVBrx085AjdJYC7qDCY5XhWiabwfWpLHAc9GwkVuGhvKS1IwDaamkDWiQirwyfcenBuTgcTGx240SdafdgQe36+ERik73xZWBxw5ZJeZFJ3GKD0I+GN1dMgmG244pKpWVyoH1K0Zl1SFrEmps/Jqp9DyPEp0sFpxZ+YzPnTiTM8ewLLWQ2Kx0ldAIjRlHCUoNWFZInTV6ITXhXbMBmQJN1kqcAPNRXb/odroGncs6l2ILZbU/PUOVEgVDLxLLQCjpOHskShQPkK5UCUozTgcHr2aW62P/sJ/R/1q7/39rx0Gb/b2ym3v+6Fy6hv5hPwCHQrSk2y0ryKt7xgvRMF7dGwl4LzpdCfPp24R4CHEs86sVfO2AF8tONaRxBd/ULV+QlFjwZ+SPUqyKyisfQHKoV+RzSpK49jazr7Qw7Iriq2Ncv7nxPpJtbo/kUW4smFrsh1yjQKlh2cni40yuMbn72PddCyNHjUOzIre+qlgRCbUtzww/BySPgt7aju5tFEkZ1sSYS9phCaqrT+2Y3hzqCpTxQpItZfX85jSAwVg13AClL20VnLxyON8ZkOF5C9U1DR1EoOYj2YHljAYbGFlxqzkTy28vdHcm5JUNIEVSa9yF2P5iUdfnZvMEB3gu8Ge0pu3+RC54Q9EeYaD1gBPMf3cfV+ROW2bgIhyMMYPyZ5JsgSh+j5475Xj8yHsF4xyNgkLw8oLBWiZ3jc8O8NiygtZJ0ov0vOzOULnZuiLwWECr0XEJtb84sh21NenkXnX3QxN9EM4+yztTx1AMJb3DSIh0gGQDWAaLeeau7WvbFzLT00tIb+Mf/QwfkMd3QL7B/n5jem+jMbb63+qSqwTIzOG1kykTCcEL2f6bb56PqwJWr7qfdrTIHdUCZZA8a6JIpYq0ILdCNlCuyBeXTtQRI/R2PsGzRcFTQPmCRRvIGtaeh/LcufPYdJjP8CXj8RlQbDIyG0ZcOC1T4wPAi2bRHgC2qnTo2cM+B8keSLaeG1VkKzqOJ9e+63Gi7bvODXaM2MWA5IYzJkHyIxiTMEuCFGdREglyEcYlwUh7JC0O70eWRiYb6HA2PTnNrVooGK3taXI8xvrtgJD5N3f4D4yXZs/YUQRjxe5VCVuQxmIsh3N0SZjQzlES5Y+2nYYmB9GW9JlWRgD57hZClqPONpkbGo9ub2ta4ZlJ2t82MVPHtHT8uK8ZJ3RfqMHdBbaGmCyLj/Y1xb2tFs7AnbMXh8dxznScg96yWoMkLTUqkpRMtUKxicBow/jIGF0i7rDdtPykxchEOR3O9dFWH9edCHEb5zl3N6Z+t9Nc3sximPTmxofkFa7LgGLB2l2epPbudCGPrRYOgCiw9XvDnZO1E8tOWgPNTpCzfgeQwLVkoLK6HFxR19gbfj0HYRBvQMmInxwyCuTekJGkqJlxdxn3sxSE15BEdHfAMpEIR3t7AB4hOg2u2guyTH2Yd8ibstjl8ibU27SWECXXCLgqcoH7Fe2Dm+gkoQm3x7SPcVOGYoD+fpkYMO2mxUDHWZYYsPmoMm4X/IuIBzMFIDadkigaZGOc0KwJdm3sVo9liStMN0xSSdoYHTSgV0nKtZBZ27OljSuYUYS2eMbDnaBOdEJQSEkAbelCxoG0M/tGtLnsOLRNgjsSLj5wIZfh6uisNJTWR9I5lsxblnQ7LGVU08RvNcZdwnd4iFcUmL4YDi0yG5al2H773/E50pD0TSts86wCKbQoxFgIZIkbxyy/3H4gtK6EZHrXzKm7dpvH+CBxRrdCHqgsjbMtoTiSBvROjFSphiYLPZWTGK22RqhK4h5Dafz9kxwU+vppzW+e1vzNk5oPsguLZzlcZpBpfNVRXmRhDU1hxLLLAUStgzuXi+isxyjhMjS3HrP1G4oN0y4+zx02dJZwdZvPWD1bxivc1GzEs3Wu3ZzOozOdR9bTRYszP5MFbd3ZtiwZIRR77NuykT3Mo7Twk+0nF2wkE7YTh8M1cGNnZBmVHA4o213Y3gIQBdpsiuH2Nc+sN7R4qEW1rlmTx3qWhDWwXiricMjXDjqIL+mKDIl8E0LI45SdVdB2nRfc8IZ2X0bVMwjpiwrjQEGe/va37mFLZ5hgyN24Dns2kh0cDus2a3eatfXDaI3RiFVKJ4ch6vISLnK30i3mJJFbnRguIxqXCSedz8KsS5vcG8BsKOeQZyK7Jnb9BLdSEErS9uXxQcACTU6BLXIqMYnqWjpWcZevDdD3zXWhH3PFzd7sR7xU7lG7u9dGNveF6zXrryowLJk1y760C3kZcysq2jcj/CRLm0Xg9n9+ICUUDFPC6DJB+V0JnI2oGKkro/tPMiP1Ykskr3ziq/SX0RFK9kNjxOxMDlnD8bsSbWMXf4pz5CMailWwX5sHRJYoYxW6o1NaleVBnUiIiLpcA98KWbC8CTf73EyBbQxGm+L5626oqs0U74u2u2SOe4394dMXUgg5MgSk2a850KH6k7jyzwggqkTIMydP1hmkZQZD7V+WmaZFmBKjiGzqoQQ95ZhtaVYA31lBnqcTJzTszJo/XGdKrZaVtru6ZvzBh60V8HLEjarb/J5lfaq2xUZWLJ4UtvT/fH/95v/mCvGhpTgZYSvSNP2i6I+9QNq2NK6pOqrtkLmtispXai9VuE9tetMXosliDdHPcFRsOmERGsmdKU9jOepE9ilp6tXUPmvXhwp2DAP0GHEqKqxmTVXu/rdFotjydPftBVhZwlutbaO1puohrZIKO53lIGLuK/SS8R1IdtaGTS2jTFHlGltxNbTxVZbesUWIx1rQMpG4Lswz9HaePTRgOFxBbqbfyGrPg1HbXp/vac3KtRNgl3B22jSMPzPu54Yf7clhT9vH3A1+9+lfIeowbc1kui2sLXqJNO20NLTAQGgOLOgdnpn1Gso0IXcfbeZ2KEEv0FOuHuukkmKZuvXuE8acK0kbspW0QjOsr1GY4N28YG18mii2p6dE277JtQzQQ5r1ZTC6kCXYPNRCkWYs01yXuZWwZ6JT9tDglKMrVJ71GDi5r7kdSvf8yg7/zVwCj2WWiyRbbqZgxHBXydRDltvG1MNZxmrxLE6ukoidnnCrlT3VM1E/UkNWZq0GXoUC+rDy25rm+VAt8BA7nsoSd5ncScmXL+N9lFlto6AwBoevNJwIZ6MXUMnMsJk1/92ZxJlY8TovUo68vjAIZ9izyStpMlqbNliWIbakgQavDODk738d+i0YdrlEbfdRF7cNSihYifGuAY2LvHQzgkVeupmdYkdzK+MMvmlGCw3SB2EWWPFG9D5XhMfYLPtTMZ5Os23Wtuw0vlVJbmkxEzYpmqx96d2mtOB2gLkTIkt29jlf09Km4ZzpUphJmsi9mVXOruk0ayw7npxYDYyfqz+VuwJ7NtysWBZXnIqSpZmVA2V6Hb2+4ynZFYNFIqygqTBCneUZuCZTzrTZJJkSywfOFkmtXJkVXl0yEloTkWJRl5nWm6jLXAvOiLIL1rOk0KBmi6+asI41XnkfoQYbDK+KzydVi+qlSlvHqadMM8xsxfBKplTMDsD1Ba6kDu8Ow4M8cdTnhMvS0CxX6Hxx2YHqvKK7tJbXtp8JB8SvJMvVOgOsTPtn9oaHjRB54tgpQG9BpG+ACgZ+0bTxxbjLGMDdgmsaD+69j5m2FkWmojq4O5bENuTbIpBEZPRKa6kStGUhUjRMFZ2RE5EZ3tctZ80xtecL/SvDQoDvpE3Dcu3huA4DqxIG+iPJUWeJ0jhHvVCUlrClXa2vL5AbrikamtNhKPTFcoVdEHT+6lkDYiRT9H68xLAR1bqlSh2yRaplzq2QmGpHDCFLfI3nWN7hamRWrB5yCx98qd0ly+Ftfi+1Y6ZNDwUnjtc+z26JI0Nxoc7pzEPkeeeJ9tEiZE6J3Wzzcc2oNGRZosWd0jHtOlqzPxad0sGQVl6iKL/SJDsl6o3GcCPOdEpUZkaSncw8aaZgpj+3v3EwfaqvYvP7U2RNchC3t+wzo8mRYpqqxzUL9YRgWeLzjjSRdXy2lNWZNTEDd8chTCXJGM/zqfE1fSdd6n2W3nf7zl+WPSUj4oDaIsymoe3JcBzGvTKz6iHt7fhgtgzjyZH50zxhKGRGpkOsWxmxPhHqVvA1s6YlvSdjuGD/FhX0Fd8/msW7IfityNeOhpsCYqB+SnJTg8OCFAyLROfsEiWdrzyD9l9oeRiJnFuuYwTy+VIdw4HZwSdRl7PBJ/QXsrRpHGJYqFHL3DJrl0K9+zRjaaCKzjzjPdLQ4zvrmCh0nRmWtJekfO0Ab7+3hS7+wkzEm6l3UZlZZVskPm980balssk7pObbuAM50Ttgh1Llz98zhpXWDc3T5sNImmk/vILBBzOmzqxkCnMbc53NMaJ0ya8iWZpp3fy+vuSy0xMGmQSq8k7DRSEsUgIX2r4A2QKFV3BMnr6rmcqOpA1tKRuqN0jzlZi5kRUvF5ZEVzZd3vFhF2ZS3cZbGR1elt3Epn0teKVIsIwTyZy172LJnKGpcgVp4jqMheifFH3qZNbO+vLPO9IKxpFBsexwOi5kzfwLDhbErKleKnes4LuSKTxRO1vGe1mEJWXShVEWa67m2lJDPTlRatXmIdY+S6sEp9ExpiCFL4yzLyxa1dnFiH0A35UbGQC2nb46N9PgjrMn52pMOWTdfGSZwzoy7sXQ6ZmO0OW8Y934rgJWOPypG4jmDkSzprr41Ad66vZVWacPgNTlJVl3z+ajzPu5nPs6c/b8Ha1Bv5xOOrhC9xwKhx2W3dk94VzvA1UeatvVRMjoTd9pRPlym+JcQNk501nKIHhAJdSgZ8pdbZX9ADetH59/SY6MX3cy3aUJE390sWjSn8S/yUCV9DA8MqC07PDKi4l3dV1KZeKC1Rh38cuo+vcycvZ4iqIbk389lXn8irB2/wP+++7KR3SmbqDrL1XNGGTe5aoxPX/H0IuYXqSNEkLRGyc5EbJEB6N2F7Rpt6AekUgoIH0dijsvZzyUgHQACfZ4nhZG0lkGsLfQlaJAh9Jdo2jfycBUEF5sizeXRO/6JuFaBW8K+isx8IIlIU2be8aLuithLelh7brr3yURcJJ3SaRzdqCSM778Tun0XlTfevwOHHzVurvLx9G2sxFdeji4crF/uY6zvBDMZ9MoL/G3cHq1hD3UokUn3fxYwqbra2XaTrZCuXvIkmtwKxAuNTkUNpMDNcPEJv51P9hB2DLuL6MtBN8DZxjI869TP4rOla713gBwad/y6RawU9bj8ujoEDFOfhaV0lTtzArf8QqUJv8IrzOfvjfLaFXGgev4DbOLpHR6xWL62tmAOro+n+nj81Ji+jgkYt+g+qxkLORoNKLjWh7XTIl1bnFoTO2DxSF3n39N3j8fvGaRGJ2B/0Cs0b1ZNh7jYjLdlYBMX1ONH8IbN42OXRvvqJLWOnfvs/qR1UDuou+Hgn7Be61S7JPvt9pRtVu+x36iagfKr5Ihs8KxPsDR7rf+0hWOb7Gx13bGL3728msHZAePBLhZgZKUDPePfc5d3Dl12fWmpg9ws1nfvH23VBL+9efbv//tZnN98/YdDjfp/otJ9Dd/+SEX/c1ffliK/vb1TS7629c359Cb8u1S1F8+vj2HpnbhcpizcJ9/un29AO/mZvGkfv7p9ubm7HwazOVsYDDPc4Da0YzF//zT7YJ1N5jrvNHj88tws2YAn1+EmzkL66XzkMH8iLuA89WO5qEuxsxctbevb75btm6InbVyiH1+7R4fd+8Wd/lf/3o31dn/FwAA//85Xi3p"
}
