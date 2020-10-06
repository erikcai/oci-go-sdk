// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperationsInsights API
//
// The OperationsInsights API for OPSI service.
//

package operationsinsights

import (
	"github.com/oracle/oci-go-sdk/v26/common"
)

// SqlInventory Inventory details.
type SqlInventory struct {

	// Total number of sqls. Example `2000`
	TotalSqls *int64 `mandatory:"true" json:"totalSqls"`

	// Total number of Databases. Example `400`
	TotalDatabases *int `mandatory:"true" json:"totalDatabases"`

	// Total number of sqls analyzed by the query. Example `120`
	SqlsAnalyzed *int64 `mandatory:"true" json:"sqlsAnalyzed"`
}

func (m SqlInventory) String() string {
	return common.PointerString(m)
}
