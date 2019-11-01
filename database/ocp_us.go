// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OcpUs The details of the available and consumed CPU cores of the Autonomous Exadata Infrastructure instance, including consumption by database workload type.
type OcpUs struct {

	// The total number of OCPUs in the Autonomous Exadata Infrastructure instance.
	TotalCpu *float32 `mandatory:"false" json:"totalCpu"`

	// The total number of consumed OCPUs in the Autonomous Exadata Infrastructure instance.
	ConsumedCpu *float32 `mandatory:"false" json:"consumedCpu"`

	ByWorkloadType *WorkloadType `mandatory:"false" json:"byWorkloadType"`
}

func (m OcpUs) String() string {
	return common.PointerString(m)
}
