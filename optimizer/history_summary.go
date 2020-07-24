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

// HistorySummary The metadata specific to the recommendation history for all the resources inside it.
type HistorySummary struct {

	// The recommendation history unique OCID.
	Id *string `mandatory:"true" json:"id"`

	// The name of the resource.
	Name *string `mandatory:"true" json:"name"`

	// The type of the resource.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The category unique OCID.
	CategoryId *string `mandatory:"true" json:"categoryId"`

	// The recommendation unique OCID.
	RecommendationId *string `mandatory:"true" json:"recommendationId"`

	// The resource unique OCID.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The resource action unique OCID.
	ResourceActionId *string `mandatory:"true" json:"resourceActionId"`

	Action *Action `mandatory:"true" json:"action"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the compartment.
	CompartmentName *string `mandatory:"true" json:"compartmentName"`

	// The lifecycleState of the recommendation history.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The status of the associated resource action.
	Status StatusEnum `mandatory:"true" json:"status"`

	// Date and time the recommendation history was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m HistorySummary) String() string {
	return common.PointerString(m)
}
