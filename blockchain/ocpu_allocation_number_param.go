// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OcpuAllocationNumberParam OCPU allocation parameter
type OcpuAllocationNumberParam struct {

	// Number of OCPU allocation
	OcpuAllocationNumber *float32 `mandatory:"false" json:"ocpuAllocationNumber"`
}

func (m OcpuAllocationNumberParam) String() string {
	return common.PointerString(m)
}
