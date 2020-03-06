// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Registry Extension API
//
// API for managing images and repositories in Oracle Cloud Infrastructure Registry.
//

package containerregistry

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UndeleteManifestDetails Undelete manifest request details
type UndeleteManifestDetails struct {

	// The digest of manifest you want to undelete.
	Digest *string `mandatory:"true" json:"digest"`

	// Optional tag to undelete a manifest with
	Tag *string `mandatory:"false" json:"tag"`
}

func (m UndeleteManifestDetails) String() string {
	return common.PointerString(m)
}
