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

// UpgradeStatus Upgrade status of the DRG.
type UpgradeStatus struct {

	// The Oracle-assigned ID of the DRG Attachment
	Status UpgradeStatusStatusEnum `mandatory:"true" json:"status"`

	// The number of connections upgraded
	UpgradedConnections *string `mandatory:"true" json:"upgradedConnections"`
}

func (m UpgradeStatus) String() string {
	return common.PointerString(m)
}

// UpgradeStatusStatusEnum Enum with underlying type: string
type UpgradeStatusStatusEnum string

// Set of constants representing the allowable values for UpgradeStatusStatusEnum
const (
	UpgradeStatusStatusNotUpgraded UpgradeStatusStatusEnum = "NOT_UPGRADED"
	UpgradeStatusStatusInProgress  UpgradeStatusStatusEnum = "IN_PROGRESS"
	UpgradeStatusStatusUpgraded    UpgradeStatusStatusEnum = "UPGRADED"
)

var mappingUpgradeStatusStatus = map[string]UpgradeStatusStatusEnum{
	"NOT_UPGRADED": UpgradeStatusStatusNotUpgraded,
	"IN_PROGRESS":  UpgradeStatusStatusInProgress,
	"UPGRADED":     UpgradeStatusStatusUpgraded,
}

// GetUpgradeStatusStatusEnumValues Enumerates the set of values for UpgradeStatusStatusEnum
func GetUpgradeStatusStatusEnumValues() []UpgradeStatusStatusEnum {
	values := make([]UpgradeStatusStatusEnum, 0)
	for _, v := range mappingUpgradeStatusStatus {
		values = append(values, v)
	}
	return values
}
