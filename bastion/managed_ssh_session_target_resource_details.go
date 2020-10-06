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

// ManagedSshSessionTargetResourceDetails details of the target resource of a managed_ssh session
type ManagedSshSessionTargetResourceDetails struct {

	// target resource port to use.
	TargetResourcePort *int `mandatory:"true" json:"targetResourcePort"`

	// name of the user to use on target resource operating system
	TargetResourceOperatingSystemUserName *string `mandatory:"true" json:"targetResourceOperatingSystemUserName"`

	// The OCID of the target resource to connect to.
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// display name of the target compute instance
	TargetResourceDisplayName *string `mandatory:"true" json:"targetResourceDisplayName"`

	// target resource private ip address.
	TargetResourcePrivateIpAddress *string `mandatory:"false" json:"targetResourcePrivateIpAddress"`
}

//GetTargetResourcePort returns TargetResourcePort
func (m ManagedSshSessionTargetResourceDetails) GetTargetResourcePort() *int {
	return m.TargetResourcePort
}

func (m ManagedSshSessionTargetResourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ManagedSshSessionTargetResourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedSshSessionTargetResourceDetails ManagedSshSessionTargetResourceDetails
	s := struct {
		DiscriminatorParam string `json:"sessionType"`
		MarshalTypeManagedSshSessionTargetResourceDetails
	}{
		"MANAGED_SSH",
		(MarshalTypeManagedSshSessionTargetResourceDetails)(m),
	}

	return json.Marshal(&s)
}
