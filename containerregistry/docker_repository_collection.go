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

// DockerRepositoryCollection Results of a docker repository search. Contains DockerRepositorySummary items along with metadata for the selected repositories.
type DockerRepositoryCollection struct {

	// collection of docker repositories
	Items []DockerRepositorySummary `mandatory:"true" json:"items"`

	// Total storage in bytes consumed by layers
	LayerSizeInBytes *int64 `mandatory:"true" json:"layerSizeInBytes"`

	// Total number of images
	ImageCount *int `mandatory:"true" json:"imageCount"`

	// Total number of layers
	LayerCount *int `mandatory:"true" json:"layerCount"`
}

func (m DockerRepositoryCollection) String() string {
	return common.PointerString(m)
}
