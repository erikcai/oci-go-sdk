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

// ComputeInstanceGroupLinearRollOutStrategy Specifies a linear rollout strategy for Compute Instance group rolling deployment stage.
type ComputeInstanceGroupLinearRollOutStrategy struct {

	// Whether the step size is specified by abosulte number or by percetage.
	BatchType ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum `mandatory:"true" json:"batchType"`

	// The number or percentage that will be used to determine how many instances will be deployed to concurrently.
	BatchLimit *int `mandatory:"true" json:"batchLimit"`
}

func (m ComputeInstanceGroupLinearRollOutStrategy) String() string {
	return common.PointerString(m)
}

// ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum Enum with underlying type: string
type ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum string

// Set of constants representing the allowable values for ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum
const (
	ComputeInstanceGroupLinearRollOutStrategyBatchTypeNumber     ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum = "NUMBER"
	ComputeInstanceGroupLinearRollOutStrategyBatchTypePercentage ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum = "PERCENTAGE"
)

var mappingComputeInstanceGroupLinearRollOutStrategyBatchType = map[string]ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum{
	"NUMBER":     ComputeInstanceGroupLinearRollOutStrategyBatchTypeNumber,
	"PERCENTAGE": ComputeInstanceGroupLinearRollOutStrategyBatchTypePercentage,
}

// GetComputeInstanceGroupLinearRollOutStrategyBatchTypeEnumValues Enumerates the set of values for ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum
func GetComputeInstanceGroupLinearRollOutStrategyBatchTypeEnumValues() []ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum {
	values := make([]ComputeInstanceGroupLinearRollOutStrategyBatchTypeEnum, 0)
	for _, v := range mappingComputeInstanceGroupLinearRollOutStrategyBatchType {
		values = append(values, v)
	}
	return values
}
