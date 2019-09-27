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

// GlossaryTreeElement Glossary Tree Element with child terms
type GlossaryTreeElement struct {

	// Unique Term key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the Term.
	Description *string `mandatory:"false" json:"description"`

	// Unique id of the parent Glossary.
	GlossaryKey *string `mandatory:"false" json:"glossaryKey"`

	// URI to the Term instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// This terms parent term key. Will be null if the term has no parent term.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// Indicates whether a term may contain child terms.
	IsAllowedToHaveChildTerms *bool `mandatory:"false" json:"isAllowedToHaveChildTerms"`

	// Absolute path of the term.
	Path *string `mandatory:"false" json:"path"`

	// The date and time the Term was created, in the format defined by RFC3339.
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The number of objects tagged with this term
	AssociatedObjectCount *int `mandatory:"false" json:"associatedObjectCount"`

	// State of the Term.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An array of child terms
	ChildTerms []GlossaryTreeElement `mandatory:"false" json:"childTerms"`
}

func (m GlossaryTreeElement) String() string {
	return common.PointerString(m)
}
