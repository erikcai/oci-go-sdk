// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v28/common"
)

// UpdateRoverEntitlementDetails The representation of UpdateRoverEntitlementDetails
type UpdateRoverEntitlementDetails struct {

	// Display name for the entitlement.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// tenant Id.
	TenantId *string `mandatory:"false" json:"tenantId"`

	// Requestor name for the entitlement.
	RequestorName *string `mandatory:"false" json:"requestorName"`

	// Requestor email for the entitlement.
	RequestorEmail *string `mandatory:"false" json:"requestorEmail"`

	// Details about the entitlement.
	EntitlementDetails *string `mandatory:"false" json:"entitlementDetails"`

	// The current state of the RoverNode.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateRoverEntitlementDetails) String() string {
	return common.PointerString(m)
}
