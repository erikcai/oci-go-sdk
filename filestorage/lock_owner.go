// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// File Storage Service API
//
// The API for the File Storage Service.
//

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LockOwner The client that owns one or more locks in the specified file system.
type LockOwner struct {

	// The OCID of the file system where the locks are held.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// The IP address of the client that holds the locks.
	ClientIpAddress *string `mandatory:"true" json:"clientIpAddress"`

	// The OCID of the mount target through which locks are held.
	MountTargetId *string `mandatory:"true" json:"mountTargetId"`

	// The OCID of the mount target's virtual cloud network (VCN).
	VcnId *string `mandatory:"true" json:"vcnId"`
}

func (m LockOwner) String() string {
	return common.PointerString(m)
}
