// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SoftwarePackageDependency A dependecy for a software package
type SoftwarePackageDependency struct {

	// the software package's dependency
	Dependency *string `mandatory:"false" json:"dependency"`

	// the type of the dependency
	DependencyType *string `mandatory:"false" json:"dependencyType"`

	// the modifier for the dependency
	DependencyModifier *string `mandatory:"false" json:"dependencyModifier"`
}

func (m SoftwarePackageDependency) String() string {
	return common.PointerString(m)
}
