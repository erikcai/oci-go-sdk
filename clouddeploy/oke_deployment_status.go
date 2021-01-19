// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// OkeDeploymentStatus Specifies OKE Deployment Stage Specific Status.
type OkeDeploymentStatus struct {

	// The current state of the Oke Deployment Stage.
	Status OkeDeploymentStatusStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A message describing the current status of Oke Deployment Stage in more details.
	Message *string `mandatory:"false" json:"message"`

	// The timestamp when current status of Oke Deployment recorded. An RFC3339 formatted datetime string
	TimeReported *common.SDKTime `mandatory:"false" json:"timeReported"`
}

func (m OkeDeploymentStatus) String() string {
	return common.PointerString(m)
}

// OkeDeploymentStatusStatusEnum Enum with underlying type: string
type OkeDeploymentStatusStatusEnum string

// Set of constants representing the allowable values for OkeDeploymentStatusStatusEnum
const (
	OkeDeploymentStatusStatusInProgress OkeDeploymentStatusStatusEnum = "IN_PROGRESS"
	OkeDeploymentStatusStatusSucceeded  OkeDeploymentStatusStatusEnum = "SUCCEEDED"
	OkeDeploymentStatusStatusFailed     OkeDeploymentStatusStatusEnum = "FAILED"
)

var mappingOkeDeploymentStatusStatus = map[string]OkeDeploymentStatusStatusEnum{
	"IN_PROGRESS": OkeDeploymentStatusStatusInProgress,
	"SUCCEEDED":   OkeDeploymentStatusStatusSucceeded,
	"FAILED":      OkeDeploymentStatusStatusFailed,
}

// GetOkeDeploymentStatusStatusEnumValues Enumerates the set of values for OkeDeploymentStatusStatusEnum
func GetOkeDeploymentStatusStatusEnumValues() []OkeDeploymentStatusStatusEnum {
	values := make([]OkeDeploymentStatusStatusEnum, 0)
	for _, v := range mappingOkeDeploymentStatusStatus {
		values = append(values, v)
	}
	return values
}
