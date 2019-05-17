// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// StorageGateway API
//
// API for interfacing with StorageGateway
//

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateStorageGatewayDetails Details to set when creating the storage gateway.
type CreateStorageGatewayDetails struct {

	// The compartment OCID containing the storage gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A non-unique name you assign to the storage gateway during creation.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The non-unique, changeable description you assign to the storage gateway during creation.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information,
	// see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateStorageGatewayDetails) String() string {
	return common.PointerString(m)
}
