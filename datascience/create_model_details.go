// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// CreateModelDetails Parameters needed to create a new model. Models are mathematical representations of the relationships between data. Models are represented by their associated metadata and artifact.
type CreateModelDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the model in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	// Example: `My Model`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the model.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateModelDetails) String() string {
	return common.PointerString(m)
}
