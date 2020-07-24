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

// CategorySummary The metadata specific to the category summary.
type CategorySummary struct {

	// The category unique OCID.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the category.
	Name *string `mandatory:"true" json:"name"`

	// The description of the category.
	Description *string `mandatory:"true" json:"description"`

	// An array of RecommendationCount grouped by the importance of the recommendation.
	RecommendationCounts []RecommendationCount `mandatory:"true" json:"recommendationCounts"`

	// An array of ResourceCount grouped by the status of the recommendation.
	ResourceCounts []ResourceCount `mandatory:"true" json:"resourceCounts"`

	// The lifecycleState of the category.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The estimated cost saving in dollars for the category.
	EstimatedCostSaving *float64 `mandatory:"true" json:"estimatedCostSaving"`

	// Date and time the category details was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the category details was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m CategorySummary) String() string {
	return common.PointerString(m)
}
