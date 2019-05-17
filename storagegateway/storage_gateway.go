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

// StorageGateway Detailed view of the storage gateway.
type StorageGateway struct {

	// The storage gateway's Oracle Cloud ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the storage gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The non-unique name assigned to the storage gateway.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Date and time the storage gateway was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The storage gateway instance's state. After creating a resource, make sure its state
	// changes to ACTIVE before using it.
	LifecycleState LifecycleState `mandatory:"true" json:"lifecycleState"`

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

func (m StorageGateway) String() string {
	return common.PointerString(m)
}
