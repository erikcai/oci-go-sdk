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

// SqlStatisticsTimeSeriesByPlanAggregationCollection SQL performance statistics by plan over the selected time window.
type SqlStatisticsTimeSeriesByPlanAggregationCollection struct {

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// The OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Array comprising of all the sampling period end timestamps in RFC 3339 format.
	EndTimestamps []common.SDKTime `mandatory:"true" json:"endTimestamps"`

	// array of SQL performance statistics by plans
	Items []SqlStatisticsTimeSeriesByPlanAggregation `mandatory:"true" json:"items"`
}

func (m SqlStatisticsTimeSeriesByPlanAggregationCollection) String() string {
	return common.PointerString(m)
}
