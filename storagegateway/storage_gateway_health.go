// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// StorageGateway API
//
// API for interfacing with StorageGateway
//

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// StorageGatewayHealth The current health of the storage gateway.
type StorageGatewayHealth struct {

	// The overall health status. ACTIVE means the storage gateway is running with no issue.
	// INACTIVE means the storage gateway instance has not been installed on premise or on
	// compute yet or the control plane has not received the heartbeat data for a while.
	// WARNING or CRITICAL means the storage gateway has warning(s) or ran into critical
	// issue(s). Specific reasons for this health status should be found from the reason
	// strings.
	Status StorageGatewayHealthStatusEnum `mandatory:"true" json:"status"`

	// The current version of storage gateway instance on-premise or on-compute.
	Version *string `mandatory:"true" json:"version"`

	// Availability of a newer version.
	IsNewerVersionAvailable *bool `mandatory:"true" json:"isNewerVersionAvailable"`

	// A timestamp when the most recent heartbeat is received.
	TimeLastHeartbeatReceived *common.SDKTime `mandatory:"true" json:"timeLastHeartbeatReceived"`

	// metrics
	Metrics *Metrics `mandatory:"true" json:"metrics"`

	// reasons
	Reasons *StatusReasons `mandatory:"false" json:"reasons"`
}

func (m StorageGatewayHealth) String() string {
	return common.PointerString(m)
}

// StorageGatewayHealthStatusEnum Enum with underlying type: string
type StorageGatewayHealthStatusEnum string

// Set of constants representing the allowable values for StorageGatewayHealthStatusEnum
const (
	StorageGatewayHealthStatusActive   StorageGatewayHealthStatusEnum = "ACTIVE"
	StorageGatewayHealthStatusInactive StorageGatewayHealthStatusEnum = "INACTIVE"
	StorageGatewayHealthStatusWarning  StorageGatewayHealthStatusEnum = "WARNING"
	StorageGatewayHealthStatusCritical StorageGatewayHealthStatusEnum = "CRITICAL"
)

var mappingStorageGatewayHealthStatus = map[string]StorageGatewayHealthStatusEnum{
	"ACTIVE":   StorageGatewayHealthStatusActive,
	"INACTIVE": StorageGatewayHealthStatusInactive,
	"WARNING":  StorageGatewayHealthStatusWarning,
	"CRITICAL": StorageGatewayHealthStatusCritical,
}

// GetStorageGatewayHealthStatusEnumValues Enumerates the set of values for StorageGatewayHealthStatusEnum
func GetStorageGatewayHealthStatusEnumValues() []StorageGatewayHealthStatusEnum {
	values := make([]StorageGatewayHealthStatusEnum, 0)
	for _, v := range mappingStorageGatewayHealthStatus {
		values = append(values, v)
	}
	return values
}
