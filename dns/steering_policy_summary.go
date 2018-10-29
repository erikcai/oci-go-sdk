// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SteeringPolicySummary A DNS steering policy.
// *Warning:* Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type SteeringPolicySummary struct {

	// The OCID of the compartment containing the steering policy.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// A user-friendly name for the steering policy.
	// Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The Time To Live for responses from the steering policy, in seconds.
	// If not specified during creation, a value of 30 seconds will be used.
	Ttl *int `mandatory:"false" json:"ttl"`

	// The OCID of the health check monitor providing health data about the answers of the
	// steering policy.
	// A steering policy answer with `rdata` matching a monitored endpoint will use the health
	// data of that endpoint.
	// A steering policy answer with `rdata` not matching any monitored endpoint will be assumed
	// healthy.
	HealthCheckMonitorId *string `mandatory:"false" json:"healthCheckMonitorId"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"false" json:"self"`

	// The OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// The date and time the resource was created in "YYYY-MM-ddThh:mmZ" format
	// with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the resource.
	LifecycleState SteeringPolicySummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m SteeringPolicySummary) String() string {
	return common.PointerString(m)
}

// SteeringPolicySummaryLifecycleStateEnum Enum with underlying type: string
type SteeringPolicySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SteeringPolicySummaryLifecycleStateEnum
const (
	SteeringPolicySummaryLifecycleStateActive   SteeringPolicySummaryLifecycleStateEnum = "ACTIVE"
	SteeringPolicySummaryLifecycleStateCreating SteeringPolicySummaryLifecycleStateEnum = "CREATING"
	SteeringPolicySummaryLifecycleStateDeleted  SteeringPolicySummaryLifecycleStateEnum = "DELETED"
	SteeringPolicySummaryLifecycleStateDeleting SteeringPolicySummaryLifecycleStateEnum = "DELETING"
)

var mappingSteeringPolicySummaryLifecycleState = map[string]SteeringPolicySummaryLifecycleStateEnum{
	"ACTIVE":   SteeringPolicySummaryLifecycleStateActive,
	"CREATING": SteeringPolicySummaryLifecycleStateCreating,
	"DELETED":  SteeringPolicySummaryLifecycleStateDeleted,
	"DELETING": SteeringPolicySummaryLifecycleStateDeleting,
}

// GetSteeringPolicySummaryLifecycleStateEnumValues Enumerates the set of values for SteeringPolicySummaryLifecycleStateEnum
func GetSteeringPolicySummaryLifecycleStateEnumValues() []SteeringPolicySummaryLifecycleStateEnum {
	values := make([]SteeringPolicySummaryLifecycleStateEnum, 0)
	for _, v := range mappingSteeringPolicySummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
