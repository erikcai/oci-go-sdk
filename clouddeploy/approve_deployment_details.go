// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// ApproveDeploymentDetails The information of stage for submitting aprove.
type ApproveDeploymentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the stage which is marked for approval.
	StageId *string `mandatory:"true" json:"stageId"`
}

func (m ApproveDeploymentDetails) String() string {
	return common.PointerString(m)
}
