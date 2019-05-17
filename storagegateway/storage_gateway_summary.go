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

// StorageGatewaySummary Summary view of the storage gateway.
type StorageGatewaySummary struct {

	// The storage gateway's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the storage gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The non-unique name assigned to the storage gateway during creation or update.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current storage gateway status
	Status StorageGatewaySummaryStatusEnum `mandatory:"true" json:"status"`

	// The current version of storage gateway instance on-premise or on-compute.
	Version *string `mandatory:"true" json:"version"`

	// Availability of a newer version.
	IsNewerVersionAvailable *bool `mandatory:"true" json:"isNewerVersionAvailable"`

	// Date and time the storage gateway was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The storage gateway's state. After creating the resource, make sure its state changes
	// to ACTIVE before using it.
	LifecycleState LifecycleState `mandatory:"true" json:"lifecycleState"`

	// The non-unique, changeable description you assign to the storage gateway during creation or update.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information,
	// see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m StorageGatewaySummary) String() string {
	return common.PointerString(m)
}

// StorageGatewaySummaryStatusEnum Enum with underlying type: string
type StorageGatewaySummaryStatusEnum string

// Set of constants representing the allowable values for StorageGatewaySummaryStatusEnum
const (
	StorageGatewaySummaryStatusActive   StorageGatewaySummaryStatusEnum = "ACTIVE"
	StorageGatewaySummaryStatusInactive StorageGatewaySummaryStatusEnum = "INACTIVE"
	StorageGatewaySummaryStatusWarning  StorageGatewaySummaryStatusEnum = "WARNING"
	StorageGatewaySummaryStatusCritical StorageGatewaySummaryStatusEnum = "CRITICAL"
)

var mappingStorageGatewaySummaryStatus = map[string]StorageGatewaySummaryStatusEnum{
	"ACTIVE":   StorageGatewaySummaryStatusActive,
	"INACTIVE": StorageGatewaySummaryStatusInactive,
	"WARNING":  StorageGatewaySummaryStatusWarning,
	"CRITICAL": StorageGatewaySummaryStatusCritical,
}

// GetStorageGatewaySummaryStatusEnumValues Enumerates the set of values for StorageGatewaySummaryStatusEnum
func GetStorageGatewaySummaryStatusEnumValues() []StorageGatewaySummaryStatusEnum {
	values := make([]StorageGatewaySummaryStatusEnum, 0)
	for _, v := range mappingStorageGatewaySummaryStatus {
		values = append(values, v)
	}
	return values
}
