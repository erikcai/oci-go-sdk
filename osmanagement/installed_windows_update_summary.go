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

// InstalledWindowsUpdateSummary A Windows update installed on the Windows managed instance.
type InstalledWindowsUpdateSummary struct {

	// Windows Update name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the Windows update. NOTE - This is not an OCID,
	// but is a unique identifier assigned by Microsoft.
	// Example: `6981d463-cd91-4a26-b7c4-ea4ded9183ed`
	Name *string `mandatory:"true" json:"name"`

	// The purpose of this update.
	UpdateType UpdateTypesEnum `mandatory:"true" json:"updateType"`
}

func (m InstalledWindowsUpdateSummary) String() string {
	return common.PointerString(m)
}
