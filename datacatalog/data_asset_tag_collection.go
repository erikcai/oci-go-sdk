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

// DataAssetTagCollection Results of a Data Asset Tag Listing. Data Asset tags represent an association of a Data Asset to a term.
type DataAssetTagCollection struct {

	// Collection of Data Asset Tags
	Items []DataAssetTagSummary `mandatory:"true" json:"items"`
}

func (m DataAssetTagCollection) String() string {
	return common.PointerString(m)
}
