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

// DockerTagCollection Results of a docker tag search
type DockerTagCollection struct {

	// Estimated number of remaining search results
	RemainingTagsCount *int `mandatory:"true" json:"remainingTagsCount"`

	// Page of matching docker tags
	Items []DockerTagSummary `mandatory:"true" json:"items"`
}

func (m DockerTagCollection) String() string {
	return common.PointerString(m)
}
