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

// TermCollection Results of a Terms Listing. Terms are defined in business glossary and are used in tagging catalog objects.
type TermCollection struct {

	// Collection of Terms
	Items []TermSummary `mandatory:"true" json:"items"`
}

func (m TermCollection) String() string {
	return common.PointerString(m)
}
