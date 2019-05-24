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

// CloudSync Detailed view of the cloud sync.
type CloudSync struct {

	// The compartment OCID containing the storage gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The storage gateway OCID to create a cloud sync.
	StorageGatewayId *string `mandatory:"true" json:"storageGatewayId"`

	// A unique name in the storage gateway you assign to the cloud sync during creation.
	Name *string `mandatory:"true" json:"name"`

	// The cloud Sync's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

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

	// Date and time the cloud sync created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the cloud sync started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The cloud sync's state.
	LifecycleState LifecycleState `mandatory:"true" json:"lifecycleState"`

	// Cloud sync-specific lifecycle substates.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The non-unique, changeable description you assign to the cloud sync during creation.
	Description *string `mandatory:"false" json:"description"`

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

func (m CloudSync) String() string {
	return common.PointerString(m)
}
