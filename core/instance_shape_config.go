// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// InstanceShapeConfig The shape configuration for an instance. The shape configuration determines
// the resources allocated to an instance.
type InstanceShapeConfig struct {

	// The total number of OCPUs available to the instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The total amount of memory available to the instance, in gigabytes.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance.
	// The following values are supported:
	// - `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
	// - `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
	BaselineOcpuUtilization InstanceShapeConfigBaselineOcpuUtilizationEnum `mandatory:"false" json:"baselineOcpuUtilization,omitempty"`

	// A short description of the instance's processor (CPU).
	ProcessorDescription *string `mandatory:"false" json:"processorDescription"`

	// The networking bandwidth available to the instance, in gigabits per second.
	NetworkingBandwidthInGbps *float32 `mandatory:"false" json:"networkingBandwidthInGbps"`

	// The maximum number of VNIC attachments for the instance.
	MaxVnicAttachments *int `mandatory:"false" json:"maxVnicAttachments"`

	// The number of GPUs available to the instance.
	Gpus *int `mandatory:"false" json:"gpus"`

	// A short description of the instance's graphics processing unit (GPU).
	// If the instance does not have any GPUs, this field is `null`.
	GpuDescription *string `mandatory:"false" json:"gpuDescription"`

	// The number of local disks available to the instance.
	LocalDisks *int `mandatory:"false" json:"localDisks"`

	// The aggregate size of all local disks, in gigabytes.
	// If the instance does not have any local disks, this field is `null`.
	LocalDisksTotalSizeInGBs *float32 `mandatory:"false" json:"localDisksTotalSizeInGBs"`

	// A short description of the local disks available to this instance.
	// If the instance does not have any local disks, this field is `null`.
	LocalDiskDescription *string `mandatory:"false" json:"localDiskDescription"`
}

func (m InstanceShapeConfig) String() string {
	return common.PointerString(m)
}

// InstanceShapeConfigBaselineOcpuUtilizationEnum Enum with underlying type: string
type InstanceShapeConfigBaselineOcpuUtilizationEnum string

// Set of constants representing the allowable values for InstanceShapeConfigBaselineOcpuUtilizationEnum
const (
	InstanceShapeConfigBaselineOcpuUtilization8 InstanceShapeConfigBaselineOcpuUtilizationEnum = "BASELINE_1_8"
	InstanceShapeConfigBaselineOcpuUtilization2 InstanceShapeConfigBaselineOcpuUtilizationEnum = "BASELINE_1_2"
)

var mappingInstanceShapeConfigBaselineOcpuUtilization = map[string]InstanceShapeConfigBaselineOcpuUtilizationEnum{
	"BASELINE_1_8": InstanceShapeConfigBaselineOcpuUtilization8,
	"BASELINE_1_2": InstanceShapeConfigBaselineOcpuUtilization2,
}

// GetInstanceShapeConfigBaselineOcpuUtilizationEnumValues Enumerates the set of values for InstanceShapeConfigBaselineOcpuUtilizationEnum
func GetInstanceShapeConfigBaselineOcpuUtilizationEnumValues() []InstanceShapeConfigBaselineOcpuUtilizationEnum {
	values := make([]InstanceShapeConfigBaselineOcpuUtilizationEnum, 0)
	for _, v := range mappingInstanceShapeConfigBaselineOcpuUtilization {
		values = append(values, v)
	}
	return values
}
