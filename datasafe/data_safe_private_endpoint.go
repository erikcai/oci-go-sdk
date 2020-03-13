// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Safe APIs
//
// APIs for using Data Safe
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DataSafePrivateEndpoint Object of DataSafePrivateEndpoint
type DataSafePrivateEndpoint struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Data safe Private Endpoints display name, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// VCN Identifier
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Subnet Identifier
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Private Endpoint Identifier
	PrivateEndpointId *string `mandatory:"true" json:"privateEndpointId"`

	// The private IP address to assign to this private endpoint's reverse connection configuration.
	// If you provide a value, it must be an available IP address in the customer's subnet.
	// If it's not available, an error is returned.
	// If you do not provide a value, an available IP address in the subnet is automatically chosen.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// The three-label FQDN to use for the private endpoint. The customer VCN's DNS records are updated with this FQDN.
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// Description of the data safe private endpoint
	Description *string `mandatory:"false" json:"description"`

	// The time the the data safe PE was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the data safe private endpoint
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A list of the OCIDs of the network security groups that the private endpoint's VNIC belongs to.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DataSafePrivateEndpoint) String() string {
	return common.PointerString(m)
}
