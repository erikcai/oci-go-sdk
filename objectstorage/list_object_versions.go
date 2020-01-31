// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ListObjectVersions To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type ListObjectVersions struct {

	// An array of object version summaries.
	ObjectVersions []ObjectVersionSummary `mandatory:"true" json:"objectVersions"`

	// Prefixes that are common to the results returned by the request if the request specified a delimiter.
	Prefixes []string `mandatory:"false" json:"prefixes"`

	// The token to use as a query parameter to obtain the next page of
	// a truncated ListObjects(Versions) response. It will be present only if 'versions' were requested.
	ContinuationToken *string `mandatory:"false" json:"continuationToken"`
}

func (m ListObjectVersions) String() string {
	return common.PointerString(m)
}