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
)

// ListDhcpServers returns a collection of DHCP servers.
func (s *dhcpService) ListDhcpServers() (*DhcpServerCollection, error) {
	u := "/api/dhcp/servers"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &DhcpServerCollection{}
	_, err = s.do(req, r)

	return r, err
}

// UpdateDhcpServer updates a DHCP server.
func (s *dhcpService) UpdateDhcpServer(server *DhcpServer) (string, error) {
	u := fmt.Sprintf("/api/dhcp/servers/%s", server.ObjectId)

	req, err := s.newRequest("PUT", u, server)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)

	return idFromResponse(resp)
}

// GetDhcpServer retrieves a DHCP server.
func (s *dhcpService) GetDhcpServer(ipAddress string) (*DhcpServer, error) {
	u := fmt.Sprintf("/api/dhcp/servers/%s", ipAddress)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &DhcpServer{}
	_, err = s.do(req, r)

	return r, err
}
