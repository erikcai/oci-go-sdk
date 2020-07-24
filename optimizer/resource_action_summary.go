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

// ResourceActionSummary The metadata specific to the resource action summary.
type ResourceActionSummary struct {

	// The resource action unique OCID.
	Id *string `mandatory:"true" json:"id"`

	// The category unique OCID.
	CategoryId *string `mandatory:"true" json:"categoryId"`

	// The recommendation unique OCID.
	RecommendationId *string `mandatory:"true" json:"recommendationId"`

	// The resource unique OCID.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The name of the resource.
	Name *string `mandatory:"true" json:"name"`

	// The type of the resource.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the compartment.
	CompartmentName *string `mandatory:"true" json:"compartmentName"`

	Action *Action `mandatory:"true" json:"action"`

	// The lifecycleState of the resource action.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The estimated cost saving in dollars for the resource action.
	EstimatedCostSaving *float64 `mandatory:"true" json:"estimatedCostSaving"`

	// The resource action status.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The date and time the current status has started. The format is defined by RFC3339.
	// For example `The status of resource action has changed from pending to current(ignored) on this date.`
	TimeStatusBegin *common.SDKTime `mandatory:"true" json:"timeStatusBegin"`

	// The date and time the current status will end. The format is defined by RFC3339.
	// For example `The current status of resource action 'postponed' will end on this date and would change to 'pending'.``
	TimeStatusEnd *common.SDKTime `mandatory:"false" json:"timeStatusEnd"`

	// Date and time the resource action details was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time the resource action details was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ResourceActionSummary) String() string {
	return common.PointerString(m)
}
