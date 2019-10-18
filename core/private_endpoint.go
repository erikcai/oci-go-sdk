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

// PrivateEndpoint A Private Endpoint is a set of private IPs in the customer VCN that act as the access point for the service that are created for.
// Every Private Endpoint has a set of secondary IPs that correspond to the LBs of the service.
// The customer can then communicate to the service by sending traffic to this IP on the registered port.
type PrivateEndpoint struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the
	// Private Endpoint.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Private Endpoint.
	Id *string `mandatory:"true" json:"id"`

	// Endpoint Service OCID that sits behind this Private Endpoint.
	EndpointServiceId *string `mandatory:"true" json:"endpointServiceId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the vcn the Private Endpoint
	// belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the Private Endpoint
	// belongs to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the PE VNIC created in the consumer VCN
	// belongs to.
	PrivateEndpointVnicId *string `mandatory:"true" json:"privateEndpointVnicId"`

	// IP which represents the access point for the associated Endpoint Service.
	PrivateEndpointIp *string `mandatory:"true" json:"privateEndpointIp"`

	// The private endpoint's current state.
	LifecycleState PrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// A user-friendly name. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of this particular Private Endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the Private Endpoint was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Service's 3 label FQDN representing the Endpoint Service. If this value is provided, it will be chosen over the FQDN that the Endpoint Service prescribes.
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// The network security group OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) that this Private Endpoint is associated with.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	ReverseConnectionConfiguration *ReverseConnectionConfiguration `mandatory:"false" json:"reverseConnectionConfiguration"`
}

func (m PrivateEndpoint) String() string {
	return common.PointerString(m)
}

// PrivateEndpointLifecycleStateEnum Enum with underlying type: string
type PrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateEndpointLifecycleStateEnum
const (
	PrivateEndpointLifecycleStateProvisioning PrivateEndpointLifecycleStateEnum = "PROVISIONING"
	PrivateEndpointLifecycleStateAvailable    PrivateEndpointLifecycleStateEnum = "AVAILABLE"
	PrivateEndpointLifecycleStateTerminating  PrivateEndpointLifecycleStateEnum = "TERMINATING"
	PrivateEndpointLifecycleStateTerminated   PrivateEndpointLifecycleStateEnum = "TERMINATED"
)

var mappingPrivateEndpointLifecycleState = map[string]PrivateEndpointLifecycleStateEnum{
	"PROVISIONING": PrivateEndpointLifecycleStateProvisioning,
	"AVAILABLE":    PrivateEndpointLifecycleStateAvailable,
	"TERMINATING":  PrivateEndpointLifecycleStateTerminating,
	"TERMINATED":   PrivateEndpointLifecycleStateTerminated,
}

// GetPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for PrivateEndpointLifecycleStateEnum
func GetPrivateEndpointLifecycleStateEnumValues() []PrivateEndpointLifecycleStateEnum {
	values := make([]PrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingPrivateEndpointLifecycleState {
		values = append(values, v)
	}
	return values
}
