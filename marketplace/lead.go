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

// Lead The model of the lead
type Lead struct {

	// Unique identifier of the lead
	Id *string `mandatory:"false" json:"id"`

	// Status message of the lead
	StatusMessage *string `mandatory:"false" json:"statusMessage"`
}

func (m Lead) String() string {
	return common.PointerString(m)
}
