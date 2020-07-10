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

// SqlPlanInsightAggregationCollection SQL plan insights response.
type SqlPlanInsightAggregationCollection struct {

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// The OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// List of SQL plan insights.
	Insights []SqlPlanInsights `mandatory:"true" json:"insights"`

	// List of SQL plan statistics.
	Items []SqlPlanInsightAggregation `mandatory:"true" json:"items"`
}

func (m SqlPlanInsightAggregationCollection) String() string {
	return common.PointerString(m)
}