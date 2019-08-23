// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PrivateAccessGateway A Private Access Gateway is a gateway on the service VCN that is required by the service to recieve and send the traffic from Private Endpoints. Once created, the service needs to update the route tables to send all PE traffic via this gateway.
type PrivateAccessGateway struct {

	// The Private Access Gateway's Oracle ID (OCID) (/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The OCID  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)  of the compartment to contain the Private Access Gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)  of the VCN to contain the Private Access Gateway.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The date and time the Private Access Gateway was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The private access gateway's current state.
	LifecycleState PrivateAccessGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m PrivateAccessGateway) String() string {
	return common.PointerString(m)
}

// PrivateAccessGatewayLifecycleStateEnum Enum with underlying type: string
type PrivateAccessGatewayLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateAccessGatewayLifecycleStateEnum
const (
	PrivateAccessGatewayLifecycleStateProvisioning PrivateAccessGatewayLifecycleStateEnum = "PROVISIONING"
	PrivateAccessGatewayLifecycleStateAvailable    PrivateAccessGatewayLifecycleStateEnum = "AVAILABLE"
	PrivateAccessGatewayLifecycleStateTerminating  PrivateAccessGatewayLifecycleStateEnum = "TERMINATING"
	PrivateAccessGatewayLifecycleStateTerminated   PrivateAccessGatewayLifecycleStateEnum = "TERMINATED"
)

var mappingPrivateAccessGatewayLifecycleState = map[string]PrivateAccessGatewayLifecycleStateEnum{
	"PROVISIONING": PrivateAccessGatewayLifecycleStateProvisioning,
	"AVAILABLE":    PrivateAccessGatewayLifecycleStateAvailable,
	"TERMINATING":  PrivateAccessGatewayLifecycleStateTerminating,
	"TERMINATED":   PrivateAccessGatewayLifecycleStateTerminated,
}

// GetPrivateAccessGatewayLifecycleStateEnumValues Enumerates the set of values for PrivateAccessGatewayLifecycleStateEnum
func GetPrivateAccessGatewayLifecycleStateEnumValues() []PrivateAccessGatewayLifecycleStateEnum {
	values := make([]PrivateAccessGatewayLifecycleStateEnum, 0)
	for _, v := range mappingPrivateAccessGatewayLifecycleState {
		values = append(values, v)
	}
	return values
}
