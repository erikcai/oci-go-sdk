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

// UploadData Model of the upload data for images, icon, etc
type UploadData struct {

	// Name of the upload data
	Name *string `mandatory:"false" json:"name"`

	// Content URL of the upload data
	ContentURL *string `mandatory:"false" json:"contentURL"`

	// Mimetype of the upload data
	MimeType *string `mandatory:"false" json:"mimeType"`

	// File extension of the upload data
	FileExtension *string `mandatory:"false" json:"fileExtension"`
}

func (m UploadData) String() string {
	return common.PointerString(m)
}
