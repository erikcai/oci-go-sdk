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

// CreateRoverClusterDetails The information about new RoverCluster.
type CreateRoverClusterDetails struct {

	// RoverCluster Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Number of nodes desired in the cluster, between 5 and 15.
	ClusterSize *string `mandatory:"true" json:"clusterSize"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// List of existing workloads that should be provisioned on the nodes.
	ClusterWorkloads []RoverWorkload `mandatory:"false" json:"clusterWorkloads"`

	// Root password for the rover cluster.
	SuperUserPassword *string `mandatory:"false" json:"superUserPassword"`

	// Password to unlock the rover cluster.
	UnlockPassphrase *string `mandatory:"false" json:"unlockPassphrase"`

	// Name of point of contact for this order if customer is picking up.
	PointOfContact *string `mandatory:"false" json:"pointOfContact"`

	// Phone number of point of contact for this order if customer is picking up.
	PointOfContactPhoneNumber *string `mandatory:"false" json:"pointOfContactPhoneNumber"`

	// Preference for device delivery.
	ShippingPreference CreateRoverClusterDetailsShippingPreferenceEnum `mandatory:"false" json:"shippingPreference,omitempty"`

	// Shipping vendor of choice for orace to customer shipping.
	ShippingVendor *string `mandatory:"false" json:"shippingVendor"`

	// Expected date when customer wants to pickup the cluster if they chose customer pickup.
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

func (m CreateRoverClusterDetails) String() string {
	return common.PointerString(m)
}

// CreateRoverClusterDetailsShippingPreferenceEnum Enum with underlying type: string
type CreateRoverClusterDetailsShippingPreferenceEnum string

// Set of constants representing the allowable values for CreateRoverClusterDetailsShippingPreferenceEnum
const (
	CreateRoverClusterDetailsShippingPreferenceOracleShipped  CreateRoverClusterDetailsShippingPreferenceEnum = "ORACLE_SHIPPED"
	CreateRoverClusterDetailsShippingPreferenceCustomerPickup CreateRoverClusterDetailsShippingPreferenceEnum = "CUSTOMER_PICKUP"
)

var mappingCreateRoverClusterDetailsShippingPreference = map[string]CreateRoverClusterDetailsShippingPreferenceEnum{
	"ORACLE_SHIPPED":  CreateRoverClusterDetailsShippingPreferenceOracleShipped,
	"CUSTOMER_PICKUP": CreateRoverClusterDetailsShippingPreferenceCustomerPickup,
}

// GetCreateRoverClusterDetailsShippingPreferenceEnumValues Enumerates the set of values for CreateRoverClusterDetailsShippingPreferenceEnum
func GetCreateRoverClusterDetailsShippingPreferenceEnumValues() []CreateRoverClusterDetailsShippingPreferenceEnum {
	values := make([]CreateRoverClusterDetailsShippingPreferenceEnum, 0)
	for _, v := range mappingCreateRoverClusterDetailsShippingPreference {
		values = append(values, v)
	}
	return values
}
