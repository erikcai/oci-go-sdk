// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastion

import (
	"github.com/oracle/oci-go-sdk/v27/common"
)

// Bastion Description of Bastion.
type Bastion struct {

	// The unique OCID of the bastion which is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// bastion Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// the target vcn OCID the bastion connects to.
	TargetVcnId *string `mandatory:"true" json:"targetVcnId"`

	// the target subnet OCID the bastion connects to.
	TargetSubnetId *string `mandatory:"true" json:"targetSubnetId"`

	// the ip ranges that the bastion has access to.
	ClientCidrBlockAllowList []string `mandatory:"true" json:"clientCidrBlockAllowList"`

	// max TTL of the sessions on the bastion.
	MaxSessionTtlInSeconds *int `mandatory:"true" json:"maxSessionTtlInSeconds"`

	// Maximum sessions allowed to create on a bastion host.
	MaxSessionsAllowed *int `mandatory:"true" json:"maxSessionsAllowed"`

	// The time the bastion was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the bastion.
	LifecycleState BastionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// availability domain of the bastion.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Private ip address of the created private endpoint
	PrivateEndpointIpAddress *string `mandatory:"false" json:"privateEndpointIpAddress"`

	// The time the bastion was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Bastion) String() string {
	return common.PointerString(m)
}
