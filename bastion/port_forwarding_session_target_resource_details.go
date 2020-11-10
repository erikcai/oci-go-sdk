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
	"github.com/oracle/oci-go-sdk/v28/common"
)

// PortForwardingSessionTargetResourceDetails details of the target resource of a port forwarding session
type PortForwardingSessionTargetResourceDetails struct {

	// target resource port to use.
	TargetResourcePort *int `mandatory:"true" json:"targetResourcePort"`

	// The OCID of the target resource to connect to.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`

	// target resource private ip address.
	TargetResourcePrivateIpAddress *string `mandatory:"false" json:"targetResourcePrivateIpAddress"`

	// display name of the target compute instance
	TargetResourceDisplayName *string `mandatory:"false" json:"targetResourceDisplayName"`
}

//GetTargetResourcePort returns TargetResourcePort
func (m PortForwardingSessionTargetResourceDetails) GetTargetResourcePort() *int {
	return m.TargetResourcePort
}

func (m PortForwardingSessionTargetResourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PortForwardingSessionTargetResourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePortForwardingSessionTargetResourceDetails PortForwardingSessionTargetResourceDetails
	s := struct {
		DiscriminatorParam string `json:"sessionType"`
		MarshalTypePortForwardingSessionTargetResourceDetails
	}{
		"PORT_FORWARDING",
		(MarshalTypePortForwardingSessionTargetResourceDetails)(m),
	}

	return json.Marshal(&s)
}
