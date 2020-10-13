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
	"github.com/oracle/oci-go-sdk/v27/common"
)

// CreateManagedSshSessionTargetResourceDetails details of the target resource of a managed_ssh session
type CreateManagedSshSessionTargetResourceDetails struct {

	// name of the user to use on target resource operating system
	TargetResourceOperatingSystemUserName *string `mandatory:"true" json:"targetResourceOperatingSystemUserName"`

	// The OCID of the target resource to connect to.
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// target resource port to use.
	TargetResourcePort *int `mandatory:"false" json:"targetResourcePort"`

	// target resource private ip address.
	TargetResourcePrivateIpAddress *string `mandatory:"false" json:"targetResourcePrivateIpAddress"`
}

//GetTargetResourcePort returns TargetResourcePort
func (m CreateManagedSshSessionTargetResourceDetails) GetTargetResourcePort() *int {
	return m.TargetResourcePort
}

func (m CreateManagedSshSessionTargetResourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateManagedSshSessionTargetResourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateManagedSshSessionTargetResourceDetails CreateManagedSshSessionTargetResourceDetails
	s := struct {
		DiscriminatorParam string `json:"sessionType"`
		MarshalTypeCreateManagedSshSessionTargetResourceDetails
	}{
		"MANAGED_SSH",
		(MarshalTypeCreateManagedSshSessionTargetResourceDetails)(m),
	}

	return json.Marshal(&s)
}
