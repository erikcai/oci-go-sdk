// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateImageDetails Either instanceId or imageSourceDetails must be provided in addition to other required parameters.
type CreateImageDetails struct {

	// The OCID of the compartment containing the instance you want to use as the basis for the image.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name for the image. It does not have to be unique, and it's changeable. You cannot use an
	// Oracle-provided image name as a custom image name.
	// Example: `My Oracle Linux image`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Details for creating an image through import
	ImageSourceDetails ImageSourceDetails `mandatory:"false" json:"imageSourceDetails"`

	// The OCID of the instance you want to use as the basis for the image.
	InstanceId *string `mandatory:"false" json:"instanceId"`
}

func (m CreateImageDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateImageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		DisplayName        *string                           `json:"displayName"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		ImageSourceDetails imagesourcedetails                `json:"imageSourceDetails"`
		InstanceId         *string                           `json:"instanceId"`
		CompartmentId      *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DefinedTags = model.DefinedTags
	m.DisplayName = model.DisplayName
	m.FreeformTags = model.FreeformTags
	nn, e := model.ImageSourceDetails.UnmarshalPolymorphicJSON(model.ImageSourceDetails.JsonData)
	if e != nil {
		return
	}
	m.ImageSourceDetails = nn.(ImageSourceDetails)
	m.InstanceId = model.InstanceId
	m.CompartmentId = model.CompartmentId
	return
}
