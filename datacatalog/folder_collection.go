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

// FolderCollection Results of a Folders Listing. Folders are external organization concept that groups entities.
type FolderCollection struct {

	// Collection of Folders
	Items []FolderSummary `mandatory:"true" json:"items"`
}

func (m FolderCollection) String() string {
	return common.PointerString(m)
}
