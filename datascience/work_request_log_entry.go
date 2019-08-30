// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Science API
//
// The Data Science service enables data science teams to organize their work, easily access data and computing resources, and build, train, deploy, and manage ML/AI models on the Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestLogEntry Log entries related to a specific work request.
type WorkRequestLogEntry struct {

	// The description of an action that occurred.
	Message *string `mandatory:"false" json:"message"`

	// The date and time the log entry occurred.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`
}

func (m WorkRequestLogEntry) String() string {
	return common.PointerString(m)
}
