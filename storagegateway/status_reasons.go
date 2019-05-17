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

// StatusReasons Reasons for the overall health status: one array of reason strings for the 'CRITICAL'
// health status and another array of reason strings for the 'WARNING' health status.
// The 'CRITICAL' health status will have both arrays if any warning exists.
// Example:
//   `{"critical": [
//       "Rejecting IO due to low cache space",
//       "File system 'fs1' failed to upload file 'f1'"
//     ],
//     "warning": [
//       "A newer version is available"
//     ]}`
type StatusReasons struct {

	// array of reasons for the critical status
	Critical []string `mandatory:"false" json:"critical"`

	// array of reasons for the warning status
	Warning []string `mandatory:"false" json:"warning"`
}

func (m StatusReasons) String() string {
	return common.PointerString(m)
}
