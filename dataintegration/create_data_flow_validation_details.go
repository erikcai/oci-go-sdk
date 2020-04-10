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

// CreateDataFlowValidationDetails The information about integration dataflow validation.
type CreateDataFlowValidationDetails struct {
	Details *DataFlowDetails `mandatory:"false" json:"details"`
}

func (m CreateDataFlowValidationDetails) String() string {
	return common.PointerString(m)
}
