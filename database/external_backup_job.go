// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ExternalBackupJob Provides all the details of a backup, along with additional details that are applicable to an external backup.
type ExternalBackupJob struct {

	// The id of the associated backup resource.
	BackupId *string `mandatory:"true" json:"backupId"`

	// The bucket where the backup should be stored.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// An indicator for the provisioning state of the resource. If TRUE, the resource is still being provisioned.
	Provisioning *bool `mandatory:"true" json:"provisioning"`

	// The Swift path to be used to take the backup.
	SwiftPath *string `mandatory:"true" json:"swiftPath"`

	// The tag to be used by RMAN for the backup.
	Tag *string `mandatory:"true" json:"tag"`

	// The Swift user name to be used to take the backup.
	UserName *string `mandatory:"true" json:"userName"`

	// The Swift password to be used to take the backup.
	SwiftPassword *string `mandatory:"false" json:"swiftPassword"`
}

func (m ExternalBackupJob) String() string {
	return common.PointerString(m)
}
