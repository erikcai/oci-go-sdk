// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AttachChildSoftwareSourceToManagedInstanceDetails Information for attaching a software source to a managed instance
type AttachChildSoftwareSourceToManagedInstanceDetails struct {

	// OCID for the Software Source
	SoftwareSourceId *string `mandatory:"true" json:"softwareSourceId"`
}

func (m AttachChildSoftwareSourceToManagedInstanceDetails) String() string {
	return common.PointerString(m)
}
