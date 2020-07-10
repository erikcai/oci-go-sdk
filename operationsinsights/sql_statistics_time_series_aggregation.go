// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperationsInsights API
//
// The OperationsInsights API for OPSI service.
//

package operationsinsights

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SqlStatisticsTimeSeriesAggregation Database details and SQL performance statistics for a given database
type SqlStatisticsTimeSeriesAggregation struct {
	DatabaseDetails *DatabaseDetails `mandatory:"true" json:"databaseDetails"`

	// SQL performance statistics for a given database
	Statistics []SqlStatisticsTimeSeries `mandatory:"true" json:"statistics"`
}

func (m SqlStatisticsTimeSeriesAggregation) String() string {
	return common.PointerString(m)
}