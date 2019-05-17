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

// CreateFileSystemDetails Details to set when creating a file system.
type CreateFileSystemDetails struct {

	// A unique file system name in the storage gateway. If an Object Storage bucket matching the
	// file system name does not exist, it will be created.
	Name *string `mandatory:"true" json:"name"`

	// The type of Object Storage tier to store data. The Standard tier is the primary default Object
	// Storage tier for storing data that is accessed frequently and requires fast and immediate access.
	StorageTier CreateFileSystemDetailsStorageTierEnum `mandatory:"true" json:"storageTier"`

	// The non-unique, changeable description you assign to the file system during creation.
	Description *string `mandatory:"false" json:"description"`

	// A comma-separated with optional whitespace list of hosts allowed to connect to the NFS export.
	// Specifying '*' allows all hosts to connect. Example: 2001:db8:9:e54::/64, 192.0.2.0/24
	NfsAllowedHosts *string `mandatory:"false" json:"nfsAllowedHosts"`

	// The NFS export options. Example: rw, sync, insecure, no_subtree_check, no_root_squash
	// Do not specify the fsid option.
	NfsExportOptions *string `mandatory:"false" json:"nfsExportOptions"`

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

func (m CreateFileSystemDetails) String() string {
	return common.PointerString(m)
}

// CreateFileSystemDetailsStorageTierEnum Enum with underlying type: string
type CreateFileSystemDetailsStorageTierEnum string

// Set of constants representing the allowable values for CreateFileSystemDetailsStorageTierEnum
const (
	CreateFileSystemDetailsStorageTierStandard CreateFileSystemDetailsStorageTierEnum = "STANDARD"
	CreateFileSystemDetailsStorageTierArchive  CreateFileSystemDetailsStorageTierEnum = "ARCHIVE"
)

var mappingCreateFileSystemDetailsStorageTier = map[string]CreateFileSystemDetailsStorageTierEnum{
	"STANDARD": CreateFileSystemDetailsStorageTierStandard,
	"ARCHIVE":  CreateFileSystemDetailsStorageTierArchive,
}

// GetCreateFileSystemDetailsStorageTierEnumValues Enumerates the set of values for CreateFileSystemDetailsStorageTierEnum
func GetCreateFileSystemDetailsStorageTierEnumValues() []CreateFileSystemDetailsStorageTierEnum {
	values := make([]CreateFileSystemDetailsStorageTierEnum, 0)
	for _, v := range mappingCreateFileSystemDetailsStorageTier {
		values = append(values, v)
	}
	return values
}
