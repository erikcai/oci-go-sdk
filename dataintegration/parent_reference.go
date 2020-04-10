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

// ParentReference A reference to the object's parent
type ParentReference struct {

	// Key of the parent object
	Parent *string `mandatory:"false" json:"parent"`
}

func (m ParentReference) String() string {
	return common.PointerString(m)
}
