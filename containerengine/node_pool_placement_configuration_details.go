// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"github.com/oracle/oci-go-sdk/common"
)

// NodePoolPlacementConfigurationDetails The location where a node pool will place nodes.
type NodePoolPlacementConfigurationDetails struct {

	// The availability domain to place nodes.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the subnet to place nodes.
	SubnetId *string `mandatory:"true" json:"subnetId"`
}

func (m NodePoolPlacementConfigurationDetails) String() string {
	return common.PointerString(m)
}
