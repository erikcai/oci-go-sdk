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

// InstanceMetricsConfig The representation of InstanceMetricsConfig
type InstanceMetricsConfig struct {

	// Whether the performance metrics gathering agent on the instance is to be running.
	// Default value is false.
	IsDisabled *bool `mandatory:"false" json:"isDisabled"`
}

func (m InstanceMetricsConfig) String() string {
	return common.PointerString(m)
}
