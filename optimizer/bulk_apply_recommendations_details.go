// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Optimizer API
//
// The API for the OCI Optimizer
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BulkApplyRecommendationsDetails Properties for optimizer resource recommendation actions in bulk
type BulkApplyRecommendationsDetails struct {

	// The OCIDs of the optimizer resource actions to apply recommendations for.
	ResourceActionIds []string `mandatory:"true" json:"resourceActionIds"`

	// The recommendation status.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The date and time the current status will end. The format is defined by RFC3339.
	// For example `The current status of recommendation 'postponed' will end on this date and would change to 'pending'.`
	TimeStatusEnd *common.SDKTime `mandatory:"false" json:"timeStatusEnd"`
}

func (m BulkApplyRecommendationsDetails) String() string {
	return common.PointerString(m)
}
