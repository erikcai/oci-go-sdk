// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PrivateIpAddressMapping The subnet IP address's assignment.
type PrivateIpAddressMapping struct {

	// Creation time of the mapping.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// OCID of the VNIC to which the IP address is mapped.
	VnicId *string `mandatory:"true" json:"vnicId"`
}

func (m PrivateIpAddressMapping) String() string {
	return common.PointerString(m)
}
