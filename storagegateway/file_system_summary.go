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

// FileSystemSummary Summary view of a file system.
type FileSystemSummary struct {

	// A unique file system name in the given storage gateway.
	Name *string `mandatory:"true" json:"name"`

	// The tyep of Object Storage tier.
	StorageTier FileSystemSummaryStorageTierEnum `mandatory:"true" json:"storageTier"`

	// True if the Object Storage bucket is connected.
	IsConnected *bool `mandatory:"true" json:"isConnected"`

	// Error count.
	ErrorCount *float32 `mandatory:"true" json:"errorCount"`

	// Warning count.
	WarnCount *float32 `mandatory:"true" json:"warnCount"`

	// Date and time the file system was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The file system's state. After creating the resource, make sure its state changes
	// to ACTIVE before using it.
	LifecycleState LifecycleState `mandatory:"true" json:"lifecycleState"`

	// File system-specific lifecycle substates: File system has one of the following substates:
	// "NONE", "CONNECTING", "DISCONNECTING", "RECLAIMING", "REFRESHING", "UPDATING".
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// True if the File System is in refresh mode.
	IsInRefreshMode *bool `mandatory:"false" json:"isInRefreshMode"`

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

func (m FileSystemSummary) String() string {
	return common.PointerString(m)
}

// FileSystemSummaryStorageTierEnum Enum with underlying type: string
type FileSystemSummaryStorageTierEnum string

// Set of constants representing the allowable values for FileSystemSummaryStorageTierEnum
const (
	FileSystemSummaryStorageTierStandard FileSystemSummaryStorageTierEnum = "STANDARD"
	FileSystemSummaryStorageTierArchive  FileSystemSummaryStorageTierEnum = "ARCHIVE"
)

var mappingFileSystemSummaryStorageTier = map[string]FileSystemSummaryStorageTierEnum{
	"STANDARD": FileSystemSummaryStorageTierStandard,
	"ARCHIVE":  FileSystemSummaryStorageTierArchive,
}

// GetFileSystemSummaryStorageTierEnumValues Enumerates the set of values for FileSystemSummaryStorageTierEnum
func GetFileSystemSummaryStorageTierEnumValues() []FileSystemSummaryStorageTierEnum {
	values := make([]FileSystemSummaryStorageTierEnum, 0)
	for _, v := range mappingFileSystemSummaryStorageTier {
		values = append(values, v)
	}
	return values
}
