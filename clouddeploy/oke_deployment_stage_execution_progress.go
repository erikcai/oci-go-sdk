// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v31/common"
)

// OkeDeploymentStageExecutionProgress Specifies the execution details for OKEDeployment stage.
type OkeDeploymentStageExecutionProgress struct {

	// The time the Stage was started executing. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Stage was finished executing. An RFC3339 formatted datetime string
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// A message describing the current state of the stage in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StatusMessages []string `mandatory:"false" json:"statusMessages"`

	// Specifies a list of Oke Deployment statuses
	OkeDeploymentStatuses []OkeDeploymentStatus `mandatory:"false" json:"okeDeploymentStatuses"`

	// The current state of the Stage.
	Status StageExecutionProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
}

//GetTimeStarted returns TimeStarted
func (m OkeDeploymentStageExecutionProgress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m OkeDeploymentStageExecutionProgress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m OkeDeploymentStageExecutionProgress) GetStatus() StageExecutionProgressStatusEnum {
	return m.Status
}

//GetStatusMessages returns StatusMessages
func (m OkeDeploymentStageExecutionProgress) GetStatusMessages() []string {
	return m.StatusMessages
}

func (m OkeDeploymentStageExecutionProgress) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OkeDeploymentStageExecutionProgress) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOkeDeploymentStageExecutionProgress OkeDeploymentStageExecutionProgress
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeOkeDeploymentStageExecutionProgress
	}{
		"OKE_DEPLOYMENT",
		(MarshalTypeOkeDeploymentStageExecutionProgress)(m),
	}

	return json.Marshal(&s)
}
