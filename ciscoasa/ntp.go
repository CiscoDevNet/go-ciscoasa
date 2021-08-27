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

type devicesetupService struct {
	*Client
}

// NtpServerObjectCollection represents a collection of ntp server objects.
type NtpServerObjectCollection struct {
	RangeInfo RangeInfo          `json:"rangeInfo"`
	Items     []*NtpServerObject `json:"items"`
	Kind      string             `json:"kind"`
	SelfLink  string             `json:"selfLink"`
}

// NtpServerObject represents a ntp server object.
type NtpServerObject struct {
	IpAddress   string        `json:"ipAddress"`
	IsPreferred bool          `json:"isPreferred,omitempty"`
	Interface   *InterfaceRef `json:"interface,omitempty"`
	Key         struct {
		Number    string `json:"number"`
		Value     string `json:"value"`
		IsTrusted bool   `json:"isTrusted,omitempty"`
	} `json:"key"`
	Kind     string `json:"kind"`
	ObjectId string `json:"objectId,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// ListNtpServers returns a collection of NTP servers.
func (s *devicesetupService) ListNtpServers() (*NtpServerObjectCollection, error) {
	u := "/api/devicesetup/ntp/servers"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &NtpServerObjectCollection{}
	_, err = s.do(req, r)

	return r, err
}

// CreateNtpServer creates a NTP server.
func (s *devicesetupService) CreateNtpServer(ipAddress string, preferred bool, iface, knumber, kvalue string, ktrusted bool) error {
	u := "/api/devicesetup/ntp/servers"

	r := &NtpServerObject{
		IpAddress:   ipAddress,
		IsPreferred: preferred,
		Kind:        "object#NTPServer",
	}

	r.Key.Number = knumber
	r.Key.Value = kvalue
	r.Key.IsTrusted = ktrusted

	if iface != "" {
		i := &InterfaceRef{
			Kind: "objectRef#Interface",
		}
		if isInterfaceObjectId(iface) {
			i.ObjectId = iface
		} else {
			i.Name = iface
		}
		r.Interface = i
	}

	req, err := s.newRequest("POST", u, r)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

// GetNtpServer retrieves a NTP server.
func (s *devicesetupService) GetNtpServer(ipAddress string) (*NtpServerObject, error) {
	u := fmt.Sprintf("/api/devicesetup/ntp/servers/%s", ipAddress)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &NtpServerObject{}
	_, err = s.do(req, r)

	return r, err
}

// UpdateNtpServer updates a NTP server.
func (s *devicesetupService) UpdateNtpServer(objectId, ipAddress string, preferred bool, iface, knumber, kvalue string, ktrusted bool) error {
	u := fmt.Sprintf("/api/devicesetup/ntp/servers/%s", objectId)

	r := &NtpServerObject{
		IpAddress:   ipAddress,
		IsPreferred: preferred,
		Kind:        "object#NTPServer",
	}

	r.Key.Number = knumber
	r.Key.Value = kvalue
	r.Key.IsTrusted = ktrusted

	if iface != "" {
		i := &InterfaceRef{
			Kind: "objectRef#Interface",
		}
		if isInterfaceObjectId(iface) {
			i.ObjectId = iface
		} else {
			i.Name = iface
		}
		r.Interface = i
	}

	req, err := s.newRequest("PUT", u, r)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

// DeleteNtpServer deletes a NTP server.
func (s *devicesetupService) DeleteNtpServer(ipAddress string) error {
	u := fmt.Sprintf("/api/devicesetup/ntp/servers/%s", ipAddress)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

func isInterfaceObjectId(iface string) bool {
	return strings.Contains(iface, "_API_SLASH_")
}
