// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Query Service
//
// Query for resources across your cloud infrastructure
//

package resourcequery

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResourceSummaryCollection Found resources summaries.
type ResourceSummaryCollection struct {

	// Resources that exist within the user's cloud infrastructure.
	Items []ResourceSummary `mandatory:"false" json:"items"`
}

func (m ResourceSummaryCollection) String() string {
	return common.PointerString(m)
}
