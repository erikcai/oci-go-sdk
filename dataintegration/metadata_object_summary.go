// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MetadataObjectSummary A summary type containing information about the object including its key, name and when/who created/updated it
type MetadataObjectSummary struct {

	// The identifying key for the object.
	Key *string `mandatory:"true" json:"key"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The user that created the object.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// The user that created the object.
	CreatedByName *string `mandatory:"false" json:"createdByName"`

	// The user that updated the object.
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// The user that updated the object.
	UpdatedByName *string `mandatory:"false" json:"updatedByName"`

	// The date and time that the object was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time that the object was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The type of the object.
	Type *string `mandatory:"false" json:"type"`

	// The version of the object, to track changes in the object instance
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The owning object key for this object.
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// The full path to identify this object.
	IdentifierPath *string `mandatory:"false" json:"identifierPath"`

	// infoFields
	InfoFields map[string]string `mandatory:"false" json:"infoFields"`

	// registryVersion
	RegistryVersion *int `mandatory:"false" json:"registryVersion"`

	// Labels are keywords or tags that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`
}

func (m MetadataObjectSummary) String() string {
	return common.PointerString(m)
}
