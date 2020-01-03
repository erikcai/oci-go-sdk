// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TypeCollection Results of a Types Listing. Types define the basic type of catalog objects and are immutable.
type TypeCollection struct {

	// Collection of Types
	Items []TypeSummary `mandatory:"true" json:"items"`
}

func (m TypeCollection) String() string {
	return common.PointerString(m)
}
