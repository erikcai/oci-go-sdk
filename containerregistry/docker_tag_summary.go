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

// DockerTagSummary Docker tag search result
type DockerTagSummary struct {

	// The digest of the image referenced by this tag, which can be used for retrieval.
	Digest *string `mandatory:"true" json:"digest"`

	// Repository name
	Repository *string `mandatory:"true" json:"repository"`

	// Tag name
	Tag *string `mandatory:"true" json:"tag"`

	// An RFC 3339 timestamp indicating when the tag was created
	TimeTagged *common.SDKTime `mandatory:"true" json:"timeTagged"`
}

func (m DockerTagSummary) String() string {
	return common.PointerString(m)
}
