// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service. Use this API to install, configure, and manage resources via the "infrastructure-as-code" model. For more information, see Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ApplyJobOperationDetailsSummary Job details that are specific to apply operations.
type ApplyJobOperationDetailsSummary struct {

	// The OCID of the plan job that contains the execution plan used for this job,
	// or `null` if no execution plan was used.
	ExecutionPlanJobId *string `mandatory:"false" json:"executionPlanJobId"`

	// Specifies the source of the execution plan to apply.
	// Use `AUTO_APPROVED` to run the job without an execution plan.
	ExecutionPlanStrategy ApplyJobOperationDetailsExecutionPlanStrategyEnum `mandatory:"true" json:"executionPlanStrategy"`
}

func (m ApplyJobOperationDetailsSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ApplyJobOperationDetailsSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApplyJobOperationDetailsSummary ApplyJobOperationDetailsSummary
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeApplyJobOperationDetailsSummary
	}{
		"APPLY",
		(MarshalTypeApplyJobOperationDetailsSummary)(m),
	}

	return json.Marshal(&s)
}
