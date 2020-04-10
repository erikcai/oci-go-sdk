// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TaskRunSummary The information about TaskRun.
type TaskRunSummary struct {
	TaskRun *TaskRunDetails `mandatory:"false" json:"taskRun"`

	RegistryInfo *MetadataObjectSummary `mandatory:"false" json:"registryInfo"`
}

func (m TaskRunSummary) String() string {
	return common.PointerString(m)
}
