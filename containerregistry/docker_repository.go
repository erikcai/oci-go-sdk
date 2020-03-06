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

// DockerRepository Docker repository
type DockerRepository struct {

	// The OCID of the compartment the repository exists in
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The id of the user or principal that created the resource
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID for the repository
	Id *string `mandatory:"true" json:"id"`

	// Public repository allows unauthenticated access
	IsPublic *bool `mandatory:"true" json:"isPublic"`

	// Total number of images
	ImageCount *int `mandatory:"true" json:"imageCount"`

	// The repository name
	Name *string `mandatory:"true" json:"name"`

	Readme *DockerRepositoryReadme `mandatory:"true" json:"readme"`

	// Total storage in bytes consumed by layers
	LayerSizeInBytes *int64 `mandatory:"true" json:"layerSizeInBytes"`

	// An RFC 3339 timestamp indicating when the repository was created
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// An RFC 3339 timestamp indicating when image was last pushed to the repository
	TimeLastPushed *common.SDKTime `mandatory:"false" json:"timeLastPushed"`
}

func (m DockerRepository) String() string {
	return common.PointerString(m)
}
