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

// UpdateIpv6Details Attributes availble for updating IPv6. Users can move the IPv6 by specifying the target VNIC ID.
// Internet access can also be enabled/disabled via isInternetAccessAllowed flag.
type UpdateIpv6Details struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Whether IPv6 is usable for intenet communication. Internet access via IPv6 will not be allowed for
	// private subnet the same way as IPv4. Internet access will be enabled by default for a public subnet.
	// If VCN has IPv6 enabled with a custom IPv6 CIDR, a different public IPv6 address will be assigned
	// for a particular IPv6.
	IsInternetAccessAllowed *bool `mandatory:"false" json:"isInternetAccessAllowed"`

	// The OCID of the VNIC to reassign the IPv6 to. The VNIC must
	// be in the same subnet as the current VNIC.
	VnicId *string `mandatory:"false" json:"vnicId"`
}

func (m UpdateIpv6Details) String() string {
	return common.PointerString(m)
}
