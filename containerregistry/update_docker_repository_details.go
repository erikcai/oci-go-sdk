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

// UpdateDockerRepositoryDetails Update repository request details
type UpdateDockerRepositoryDetails struct {

	// Public repository allows unauthenticated access
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	Readme *DockerRepositoryReadme `mandatory:"false" json:"readme"`
}

func (m UpdateDockerRepositoryDetails) String() string {
	return common.PointerString(m)
}
