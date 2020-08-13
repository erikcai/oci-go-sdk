// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TestOutboundConnectorDetails Details for testing outbound connector.
type TestOutboundConnectorDetails struct {

	// OCID of outbound connector which needs to be tested.
	OutboundConnectorId *string `mandatory:"true" json:"outboundConnectorId"`
}

func (m TestOutboundConnectorDetails) String() string {
	return common.PointerString(m)
}
