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

// UpdateNotebookSessionDetails Details for updating a notebook session.  `notebookSessionConfigurationDetails` can only be updated while the notebook session is in the INACTIVE state.
// Changes to the `notebookSessionConfigurationDetails` will take effect the next time the `ActivateNotebookSession` action is invoked on the notebook session resource.
type UpdateNotebookSessionDetails struct {

	// A user-friendly name. Does not have to be unique and is changeable.
	// Avoid entering confidential information.
	// Example: `My NotebookSession`
	DisplayName *string `mandatory:"false" json:"displayName"`

	NotebookSessionConfigurationDetails *NotebookSessionConfigurationDetails `mandatory:"false" json:"notebookSessionConfigurationDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateNotebookSessionDetails) String() string {
	return common.PointerString(m)
}
