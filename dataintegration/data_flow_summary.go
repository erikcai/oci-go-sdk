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

// DataFlowSummary The data flow summary type contains the audit summary information and the definition of the data flow.
type DataFlowSummary struct {
	Details *DataFlowDetails `mandatory:"false" json:"details"`

	Summary *MetadataObjectSummary `mandatory:"false" json:"summary"`
}

func (m DataFlowSummary) String() string {
	return common.PointerString(m)
}
