// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastion

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// CreateBastionDetails The information about new bastion.
type CreateBastionDetails struct {

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// the target subnet OCID the bastion connects to.
	TargetSubnetId *string `mandatory:"true" json:"targetSubnetId"`

	// the ip ranges that the bastion has access to.
	ClientCidrBlockAllowList []string `mandatory:"true" json:"clientCidrBlockAllowList"`

	// Bastion Identifier, not unique
	DisplayName *string `mandatory:"false" json:"displayName"`

	// availability domain of the bastion.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// max TTL of the sessions on the bastion.
	MaxSessionTtlInSeconds *int `mandatory:"false" json:"maxSessionTtlInSeconds"`

	// Maximum sessions allowed to create on a bastion host.
	MaxSessionsAllowed *int `mandatory:"false" json:"maxSessionsAllowed"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBastionDetails) String() string {
	return common.PointerString(m)
}
