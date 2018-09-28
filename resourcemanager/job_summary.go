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

// JobSummary The representation of JobSummary
type JobSummary struct {
	Id *string `mandatory:"false" json:"id"`

	StackId *string `mandatory:"false" json:"stackId"`

	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of job executing
	Operation JobOperationEnum `mandatory:"false" json:"operation,omitempty"`

	ApplyJobPlanResolution *ApplyJobPlanResolution `mandatory:"false" json:"applyJobPlanResolution"`

	// The plan job OCID that was used (if this was an APPLY job and not auto approved).
	ResolvedPlanJobId *string `mandatory:"false" json:"resolvedPlanJobId"`

	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the job succeeded or failed.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m JobSummary) String() string {
	return common.PointerString(m)
}
