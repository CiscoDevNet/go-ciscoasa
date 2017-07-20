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

import (
	"fmt"
	"net/http"
	"testing"
)

func TestObjectFromService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/objects/networkservices/443", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
  			"kind": "object#TcpUdpServiceObj",
  			"selfLink": "https://localhost/api/objects/networkservices/443",
  			"name": "443",
  			"value": "tcp/443",
			"objectId": "443"
		}`)
	})

	mux.HandleFunc("/api/objects/networkservicegroups/28ec2bb3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
  			"kind": "object#NetworkServiceGroup",
  			"selfLink": "https://localhost/api/objects/networkservicegroups/28ec2bb3",
  			"name": "28ec2bb3",
  			"members": [
  			{
  				"kind": "TcpUdpService",
  				"value": "tcp/2121"
  			},
  			{
  				"kind": "TcpUdpService",
  				"value": "udp/123"
  			},
  			{
  				"kind": "TcpUdpService",
  				"value": "tcp/https"
  			}],
			"objectId": "28ec2bb3"
		}`)
	})

	_, err := client.Objects.objectFromService("443")
	if err != nil {
		t.Errorf("Failed to create ServiceObject for ServiceObject '443': %s", err)
	}

	_, err = client.Objects.objectFromService("28ec2bb3")
	if err != nil {
		t.Errorf("Failed to create ServiceObject for ServiceGroup '28ec2bb3': %s", err)
	}

	o, err := client.Objects.objectFromService("tcp/443")
	if err != nil {
		t.Errorf("Failed to create ServiceObject for service 'tcp/443': %s", err)
	}
	if o.Value != "tcp/443" || o.Kind != "TcpUdpService" {
		t.Errorf("Wrong values while createing Service object, expected 'tcp/443' and 'TcpUdpService', got %s and %s", o.Value, o.Kind)
	}
}
