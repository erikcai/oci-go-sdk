// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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

// PipelineParameterCollection Specifies list of parameters present in Pipeline.
// In case of UPDATE operation, replaces existing parameters list. Merging with existing parameters is not supported.
type PipelineParameterCollection struct {

	// List of Parameters defined for a Pipeline.
	Items []PipelineParameter `mandatory:"true" json:"items"`
}

func (m PipelineParameterCollection) String() string {
	return common.PointerString(m)
}
