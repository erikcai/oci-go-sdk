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

// FunctionDeploymentStatus Specifies Function Deployment Stage Specific Status.
type FunctionDeploymentStatus struct {

	// The current state of the Function Deployment Stage.
	Status FunctionDeploymentStatusStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A message describing the current status of Function Deployment Stage in more details.
	Message *string `mandatory:"false" json:"message"`

	// The timestamp when current status of Function Deployment recorded. An RFC3339 formatted datetime string
	TimeReported *common.SDKTime `mandatory:"false" json:"timeReported"`
}

func (m FunctionDeploymentStatus) String() string {
	return common.PointerString(m)
}

// FunctionDeploymentStatusStatusEnum Enum with underlying type: string
type FunctionDeploymentStatusStatusEnum string

// Set of constants representing the allowable values for FunctionDeploymentStatusStatusEnum
const (
	FunctionDeploymentStatusStatusInProgress FunctionDeploymentStatusStatusEnum = "IN_PROGRESS"
	FunctionDeploymentStatusStatusSucceeded  FunctionDeploymentStatusStatusEnum = "SUCCEEDED"
	FunctionDeploymentStatusStatusFailed     FunctionDeploymentStatusStatusEnum = "FAILED"
)

var mappingFunctionDeploymentStatusStatus = map[string]FunctionDeploymentStatusStatusEnum{
	"IN_PROGRESS": FunctionDeploymentStatusStatusInProgress,
	"SUCCEEDED":   FunctionDeploymentStatusStatusSucceeded,
	"FAILED":      FunctionDeploymentStatusStatusFailed,
}

// GetFunctionDeploymentStatusStatusEnumValues Enumerates the set of values for FunctionDeploymentStatusStatusEnum
func GetFunctionDeploymentStatusStatusEnumValues() []FunctionDeploymentStatusStatusEnum {
	values := make([]FunctionDeploymentStatusStatusEnum, 0)
	for _, v := range mappingFunctionDeploymentStatusStatus {
		values = append(values, v)
	}
	return values
}
