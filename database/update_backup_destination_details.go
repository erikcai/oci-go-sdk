// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateBackupDestinationDetails Used for adding Recovery Appliance backup destination users or backup destination in case of Recovery Appliance , NFS if no database is attached yet.
type UpdateBackupDestinationDetails struct {

	// The Virtual Private Catalog users that will be used to access the Zero Data Loss Recovery Appliance.
	VpcUsers []string `mandatory:"false" json:"vpcUsers"`

	// The connection string that is used to connect to the Zero Data Loss Recovery Appliance.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The network file path of the NFS device to be mounted.In the format server:/directory/folder.
	Path *string `mandatory:"false" json:"path"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateBackupDestinationDetails) String() string {
	return common.PointerString(m)
}
