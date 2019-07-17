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

// CreateTransferApplianceDetails The representation of CreateTransferApplianceDetails
type CreateTransferApplianceDetails struct {
	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`
}

func (m CreateTransferApplianceDetails) String() string {
	return common.PointerString(m)
}
