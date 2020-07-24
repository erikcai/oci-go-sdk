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

// Profile The metadata specific to the profile summary.
type Profile struct {

	// The profile unique OCID.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the profile.
	Name *string `mandatory:"true" json:"name"`

	// The description of the profile.
	Description *string `mandatory:"true" json:"description"`

	LevelsConfiguration *LevelsConfiguration `mandatory:"true" json:"levelsConfiguration"`

	// The lifecycleState of the profile.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Date and time the profile was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the profile was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m Profile) String() string {
	return common.PointerString(m)
}
