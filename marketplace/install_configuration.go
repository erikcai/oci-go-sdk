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

// InstallConfiguration The model of install configuration
type InstallConfiguration struct {

	// The user properties
	UserProperties []ActionProperty `mandatory:"false" json:"userProperties"`

	// Reference links
	Links []Link `mandatory:"false" json:"links"`
}

func (m InstallConfiguration) String() string {
	return common.PointerString(m)
}
