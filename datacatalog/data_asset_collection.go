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

// DataAssetCollection Results of a Data Assets Listing. A Data Asset is often synonymous with a 'System', such as a Database, or may be a file container, or a message stream
type DataAssetCollection struct {

	// Collection of Data Asset Summaries
	Items []DataAssetSummary `mandatory:"true" json:"items"`
}

func (m DataAssetCollection) String() string {
	return common.PointerString(m)
}
