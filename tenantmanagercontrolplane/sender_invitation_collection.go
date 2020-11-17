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

// SenderInvitationCollection Result of a query request for a list of sender invitations. Contains SenderInvitationSummary items.
type SenderInvitationCollection struct {

	// Array containing SenderInvitationSummary items.
	Items []SenderInvitationSummary `mandatory:"true" json:"items"`
}

func (m SenderInvitationCollection) String() string {
	return common.PointerString(m)
}
