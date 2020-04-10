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

// DataAssetSummaryCollection This is the collection of data asset summaries, it may be a collection of lightweight details or full definitions.
type DataAssetSummaryCollection struct {

	// The array of DataAsset summaries
	Items []DataAssetSummary `mandatory:"true" json:"items"`
}

func (m DataAssetSummaryCollection) String() string {
	return common.PointerString(m)
}
