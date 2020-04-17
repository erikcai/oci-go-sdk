// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CloudExadataInfrastructure CloudExadataInfrastructure
type CloudExadataInfrastructure struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the cloud Exadata infrastructure.
	LifecycleState CloudExadataInfrastructureLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the Cloud Exadata infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The model name of the Cloud Exadata infrastructure.
	Shape *string `mandatory:"true" json:"shape"`

	// The name of the availability domain that the cloud Exadata infrastructure is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The additional number of compute servers for the cloud Exadata infrastructure.
	AdditionalComputeCount *int64 `mandatory:"false" json:"additionalComputeCount"`

	// The additional number of storage servers for the cloud Exadata infrastructure.
	AdditionalStorageCount *int64 `mandatory:"false" json:"additionalStorageCount"`

	// The date and time the cloud Exadata infrastructure was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CloudExadataInfrastructure) String() string {
	return common.PointerString(m)
}

// CloudExadataInfrastructureLifecycleStateEnum Enum with underlying type: string
type CloudExadataInfrastructureLifecycleStateEnum string

// Set of constants representing the allowable values for CloudExadataInfrastructureLifecycleStateEnum
const (
	CloudExadataInfrastructureLifecycleStateProvisioning CloudExadataInfrastructureLifecycleStateEnum = "PROVISIONING"
	CloudExadataInfrastructureLifecycleStateAvailable    CloudExadataInfrastructureLifecycleStateEnum = "AVAILABLE"
	CloudExadataInfrastructureLifecycleStateUpdating     CloudExadataInfrastructureLifecycleStateEnum = "UPDATING"
	CloudExadataInfrastructureLifecycleStateTerminating  CloudExadataInfrastructureLifecycleStateEnum = "TERMINATING"
	CloudExadataInfrastructureLifecycleStateTerminated   CloudExadataInfrastructureLifecycleStateEnum = "TERMINATED"
	CloudExadataInfrastructureLifecycleStateFailed       CloudExadataInfrastructureLifecycleStateEnum = "FAILED"
)

var mappingCloudExadataInfrastructureLifecycleState = map[string]CloudExadataInfrastructureLifecycleStateEnum{
	"PROVISIONING": CloudExadataInfrastructureLifecycleStateProvisioning,
	"AVAILABLE":    CloudExadataInfrastructureLifecycleStateAvailable,
	"UPDATING":     CloudExadataInfrastructureLifecycleStateUpdating,
	"TERMINATING":  CloudExadataInfrastructureLifecycleStateTerminating,
	"TERMINATED":   CloudExadataInfrastructureLifecycleStateTerminated,
	"FAILED":       CloudExadataInfrastructureLifecycleStateFailed,
}

// GetCloudExadataInfrastructureLifecycleStateEnumValues Enumerates the set of values for CloudExadataInfrastructureLifecycleStateEnum
func GetCloudExadataInfrastructureLifecycleStateEnumValues() []CloudExadataInfrastructureLifecycleStateEnum {
	values := make([]CloudExadataInfrastructureLifecycleStateEnum, 0)
	for _, v := range mappingCloudExadataInfrastructureLifecycleState {
		values = append(values, v)
	}
	return values
}
