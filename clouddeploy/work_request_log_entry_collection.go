// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// WorkRequestLogEntryCollection Results of a list of log message from the execution of a work request.
type WorkRequestLogEntryCollection struct {

	// WorkRequestLogEntry items found for the listing.
	Items []WorkRequestLogEntry `mandatory:"true" json:"items"`
}

func (m WorkRequestLogEntryCollection) String() string {
	return common.PointerString(m)
}
