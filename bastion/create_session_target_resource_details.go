// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastion

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v25/common"
)

// CreateSessionTargetResourceDetails details of the target resource of a session
type CreateSessionTargetResourceDetails interface {

	// target resource port to use.
	GetTargetResourcePort() *int
}

type createsessiontargetresourcedetails struct {
	JsonData           []byte
	TargetResourcePort *int   `mandatory:"false" json:"targetResourcePort"`
	SessionType        string `json:"sessionType"`
}

// UnmarshalJSON unmarshals json
func (m *createsessiontargetresourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatesessiontargetresourcedetails createsessiontargetresourcedetails
	s := struct {
		Model Unmarshalercreatesessiontargetresourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TargetResourcePort = s.Model.TargetResourcePort
	m.SessionType = s.Model.SessionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createsessiontargetresourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SessionType {
	case "MANAGED_SSH":
		mm := CreateManagedSshSessionTargetResourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PORT_FORWARDING":
		mm := CreatePortForwardingSessionTargetResourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetTargetResourcePort returns TargetResourcePort
func (m createsessiontargetresourcedetails) GetTargetResourcePort() *int {
	return m.TargetResourcePort
}

func (m createsessiontargetresourcedetails) String() string {
	return common.PointerString(m)
}
