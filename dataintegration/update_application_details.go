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

// UpdateApplicationDetails Properties used in application create operations.
type UpdateApplicationDetails struct {
	Details *ApplicationDetails `mandatory:"false" json:"details"`

	RegistryInfo *RegistryInfo `mandatory:"false" json:"registryInfo"`
}

func (m UpdateApplicationDetails) String() string {
	return common.PointerString(m)
}
