// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// StorageGateway API
//
// API for interfacing with StorageGateway
//

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeStorageGatewayCompartmentDetails Contains details indicating which compartment the resource should move to
type ChangeStorageGatewayCompartmentDetails struct {

	// The OCID of the new compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeStorageGatewayCompartmentDetails) String() string {
	return common.PointerString(m)
}
