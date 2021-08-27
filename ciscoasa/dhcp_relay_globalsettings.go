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

// UpdateDhcpRelayGlobalsettings updates a DHCP Relay Global Settings.
func (s *dhcpService) UpdateDhcpRelayGlobalsettings(gs *DhcpRelayGS) error {
	u := "/api/dhcp/relay/servers/globalsettings"

	req, err := s.newRequest("PUT", u, gs)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

// GetDhcpRelayGlobalsettings retrieves a DHCP Relay Global Settings.
func (s *dhcpService) GetDhcpRelayGlobalsettings() (*DhcpRelayGS, error) {
	u := "/api/dhcp/relay/servers/globalsettings"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &DhcpRelayGS{}
	_, err = s.do(req, r)

	return r, err
}
