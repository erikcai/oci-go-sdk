// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// BulkEditTagsResourceType The representation of BulkEditTagsResourceType
type BulkEditTagsResourceType struct {

	// The unique name of the resource type.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Metadata keys required to identify the resource.
	// E.g. for bucket, metadataKeys will be ["namespaceName", "bucketName"].
	// This informatino will match the public API document:
	// https://docs.cloud.oracle.com/en-us/iaas/api/#/en/objectstorage/20160918/Bucket/DeleteBucket
	// https://docs.cloud.oracle.com/en-us/iaas/api/#/en/objectstorage/20160918/Bucket/UpdateBucket
	MetadataKeys []string `mandatory:"false" json:"metadataKeys"`
}

func (m BulkEditTagsResourceType) String() string {
	return common.PointerString(m)
}
