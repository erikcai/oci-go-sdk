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

// User User of the application instance
type User struct {

	// Unique identifier of the user
	Id *string `mandatory:"false" json:"id"`

	// Name of the user
	Name *string `mandatory:"false" json:"name"`
}

func (m User) String() string {
	return common.PointerString(m)
}
