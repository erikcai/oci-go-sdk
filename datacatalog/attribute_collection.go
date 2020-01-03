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

// AttributeCollection Results of a Attributes Listing. Attributes describe an iten of data with name and datatype
type AttributeCollection struct {

	// Collection of Attributes
	Items []AttributeSummary `mandatory:"true" json:"items"`
}

func (m AttributeCollection) String() string {
	return common.PointerString(m)
}
