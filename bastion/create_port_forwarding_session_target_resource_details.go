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
	"github.com/oracle/oci-go-sdk/v26/common"
)

// CreatePortForwardingSessionTargetResourceDetails details of the target resource of a port forwarding session
type CreatePortForwardingSessionTargetResourceDetails struct {

	// target resource port to use.
	TargetResourcePort *int `mandatory:"false" json:"targetResourcePort"`

	// The OCID of the target resource to connect to.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`

	// target resource private ip address.
	TargetResourcePrivateIpAddress *string `mandatory:"false" json:"targetResourcePrivateIpAddress"`
}

//GetTargetResourcePort returns TargetResourcePort
func (m CreatePortForwardingSessionTargetResourceDetails) GetTargetResourcePort() *int {
	return m.TargetResourcePort
}

func (m CreatePortForwardingSessionTargetResourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreatePortForwardingSessionTargetResourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreatePortForwardingSessionTargetResourceDetails CreatePortForwardingSessionTargetResourceDetails
	s := struct {
		DiscriminatorParam string `json:"sessionType"`
		MarshalTypeCreatePortForwardingSessionTargetResourceDetails
	}{
		"PORT_FORWARDING",
		(MarshalTypeCreatePortForwardingSessionTargetResourceDetails)(m),
	}

	return json.Marshal(&s)
}
