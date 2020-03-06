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

// DockerImageLayer Docker image layer
type DockerImageLayer struct {

	// The sha of the image layer
	Digest *string `mandatory:"true" json:"digest"`

	// Layer size in bytes
	SizeInBytes *int64 `mandatory:"true" json:"sizeInBytes"`

	// An RFC 3339 timestamp indicating when the layer was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m DockerImageLayer) String() string {
	return common.PointerString(m)
}
