// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// WorkRequestCollection Results of a list of WorkRequest.
type WorkRequestCollection struct {

	// WorkRequest items found for the search.
	Items []WorkRequestSummary `mandatory:"true" json:"items"`
}

func (m WorkRequestCollection) String() string {
	return common.PointerString(m)
}