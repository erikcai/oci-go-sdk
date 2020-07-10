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

// SqlResponseTimeDistributionAggregation SQL Response time distribution entry.
type SqlResponseTimeDistributionAggregation struct {

	// Response time bucket id
	BucketId *string `mandatory:"true" json:"bucketId"`

	// Total number of SQL executions
	ExecutionsCount *int64 `mandatory:"true" json:"executionsCount"`
}

func (m SqlResponseTimeDistributionAggregation) String() string {
	return common.PointerString(m)
}