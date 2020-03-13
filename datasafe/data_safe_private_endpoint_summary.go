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

// DataSafePrivateEndpointSummary Summary of the data safe PrivateEndpoint.
type DataSafePrivateEndpointSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Data safe private endpoint name, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// VCN Identifier
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Subnet Identifier
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Private Endpoint Identifier
	PrivateEndpointId *string `mandatory:"true" json:"privateEndpointId"`

	// Description of the data safe private endpoint
	Description *string `mandatory:"false" json:"description"`

	// The time the the data safe private endpoint was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the data safe private endpoint
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m DataSafePrivateEndpointSummary) String() string {
	return common.PointerString(m)
}
