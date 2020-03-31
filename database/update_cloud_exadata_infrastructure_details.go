// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateCloudExadataInfrastructureDetails Updates the cloud Exadata infrastructure.
type UpdateCloudExadataInfrastructureDetails struct {

	// The user-friendly name for the cloud Exadata infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The additional number of compute servers for the cloud Exadata infrastructure.
	AdditionalComputeCount *int64 `mandatory:"false" json:"additionalComputeCount"`

	// The additional number of storage servers for the cloud Exadata infrastructure.
	AdditionalStorageCount *int64 `mandatory:"false" json:"additionalStorageCount"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCloudExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}