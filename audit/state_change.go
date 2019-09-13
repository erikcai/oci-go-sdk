// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. Use this API for compliance monitoring in your tenancy.
// For more information, see Overview of Audit (https://docs.cloud.oracle.com/iaas/Content/Audit/Concepts/auditoverview.htm).
// **Tip**: This API is good for queries, but not bulk-export operations.
//

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
)

// StateChange The representation of StateChange
type StateChange struct {
	Previous map[string]interface{} `mandatory:"false" json:"previous"`

	Current map[string]interface{} `mandatory:"false" json:"current"`
}

func (m StateChange) String() string {
	return common.PointerString(m)
}
