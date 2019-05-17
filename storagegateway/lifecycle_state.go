// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// StorageGateway API
//
// API for interfacing with StorageGateway
//

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LifecycleState The resource's life-cycle state. After creating the resource, make sure its state changes to
// ACTIVE before using it.
type LifecycleState struct {
}

func (m LifecycleState) String() string {
	return common.PointerString(m)
}
