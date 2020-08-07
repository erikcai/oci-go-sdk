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

// HistorySummary The metadata associated with the recommendation history and its related resources.
// **Caution:** Avoid using any confidential information when you use the API to supply string values.
type HistorySummary struct {

	// The unique OCID associated with the recommendation history.
	Id *string `mandatory:"true" json:"id"`

	// The name assigned to the resource.
	Name *string `mandatory:"true" json:"name"`

	// The kind of resource.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The unique OCID associated with the category.
	CategoryId *string `mandatory:"true" json:"categoryId"`

	// The unique OCID associated with the recommendation.
	RecommendationId *string `mandatory:"true" json:"recommendationId"`

	// The unique OCID associated with the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The unique OCID associated with the resource action.
	ResourceActionId *string `mandatory:"true" json:"resourceActionId"`

	Action *Action `mandatory:"true" json:"action"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name assigned to the compartment.
	CompartmentName *string `mandatory:"true" json:"compartmentName"`

	// The recommendation history's current state.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current status of the resource action.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The date and time the recommendation history was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m HistorySummary) String() string {
	return common.PointerString(m)
}
