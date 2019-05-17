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

// CreateCloudSyncDetails Details to set when creating a cloud sync.
type CreateCloudSyncDetails struct {

	// A unique name in the storage gateway you assign to the cloud sync during creation.
	Name *string `mandatory:"true" json:"name"`

	// A path to a directory or file of the cloud sync source. Note that the file system on
	// premise needs to be mounted under /cloudsync/mounts.
	// In the case of the upload from the on-premise to the cloud, the source & target paths
	// look like this:
	// Source Path: /cloudsync/mounts/<user_mount>[/<path_to_folder>] and
	// Target Path: <storage_gateway_file_system>[/<path_to_folder>].
	// In the case of the download from the cloud to the on-premise, the source & target paths
	// look like this:
	// Source Path: <storage_gateway_file_system>[/<path_to_folder>]
	// Target Path: /cloudsync/mounts/<user_mount>[/<path_to_folder>]
	SourcePath *string `mandatory:"true" json:"sourcePath"`

	// A path to a directory or file of the cloud sync target.
	TargetPath *string `mandatory:"true" json:"targetPath"`

	// The non-unique, changeable description you assign to the cloud sync during creation.
	Description *string `mandatory:"false" json:"description"`

	// True if cloud sync needs to automatically delete files from the target when files are
	// deleted from the source or source files have been renamed.
	IsAutoDeletionEnabled *bool `mandatory:"false" json:"isAutoDeletionEnabled"`

	// A path to the file that includes a set of files to sync to the target. If not provided,
	// the service syncs all files. The file should be a file under the /cloudsync/ directory,
	// for example, "/cloudsync/files.list" on the VM running the storage gateway instance.
	FilesFrom *string `mandatory:"false" json:"filesFrom"`

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

func (m CreateCloudSyncDetails) String() string {
	return common.PointerString(m)
}
