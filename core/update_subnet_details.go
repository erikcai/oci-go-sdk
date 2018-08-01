// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateSubnetDetails The representation of UpdateSubnetDetails
type UpdateSubnetDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID of the set of DHCP options the subnet will use.
	DhcpOptionsId *string `mandatory:"false" json:"dhcpOptionsId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the route table the subnet will use.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// OCIDs for the security lists to associate with the subnet. This
	// replaces the current set of associated security lists. Remember that
	// security lists are associated at the subnet level, but the rules are
	// applied to the individual VNICs in the subnet.
	SecurityListIds []string `mandatory:"false" json:"securityListIds"`
}

func (m UpdateSubnetDetails) String() string {
	return common.PointerString(m)
}
