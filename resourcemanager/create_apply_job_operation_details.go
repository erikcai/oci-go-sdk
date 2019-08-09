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

// CreateApplyJobOperationDetails Job details that are specific to apply operations.
type CreateApplyJobOperationDetails struct {

	// The OCID of a plan job, for use when specifying `FROM_PLAN_JOB_ID` as the `executionPlanStrategy`.
	ExecutionPlanJobId *string `mandatory:"false" json:"executionPlanJobId"`

	// Specifies the source of the execution plan to apply.
	// Use `AUTO_APPROVED` to run the job without an execution plan.
	ExecutionPlanStrategy ApplyJobOperationDetailsExecutionPlanStrategyEnum `mandatory:"false" json:"executionPlanStrategy,omitempty"`
}

func (m CreateApplyJobOperationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateApplyJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateApplyJobOperationDetails CreateApplyJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeCreateApplyJobOperationDetails
	}{
		"APPLY",
		(MarshalTypeCreateApplyJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}
