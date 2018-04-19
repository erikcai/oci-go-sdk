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

// CreateBootVolumeDetails The representation of CreateBootVolumeDetails
type CreateBootVolumeDetails struct {

	// The Availability Domain of the boot volume.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment that contains the boot volume.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Specifies the boot volume source details for a new boot volume. The volume source is either another boot volume in the same Availability Domain or a boot volume backup.
	// This is a mandatory field for a boot volume.
	SourceDetails BootVolumeSourceDetails `mandatory:"true" json:"sourceDetails"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The size of the volume in GBs.
	SizeInGBs *int `mandatory:"false" json:"sizeInGBs"`
}

func (m CreateBootVolumeDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateBootVolumeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		DisplayName        *string                           `json:"displayName"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		SizeInGBs          *int                              `json:"sizeInGBs"`
		AvailabilityDomain *string                           `json:"availabilityDomain"`
		CompartmentId      *string                           `json:"compartmentId"`
		SourceDetails      bootvolumesourcedetails           `json:"sourceDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DefinedTags = model.DefinedTags
	m.DisplayName = model.DisplayName
	m.FreeformTags = model.FreeformTags
	m.SizeInGBs = model.SizeInGBs
	m.AvailabilityDomain = model.AvailabilityDomain
	m.CompartmentId = model.CompartmentId
	nn, e := model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	m.SourceDetails = nn.(BootVolumeSourceDetails)
	return
}
