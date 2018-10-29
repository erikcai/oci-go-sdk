// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace REST API endpoint
//
// Marketplace REST API for loom plugin
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DeviceType The model of device type
type DeviceType struct {

	// The name of the device type
	Name *string `mandatory:"false" json:"name"`

	// The code of the device type
	Code *string `mandatory:"false" json:"code"`
}

func (m DeviceType) String() string {
	return common.PointerString(m)
}
