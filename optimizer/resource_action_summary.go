// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResourceActionSummary The metadata associated with the resource action summary.
// **Caution:** Avoid using any confidential information when you use the API to supply string values.
type ResourceActionSummary struct {

	// The unique OCID associated with the resource action.
	Id *string `mandatory:"true" json:"id"`

	// The unique OCID associated with the category.
	CategoryId *string `mandatory:"true" json:"categoryId"`

	// The unique OCID associated with the recommendation.
	RecommendationId *string `mandatory:"true" json:"recommendationId"`

	// The unique OCID associated with the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The name assigned to the resource.
	Name *string `mandatory:"true" json:"name"`

	// The kind of resource.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name associated with the compartment.
	CompartmentName *string `mandatory:"true" json:"compartmentName"`

	Action *Action `mandatory:"true" json:"action"`

	// The resource action's current state.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The estimated cost savings, in dollars, for the resource action.
	EstimatedCostSaving *float64 `mandatory:"true" json:"estimatedCostSaving"`

	// The current status of the resource action.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The date and time that the resource action entered its current status. The format is defined by RFC3339.
	// For example, "The status of the resource action changed from `pending` to `current(ignored)` on this date and time."
	TimeStatusBegin *common.SDKTime `mandatory:"true" json:"timeStatusBegin"`

	// The date and time the current status will change. The format is defined by RFC3339.
	// For example, "The current `postponed` status of the resource action will end and change to `pending` on this
	// date and time."
	TimeStatusEnd *common.SDKTime `mandatory:"false" json:"timeStatusEnd"`

	// The date and time the resource action details were created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the resource action details were last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ResourceActionSummary) String() string {
	return common.PointerString(m)
}
