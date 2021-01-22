// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Synthetic
//
// API for APM Synthetic service. Use this API to query Synthethetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// MonitorResult The  monitor result for a specific execution.
type MonitorResult struct {

	// Type of the result content like zip or raw.
	ResultContentType *string `mandatory:"true" json:"resultContentType"`

	// Type of the result like har or screenshot or log.
	ResultType *string `mandatory:"false" json:"resultType"`

	// Monitor result data set.
	ResultDataSet []MonitorResultData `mandatory:"false" json:"resultDataSet"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the monitor.
	MonitorId *string `mandatory:"false" json:"monitorId"`

	// The vantagePoint name.
	VantagePoint *string `mandatory:"false" json:"vantagePoint"`

	// The specific point of time when result of an execution collected.
	ExecutionTime *string `mandatory:"false" json:"executionTime"`
}

func (m MonitorResult) String() string {
	return common.PointerString(m)
}
