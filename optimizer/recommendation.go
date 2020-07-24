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

// Recommendation The metadata specific to the recommendation.
type Recommendation struct {

	// The recommendation unique OCID.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The category unique OCID.
	CategoryId *string `mandatory:"true" json:"categoryId"`

	// The name of the recommendation.
	Name *string `mandatory:"true" json:"name"`

	// The description of the recommendation.
	Description *string `mandatory:"true" json:"description"`

	// The Importance of the recommendation.
	Importance ImportanceEnum `mandatory:"true" json:"importance"`

	// An array of ResourceCount grouped by the status of the resource actions.
	ResourceCounts []ResourceCount `mandatory:"true" json:"resourceCounts"`

	// The lifecycleState of the recommendation.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The estimated cost saving in dollars for the recommendation.
	EstimatedCostSaving *float64 `mandatory:"true" json:"estimatedCostSaving"`

	// The recommendation status.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The date and time the current status has started. The format is defined by RFC3339.
	// For example `The status of recommendation has changed from pending to current(ignored) on this date.`
	TimeStatusBegin *common.SDKTime `mandatory:"true" json:"timeStatusBegin"`

	// The date and time the current status will end. The format is defined by RFC3339.
	// For example `The current status of recommendation 'postponed' will end on this date and would change to 'pending'.`
	TimeStatusEnd *common.SDKTime `mandatory:"false" json:"timeStatusEnd"`

	// Date and time the recommendation details was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time the recommendation details was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	SupportedLevels *SupportedLevels `mandatory:"false" json:"supportedLevels"`
}

func (m Recommendation) String() string {
	return common.PointerString(m)
}
