// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// CreateRoverEntitlementDetails The representation of CreateRoverEntitlementDetails
type CreateRoverEntitlementDetails struct {

	// Compartment id for the entitlement.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Display name for the entitlement.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Requestor name for the entitlement.
	RequestorName *string `mandatory:"false" json:"requestorName"`

	// Requestor email for the entitlement.
	RequestorEmail *string `mandatory:"false" json:"requestorEmail"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRoverEntitlementDetails) String() string {
	return common.PointerString(m)
}
