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

// SupportInfo Model of the support information
type SupportInfo struct {

	// Supporter contact
	Contacts []SupportContact `mandatory:"false" json:"contacts"`

	// Reference links
	Links []SupportLink `mandatory:"false" json:"links"`
}

func (m SupportInfo) String() string {
	return common.PointerString(m)
}
