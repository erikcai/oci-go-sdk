// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// CreateManagedCompartmentDetails The representation of CreateManagedCompartmentDetails
type CreateManagedCompartmentDetails struct {

	// The OCID of the tenancy containing the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the compartment during creation. The name must be unique across all compartments
	// in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the compartment during creation. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// The name of the internal service caller.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The service entitlement Id.
	ServiceEntitlementId *string `mandatory:"true" json:"serviceEntitlementId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateManagedCompartmentDetails) String() string {
	return common.PointerString(m)
}
