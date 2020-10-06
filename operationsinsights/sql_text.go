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

// SqlText SQL Text type object.
type SqlText struct {

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// Collection timestamp
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// SQL command
	SqlCommand *string `mandatory:"true" json:"sqlCommand"`

	// Full SQL Text
	SqlFullText *string `mandatory:"true" json:"sqlFullText"`

	// Version
	Version *float32 `mandatory:"false" json:"version"`

	// Exact matching signature
	ExactMatchingSignature *string `mandatory:"false" json:"exactMatchingSignature"`

	// Force matching signature
	ForceMatchingSignature *string `mandatory:"false" json:"forceMatchingSignature"`
}

func (m SqlText) String() string {
	return common.PointerString(m)
}
