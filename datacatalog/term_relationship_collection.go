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

// TermRelationshipCollection Results of a Terms Relationship Listing. Term Relationships are associations between two terms in business glossary.
type TermRelationshipCollection struct {

	// Collection of Term Relationships
	Items []TermRelationshipSummary `mandatory:"true" json:"items"`
}

func (m TermRelationshipCollection) String() string {
	return common.PointerString(m)
}
