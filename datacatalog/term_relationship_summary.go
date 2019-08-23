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

// TermRelationshipSummary Summary of a Term Relationship. Business term relationship between two terms in a business glossary.
type TermRelationshipSummary struct {

	// Unique Term Relationship key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the Term.
	Description *string `mandatory:"false" json:"description"`

	// Unique id of the related term.
	RelatedTermKey *string `mandatory:"false" json:"relatedTermKey"`

	// URI to the Term Relationship instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// This relationships parent term key.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// The date and time the Term Relationship was created, in the format defined by RFC3339.
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// State of the Term Relationship.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m TermRelationshipSummary) String() string {
	return common.PointerString(m)
}
