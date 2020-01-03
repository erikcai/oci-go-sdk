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

// GlossaryCollection Results of a Glossaries Listing.  Glossary is an organizing concept for business terms to provide a unified semantic model across disparate data assets
type GlossaryCollection struct {

	// Collection of Glossaries
	Items []GlossarySummary `mandatory:"true" json:"items"`
}

func (m GlossaryCollection) String() string {
	return common.PointerString(m)
}
