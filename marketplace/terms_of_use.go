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

// TermsOfUse Terms of Use Model
type TermsOfUse struct {

	// Reference links
	Links []Link `mandatory:"false" json:"links"`

	// Unique Identifier of the Terms Of Use
	EntityId *string `mandatory:"false" json:"entityId"`

	// Types of Terms Of Use
	Type *TermsEnum `mandatory:"false" json:"type"`

	// Unique Identifier of the terms of use
	Id *string `mandatory:"false" json:"id"`

	// The content URL of the terms of use
	ContentURL *string `mandatory:"false" json:"contentURL"`

	// The mimeTypes of the terms of use
	MimeType *string `mandatory:"false" json:"mimeType"`

	// The file extension fo the terms of use
	FileExtension *string `mandatory:"false" json:"fileExtension"`

	// Click to Accept Term message
	ClickToAcceptText *string `mandatory:"false" json:"clickToAcceptText"`
}

func (m TermsOfUse) String() string {
	return common.PointerString(m)
}
