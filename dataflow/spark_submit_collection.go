// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// SparkSubmitCollection The results of a query for a list of spark submit objects. It contains SparkSubmitSummary items.
type SparkSubmitCollection struct {

	// A list of spark submits.
	Items []SparkSubmitSummary `mandatory:"true" json:"items"`
}

func (m SparkSubmitCollection) String() string {
	return common.PointerString(m)
}
