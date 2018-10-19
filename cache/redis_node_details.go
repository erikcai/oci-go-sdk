// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// Oracle Caching Service Public API
//

package cache

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RedisNodeDetails The availability domain and subnet details for each redis node.
type RedisNodeDetails struct {

	// The availability domain of this redis node.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The subnet id of this redis node.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A flag that indicates the primary node. Default value is 'false'.
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`
}

func (m RedisNodeDetails) String() string {
	return common.PointerString(m)
}
