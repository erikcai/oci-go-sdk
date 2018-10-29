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

// Resource Model of primary resource for the package
type Resource struct {

	// Type of the service
	ServiceType *string `mandatory:"false" json:"serviceType"`

	// Type of the resource
	ResourceType *string `mandatory:"false" json:"resourceType"`
}

func (m Resource) String() string {
	return common.PointerString(m)
}
