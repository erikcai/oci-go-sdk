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

// ConnectionValidationSummary The information about connection validation
type ConnectionValidationSummary struct {
	Details *ConnectionValidationObject `mandatory:"false" json:"details"`

	Summary *MetadataObjectSummary `mandatory:"false" json:"summary"`
}

func (m ConnectionValidationSummary) String() string {
	return common.PointerString(m)
}
