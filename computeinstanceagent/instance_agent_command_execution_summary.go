// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// InstanceAgentCommandExecutionSummary A command's execution summary.
type InstanceAgentCommandExecutionSummary struct {

	// The OCID of the instance
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Specifies the command delivery state.
	//  * `VISIBLE` - The command is visible to instance.
	//  * `PENDING` - The command is pending ack from the instance.
	//  * `ACKED` - The command has been received and acked by the instance.
	//  * `EXPIRED` - The instance has not requested for commands and its delivery has expired.
	DeliveryState InstanceAgentCommandExecutionSummaryDeliveryStateEnum `mandatory:"true" json:"deliveryState"`

	// command execution life cycle state.
	// * `ACCEPTED` - The command execution has been accepted to run.
	// * `IN_PROGRESS` - The command execution is in progress.
	// * `SUCCEEDED` - The command execution is successful.
	// * `FAILED` - The command execution has failed.
	// * `TIMED_OUT` - The command execution has timedout.
	// * `CANCELED` - The command execution has canceled.
	LifecycleState InstanceAgentCommandExecutionSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Content InstanceAgentCommandExecutionOutputContent `mandatory:"true" json:"content"`

	// The OCID of the command
	InstanceAgentCommandId *string `mandatory:"false" json:"instanceAgentCommandId"`
}

func (m InstanceAgentCommandExecutionSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *InstanceAgentCommandExecutionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InstanceAgentCommandId *string                                                `json:"instanceAgentCommandId"`
		InstanceId             *string                                                `json:"instanceId"`
		DeliveryState          InstanceAgentCommandExecutionSummaryDeliveryStateEnum  `json:"deliveryState"`
		LifecycleState         InstanceAgentCommandExecutionSummaryLifecycleStateEnum `json:"lifecycleState"`
		Content                instanceagentcommandexecutionoutputcontent             `json:"content"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.InstanceAgentCommandId = model.InstanceAgentCommandId

	m.InstanceId = model.InstanceId

	m.DeliveryState = model.DeliveryState

	m.LifecycleState = model.LifecycleState

	nn, e = model.Content.UnmarshalPolymorphicJSON(model.Content.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Content = nn.(InstanceAgentCommandExecutionOutputContent)
	} else {
		m.Content = nil
	}

	return
}

// InstanceAgentCommandExecutionSummaryDeliveryStateEnum Enum with underlying type: string
type InstanceAgentCommandExecutionSummaryDeliveryStateEnum string

// Set of constants representing the allowable values for InstanceAgentCommandExecutionSummaryDeliveryStateEnum
const (
	InstanceAgentCommandExecutionSummaryDeliveryStateVisible InstanceAgentCommandExecutionSummaryDeliveryStateEnum = "VISIBLE"
	InstanceAgentCommandExecutionSummaryDeliveryStatePending InstanceAgentCommandExecutionSummaryDeliveryStateEnum = "PENDING"
	InstanceAgentCommandExecutionSummaryDeliveryStateAcked   InstanceAgentCommandExecutionSummaryDeliveryStateEnum = "ACKED"
	InstanceAgentCommandExecutionSummaryDeliveryStateExpired InstanceAgentCommandExecutionSummaryDeliveryStateEnum = "EXPIRED"
)

var mappingInstanceAgentCommandExecutionSummaryDeliveryState = map[string]InstanceAgentCommandExecutionSummaryDeliveryStateEnum{
	"VISIBLE": InstanceAgentCommandExecutionSummaryDeliveryStateVisible,
	"PENDING": InstanceAgentCommandExecutionSummaryDeliveryStatePending,
	"ACKED":   InstanceAgentCommandExecutionSummaryDeliveryStateAcked,
	"EXPIRED": InstanceAgentCommandExecutionSummaryDeliveryStateExpired,
}

// GetInstanceAgentCommandExecutionSummaryDeliveryStateEnumValues Enumerates the set of values for InstanceAgentCommandExecutionSummaryDeliveryStateEnum
func GetInstanceAgentCommandExecutionSummaryDeliveryStateEnumValues() []InstanceAgentCommandExecutionSummaryDeliveryStateEnum {
	values := make([]InstanceAgentCommandExecutionSummaryDeliveryStateEnum, 0)
	for _, v := range mappingInstanceAgentCommandExecutionSummaryDeliveryState {
		values = append(values, v)
	}
	return values
}

// InstanceAgentCommandExecutionSummaryLifecycleStateEnum Enum with underlying type: string
type InstanceAgentCommandExecutionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for InstanceAgentCommandExecutionSummaryLifecycleStateEnum
const (
	InstanceAgentCommandExecutionSummaryLifecycleStateAccepted   InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "ACCEPTED"
	InstanceAgentCommandExecutionSummaryLifecycleStateInProgress InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "IN_PROGRESS"
	InstanceAgentCommandExecutionSummaryLifecycleStateSucceeded  InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "SUCCEEDED"
	InstanceAgentCommandExecutionSummaryLifecycleStateFailed     InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "FAILED"
	InstanceAgentCommandExecutionSummaryLifecycleStateTimedOut   InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "TIMED_OUT"
	InstanceAgentCommandExecutionSummaryLifecycleStateCanceled   InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "CANCELED"
)

var mappingInstanceAgentCommandExecutionSummaryLifecycleState = map[string]InstanceAgentCommandExecutionSummaryLifecycleStateEnum{
	"ACCEPTED":    InstanceAgentCommandExecutionSummaryLifecycleStateAccepted,
	"IN_PROGRESS": InstanceAgentCommandExecutionSummaryLifecycleStateInProgress,
	"SUCCEEDED":   InstanceAgentCommandExecutionSummaryLifecycleStateSucceeded,
	"FAILED":      InstanceAgentCommandExecutionSummaryLifecycleStateFailed,
	"TIMED_OUT":   InstanceAgentCommandExecutionSummaryLifecycleStateTimedOut,
	"CANCELED":    InstanceAgentCommandExecutionSummaryLifecycleStateCanceled,
}

// GetInstanceAgentCommandExecutionSummaryLifecycleStateEnumValues Enumerates the set of values for InstanceAgentCommandExecutionSummaryLifecycleStateEnum
func GetInstanceAgentCommandExecutionSummaryLifecycleStateEnumValues() []InstanceAgentCommandExecutionSummaryLifecycleStateEnum {
	values := make([]InstanceAgentCommandExecutionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingInstanceAgentCommandExecutionSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
