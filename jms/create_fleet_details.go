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

// CreateFleetDetails Attributes to create a fleet
type CreateFleetDetails struct {

	// The name of the fleet.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the fleet
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The tagging key-value pairs that are used to associate managed instances to the fleet.
	// A report submitted by a managed instance is associated to the fleet if the managed instance has at least one freeform tag from the fleets key-value pairs.
	// The set of tagging key-value pairs serves as a natural key within a compartment, thus there can't be two fleets in a compartment with the same tagging configuration.
	ManagedInstanceTaggingConfigurations []ManagedInstanceTaggingConfiguration `mandatory:"true" json:"managedInstanceTaggingConfigurations"`

	// The fleet's description. If nothing is provided, the fleet description will be null.
	Description *string `mandatory:"false" json:"description"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateFleetDetails) String() string {
	return common.PointerString(m)
}
