// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ProblemTrendAggregation Provides the dimensions and their corresponding time and count.
type ProblemTrendAggregation struct {

	// The key-value pairs of dimensions and their names.
	DimensionsMap map[string]string `mandatory:"false" json:"dimensionsMap"`

	// Start Time
	StartTimestamp *float32 `mandatory:"false" json:"startTimestamp"`

	// Duration
	DurationInSeconds *int `mandatory:"false" json:"durationInSeconds"`

	// The number of occurences with for the corresponding time range and dimensions.
	Count *int `mandatory:"false" json:"count"`
}

func (m ProblemTrendAggregation) String() string {
	return common.PointerString(m)
}
