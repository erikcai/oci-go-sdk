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

// SearchContext Contains search context like highlights for found resource.
type SearchContext struct {

	// Describes what matched the query per each field. The list of string represents fragments of field value which has query matching. Highlighted values are wrapped with <hl>..</hl> tags. All values are HTML encoded (except <hl> tags). Works only for Free Text Search or MATCHING clause in structured query.
	Highlights map[string][]string `mandatory:"false" json:"highlights"`
}

func (m SearchContext) String() string {
	return common.PointerString(m)
}
