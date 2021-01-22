// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v33/common"
)

// IntelVmLaunchInstancePlatformConfig The platform configuration used when launching a virtual machine instance specific to the Intel platform.
type IntelVmLaunchInstancePlatformConfig struct {

	// Whether the Secure Boot is to be enabled on the instance
	IsSecureBootEnabled *bool `mandatory:"false" json:"isSecureBootEnabled"`

	// Whether the Trusted Platform Module (TPM) is to be enabled on the instance
	IsTrustedPlatformModuleEnabled *bool `mandatory:"false" json:"isTrustedPlatformModuleEnabled"`

	// Whether the Measured Boot is to be enabled on the instance
	IsMeasuredBootEnabled *bool `mandatory:"false" json:"isMeasuredBootEnabled"`
}

//GetIsSecureBootEnabled returns IsSecureBootEnabled
func (m IntelVmLaunchInstancePlatformConfig) GetIsSecureBootEnabled() *bool {
	return m.IsSecureBootEnabled
}

//GetIsTrustedPlatformModuleEnabled returns IsTrustedPlatformModuleEnabled
func (m IntelVmLaunchInstancePlatformConfig) GetIsTrustedPlatformModuleEnabled() *bool {
	return m.IsTrustedPlatformModuleEnabled
}

//GetIsMeasuredBootEnabled returns IsMeasuredBootEnabled
func (m IntelVmLaunchInstancePlatformConfig) GetIsMeasuredBootEnabled() *bool {
	return m.IsMeasuredBootEnabled
}

func (m IntelVmLaunchInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m IntelVmLaunchInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIntelVmLaunchInstancePlatformConfig IntelVmLaunchInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIntelVmLaunchInstancePlatformConfig
	}{
		"INTEL_VM",
		(MarshalTypeIntelVmLaunchInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}
