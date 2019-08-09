// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AttributeSummary Summary of an Entity Attribute.
type AttributeSummary struct {

	// Unique Attribute key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The unique key of the parent Entity.
	EntityKey *string `mandatory:"false" json:"entityKey"`

	// Unique external key of this attribute in the external source system
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// URI to the Attribute instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// State of the Attribute.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the Attribute was created, in the format defined by RFC3339.
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Data type of the attribute as defined in the external source system
	ExternalDataType *string `mandatory:"false" json:"externalDataType"`
}

func (m AttributeSummary) String() string {
	return common.PointerString(m)
}
