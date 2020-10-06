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

// SqlPlanLine SQL Plan Line type object.
type SqlPlanLine struct {

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// Plan hash value for the SQL Execution Plan
	PlanHash *int64 `mandatory:"true" json:"planHash"`

	// Collection time stamp
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Operation
	Operation *string `mandatory:"true" json:"operation"`

	// Identifier
	Identifier *int64 `mandatory:"true" json:"identifier"`

	// Version
	Version *float32 `mandatory:"false" json:"version"`

	// Remark
	Remark *string `mandatory:"false" json:"remark"`

	// Options
	Options *string `mandatory:"false" json:"options"`

	// Object Node
	ObjectNode *string `mandatory:"false" json:"objectNode"`

	// Object Owner
	ObjectOwner *string `mandatory:"false" json:"objectOwner"`

	// Object Name
	ObjectName *string `mandatory:"false" json:"objectName"`

	// Object Alias
	ObjectAlias *string `mandatory:"false" json:"objectAlias"`

	// Object Instance
	ObjectInstance *int64 `mandatory:"false" json:"objectInstance"`

	// Object Type
	ObjectType *string `mandatory:"false" json:"objectType"`

	// Optimizer
	Optimizer *string `mandatory:"false" json:"optimizer"`

	// Search Columns
	SearchColumns *int64 `mandatory:"false" json:"searchColumns"`

	// Parent Identifier
	ParentIdentifier *int64 `mandatory:"false" json:"parentIdentifier"`

	// Depth
	Depth *int64 `mandatory:"false" json:"depth"`

	// Position
	Position *int64 `mandatory:"false" json:"position"`

	// Cost
	Cost *int64 `mandatory:"false" json:"cost"`

	// Cardinality
	Cardinality *int64 `mandatory:"false" json:"cardinality"`

	// Bytes
	Bytes *int64 `mandatory:"false" json:"bytes"`

	// Other
	Other *string `mandatory:"false" json:"other"`

	// Other Tag
	OtherTag *string `mandatory:"false" json:"otherTag"`

	// Partition start
	PartitionStart *string `mandatory:"false" json:"partitionStart"`

	// Partition stop
	PartitionStop *string `mandatory:"false" json:"partitionStop"`

	// Partition identifier
	PartitionIdentifier *int64 `mandatory:"false" json:"partitionIdentifier"`

	// Distribution
	Distribution *string `mandatory:"false" json:"distribution"`

	// CPU cost
	CpuCost *int64 `mandatory:"false" json:"cpuCost"`

	// IO cost
	IoCost *int64 `mandatory:"false" json:"ioCost"`

	// Time space
	TempSpace *int64 `mandatory:"false" json:"tempSpace"`

	// Access predicates
	AccessPredicates *string `mandatory:"false" json:"accessPredicates"`

	// Filter predicates
	FilterPredicates *string `mandatory:"false" json:"filterPredicates"`

	// Projection
	Projection *string `mandatory:"false" json:"projection"`

	// Qblock Name
	QblockName *string `mandatory:"false" json:"qblockName"`

	// Total elapsed time
	ElapsedTimeInSec *float32 `mandatory:"false" json:"elapsedTimeInSec"`

	// Other SQL
	OtherXML *string `mandatory:"false" json:"otherXML"`
}

func (m SqlPlanLine) String() string {
	return common.PointerString(m)
}
