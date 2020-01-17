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

// CpeDeviceShapeDetail The representation of CpeDeviceShapeDetail
type CpeDeviceShapeDetail struct {

	// The CPE device type's unique identifier.
	CpeDeviceShapeId *string `mandatory:"false" json:"cpeDeviceShapeId"`

	// customer premise equipment hardware information
	CpeDeviceInfo *CpeDeviceInfo `mandatory:"false" json:"cpeDeviceInfo"`

	// list of questions to ask to cusomter regarding their cpe device in order to generate their cpe device config
	Parameters []CpeDeviceConfigQuestion `mandatory:"false" json:"parameters"`

	// the template that will be combined together with customer input to render customer cpe device configuration
	Template *string `mandatory:"false" json:"template"`
}

func (m CpeDeviceShapeDetail) String() string {
	return common.PointerString(m)
}
