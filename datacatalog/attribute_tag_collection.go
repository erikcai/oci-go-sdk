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

// AttributeTagCollection Results of a Attribute Tags Listing. Attribnute tags allow association of business terms with attributes.
type AttributeTagCollection struct {

	// Collection of Attribute Tags
	Items []AttributeTagSummary `mandatory:"true" json:"items"`
}

func (m AttributeTagCollection) String() string {
	return common.PointerString(m)
}
