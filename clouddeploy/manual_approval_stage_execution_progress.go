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
	"github.com/oracle/oci-go-sdk/v29/common"
)

// ManualApprovalStageExecutionProgress Specifies the manual approval stage specific execution details.
type ManualApprovalStageExecutionProgress struct {

	// The time the Stage was started executing. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Stage was finished executing. An RFC3339 formatted datetime string
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// A message describing the current state of the stage in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StatusMessages []string `mandatory:"false" json:"statusMessages"`

	// The current state of the Stage.
	Status StageExecutionProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
}

//GetTimeStarted returns TimeStarted
func (m ManualApprovalStageExecutionProgress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m ManualApprovalStageExecutionProgress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m ManualApprovalStageExecutionProgress) GetStatus() StageExecutionProgressStatusEnum {
	return m.Status
}

//GetStatusMessages returns StatusMessages
func (m ManualApprovalStageExecutionProgress) GetStatusMessages() []string {
	return m.StatusMessages
}

func (m ManualApprovalStageExecutionProgress) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ManualApprovalStageExecutionProgress) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManualApprovalStageExecutionProgress ManualApprovalStageExecutionProgress
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeManualApprovalStageExecutionProgress
	}{
		"MANUAL_APPROVAL",
		(MarshalTypeManualApprovalStageExecutionProgress)(m),
	}

	return json.Marshal(&s)
}
