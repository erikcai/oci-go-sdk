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

// DockerRepositoryReadme Docker repository readme
type DockerRepositoryReadme struct {

	// Readme content. Avoid entering confidential information.
	Content *string `mandatory:"true" json:"content"`

	// Readme format. Supported formats are text/plain and text/markdown.
	Format DockerRepositoryReadmeFormatEnum `mandatory:"true" json:"format"`
}

func (m DockerRepositoryReadme) String() string {
	return common.PointerString(m)
}

// DockerRepositoryReadmeFormatEnum Enum with underlying type: string
type DockerRepositoryReadmeFormatEnum string

// Set of constants representing the allowable values for DockerRepositoryReadmeFormatEnum
const (
	DockerRepositoryReadmeFormatPlain    DockerRepositoryReadmeFormatEnum = "TEXT_PLAIN"
	DockerRepositoryReadmeFormatMarkdown DockerRepositoryReadmeFormatEnum = "TEXT_MARKDOWN"
)

var mappingDockerRepositoryReadmeFormat = map[string]DockerRepositoryReadmeFormatEnum{
	"TEXT_PLAIN":    DockerRepositoryReadmeFormatPlain,
	"TEXT_MARKDOWN": DockerRepositoryReadmeFormatMarkdown,
}

// GetDockerRepositoryReadmeFormatEnumValues Enumerates the set of values for DockerRepositoryReadmeFormatEnum
func GetDockerRepositoryReadmeFormatEnumValues() []DockerRepositoryReadmeFormatEnum {
	values := make([]DockerRepositoryReadmeFormatEnum, 0)
	for _, v := range mappingDockerRepositoryReadmeFormat {
		values = append(values, v)
	}
	return values
}
