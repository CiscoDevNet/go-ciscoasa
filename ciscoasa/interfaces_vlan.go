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

// VlanInterfaceCollection represents a collection of vlan interfaces.
type VlanInterfaceCollection struct {
	RangeInfo RangeInfo        `json:"rangeInfo"`
	Items     []*VlanInterface `json:"items"`
	Kind      string           `json:"kind"`
	SelfLink  string           `json:"selfLink"`
}

// VlanInterface represents an vlan interface.
type VlanInterface struct {
	HardwareID        string     `json:"hardwareID"`
	InterfaceDesc     string     `json:"interfaceDesc"`
	ForwardTrafficCX  bool       `json:"forwardTrafficCX"`
	ForwardTrafficSFR bool       `json:"forwardTrafficSFR"`
	ActiveMacAddress  string     `json:"activeMacAddress"`
	StandByMacAddress string     `json:"standByMacAddress"`
	ManagementOnly    bool       `json:"managementOnly"`
	Mtu               int        `json:"mtu"`
	Name              string     `json:"name"`
	SecurityLevel     int        `json:"securityLevel"`
	Shutdown          bool       `json:"shutdown"`
	VlanID            int        `json:"vlanID"`
	IPAddress         *IPAddress `json:"ipAddress"`
	Ipv6Info          *IPv6Info  `json:"ipv6Info"`
	Kind              string     `json:"kind"`
	ObjectID          string     `json:"objectId"`
	SelfLink          string     `json:"selfLink"`
}

// ListVlanInterfaces returns a collection of interfaces.
func (s *interfaceService) ListVlanInterfaces() (*VlanInterfaceCollection, error) {
	u := "/api/interfaces/vlan"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	e := &VlanInterfaceCollection{}
	_, err = s.do(req, e)

	return e, err
}
