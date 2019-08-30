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

// CreateModelDetails Parameters needed to create a new model. Models are mathematical representations of the relationships between data. Models are represented by their associated metadata and artifact.
type CreateModelDetails struct {

	// The OCID of the compartment to create the model in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the project to associate with the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// A user-friendly name. Does not have to be unique and is changeable.
	// Avoid entering confidential information.
	// Example: `My Model`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short blurb describing the model.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateModelDetails) String() string {
	return common.PointerString(m)
}
