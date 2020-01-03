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

// FolderTagCollection Results of a Folders tag Listing. Folder tags allow association of Folder objects to business terms.
type FolderTagCollection struct {

	// Collection of Folder Tags
	Items []FolderTagSummary `mandatory:"true" json:"items"`
}

func (m FolderTagCollection) String() string {
	return common.PointerString(m)
}
