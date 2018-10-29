// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace REST API endpoint
//
// Marketplace REST API for loom plugin
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ApplicationDoc The model of related document for the application
type ApplicationDoc struct {

	// The name of the document
	Name *string `mandatory:"false" json:"name"`

	// The description of the document
	Description *string `mandatory:"false" json:"description"`

	// The content type of the document
	ContentType *Item `mandatory:"false" json:"contentType"`

	// The source type of the document
	SourceType *string `mandatory:"false" json:"sourceType"`

	// The source URL of the document
	SourceURL *string `mandatory:"false" json:"sourceURL"`

	// The Mimetype of the document
	MimeType *string `mandatory:"false" json:"mimeType"`

	// The file extension of the document
	FileExtension *string `mandatory:"false" json:"fileExtension"`

	// Defines whether the document is for Oracle users only
	IsForOracleUsersOnly *bool `mandatory:"false" json:"isForOracleUsersOnly"`
}

func (m ApplicationDoc) String() string {
	return common.PointerString(m)
}
