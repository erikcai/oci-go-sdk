// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
//
// Cloud Incident Management Service
//

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Tenant Customer Tenant
type Tenant struct {
	CustomerSupportIdentifier *string `mandatory:"true" json:"customerSupportIdentifier"`

	TenantId *string `mandatory:"true" json:"tenantId"`
}

func (m Tenant) String() string {
	return common.PointerString(m)
}
