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

// SqlStatisticAggregation SQL Statistics
type SqlStatisticAggregation struct {

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	DatabaseDetails *DatabaseDetails `mandatory:"true" json:"databaseDetails"`

	// SQL belongs to one or more categories based on the insights.
	Category []string `mandatory:"true" json:"category"`

	Statistics *SqlStatistics `mandatory:"false" json:"statistics"`
}

func (m SqlStatisticAggregation) String() string {
	return common.PointerString(m)
}
