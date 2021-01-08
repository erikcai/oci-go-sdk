// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// Pipeline Pipeline consists of Stages and metadata for Deployments.
type Pipeline struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	PipelineParameters *PipelineParameterCollection `mandatory:"true" json:"pipelineParameters"`

	// Pipeline Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Type of the Pipeline.
	PipelineType *string `mandatory:"false" json:"pipelineType"`

	// The time the the Pipeline was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Pipeline was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Pipeline.
	LifecycleState PipelineLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

func (m Pipeline) String() string {
	return common.PointerString(m)
}

// PipelineLifecycleStateEnum Enum with underlying type: string
type PipelineLifecycleStateEnum string

// Set of constants representing the allowable values for PipelineLifecycleStateEnum
const (
	PipelineLifecycleStateCreating PipelineLifecycleStateEnum = "CREATING"
	PipelineLifecycleStateUpdating PipelineLifecycleStateEnum = "UPDATING"
	PipelineLifecycleStateActive   PipelineLifecycleStateEnum = "ACTIVE"
	PipelineLifecycleStateDeleting PipelineLifecycleStateEnum = "DELETING"
	PipelineLifecycleStateDeleted  PipelineLifecycleStateEnum = "DELETED"
	PipelineLifecycleStateFailed   PipelineLifecycleStateEnum = "FAILED"
)

var mappingPipelineLifecycleState = map[string]PipelineLifecycleStateEnum{
	"CREATING": PipelineLifecycleStateCreating,
	"UPDATING": PipelineLifecycleStateUpdating,
	"ACTIVE":   PipelineLifecycleStateActive,
	"DELETING": PipelineLifecycleStateDeleting,
	"DELETED":  PipelineLifecycleStateDeleted,
	"FAILED":   PipelineLifecycleStateFailed,
}

// GetPipelineLifecycleStateEnumValues Enumerates the set of values for PipelineLifecycleStateEnum
func GetPipelineLifecycleStateEnumValues() []PipelineLifecycleStateEnum {
	values := make([]PipelineLifecycleStateEnum, 0)
	for _, v := range mappingPipelineLifecycleState {
		values = append(values, v)
	}
	return values
}
