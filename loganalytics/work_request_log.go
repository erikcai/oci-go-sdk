// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// WorkRequestLog This is a log message from the execution of a work request.
type WorkRequestLog struct {

	// This is a human-readable log message.
	Message *string `mandatory:"true" json:"message"`

	// This is the time the log message was written. An RFC3339 formatted datetime string
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestLog) String() string {
	return common.PointerString(m)
}
