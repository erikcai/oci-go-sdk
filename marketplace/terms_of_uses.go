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

// TermsOfUses List of terms of use
type TermsOfUses struct {

	// The instruction to display before users start to install the app
	PreInstallInstructions *string `mandatory:"false" json:"preInstallInstructions"`

	// List of terms of use
	TermsOfUses []TermsOfUse `mandatory:"false" json:"termsOfUses"`

	// Reference links
	Links []Link `mandatory:"false" json:"links"`
}

func (m TermsOfUses) String() string {
	return common.PointerString(m)
}
