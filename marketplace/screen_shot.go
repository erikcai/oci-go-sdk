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

// ScreenShot Screenshot model
type ScreenShot struct {

	// Name of the screenshot
	Name *string `mandatory:"false" json:"name"`

	// Description of the screenshot
	Description *string `mandatory:"false" json:"description"`

	// Content URL of the screenshot
	ContentURL *string `mandatory:"false" json:"contentURL"`

	// MimeType of the screenshot
	MimeType *string `mandatory:"false" json:"mimeType"`

	// File extension of the screenshot
	FileExtension *string `mandatory:"false" json:"fileExtension"`
}

func (m ScreenShot) String() string {
	return common.PointerString(m)
}
