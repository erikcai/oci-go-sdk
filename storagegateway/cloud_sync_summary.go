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

// CloudSyncSummary Summary view of the cloud sync.
type CloudSyncSummary struct {

	// A unique cloud sync name in the given storage gateway.
	Name *string `mandatory:"true" json:"name"`

	// A path to a directory or file of the cloud sync source.
	SourcePath *string `mandatory:"true" json:"sourcePath"`

	// A path to a directory or file of the cloud sync target.
	TargetPath *string `mandatory:"true" json:"targetPath"`

	// Set to true if the cloud sync uploads data to the cloud.
	IsUpload *bool `mandatory:"true" json:"isUpload"`

	// True if cloud sync needs to automatically delete files from the target when files are
	// deleted from the source or source files have been renamed.
	IsAutoDeletionEnabled *bool `mandatory:"true" json:"isAutoDeletionEnabled"`

	// A path to the file that includes a set of files to sync to the target. If not provided,
	// the service syncs all files. The file should be a file under the /cloudsync/ directory,
	// for example, "/cloudsync/files.list" on the VM running the storage gateway instance.
	FilesFrom *string `mandatory:"true" json:"filesFrom"`

	// Error count.
	ErrorCount *float32 `mandatory:"true" json:"errorCount"`

	// Warning count.
	WarnCount *float32 `mandatory:"true" json:"warnCount"`

	// Date and time the cloud sync created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the cloud sync started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The cloud sync's state. After creating the resource, make sure its state changes
	// to ACTIVE before using it.
	LifecycleState LifecycleState `mandatory:"true" json:"lifecycleState"`

	// Cloud sync-specific lifecycle substates: Cloud sync has one of the following substates:
	// "NONE", "CREATED", "RUN", "RUNNING", "COMPLETED", "FAILED", "CANCELING", "CANCELED",
	// "UPDATING"
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Date and time the cloud sync was completed, canceled, or failed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

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

func (m CloudSyncSummary) String() string {
	return common.PointerString(m)
}
