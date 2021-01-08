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

// InstanceDeploymentStatus Describe the deployment status of one instance.
type InstanceDeploymentStatus struct {

	// the OCID of the instance
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// The current status of the deployment.
	Status InstanceDeploymentStatusStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A more detailed message about the current status of the deployment on the instance.
	Message *string `mandatory:"false" json:"message"`
}

func (m InstanceDeploymentStatus) String() string {
	return common.PointerString(m)
}

// InstanceDeploymentStatusStatusEnum Enum with underlying type: string
type InstanceDeploymentStatusStatusEnum string

// Set of constants representing the allowable values for InstanceDeploymentStatusStatusEnum
const (
	InstanceDeploymentStatusStatusInProgress InstanceDeploymentStatusStatusEnum = "IN_PROGRESS"
	InstanceDeploymentStatusStatusSucceeded  InstanceDeploymentStatusStatusEnum = "SUCCEEDED"
	InstanceDeploymentStatusStatusFailed     InstanceDeploymentStatusStatusEnum = "FAILED"
)

var mappingInstanceDeploymentStatusStatus = map[string]InstanceDeploymentStatusStatusEnum{
	"IN_PROGRESS": InstanceDeploymentStatusStatusInProgress,
	"SUCCEEDED":   InstanceDeploymentStatusStatusSucceeded,
	"FAILED":      InstanceDeploymentStatusStatusFailed,
}

// GetInstanceDeploymentStatusStatusEnumValues Enumerates the set of values for InstanceDeploymentStatusStatusEnum
func GetInstanceDeploymentStatusStatusEnumValues() []InstanceDeploymentStatusStatusEnum {
	values := make([]InstanceDeploymentStatusStatusEnum, 0)
	for _, v := range mappingInstanceDeploymentStatusStatus {
		values = append(values, v)
	}
	return values
}
