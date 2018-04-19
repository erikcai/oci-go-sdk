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

// ConnectLocalPeeringGatewaysDetails Contains details indicating the local peering gateway with which you wish to establish a peering relationship.
type ConnectLocalPeeringGatewaysDetails struct {

	// The OCID of the other local peering gateway.
	PeerId *string `mandatory:"true" json:"peerId"`
}

func (m ConnectLocalPeeringGatewaysDetails) String() string {
	return common.PointerString(m)
}
