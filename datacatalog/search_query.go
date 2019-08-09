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

// SearchQuery Search Query object that allows complex search predicates that cannot be expressed through simple query params.
type SearchQuery struct {
	Predicate *SearchQueryPredicate `mandatory:"false" json:"predicate"`
}

func (m SearchQuery) String() string {
	return common.PointerString(m)
}
