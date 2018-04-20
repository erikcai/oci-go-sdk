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

// Volume A detachable block volume device that allows you to dynamically expand
// the storage capacity of an instance. For more information, see
// Overview of Cloud Volume Storage (https://docs.us-phoenix-1.oraclecloud.com/Content/Block/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
type Volume struct {

	// The Availability Domain of the volume.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment that contains the volume.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The volume's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The current state of a volume.
	LifecycleState VolumeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The size of the volume in MBs. This field is deprecated. Please use sizeInGBs.
	SizeInMBs *int `mandatory:"true" json:"sizeInMBs"`

	// The date and time the volume was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Boolean flag that tells if the volume's data is hydrated completely from the source.
	IsHydrated *bool `mandatory:"false" json:"isHydrated"`

	// The OCID of the KMS key which is the master encryption key for the volume.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The size of the volume in GBs.
	SizeInGBs *int `mandatory:"false" json:"sizeInGBs"`

	// Describes the source of this volume. This could be either another volume in the same AD or a volume backup.
	// If it is null, no source is used to create this volume
	SourceDetails VolumeSourceDetails `mandatory:"false" json:"sourceDetails"`

	// The OCID of the source volume group.
	VolumeGroupId *string `mandatory:"false" json:"volumeGroupId"`
}

func (m Volume) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Volume) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		IsHydrated         *bool                             `json:"isHydrated"`
		KmsKeyId           *string                           `json:"kmsKeyId"`
		SizeInGBs          *int                              `json:"sizeInGBs"`
		SourceDetails      volumesourcedetails               `json:"sourceDetails"`
		VolumeGroupId      *string                           `json:"volumeGroupId"`
		AvailabilityDomain *string                           `json:"availabilityDomain"`
		CompartmentId      *string                           `json:"compartmentId"`
		DisplayName        *string                           `json:"displayName"`
		Id                 *string                           `json:"id"`
		LifecycleState     VolumeLifecycleStateEnum          `json:"lifecycleState"`
		SizeInMBs          *int                              `json:"sizeInMBs"`
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DefinedTags = model.DefinedTags
	m.FreeformTags = model.FreeformTags
	m.IsHydrated = model.IsHydrated
	m.KmsKeyId = model.KmsKeyId
	m.SizeInGBs = model.SizeInGBs
	nn, e := model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	m.SourceDetails = nn.(VolumeSourceDetails)
	m.VolumeGroupId = model.VolumeGroupId
	m.AvailabilityDomain = model.AvailabilityDomain
	m.CompartmentId = model.CompartmentId
	m.DisplayName = model.DisplayName
	m.Id = model.Id
	m.LifecycleState = model.LifecycleState
	m.SizeInMBs = model.SizeInMBs
	m.TimeCreated = model.TimeCreated
	return
}

// VolumeLifecycleStateEnum Enum with underlying type: string
type VolumeLifecycleStateEnum string

// Set of constants representing the allowable values for VolumeLifecycleState
const (
	VolumeLifecycleStateProvisioning VolumeLifecycleStateEnum = "PROVISIONING"
	VolumeLifecycleStateRestoring    VolumeLifecycleStateEnum = "RESTORING"
	VolumeLifecycleStateAvailable    VolumeLifecycleStateEnum = "AVAILABLE"
	VolumeLifecycleStateTerminating  VolumeLifecycleStateEnum = "TERMINATING"
	VolumeLifecycleStateTerminated   VolumeLifecycleStateEnum = "TERMINATED"
	VolumeLifecycleStateFaulty       VolumeLifecycleStateEnum = "FAULTY"
)

var mappingVolumeLifecycleState = map[string]VolumeLifecycleStateEnum{
	"PROVISIONING": VolumeLifecycleStateProvisioning,
	"RESTORING":    VolumeLifecycleStateRestoring,
	"AVAILABLE":    VolumeLifecycleStateAvailable,
	"TERMINATING":  VolumeLifecycleStateTerminating,
	"TERMINATED":   VolumeLifecycleStateTerminated,
	"FAULTY":       VolumeLifecycleStateFaulty,
}

// GetVolumeLifecycleStateEnumValues Enumerates the set of values for VolumeLifecycleState
func GetVolumeLifecycleStateEnumValues() []VolumeLifecycleStateEnum {
	values := make([]VolumeLifecycleStateEnum, 0)
	for _, v := range mappingVolumeLifecycleState {
		values = append(values, v)
	}
	return values
}
