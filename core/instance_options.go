// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
)

// InstanceOptions Optional mutable instance options
type InstanceOptions struct {

	// Whether to disable the legacy (/v1) instance metadata service endpoints.
	// Customers who have migrated to /v2 should set this to true for added security.
	// Default is false.
	AreLegacyImdsEndpointsDisabled *bool `mandatory:"false" json:"areLegacyImdsEndpointsDisabled"`
}

func (m InstanceOptions) String() string {
	return common.PointerString(m)
}
