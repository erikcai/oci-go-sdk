// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DbSystemAvailabilityPolicy Details of high-availability policy.
type DbSystemAvailabilityPolicy struct {

	// If the policy is to enable high availability of the instance, by
	// maintaining secondary/failover capacity as necessary.
	IsHighlyAvailable *bool `mandatory:"true" json:"isHighlyAvailable"`

	// Optional. If set, defines the *preferred* availability-domain and
	// fault-domains of candidate instances, as an ordered list.
	PreferredPlacementOrder []AdFd `mandatory:"false" json:"preferredPlacementOrder"`
}

func (m DbSystemAvailabilityPolicy) String() string {
	return common.PointerString(m)
}
