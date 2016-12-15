//
// Copyright 2017, Rutger te Nijenhuis & Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ciscoasa

var icmpType = map[string]string{
	"echo-reply":           "0",
	"unreachable":          "3",
	"source-quench":        "4",
	"redirect":             "5",
	"alternate-address":    "6",
	"echo":                 "8",
	"router-advertisement": "9",
	"router-solicitation":  "10",
	"time-exceeded":        "11",
	"parameter-problem":    "12",
	"timestamp-request":    "13",
	"timestamp-reply":      "14",
	"information-request":  "15",
	"information-reply":    "16",
	"mask-request":         "17",
	"mask-reply":           "18",
	"conversion-error":     "31",
	"mobile-redirect":      "32",
}
