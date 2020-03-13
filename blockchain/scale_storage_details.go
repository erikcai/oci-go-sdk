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

// ScaleStorageDetails storage size to increase
type ScaleStorageDetails struct {

	// Storage size in TBs
	StorageSizeInTBs *int `mandatory:"false" json:"storageSizeInTBs"`
}

func (m ScaleStorageDetails) String() string {
	return common.PointerString(m)
}
