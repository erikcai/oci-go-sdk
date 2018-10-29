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

// Link The model of links
type Link struct {

	// Anchor tag
	Href *string `mandatory:"false" json:"href"`
}

func (m Link) String() string {
	return common.PointerString(m)
}
