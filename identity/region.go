// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Region The representation of Region
type Region struct {

	// The key of the region such as PHX, IAD.
	Key *string `mandatory:"false" json:"key"`

	// The name of the region such as us-az-phoenix.
	Name *string `mandatory:"false" json:"name"`
}

func (m Region) String() string {
	return common.PointerString(m)
}
