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

// SqlSearchCollection Search SQL response.
type SqlSearchCollection struct {

	// List of Databases executing the sql.
	Items []SqlSearchSummary `mandatory:"true" json:"items"`

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"false" json:"sqlIdentifier"`

	// SQL Statement Text
	SqlText *string `mandatory:"false" json:"sqlText"`
}

func (m SqlSearchCollection) String() string {
	return common.PointerString(m)
}