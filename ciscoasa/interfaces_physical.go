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

import "encoding/json"

type interfaceService struct {
	*Client
}

// InterfaceCollection represents a collection of interfaces.
type InterfaceCollection struct {
	RangeInfo RangeInfo    `json:"rangeInfo"`
	Items     []*Interface `json:"items"`
	Kind      string       `json:"kind"`
	SelfLink  string       `json:"selfLink"`
}

// Interface represents an interface.
type Interface struct {
	HardwareID        string     `json:"hardwareID"`
	InterfaceDesc     string     `json:"interfaceDesc"`
	ChannelGroupID    string     `json:"channelGroupID"`
	ChannelGroupMode  string     `json:"channelGroupMode"`
	Duplex            string     `json:"duplex"`
	FlowcontrolOn     bool       `json:"flowcontrolOn"`
	FlowcontrolHigh   int        `json:"flowcontrolHigh"`
	FlowcontrolLow    int        `json:"flowcontrolLow"`
	FlowcontrolPeriod int        `json:"flowcontrolPeriod"`
	ForwardTrafficCX  bool       `json:"forwardTrafficCX"`
	ForwardTrafficSFR bool       `json:"forwardTrafficSFR"`
	LacpPriority      int        `json:"lacpPriority"`
	ActiveMacAddress  string     `json:"activeMacAddress"`
	StandByMacAddress string     `json:"standByMacAddress"`
	ManagementOnly    bool       `json:"managementOnly"`
	Mtu               int        `json:"mtu"`
	Name              string     `json:"name"`
	SecurityLevel     int        `json:"securityLevel"`
	Shutdown          bool       `json:"shutdown"`
	Speed             string     `json:"speed"`
	IPAddress         *IPAddress `json:"ipAddress"`
	Ipv6Info          *IPv6Info  `json:"ipv6Info"`
	Kind              string     `json:"kind"`
	ObjectID          string     `json:"objectId"`
	SelfLink          string     `json:"selfLink"`
}

// IPAddress represents an IPv4 address.
type IPAddress struct {
	IP struct {
		Kind  string `json:"kind"`
		Value string `json:"value"`
	} `json:"ip"`
	NetMask struct {
		Kind  string `json:"kind"`
		Value string `json:"value"`
	} `json:"netMask"`
	Kind string `json:"kind"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ip *IPAddress) UnmarshalJSON(b []byte) error {
	type alias IPAddress
	if err := json.Unmarshal(b, (*alias)(ip)); err != nil {
		ip = nil
	}
	return nil
}

// IPv6Info represents an IPv6 address.
type IPv6Info struct {
	Enabled                  bool     `json:"enabled"`
	AutoConfig               bool     `json:"autoConfig"`
	EnforceEUI64             bool     `json:"enforceEUI64"`
	ManagedAddressConfig     bool     `json:"managedAddressConfig"`
	NsInterval               int      `json:"nsInterval"`
	DadAttempts              int      `json:"dadAttempts"`
	NDiscoveryPrefixList     []string `json:"nDiscoveryPrefixList"`
	OtherStatefulConfig      bool     `json:"otherStatefulConfig"`
	RouterAdvertInterval     int      `json:"routerAdvertInterval"`
	RouterAdvertIntervalUnit string   `json:"routerAdvertIntervalUnit"`
	RouterAdvertLifetime     int      `json:"routerAdvertLifetime"`
	SuppressRouterAdvert     bool     `json:"suppressRouterAdvert"`
	ReachableTime            int      `json:"reachableTime"`
	Ipv6Addresses            []string `json:"ipv6Addresses"`
	Kind                     string   `json:"kind"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ip *IPv6Info) UnmarshalJSON(b []byte) error {
	type alias IPv6Info
	if err := json.Unmarshal(b, (*alias)(ip)); err != nil {
		ip = nil
	}
	return nil
}

// ListPhysicalInterfaces returns a collection of interfaces.
func (s *interfaceService) ListPhysicalInterfaces() (*InterfaceCollection, error) {
	u := "/api/interfaces/physical"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	e := &InterfaceCollection{}
	_, err = s.do(req, e)

	return e, err
}
