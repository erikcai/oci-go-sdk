// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RoverEntitlement The representation of RoverEntitlement
type RoverEntitlement struct {

	// The compartment Id for the entitlement.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	LifecycleState RoverEntitlementLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A property that can uniquely identify the rover entitlement.
	Id *string `mandatory:"false" json:"id"`

	// The display name for the entitlement.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Requestor name for the entitlement.
	RequestorName *string `mandatory:"false" json:"requestorName"`

	// Requestor email for the entitlement.
	RequestorEmail *string `mandatory:"false" json:"requestorEmail"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time of creation for the entitlement.
	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	// Time when the entitlement was last updated.
	UpdateTime *common.SDKTime `mandatory:"false" json:"updateTime"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m RoverEntitlement) String() string {
	return common.PointerString(m)
}

// RoverEntitlementLifecycleStateEnum Enum with underlying type: string
type RoverEntitlementLifecycleStateEnum string

// Set of constants representing the allowable values for RoverEntitlementLifecycleStateEnum
const (
	RoverEntitlementLifecycleStateCreating RoverEntitlementLifecycleStateEnum = "CREATING"
	RoverEntitlementLifecycleStateActive   RoverEntitlementLifecycleStateEnum = "ACTIVE"
	RoverEntitlementLifecycleStateInactive RoverEntitlementLifecycleStateEnum = "INACTIVE"
	RoverEntitlementLifecycleStateDeleted  RoverEntitlementLifecycleStateEnum = "DELETED"
)

var mappingRoverEntitlementLifecycleState = map[string]RoverEntitlementLifecycleStateEnum{
	"CREATING": RoverEntitlementLifecycleStateCreating,
	"ACTIVE":   RoverEntitlementLifecycleStateActive,
	"INACTIVE": RoverEntitlementLifecycleStateInactive,
	"DELETED":  RoverEntitlementLifecycleStateDeleted,
}

// GetRoverEntitlementLifecycleStateEnumValues Enumerates the set of values for RoverEntitlementLifecycleStateEnum
func GetRoverEntitlementLifecycleStateEnumValues() []RoverEntitlementLifecycleStateEnum {
	values := make([]RoverEntitlementLifecycleStateEnum, 0)
	for _, v := range mappingRoverEntitlementLifecycleState {
		values = append(values, v)
	}
	return values
}
