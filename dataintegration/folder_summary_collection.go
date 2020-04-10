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

// FolderSummaryCollection This is the collection of folder summaries, it may be a collection of lightweight details or full definitions.
type FolderSummaryCollection struct {

	// The array of Folder summaries
	Items []FolderSummary `mandatory:"true" json:"items"`
}

func (m FolderSummaryCollection) String() string {
	return common.PointerString(m)
}
