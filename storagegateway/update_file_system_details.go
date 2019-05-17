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

// UpdateFileSystemDetails Details to set when updating the file system.
type UpdateFileSystemDetails struct {

	// A non-unique, changeable description for the file system.
	Description *string `mandatory:"false" json:"description"`

	// A comma-separated with optional whitespace list of hosts allowed to connect to the NFS export.
	// Specifying '*' allows all hosts to connect. Example: 2001:db8:9:e54::/64, 192.0.2.0/24
	NfsAllowedHosts *string `mandatory:"false" json:"nfsAllowedHosts"`

	// The NFS export options. Example: rw, sync, insecure, no_subtree_check, no_root_squash
	// Do not specify the fsid option.
	NfsExportOptions *string `mandatory:"false" json:"nfsExportOptions"`

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

func (m UpdateFileSystemDetails) String() string {
	return common.PointerString(m)
}
