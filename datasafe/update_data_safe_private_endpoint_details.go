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

// UpdateDataSafePrivateEndpointDetails The information to be updated for an existing private endpoint.
type UpdateDataSafePrivateEndpointDetails struct {

	// data safe private endpoint Name
	DisplayName *string `mandatory:"true" json:"displayName"`

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

	// The current state of the data safe private endpoint
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m UpdateDataSafePrivateEndpointDetails) String() string {
	return common.PointerString(m)
}
