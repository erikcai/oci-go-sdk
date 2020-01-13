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

// TableSchema Schema with estimated memory footprints for each MySQL user table
// of the schema when loaded to Analytics Cluster memory.
type TableSchema struct {

	// The name of the schema.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// Estimated memory footprints for MySQL user tables of the schema
	// when loaded to Analytics Cluster memory.
	PerTableEstimates []AnalyticsClusterTableMemoryEstimate `mandatory:"true" json:"perTableEstimates"`
}

func (m TableSchema) String() string {
	return common.PointerString(m)
}
