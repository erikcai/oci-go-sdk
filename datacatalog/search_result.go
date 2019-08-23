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

// SearchResult The search result object is the definition of an element that is returned as part of search. It contains basic
// information about the object such as key, name and description. The search result also contains the list of tags
// for each object along with other contextual information like the Data Asset root, folder or entity parents.
type SearchResult struct {

	// Unique key of the object returned as part of the search result.
	Key *string `mandatory:"false" json:"key"`

	// Name of the object
	Name *string `mandatory:"false" json:"name"`

	// Detailed description of the object.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the result object was created, in the format defined by RFC3339.
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Array of the tags associated with this object
	TagSummary []SearchTagSummary `mandatory:"false" json:"tagSummary"`

	// Array of the terms associated with this object
	TermSummary []SearchTermSummary `mandatory:"false" json:"termSummary"`

	// Name of the object type
	TypeName *string `mandatory:"false" json:"typeName"`

	// Data type of the object if the object is an attribute. Null otherwise
	ExternalDataType *string `mandatory:"false" json:"externalDataType"`

	// Unique key of the Data Asset that is the root parent of this object
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// Type name of the Data Asset. For example, Oracle, MySQL or Oracle Object Storage
	DataAssetType *string `mandatory:"false" json:"dataAssetType"`

	// Unique key of the folder object if this object is a sub folder, entity or attribute
	FolderKey *string `mandatory:"false" json:"folderKey"`

	// Unique key of the entity object if this object is an attribute
	Entitykey *string `mandatory:"false" json:"entitykey"`

	// Name of the parent entity object if this object is an attribute
	EntityName *string `mandatory:"false" json:"entityName"`

	// Unique id of the parent Glossary.
	GlossaryKey *string `mandatory:"false" json:"glossaryKey"`

	// Name of the parent glossary if this object is a term
	GlossaryName *string `mandatory:"false" json:"glossaryName"`

	// This terms parent term key. Will be null if the term has no parent term.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// Name of the parent term . Will be null if the term has no parent term.
	ParentTermName *string `mandatory:"false" json:"parentTermName"`
}

func (m SearchResult) String() string {
	return common.PointerString(m)
}
