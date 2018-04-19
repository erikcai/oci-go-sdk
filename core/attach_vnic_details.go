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

// AttachVnicDetails The representation of AttachVnicDetails
type AttachVnicDetails struct {

	// Details for creating new VNIC.
	CreateVnicDetails *CreateVnicDetails `mandatory:"true" json:"createVnicDetails"`

	// The OCID of the instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// A user-friendly name for attachment. Does not have to be unique, and it cannot be changed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Selects the physical network port on which to attach the VNIC. Defaults to 0.
	NicIndex *int `mandatory:"false" json:"nicIndex"`
}

func (m AttachVnicDetails) String() string {
	return common.PointerString(m)
}
