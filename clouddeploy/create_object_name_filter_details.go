// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v32/common"
)

// CreateObjectNameFilterDetails Specifies a filter that compares object names to a set of prefixes or patterns
type CreateObjectNameFilterDetails struct {

	// An array of glob patterns to match the object names to exclude. An empty array is ignored.
	InclusionPatterns []string `mandatory:"false" json:"inclusionPatterns"`

	// An array of glob patterns to match the object names to include. An empty array includes all objects in the bucket.
	ExclusionPatterns []string `mandatory:"false" json:"exclusionPatterns"`

	// An array of object name prefixes. An empty array means to include all objects.
	InclusionPrefixes []string `mandatory:"false" json:"inclusionPrefixes"`
}

func (m CreateObjectNameFilterDetails) String() string {
	return common.PointerString(m)
}
