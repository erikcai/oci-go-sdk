// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Storage Gateway API
//
// API for the Storage Gateway service. Use this API to manage storage gateways and related items. For more
// information, see Overview of Storage Gateway (https://docs.cloud.oracle.com/iaas/Content/StorageGateway/Concepts/storagegatewayoverview.htm).
//

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/v32/common"
)

// CreateCloudSyncDetails Configuration details for creating a cloud sync.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateCloudSyncDetails struct {

	// A name for the cloud sync. It must be unique within the storage gateway, and it cannot be changed.
	// Example: `cloud_sync_52019`
	Name *string `mandatory:"true" json:"name"`

	// The path to a source directory or file for the cloud sync.
	// The path configuration depends on the direction of the cloud sync. To upload from an on-premises system to
	// the cloud, the source path resembles the following:
	// /cloudsync/mounts/<var>&lt;user_mount&gt;</var>/<var>&lt;path_to_directory&gt;</var>
	// To download from the cloud to an on-premises system, the source path resembles the following:
	// <var>&lt;storage_gateway_file_system&gt;</var>/<var>&lt;path_to_directory&gt;</var>
	// **Note:** To configure a cloud sync, the file system on an on-premises storage gateway must be mounted under
	// `/cloudsync/mounts`.
	SourcePath *string `mandatory:"true" json:"sourcePath"`

	// The path to a target directory or file for the cloud sync.
	// The path configuration depends on the direction of the cloud sync. To upload from an on-premises system to
	// the cloud, the target path resembles the following:
	// <var>&lt;storage_gateway_file_system&gt;</var>/<var>&lt;path_to_directory&gt;</var>
	// To download from the cloud to an on-premises system, the target path resembles the following:
	// /cloudsync/mounts/<var>&lt;user_mount&gt;</var>/<var>&lt;path_to_directory&gt;</var>
	// **Note:** To configure a cloud sync, the file system on an on-premises storage gateway must be mounted under
	// `/cloudsync/mounts`.
	TargetPath *string `mandatory:"true" json:"targetPath"`

	// A description of the cloud sync. It does not have to be unique, and it is changeable.
	// Example: `my first cloud sync`
	Description *string `mandatory:"false" json:"description"`

	// Whether the cloud sync automatically deletes files from the target when source files are renamed or deleted.
	// If "true", the cloud sync automatically deletes files from the target.
	// Example: `true`
	IsAutoDeletionEnabled *bool `mandatory:"false" json:"isAutoDeletionEnabled"`

	// The path to a file that lists a set of files to sync to the target. If you do not specify a file list, the
	// service syncs all files. The list file should reside under the `/cloudsync/` directory on the machine running
	// the storage gateway instance.
	// Example: `/cloudsync/files.list`
	FilesFrom *string `mandatory:"false" json:"filesFrom"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information,
	// see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCloudSyncDetails) String() string {
	return common.PointerString(m)
}
