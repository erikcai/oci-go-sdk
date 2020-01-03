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

// EntityCollection Results of a Entities Listing. Entities are representation of a dataset with a set of attributes.
type EntityCollection struct {

	// Collection of Entities
	Items []EntitySummary `mandatory:"true" json:"items"`
}

func (m EntityCollection) String() string {
	return common.PointerString(m)
}
