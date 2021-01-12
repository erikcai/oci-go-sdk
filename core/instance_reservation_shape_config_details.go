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
	"github.com/oracle/oci-go-sdk/v32/common"
)

// InstanceReservationShapeConfigDetails The shape configuration requested when launching instances in a compute capacity reservation.
// If the parameter is provided, the reservation is created with the resources that you specify. If some
// properties are missing or the parameter is not provided, the reservation is created
// with the default configuration values for the `shape` that you specify.
// Each shape only supports certain configurable values. If the values that you provide are not valid for the
// specified `shape`, an error is returned.
// For more information about customizing the resources that are allocated to flexible shapes,
// see Flexible Shapes (https://docs.cloud.oracle.com/Content/Compute/References/computeshapes.htm#flexible).
type InstanceReservationShapeConfigDetails struct {

	// The total number of OCPUs available to the instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The total amount of memory available to the instance, in gigabytes.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`
}

func (m InstanceReservationShapeConfigDetails) String() string {
	return common.PointerString(m)
}
