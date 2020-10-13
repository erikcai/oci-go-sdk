// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperationsInsights API
//
// The OperationsInsights API for OPSI service.
//

package operationsinsights

import (
	"github.com/oracle/oci-go-sdk/v27/common"
)

// SqlBucket Sql bucket type object.
type SqlBucket struct {

	// Collection timestamp
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// Plan hash value for the SQL Execution Plan
	PlanHash *int64 `mandatory:"true" json:"planHash"`

	// SQL Bucket ID, examples <= 3 secs, 3-10 secs, 10-60 secs, 1-5 min, > 5 min
	BucketId *string `mandatory:"true" json:"bucketId"`

	// Version
	Version *float32 `mandatory:"false" json:"version"`

	// Operations Insights internal representation of the database type.
	DatabaseType *string `mandatory:"false" json:"databaseType"`

	// Total number of executions
	ExecutionsCount *int `mandatory:"false" json:"executionsCount"`

	// Total CPU time
	CpuTimeInSec *float32 `mandatory:"false" json:"cpuTimeInSec"`

	// Total IO time
	IoTimeInSec *float32 `mandatory:"false" json:"ioTimeInSec"`

	// Total other wait time
	OtherWaitTimeInSec *float32 `mandatory:"false" json:"otherWaitTimeInSec"`

	// Total time
	TotalTimeInSec *float32 `mandatory:"false" json:"totalTimeInSec"`
}

func (m SqlBucket) String() string {
	return common.PointerString(m)
}
