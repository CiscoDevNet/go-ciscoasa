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
	"strings"
)

// NetworkObjectGroupCollection represents a collection of network object groups.
type NetworkObjectGroupCollection struct {
	RangeInfo RangeInfo             `json:"rangeInfo"`
	Items     []*NetworkObjectGroup `json:"items"`
	Kind      string                `json:"kind"`
	SelfLink  string                `json:"selfLink"`
}

// NetworkObjectGroup represents a network object group.
type NetworkObjectGroup struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Members     []*AddressObject `json:"members"`
	Kind        string           `json:"kind"`
	ObjectID    string           `json:"objectId,omitempty"`
	SelfLink    string           `json:"selfLink,omitempty"`
}

// AddressObject represents an address object.
type AddressObject struct {
	Value    string `json:"value,omitempty"`
	Kind     string `json:"kind"`
	ObjectID string `json:"objectId,omitempty"`
	Reflink  string `json:"refLink,omitempty"`
}

// String returns the objects address
func (o *AddressObject) String() string {
	if strings.HasPrefix(o.Kind, "objectRef#") {
		return o.ObjectID
	}
	if strings.Contains(o.Kind, "IPv4Address") {
		return o.Value + "/32"
	}
	if strings.Contains(o.Kind, "IPv6Address") {
		return o.Value + "/128"
	}
	return o.Value
}

// ListNetworkObjectGroups returns a collection of network object groups.
func (s *objectsService) ListNetworkObjectGroups() (*NetworkObjectGroupCollection, error) {
	u := "/api/objects/networkobjectgroups"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	n := &NetworkObjectGroupCollection{}
	_, err = s.do(req, n)

	return n, err
}

// CreateNetworkObjectGroup creates a new network object group.
func (s *objectsService) CreateNetworkObjectGroup(name, description string, addresses []string) (*NetworkObjectGroup, error) {
	u := "/api/objects/networkobjectgroups"

	n := NetworkObjectGroup{
		Name:        name,
		Description: description,
		Kind:        "object#NetworkObjGroup",
	}

	for _, address := range addresses {
		o, err := s.objectFromAddress(address)
		if err != nil {
			return nil, err
		}

		n.Members = append(n.Members, o)
	}

	req, err := s.newRequest("POST", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetNetworkObjectGroup(name)
}

// GetNetworkObjectGroup retrieves a network object group.
func (s *objectsService) GetNetworkObjectGroup(name string) (*NetworkObjectGroup, error) {
	u := fmt.Sprintf("/api/objects/networkobjectgroups/%s", name)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	n := &NetworkObjectGroup{}
	_, err = s.do(req, n)

	return n, err
}

// UpdateNetworkObjectGroup updates a network object group.
func (s *objectsService) UpdateNetworkObjectGroup(name, description string, addresses []string) (*NetworkObjectGroup, error) {
	u := fmt.Sprintf("/api/objects/networkobjectgroups/%s", name)

	n := NetworkObjectGroup{
		Name:        name,
		Description: description,
		Kind:        "object#NetworkObjGroup",
	}

	for _, address := range addresses {
		o, err := s.objectFromAddress(address)
		if err != nil {
			return nil, err
		}

		n.Members = append(n.Members, o)
	}

	req, err := s.newRequest("PUT", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetNetworkObjectGroup(name)
}

// DeleteNetworkObjectGroup deletes a network object group.
func (s *objectsService) DeleteNetworkObjectGroup(name string) error {
	u := fmt.Sprintf("/api/objects/networkobjectgroups/%s", name)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
