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

// ConnectionCollection Results of a Connections Listing. Each member of the result is a summary representation of a connection to a Data Asset.
type ConnectionCollection struct {

	// Collection of Connection Summaries
	Items []ConnectionSummary `mandatory:"true" json:"items"`
}

func (m ConnectionCollection) String() string {
	return common.PointerString(m)
}
