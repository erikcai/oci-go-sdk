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

// GetPublicIpByPrivateIpIdDetails Private IP details to fetch the public IP which is assigned to it.
type GetPublicIpByPrivateIpIdDetails struct {

	// OCID of the private IP.
	PrivateIpId *string `mandatory:"true" json:"privateIpId"`
}

func (m GetPublicIpByPrivateIpIdDetails) String() string {
	return common.PointerString(m)
}
