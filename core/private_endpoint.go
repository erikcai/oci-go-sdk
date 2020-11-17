// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
	"github.com/oracle/oci-go-sdk/v29/common"
)

// PrivateEndpoint A *private endpoint* (PE) is a way for a service provider to give a customer a private access point for
// the service within the customer's VCN. The PE consists of a VNIC with a private IP
// in the customer's VCN. The PE is associated with an
// EndpointService
// that the service provider has previously registered. The customer accesses the service
// by sending traffic to the PE's IP address on the registered port. That traffic is then
// sent to the PrivateAccessGateway on the service VCN.
// The gateway then sends the traffic to the PE's associated `EndpointService` for
// processing.
// **Regarding DNS for the private endpoint:** You may optionally have the private endpoint
// service assign a fully qualified domain name (FQDN) to the private endpoint. If you do, the
// private endpoint service creates the related DNS zone and record in the customer's VCN. This
// enables the customer to use the FQDN instead of the PE's private IP address to access the
// service. Here are details about how the private endpoint service determines the value to use
// for the PE's FQDN:
//   * Both the EndpointService object and the
//     CreatePrivateEndpointDetails
//     object have an `endpointFqdn` attribute.
//   * If you don't specify an FQDN for `CreatePrivateEndpointDetails` during PE creation, the
//      endpoint service's `endpointFqdn` is used for the PE's `endpointFqdn`.
//   * If you specify an FQDN for `CreatePrivateEndpointDetails` during PE creation, that value is used.
//     It always takes precedence over any value set in the `EndpointService` object.
//   * If the `EndpointService` object does not have an FQDN value set, and you don't provide a value
//     in `CreatePrivateEndpointDetails` during creation, the PE does not get an FQDN.
//   * **Special scenario:**  If the endpoint service allows multiple PE's to be created per customer VCN
//     (see the `areMultiplePrivateEndpointsPerVcnAllowed` attribute in the `EndpointService`),
//     the `EndpointService` is prohibited from also having an `endpointFqdn` value. This restriction ensures
//     that each FQDN in the customer's VCN resolves to a single PE. Therefore, for this particular
//     scenario, you must assign each PE a unique FQDN within the scope of the customer's VCN.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you
// supply string values using the API.
type PrivateEndpoint struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the
	// private endpoint.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the endpoint service that is associated
	// with the private endpoint.
	EndpointServiceId *string `mandatory:"true" json:"endpointServiceId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer VCN that the private
	// endpoint belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet that the private endpoint
	// belongs to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint's VNIC, which
	// resides in the customer's VCN.
	PrivateEndpointVnicId *string `mandatory:"true" json:"privateEndpointVnicId"`

	// The private IP address (in the customer's VCN) that represents the access point for the
	// associated endpoint service.
	PrivateEndpointIp *string `mandatory:"true" json:"privateEndpointIp"`

	// The private endpoint's current lifecycle state.
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
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of this private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the private endpoint was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The three-label FQDN to use for the private endpoint. The customer VCN's DNS records are
	// updated with this FQDN.
	// If you provide a value for this attribute, it overrides the `endpointFqdn` in the associated
	// EndpointService. For more information, see the discussion
	// of DNS and FQDNs in PrivateEndpoint.
	// You can change the PE's FQDN (see
	// UpdatePrivateEndpointDetails).
	// Example: `xyz.oraclecloud.com`
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// A list of the OCIDs of the network security groups that the private endpoint's VNIC belongs to.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
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
	PrivateEndpointLifecycleStateUpdating     PrivateEndpointLifecycleStateEnum = "UPDATING"
	PrivateEndpointLifecycleStateFailed       PrivateEndpointLifecycleStateEnum = "FAILED"
)

var mappingPrivateEndpointLifecycleState = map[string]PrivateEndpointLifecycleStateEnum{
	"PROVISIONING": PrivateEndpointLifecycleStateProvisioning,
	"AVAILABLE":    PrivateEndpointLifecycleStateAvailable,
	"TERMINATING":  PrivateEndpointLifecycleStateTerminating,
	"TERMINATED":   PrivateEndpointLifecycleStateTerminated,
	"UPDATING":     PrivateEndpointLifecycleStateUpdating,
	"FAILED":       PrivateEndpointLifecycleStateFailed,
}

// GetPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for PrivateEndpointLifecycleStateEnum
func GetPrivateEndpointLifecycleStateEnumValues() []PrivateEndpointLifecycleStateEnum {
	values := make([]PrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingPrivateEndpointLifecycleState {
		values = append(values, v)
	}
	return values
}
