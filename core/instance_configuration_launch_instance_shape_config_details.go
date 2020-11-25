// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm),
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm), and
// Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as virtual cloud networks (VCNs),
// compute instances, block storage volumes, and container images.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// InstanceConfigurationLaunchInstanceShapeConfigDetails The shape configuration requested for the instance.
// If the parameter is provided, the instance is created
// with the resources that you specify. If some properties are missing or
// the entire parameter is not provided, the instance is created with the default
// configuration values for the `shape` that you specify.
// Each shape only supports certain configurable values. If the values that you provid are not valid for the
// specified `shape`, an error is returned.
// For more information about customizing the resources that are allocated to a flexible shapes, see
// Flexible Shapes (https://docs.cloud.oracle.com/Content/Compute/References/computeshapes.htm#flexible).
type InstanceConfigurationLaunchInstanceShapeConfigDetails struct {

	// The total number of OCPUs available to the instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The total amount of memory available to the instance, in gigabytes.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance.
	// The following values are supported:
	// - `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
	// - `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
	BaselineOcpuUtilization InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum `mandatory:"false" json:"baselineOcpuUtilization,omitempty"`
}

func (m InstanceConfigurationLaunchInstanceShapeConfigDetails) String() string {
	return common.PointerString(m)
}

// InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum Enum with underlying type: string
type InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum string

// Set of constants representing the allowable values for InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum
const (
	InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilization8 InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum = "BASELINE_1_8"
	InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilization2 InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum = "BASELINE_1_2"
)

var mappingInstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilization = map[string]InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum{
	"BASELINE_1_8": InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilization8,
	"BASELINE_1_2": InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilization2,
}

// GetInstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnumValues Enumerates the set of values for InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum
func GetInstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnumValues() []InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum {
	values := make([]InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum, 0)
	for _, v := range mappingInstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilization {
		values = append(values, v)
	}
	return values
}
