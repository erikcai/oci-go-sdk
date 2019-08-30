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

// ModelSummary Summary information for a model.
type ModelSummary struct {

	// The OCID of the model’s compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the project associated with the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique and is changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the user who created the model.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// Time resource was created, expressed in timestamp format
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the model.
	LifecycleState ModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ModelSummary) String() string {
	return common.PointerString(m)
}
