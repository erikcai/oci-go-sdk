// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// Endpoint Combination of server dns or ip and port
type Endpoint struct {

	// String consisting of domain name server for endpoint
	Hostname *string `mandatory:"true" json:"hostname"`

	// integer value of server endpoint
	Port *int64 `mandatory:"true" json:"port"`
}

func (m Endpoint) String() string {
	return common.PointerString(m)
}
