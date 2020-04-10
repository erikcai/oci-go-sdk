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

// DataFlowValidationSummary The information about dataflow validation
type DataFlowValidationSummary struct {
	Details *DataFlowValidationDetails `mandatory:"false" json:"details"`

	Summary *MetadataObjectSummary `mandatory:"false" json:"summary"`
}

func (m DataFlowValidationSummary) String() string {
	return common.PointerString(m)
}
