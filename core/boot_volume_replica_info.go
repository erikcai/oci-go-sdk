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
	"github.com/oracle/oci-go-sdk/v33/common"
)

// BootVolumeReplicaInfo Information about the boot volume replica in the destination availability domain.
type BootVolumeReplicaInfo struct {

	// The display name of the boot volume replica
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The boot volume replica's Oracle ID (OCID).
	BootVolumeReplicaId *string `mandatory:"true" json:"bootVolumeReplicaId"`

	// The availability domain of the boot volume replica.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`
}

func (m BootVolumeReplicaInfo) String() string {
	return common.PointerString(m)
}
