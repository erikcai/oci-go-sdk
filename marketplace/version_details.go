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

// VersionDetails The model of version details
type VersionDetails struct {

	// The version number
	VersionNumber *string `mandatory:"false" json:"versionNumber"`

	// The date when the version was released
	ReleaseDate *common.SDKTime `mandatory:"false" json:"releaseDate"`

	// The desription of the version
	Description *string `mandatory:"false" json:"description"`
}

func (m VersionDetails) String() string {
	return common.PointerString(m)
}
