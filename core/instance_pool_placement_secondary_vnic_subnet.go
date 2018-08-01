// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// InstancePoolPlacementSecondaryVnicSubnet The secondary VNIC object for the placement configuration for an instance pool.
type InstancePoolPlacementSecondaryVnicSubnet struct {

	// The subnet OCID for the secondary vnic
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Which physical network interface card (NIC) the VNIC will use. Defaults to 0.
	// Certain bare metal instance shapes have two active physical NICs (0 and 1). If
	// you add a secondary VNIC to one of these instances, you can specify which NIC
	// the VNIC will use. For more information, see
	// Virtual Network Interface Cards (VNICs) (https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVNICs.htm).
	NicIndex *int `mandatory:"false" json:"nicIndex"`
}

func (m InstancePoolPlacementSecondaryVnicSubnet) String() string {
	return common.PointerString(m)
}
