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

// EntityTagCollection Results of a Entity Tags Listing. Entity Tags allow assciation of business terms with Entities.
type EntityTagCollection struct {

	// Collection of Entity Tags
	Items []EntityTagSummary `mandatory:"true" json:"items"`
}

func (m EntityTagCollection) String() string {
	return common.PointerString(m)
}
