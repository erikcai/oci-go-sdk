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

// DatabaseInsights Logical grouping used for Operations Insights database-targeted operations.
type DatabaseInsights struct {

	// Database Insights Object.
	DatabaseInsights *interface{} `mandatory:"false" json:"databaseInsights"`
}

func (m DatabaseInsights) String() string {
	return common.PointerString(m)
}
