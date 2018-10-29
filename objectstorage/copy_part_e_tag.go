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

// CopyPartETag The representation of CopyPartETag
type CopyPartETag struct {

	// The ETag of the new part.
	ETag *string `mandatory:"true" json:"ETag"`
}

func (m CopyPartETag) String() string {
	return common.PointerString(m)
}
