// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScaleReplicaInfoDetails Number of replicas of service components like Rest Proxy, CA and Console
type ScaleReplicaInfoDetails struct {

	// Number of REST proxy replicas
	ProxyCount *int `mandatory:"false" json:"proxyCount"`

	// Number of CA replicas
	CaCount *int `mandatory:"false" json:"caCount"`

	// Number of console replicas
	ConsoleCount *int `mandatory:"false" json:"consoleCount"`
}

func (m ScaleReplicaInfoDetails) String() string {
	return common.PointerString(m)
}
