// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateRegionSubscriptionDetails The representation of CreateRegionSubscriptionDetails
type CreateRegionSubscriptionDetails struct {

	// The key of the region such as PHX, IAD.
	RegionKey *string `mandatory:"true" json:"regionKey"`
}

func (m CreateRegionSubscriptionDetails) String() string {
	return common.PointerString(m)
}
