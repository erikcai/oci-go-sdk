// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// API for the Data Caching Service. Use this service to manage Redis replicated caches.
//

package cache

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// EndPoint An endpoint for accessing a Redis replicated cache.
type EndPoint struct {

	// The IP of the endpoint.
	Ip *string `mandatory:"true" json:"ip"`

	// The port of the endpoint.
	Port *int `mandatory:"true" json:"port"`

	// A flag that indicates the primary Redis node.
	IsPrimary *bool `mandatory:"true" json:"isPrimary"`

	// The subnet id of this Redis node.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The availability domain of this Redis node.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The hostname of the endpoint.
	Hostname *string `mandatory:"false" json:"hostname"`
}

func (m EndPoint) String() string {
	return common.PointerString(m)
}
