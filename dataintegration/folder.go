// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Folder The folder type contains the audit summary information and the definition of the folder.
type Folder struct {
	Details *FolderDetails `mandatory:"false" json:"details"`

	Summary *MetadataObjectSummary `mandatory:"false" json:"summary"`

	// keyMappingFromInput
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m Folder) String() string {
	return common.PointerString(m)
}
