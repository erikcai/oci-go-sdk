// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// StagePredecessor Metadata for defining a stage's predecessor.
type StagePredecessor struct {

	// The id of the predecessor stage. If a stage is the first stage in the pipeline, then the id is the pipeline's id.
	Id *string `mandatory:"true" json:"id"`
}

func (m StagePredecessor) String() string {
	return common.PointerString(m)
}
