// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v27/common"
)

// CreateRoverNodeDetails The information about new RoverNode.
type CreateRoverNodeDetails struct {

	// RoverNode Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// List of existing workloads that should be provisioned on the node.
	NodeWorkloads []RoverWorkload `mandatory:"false" json:"nodeWorkloads"`

	// Root password for the rover node.
	SuperUserPassword *string `mandatory:"false" json:"superUserPassword"`

	// Passphrase to unlock the rover node.
	UnlockPassphrase *string `mandatory:"false" json:"unlockPassphrase"`

	// Name of point of contact for this order if customer is picking up.
	PointOfContact *string `mandatory:"false" json:"pointOfContact"`

	// Phone number of point of contact for this order if customer is picking up.
	PointOfContactPhoneNumber *string `mandatory:"false" json:"pointOfContactPhoneNumber"`

	// Preference for device delivery.
	ShippingPreference CreateRoverNodeDetailsShippingPreferenceEnum `mandatory:"false" json:"shippingPreference,omitempty"`

	// Shipping vendor of choice for orace to customer shipping.
	ShippingVendor *string `mandatory:"false" json:"shippingVendor"`

	// Expected date when customer wants to pickup the device if they chose customer pickup.
	ExpectedPickupDate *common.SDKTime `mandatory:"false" json:"expectedPickupDate"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRoverNodeDetails) String() string {
	return common.PointerString(m)
}

// CreateRoverNodeDetailsShippingPreferenceEnum Enum with underlying type: string
type CreateRoverNodeDetailsShippingPreferenceEnum string

// Set of constants representing the allowable values for CreateRoverNodeDetailsShippingPreferenceEnum
const (
	CreateRoverNodeDetailsShippingPreferenceOracleShipped  CreateRoverNodeDetailsShippingPreferenceEnum = "ORACLE_SHIPPED"
	CreateRoverNodeDetailsShippingPreferenceCustomerPickup CreateRoverNodeDetailsShippingPreferenceEnum = "CUSTOMER_PICKUP"
)

var mappingCreateRoverNodeDetailsShippingPreference = map[string]CreateRoverNodeDetailsShippingPreferenceEnum{
	"ORACLE_SHIPPED":  CreateRoverNodeDetailsShippingPreferenceOracleShipped,
	"CUSTOMER_PICKUP": CreateRoverNodeDetailsShippingPreferenceCustomerPickup,
}

// GetCreateRoverNodeDetailsShippingPreferenceEnumValues Enumerates the set of values for CreateRoverNodeDetailsShippingPreferenceEnum
func GetCreateRoverNodeDetailsShippingPreferenceEnumValues() []CreateRoverNodeDetailsShippingPreferenceEnum {
	values := make([]CreateRoverNodeDetailsShippingPreferenceEnum, 0)
	for _, v := range mappingCreateRoverNodeDetailsShippingPreference {
		values = append(values, v)
	}
	return values
}
