// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Science API
//
// The Data Science service enables data science teams to organize their work, easily access data and computing resources, and build, train, deploy, and manage ML/AI models on the Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ProjectSummary Summary information for a project.
type ProjectSummary struct {

	// The OCID
	Id *string `mandatory:"true" json:"id"`

	// Time resource was created, expressed in timestamp format
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name. Does not have to be unique and is changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the project's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the user who created this project.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The state of the project.
	LifecycleState ProjectLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A short blurb describing the project.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ProjectSummary) String() string {
	return common.PointerString(m)
}
