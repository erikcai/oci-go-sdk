// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// API for the Data Caching Service. Use this service to manage Redis replicated caches.
//

package cache

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// WorkRequestLogEntry A log entry of a work request.
type WorkRequestLogEntry struct {

	// The log message.
	Message *string `mandatory:"true" json:"message"`

	// The time the log message was written.
	TimeStamp *common.SDKTime `mandatory:"true" json:"timeStamp"`
}

func (m WorkRequestLogEntry) String() string {
	return common.PointerString(m)
}
