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

// NotebookSession Notebook sessions are interactive coding environments for data scientists.
type NotebookSession struct {

	// The OCID
	Id *string `mandatory:"true" json:"id"`

	// Time resource was created, expressed in timestamp format
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name. Does not have to be unique and is changeable.
	// Example: `My NotebookSession`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the project associated with the notebook session.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the user who created the notebook session.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID of the notebook session's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The state of the notebook session.
	LifecycleState NotebookSessionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	NotebookSessionConfigurationDetails *NotebookSessionConfigurationDetails `mandatory:"false" json:"notebookSessionConfigurationDetails"`

	// The URL to interact with the notebook session.
	NotebookSessionUrl *string `mandatory:"false" json:"notebookSessionUrl"`

	// Details about the state of the notebook session.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m NotebookSession) String() string {
	return common.PointerString(m)
}
