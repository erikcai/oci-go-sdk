// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ModelDeploymentSummary Summary information for a model deployment.
type ModelDeploymentSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created, in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: 2019-08-25T21:10:29.41Z
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information.
	// Example: `My ModelDeployment`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model deployment.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model deployment.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The URL to interact with the model deployment.
	ModelDeploymentUrl *string `mandatory:"true" json:"modelDeploymentUrl"`

	// The state of the model deployment.
	LifecycleState ModelDeploymentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	ModelDeploymentConfigurationDetails ModelDeploymentConfigurationDetails `mandatory:"false" json:"modelDeploymentConfigurationDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ModelDeploymentSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ModelDeploymentSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ModelDeploymentConfigurationDetails modeldeploymentconfigurationdetails `json:"modelDeploymentConfigurationDetails"`
		FreeformTags                        map[string]string                   `json:"freeformTags"`
		DefinedTags                         map[string]map[string]interface{}   `json:"definedTags"`
		SystemTags                          map[string]map[string]interface{}   `json:"systemTags"`
		Id                                  *string                             `json:"id"`
		TimeCreated                         *common.SDKTime                     `json:"timeCreated"`
		DisplayName                         *string                             `json:"displayName"`
		ProjectId                           *string                             `json:"projectId"`
		CreatedBy                           *string                             `json:"createdBy"`
		CompartmentId                       *string                             `json:"compartmentId"`
		ModelDeploymentUrl                  *string                             `json:"modelDeploymentUrl"`
		LifecycleState                      ModelDeploymentLifecycleStateEnum   `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ModelDeploymentConfigurationDetails.UnmarshalPolymorphicJSON(model.ModelDeploymentConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelDeploymentConfigurationDetails = nn.(ModelDeploymentConfigurationDetails)
	} else {
		m.ModelDeploymentConfigurationDetails = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.TimeCreated = model.TimeCreated

	m.DisplayName = model.DisplayName

	m.ProjectId = model.ProjectId

	m.CreatedBy = model.CreatedBy

	m.CompartmentId = model.CompartmentId

	m.ModelDeploymentUrl = model.ModelDeploymentUrl

	m.LifecycleState = model.LifecycleState

	return
}
