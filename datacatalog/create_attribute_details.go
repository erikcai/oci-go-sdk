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

// CreateAttributeDetails Properties used in Attribute create operations.
type CreateAttributeDetails struct {

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Data type of the attribute as defined in the external system
	ExternalDataType *string `mandatory:"true" json:"externalDataType"`

	// Last modified timestamp of this object in the external system
	TimeExternal *common.SDKTime `mandatory:"true" json:"timeExternal"`

	// Property that identifies if this attribute can be used as a watermark to extract incremental data
	IsIncrementalData *bool `mandatory:"false" json:"isIncrementalData"`

	// Max allowed length of the attribute value
	Length *int64 `mandatory:"false" json:"length"`

	// Position of the attribute in the record definition
	Position *int `mandatory:"false" json:"position"`

	// Precision of the attribute value usually applies to float data type
	Precision *int `mandatory:"false" json:"precision"`

	// Scale of the attribute value usually applies to float data type
	Scale *int `mandatory:"false" json:"scale"`

	// A map of maps which contains the properties which are specific to the Attribute type. Each Attribute type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// Attributes have required properties within the "default" category. To determine the set of required and
	// optional properties for an Attribute type, a query can be done on '/types?type=attribute' which returns a
	// collection of all Attribute types. The appropriate Attribute type, which will include definitions of all
	// of it's properties, can be identified from this collection.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m CreateAttributeDetails) String() string {
	return common.PointerString(m)
}
