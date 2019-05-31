// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Bdtstest1 API
//
// A description of the Bdtstest1 API
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MultipleTransferDevices The representation of MultipleTransferDevices
type MultipleTransferDevices struct {

	// List of TransferDeviceObject's
	TransferDeviceObjects []TransferDeviceSummary `mandatory:"false" json:"transferDeviceObjects"`
}

func (m MultipleTransferDevices) String() string {
	return common.PointerString(m)
}
