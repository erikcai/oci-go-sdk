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

// SupportLink Model of support links
type SupportLink struct {

	// Name of support link
	Name *string `mandatory:"false" json:"name"`

	// URL of the support link
	Url *string `mandatory:"false" json:"url"`
}

func (m SupportLink) String() string {
	return common.PointerString(m)
}
