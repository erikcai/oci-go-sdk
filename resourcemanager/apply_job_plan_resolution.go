// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Resource Manager
//
// Oracle Resource Manager API.
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ApplyJobPlanResolution Specifies which Plan job provides an execution plan for input to the Apply job.
// You can set only one of the three job properties.
type ApplyJobPlanResolution struct {

	// OCID of the most recently executed Plan job.
	PlanJobId *string `mandatory:"false" json:"planJobId"`

	// Specifies whether to use the OCID of the most recently run Plan job.
	// `True` if using the latest job OCID. Must be a Plan job that completed successfully.
	IsUseLatestJobId *bool `mandatory:"false" json:"isUseLatestJobId"`

	// Specifies whether to use the configuration directly, without reference to a Plan job.
	// `True` if using the configuration directly. Note that it is not necessary
	// for a Plan job to have run successfully.
	IsAutoApproved *bool `mandatory:"false" json:"isAutoApproved"`
}

func (m ApplyJobPlanResolution) String() string {
	return common.PointerString(m)
}
