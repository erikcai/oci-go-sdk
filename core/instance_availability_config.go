// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm),
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm), and
// Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as virtual cloud networks (VCNs),
// compute instances, block storage volumes, and container images.
//

package core

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// InstanceAvailabilityConfig Options for defining the availabiity of a VM instance after a maintenance event that impacts the underlying hardware.
type InstanceAvailabilityConfig struct {

	// The lifecycle state for an instance when it is recovered after infrastructure maintenance.
	// * `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event.
	// If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
	// * `STOP_INSTANCE` - The instance is recovered in the stopped state.
	RecoveryAction InstanceAvailabilityConfigRecoveryActionEnum `mandatory:"false" json:"recoveryAction,omitempty"`
}

func (m InstanceAvailabilityConfig) String() string {
	return common.PointerString(m)
}

// InstanceAvailabilityConfigRecoveryActionEnum Enum with underlying type: string
type InstanceAvailabilityConfigRecoveryActionEnum string

// Set of constants representing the allowable values for InstanceAvailabilityConfigRecoveryActionEnum
const (
	InstanceAvailabilityConfigRecoveryActionRestoreInstance InstanceAvailabilityConfigRecoveryActionEnum = "RESTORE_INSTANCE"
	InstanceAvailabilityConfigRecoveryActionStopInstance    InstanceAvailabilityConfigRecoveryActionEnum = "STOP_INSTANCE"
)

var mappingInstanceAvailabilityConfigRecoveryAction = map[string]InstanceAvailabilityConfigRecoveryActionEnum{
	"RESTORE_INSTANCE": InstanceAvailabilityConfigRecoveryActionRestoreInstance,
	"STOP_INSTANCE":    InstanceAvailabilityConfigRecoveryActionStopInstance,
}

// GetInstanceAvailabilityConfigRecoveryActionEnumValues Enumerates the set of values for InstanceAvailabilityConfigRecoveryActionEnum
func GetInstanceAvailabilityConfigRecoveryActionEnumValues() []InstanceAvailabilityConfigRecoveryActionEnum {
	values := make([]InstanceAvailabilityConfigRecoveryActionEnum, 0)
	for _, v := range mappingInstanceAvailabilityConfigRecoveryAction {
		values = append(values, v)
	}
	return values
}
