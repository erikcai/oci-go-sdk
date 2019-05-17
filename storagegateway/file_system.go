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

// FileSystem Detailed view of the file system.
type FileSystem struct {

	// The compartment OCID containing the storage gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The storage gateway OCID containing the file system.
	StorageGatewayId *string `mandatory:"true" json:"storageGatewayId"`

	// A unique file system in the storage gateway.
	Name *string `mandatory:"true" json:"name"`

	// The type of Object Storage tier.
	StorageTier FileSystemStorageTierEnum `mandatory:"true" json:"storageTier"`

	// True if the Object Storage bucket is connected.
	IsConnected *bool `mandatory:"true" json:"isConnected"`

	// A comma-separated with optional whitespace list of hosts allowed to connect to the NFS export.
	// Specifying '*' allows all hosts to connect. Example: 2001:db8:9:e54::/64, 192.0.2.0/24
	NfsAllowedHosts *string `mandatory:"true" json:"nfsAllowedHosts"`

	// The NFS export options. Example: rw, sync, insecure, no_subtree_check, no_root_squash
	// Do not specify the fsid option.
	NfsExportOptions *string `mandatory:"true" json:"nfsExportOptions"`

	// Date and time the file system was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The file system's state.
	LifecycleState LifecycleState `mandatory:"true" json:"lifecycleState"`

	// File system-specific lifecycle substates.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The non-unique, changeable description you assign to the file system.
	Description *string `mandatory:"false" json:"description"`

	// True if the File System is in refresh mode.
	IsInRefreshMode *bool `mandatory:"false" json:"isInRefreshMode"`

	// Reclaim an Object Storage bucket that has been already owned by another file system.
	IsReclaimAttempt *bool `mandatory:"false" json:"isReclaimAttempt"`

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

func (m FileSystem) String() string {
	return common.PointerString(m)
}

// FileSystemStorageTierEnum Enum with underlying type: string
type FileSystemStorageTierEnum string

// Set of constants representing the allowable values for FileSystemStorageTierEnum
const (
	FileSystemStorageTierStandard FileSystemStorageTierEnum = "STANDARD"
	FileSystemStorageTierArchive  FileSystemStorageTierEnum = "ARCHIVE"
)

var mappingFileSystemStorageTier = map[string]FileSystemStorageTierEnum{
	"STANDARD": FileSystemStorageTierStandard,
	"ARCHIVE":  FileSystemStorageTierArchive,
}

// GetFileSystemStorageTierEnumValues Enumerates the set of values for FileSystemStorageTierEnum
func GetFileSystemStorageTierEnumValues() []FileSystemStorageTierEnum {
	values := make([]FileSystemStorageTierEnum, 0)
	for _, v := range mappingFileSystemStorageTier {
		values = append(values, v)
	}
	return values
}
