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

// Status The status of the request, specifically for installing application instances request
type Status struct {

	// The display status
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m Status) String() string {
	return common.PointerString(m)
}
