// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateCrossRegionBackupDetails The representation of UpdateCrossRegionBackupDetails
type UpdateCrossRegionBackupDetails struct {

	// whether the encryption keys are encrypted in the request
	AreKeysEncrypted *bool `mandatory:"false" json:"areKeysEncrypted"`

	// The state of the remote backup to update to.
	LifecycleState UpdateCrossRegionBackupDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// volume encryption keys
	EncryptionKeys *CrossRegionBackupEncryptionKeys `mandatory:"false" json:"encryptionKeys"`

	// whether or not to enable metering (defaults to false)
	EnableMetering *bool `mandatory:"false" json:"enableMetering"`
}

func (m UpdateCrossRegionBackupDetails) String() string {
	return common.PointerString(m)
}

// UpdateCrossRegionBackupDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateCrossRegionBackupDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateCrossRegionBackupDetailsLifecycleState
const (
	UpdateCrossRegionBackupDetailsLifecycleStateCopying   UpdateCrossRegionBackupDetailsLifecycleStateEnum = "COPYING"
	UpdateCrossRegionBackupDetailsLifecycleStateAvailable UpdateCrossRegionBackupDetailsLifecycleStateEnum = "AVAILABLE"
	UpdateCrossRegionBackupDetailsLifecycleStateFaulty    UpdateCrossRegionBackupDetailsLifecycleStateEnum = "FAULTY"
)

var mappingUpdateCrossRegionBackupDetailsLifecycleState = map[string]UpdateCrossRegionBackupDetailsLifecycleStateEnum{
	"COPYING":   UpdateCrossRegionBackupDetailsLifecycleStateCopying,
	"AVAILABLE": UpdateCrossRegionBackupDetailsLifecycleStateAvailable,
	"FAULTY":    UpdateCrossRegionBackupDetailsLifecycleStateFaulty,
}

// GetUpdateCrossRegionBackupDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateCrossRegionBackupDetailsLifecycleState
func GetUpdateCrossRegionBackupDetailsLifecycleStateEnumValues() []UpdateCrossRegionBackupDetailsLifecycleStateEnum {
	values := make([]UpdateCrossRegionBackupDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateCrossRegionBackupDetailsLifecycleState {
		values = append(values, v)
	}
	return values
}
