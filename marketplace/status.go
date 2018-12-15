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

// Status The status of the application packages
type Status struct {

	// The name of the status
	Name *string `mandatory:"false" json:"name"`
}

func (m Status) String() string {
	return common.PointerString(m)
}
