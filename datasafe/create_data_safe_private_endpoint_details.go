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

// CreateDataSafePrivateEndpointDetails The information about new data safe private endpoint.
type CreateDataSafePrivateEndpointDetails struct {

	// data safe private endpoint name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// data safe private endpoint Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// VCN Identifier
	VcnId *string `mandatory:"true" json:"vcnId"`

	// data safe private endpoint subnet Identifier
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The private IP address to assign to this private endpoint's Reverse Connection source IP
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// data safe private endpoint Description
	Description *string `mandatory:"false" json:"description"`

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

func (m CreateDataSafePrivateEndpointDetails) String() string {
	return common.PointerString(m)
}
