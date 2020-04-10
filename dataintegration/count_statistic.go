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

// CountStatistic A count statistics
type CountStatistic struct {

	// The array of statistics
	ObjectTypeCountList []CountStatisticSummary `mandatory:"true" json:"objectTypeCountList"`
}

func (m CountStatistic) String() string {
	return common.PointerString(m)
}
