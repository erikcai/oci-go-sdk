// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"encoding/json"
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// DeploymentExecutionProgress The execution progress details of a Deployment.
type DeploymentExecutionProgress struct {

	// The time the the Deployment is started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Deployment is finished. An RFC3339 formatted datetime string
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Map of stage OCIDs to StageExecutionProgress model.
	StageExecutionProgress map[string]StageExecutionProgress `mandatory:"false" json:"stageExecutionProgress"`
}

func (m DeploymentExecutionProgress) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DeploymentExecutionProgress) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeStarted            *common.SDKTime                   `json:"timeStarted"`
		TimeFinished           *common.SDKTime                   `json:"timeFinished"`
		StageExecutionProgress map[string]stageexecutionprogress `json:"stageExecutionProgress"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.StageExecutionProgress = make(map[string]StageExecutionProgress)
	for k, v := range model.StageExecutionProgress {
		nn, e = v.UnmarshalPolymorphicJSON(v.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.StageExecutionProgress[k] = nn.(StageExecutionProgress)
		} else {
			m.StageExecutionProgress[k] = nil
		}
	}

	return
}
