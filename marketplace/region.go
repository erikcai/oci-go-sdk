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

// Region The model of regions of an application
type Region struct {

	// The name of the region
	Name *string `mandatory:"false" json:"name"`

	// Code of the region
	Code *string `mandatory:"false" json:"code"`

	// Countries of the region
	Countries []Item `mandatory:"false" json:"countries"`
}

func (m Region) String() string {
	return common.PointerString(m)
}
