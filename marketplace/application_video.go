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

// ApplicationVideo The model of videos of the application
type ApplicationVideo struct {

	// Title of the video
	Title *string `mandatory:"false" json:"title"`

	// The URL of the video
	Url *string `mandatory:"false" json:"url"`

	// The mimeType of the application
	MimeType *string `mandatory:"false" json:"mimeType"`

	// The file extension of the application
	FileExtension *string `mandatory:"false" json:"fileExtension"`
}

func (m ApplicationVideo) String() string {
	return common.PointerString(m)
}
