// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DTS API
//
// A description of the DTS API
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ShippingVendors The representation of ShippingVendors
type ShippingVendors struct {

	// List of available shipping vendors for package delivery
	Vendors []string `mandatory:"false" json:"vendors"`
}

func (m ShippingVendors) String() string {
	return common.PointerString(m)
}
