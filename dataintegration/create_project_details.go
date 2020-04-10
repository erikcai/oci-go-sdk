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

// CreateProjectDetails Properties used in project create operations.
type CreateProjectDetails struct {
	Details *ProjectDetailsForCreate `mandatory:"false" json:"details"`

	RegistryInfo *RegistryInfo `mandatory:"false" json:"registryInfo"`
}

func (m CreateProjectDetails) String() string {
	return common.PointerString(m)
}
