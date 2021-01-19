// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// Deployment Description of Deployment.
type Deployment struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Deployment Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the the Deployment was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Deployment was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Deployment.
	LifecycleState DeploymentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	DeploymentArguments *DeploymentArgumentCollection `mandatory:"false" json:"deploymentArguments"`

	DeploymentExecutionProgress *DeploymentExecutionProgress `mandatory:"false" json:"deploymentExecutionProgress"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Deployment) String() string {
	return common.PointerString(m)
}

// DeploymentLifecycleStateEnum Enum with underlying type: string
type DeploymentLifecycleStateEnum string

// Set of constants representing the allowable values for DeploymentLifecycleStateEnum
const (
	DeploymentLifecycleStateAccepted   DeploymentLifecycleStateEnum = "ACCEPTED"
	DeploymentLifecycleStateInProgress DeploymentLifecycleStateEnum = "IN_PROGRESS"
	DeploymentLifecycleStateFailed     DeploymentLifecycleStateEnum = "FAILED"
	DeploymentLifecycleStateSucceeded  DeploymentLifecycleStateEnum = "SUCCEEDED"
	DeploymentLifecycleStateCanceling  DeploymentLifecycleStateEnum = "CANCELING"
	DeploymentLifecycleStateCanceled   DeploymentLifecycleStateEnum = "CANCELED"
)

var mappingDeploymentLifecycleState = map[string]DeploymentLifecycleStateEnum{
	"ACCEPTED":    DeploymentLifecycleStateAccepted,
	"IN_PROGRESS": DeploymentLifecycleStateInProgress,
	"FAILED":      DeploymentLifecycleStateFailed,
	"SUCCEEDED":   DeploymentLifecycleStateSucceeded,
	"CANCELING":   DeploymentLifecycleStateCanceling,
	"CANCELED":    DeploymentLifecycleStateCanceled,
}

// GetDeploymentLifecycleStateEnumValues Enumerates the set of values for DeploymentLifecycleStateEnum
func GetDeploymentLifecycleStateEnumValues() []DeploymentLifecycleStateEnum {
	values := make([]DeploymentLifecycleStateEnum, 0)
	for _, v := range mappingDeploymentLifecycleState {
		values = append(values, v)
	}
	return values
}
