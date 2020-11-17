// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// TenantManager API
//
// A description of the TenantManager API
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// LinkCollection Result of a query request for a list of links. Contains Link items.
type LinkCollection struct {

	// Array containing LinkSummary items.
	Items []LinkSummary `mandatory:"true" json:"items"`
}

func (m LinkCollection) String() string {
	return common.PointerString(m)
}
