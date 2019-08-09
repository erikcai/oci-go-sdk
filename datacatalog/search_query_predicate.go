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

// SearchQueryPredicate Query that conforms to the Query DSL in ES
type SearchQueryPredicate struct {

	// The set of attributes that is part of the search query criteria
	Fields []string `mandatory:"false" json:"fields"`

	// Qualifier query predicate expression. For example, this AND that OR thus
	Query *string `mandatory:"false" json:"query"`
}

func (m SearchQueryPredicate) String() string {
	return common.PointerString(m)
}
