// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// Fleet A Fleet is the primary administrative collection that an JMS customer interact with in order to perform functions across a subset of managed instances.
type Fleet struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the fleet
	Id *string `mandatory:"true" json:"id"`

	// The name of the fleet.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The fleet's description
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the fleet
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The approximate count of all distinct Java Runtimes in the fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource etag.
	ApproximateJreCount *int `mandatory:"true" json:"approximateJreCount"`

	// The approximate count of all distinct applications in the fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource etag.
	ApproximateApplicationCount *int `mandatory:"true" json:"approximateApplicationCount"`

	// The RFC3339 compatible creation time of the fleet
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycle state of the fleet
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The tagging key-value pairs that are used to associate managed instances to the fleet.
	// A report submitted by a managed instance is associated to the fleet if the managed instance has at least one freeform tag from the fleets key-value pairs.
	// The set of tagging configurations for a fleet is unique for the compartment.
	ManagedInstanceTaggingConfigurations []ManagedInstanceTaggingConfiguration `mandatory:"false" json:"managedInstanceTaggingConfigurations"`
}

func (m Fleet) String() string {
	return common.PointerString(m)
}
