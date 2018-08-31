// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// The Object and Archive Storage APIs for managing buckets and objects.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PostObjectMetadataResponse Updated object info after user-metadata Update.
type PostObjectMetadataResponse struct {

	// The new entity tag for the object.
	ETag *string `mandatory:"true" json:"ETag"`

	// The time the object was modified, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.29.
	TimeModified *common.SDKTime `mandatory:"true" json:"timeModified"`
}

func (m PostObjectMetadataResponse) String() string {
	return common.PointerString(m)
}
