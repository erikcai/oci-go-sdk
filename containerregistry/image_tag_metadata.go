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

// ImageTagMetadata Metadata of an image tag
type ImageTagMetadata struct {

	// The image tag
	Tag *string `mandatory:"true" json:"tag"`

	// The id of the user or principal the image was tagged by
	TaggedBy *string `mandatory:"true" json:"taggedBy"`

	// The timestamp when the image was tagged
	TimeTagged *common.SDKTime `mandatory:"true" json:"timeTagged"`
}

func (m ImageTagMetadata) String() string {
	return common.PointerString(m)
}
