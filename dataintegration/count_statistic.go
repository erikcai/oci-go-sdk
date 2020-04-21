// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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