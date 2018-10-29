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

// Terms Model of terms
type Terms struct {

	// Type of terms
	Type *TermsEnum `mandatory:"false" json:"type"`

	// Unique identifier of the term
	Id *string `mandatory:"false" json:"id"`

	// Content URL of the term
	ContentURL *string `mandatory:"false" json:"contentURL"`

	// Mimetype of the term
	MimeType *string `mandatory:"false" json:"mimeType"`

	// File extension of the term
	FileExtension *string `mandatory:"false" json:"fileExtension"`

	// Click to Accept Term message
	ClickToAcceptText *string `mandatory:"false" json:"clickToAcceptText"`
}

func (m Terms) String() string {
	return common.PointerString(m)
}
