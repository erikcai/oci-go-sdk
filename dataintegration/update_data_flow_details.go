// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateDataFlowDetails Properties used in data flow update operations.
type UpdateDataFlowDetails struct {
	Details *DataFlowDetails `mandatory:"false" json:"details"`

	RegistryInfo *RegistryInfo `mandatory:"false" json:"registryInfo"`
}

func (m UpdateDataFlowDetails) String() string {
	return common.PointerString(m)
}
