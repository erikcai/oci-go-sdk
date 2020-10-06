// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v26/common"
)

// BulkEditResource The representation of BulkEditResource
type BulkEditResource struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Type of resource. See BulkEditResourceTypes
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Additional information that helps to identify the resource for bulk edit of Tags.
	// UPDATE APIs for most resource types only require the resource identifier(ocid).
	// But additional metadata is required for some resource types.
	// This information is provided in the resource's public API document. It is also
	// available through the ListTaggingSupportedResourceTypes API.
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m BulkEditResource) String() string {
	return common.PointerString(m)
}
