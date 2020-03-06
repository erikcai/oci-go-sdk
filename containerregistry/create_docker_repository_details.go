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

// CreateDockerRepositoryDetails Create repository request details
type CreateDockerRepositoryDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment into which the resource should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Repository name
	Name *string `mandatory:"true" json:"name"`

	// Public repository allows unauthenticated access
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	Readme *DockerRepositoryReadme `mandatory:"false" json:"readme"`
}

func (m CreateDockerRepositoryDetails) String() string {
	return common.PointerString(m)
}
