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

// CreateVolumeDetails The representation of CreateVolumeDetails
type CreateVolumeDetails struct {

	// The Availability Domain of the volume.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment that contains the volume.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// If provided, specifies the ID of the volume backup policy to assign to the newly
	// created volume. If omitted, no policy will be assigned.
	BackupPolicyId *string `mandatory:"false" json:"backupPolicyId"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the KMS key to be used as the master encryption key for the volume.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The size of the volume in GBs.
	SizeInGBs *int `mandatory:"false" json:"sizeInGBs"`

	// The size of the volume in MBs. The value must be a multiple of 1024.
	// This field is deprecated. Please use sizeInGBs.
	SizeInMBs *int `mandatory:"false" json:"sizeInMBs"`

	// Encapsulates the source of this volume. This could be either another volume in the same AD or a volume backup.
	// When specified newly created volume will contain data from the source.
	// This is an optional field. If left unspecified or set to null, empty volume is created.
	SourceDetails VolumeSourceDetails `mandatory:"false" json:"sourceDetails"`

	// The OCID of the volume backup from which the data should be restored on the newly created volume.
	// This field is deprecated. Please consider using sourceDetails field to specify
	// backup as the source of this volume.
	VolumeBackupId *string `mandatory:"false" json:"volumeBackupId"`
}

func (m CreateVolumeDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateVolumeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BackupPolicyId     *string                           `json:"backupPolicyId"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		DisplayName        *string                           `json:"displayName"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		KmsKeyId           *string                           `json:"kmsKeyId"`
		SizeInGBs          *int                              `json:"sizeInGBs"`
		SizeInMBs          *int                              `json:"sizeInMBs"`
		SourceDetails      volumesourcedetails               `json:"sourceDetails"`
		VolumeBackupId     *string                           `json:"volumeBackupId"`
		AvailabilityDomain *string                           `json:"availabilityDomain"`
		CompartmentId      *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.BackupPolicyId = model.BackupPolicyId
	m.DefinedTags = model.DefinedTags
	m.DisplayName = model.DisplayName
	m.FreeformTags = model.FreeformTags
	m.KmsKeyId = model.KmsKeyId
	m.SizeInGBs = model.SizeInGBs
	m.SizeInMBs = model.SizeInMBs
	nn, e := model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	m.SourceDetails = nn.(VolumeSourceDetails)
	m.VolumeBackupId = model.VolumeBackupId
	m.AvailabilityDomain = model.AvailabilityDomain
	m.CompartmentId = model.CompartmentId
	return
}
