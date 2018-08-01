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

// VolumeHealth A point-in-time representation of the health of a volume. For internal consumption.
type VolumeHealth struct {

	// The current health status of a volume.
	Status VolumeHealthStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the volume.
	VolumeId *string `mandatory:"true" json:"volumeId"`
}

func (m VolumeHealth) String() string {
	return common.PointerString(m)
}

// VolumeHealthStatusEnum Enum with underlying type: string
type VolumeHealthStatusEnum string

// Set of constants representing the allowable values for VolumeHealthStatus
const (
	VolumeHealthStatusHealthy   VolumeHealthStatusEnum = "HEALTHY"
	VolumeHealthStatusUnhealthy VolumeHealthStatusEnum = "UNHEALTHY"
	VolumeHealthStatusUnknown   VolumeHealthStatusEnum = "UNKNOWN"
)

var mappingVolumeHealthStatus = map[string]VolumeHealthStatusEnum{
	"HEALTHY":   VolumeHealthStatusHealthy,
	"UNHEALTHY": VolumeHealthStatusUnhealthy,
	"UNKNOWN":   VolumeHealthStatusUnknown,
}

// GetVolumeHealthStatusEnumValues Enumerates the set of values for VolumeHealthStatus
func GetVolumeHealthStatusEnumValues() []VolumeHealthStatusEnum {
	values := make([]VolumeHealthStatusEnum, 0)
	for _, v := range mappingVolumeHealthStatus {
		values = append(values, v)
	}
	return values
}
