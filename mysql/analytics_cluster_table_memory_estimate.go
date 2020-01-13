// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AnalyticsClusterTableMemoryEstimate Estimated memory footprint for a MySQL user table
// when loaded to the Analytics Cluster memory.
type AnalyticsClusterTableMemoryEstimate struct {

	// The table name.
	TableName *string `mandatory:"true" json:"tableName"`

	// The number of columns to be loaded to Analytics Cluster memory.
	// These columns contribute to the analytical memory footprint.
	ColumnToLoadCount *int `mandatory:"true" json:"columnToLoadCount"`

	// The estimated number of rows in the table. This number was used to
	// derive the analytical memory footprint.
	EstimatedTableRowCount *int64 `mandatory:"true" json:"estimatedTableRowCount"`

	// The estimated memory footprint of the table in MBs when loaded to
	// Analytics Cluster memory (null if the table cannot be loaded to the
	// Analytics Cluster).
	AnalyticalFootprintInMbs *int64 `mandatory:"true" json:"analyticalFootprintInMbs"`

	// Error comment (empty string if no errors occured).
	ErrorComment *string `mandatory:"true" json:"errorComment"`
}

func (m AnalyticsClusterTableMemoryEstimate) String() string {
	return common.PointerString(m)
}
