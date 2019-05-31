// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Bdtstest1 API
//
// A description of the Bdtstest1 API
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateTransferApplianceEntitlementDetails The representation of CreateTransferApplianceEntitlementDetails
type CreateTransferApplianceEntitlementDetails struct {
	TenantId *string `mandatory:"false" json:"tenantId"`

	RequestorName *string `mandatory:"false" json:"requestorName"`

	RequestorEmail *string `mandatory:"false" json:"requestorEmail"`
}

func (m CreateTransferApplianceEntitlementDetails) String() string {
	return common.PointerString(m)
}
