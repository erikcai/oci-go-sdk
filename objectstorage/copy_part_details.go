// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CopyPartDetails To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
type CopyPartDetails struct {

	// The Object Storage namespace of the object that will be copied from.
	SourceNamespace *string `mandatory:"true" json:"sourceNamespace"`

	// The bucket of the object that will be copied from.
	SourceBucket *string `mandatory:"true" json:"sourceBucket"`

	// The name of the object that will be copied from.
	SourceObject *string `mandatory:"true" json:"sourceObject"`

	// Optional byte range to copy. Format is similar to standard range header as described in
	// RFC 7233 (https://tools.ietf.org/rfc/rfc7233), section 2.1. Note that only a single range of bytes
	// is supported.
	Range *string `mandatory:"false" json:"range"`
}

func (m CopyPartDetails) String() string {
	return common.PointerString(m)
}
