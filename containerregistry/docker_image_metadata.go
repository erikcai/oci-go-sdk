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

// DockerImageMetadata Docker image metadata
type DockerImageMetadata struct {

	// The id of the user or principal that created the resource
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// A digest
	Digest *string `mandatory:"true" json:"digest"`

	// Total number of pulls
	PullCount *int64 `mandatory:"true" json:"pullCount"`

	// Layers of which image is composed, orded by the layer manifest digest
	LayerDetails []DockerImageLayer `mandatory:"true" json:"layerDetails"`

	// Layer size in bytes
	LayerSizeInBytes *int64 `mandatory:"true" json:"layerSizeInBytes"`

	// Manifest size in bytes
	ManifestSizeInBytes *int `mandatory:"true" json:"manifestSizeInBytes"`

	// Docker tags associated with this image
	Tags []ImageTagMetadata `mandatory:"true" json:"tags"`

	// An RFC 3339 timestamp indicating when the image was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// An RFC 3339 timestamp indicating when the image was last pulled
	TimeLastPulled *common.SDKTime `mandatory:"false" json:"timeLastPulled"`
}

func (m DockerImageMetadata) String() string {
	return common.PointerString(m)
}
