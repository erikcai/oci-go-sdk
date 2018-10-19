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

// EndPoint Represents an endpoint through which the cache can be accessed
type EndPoint struct {

	// The IP of the endpoint.
	Ip *string `mandatory:"true" json:"ip"`

	// The port of the endpoint.
	Port *int `mandatory:"true" json:"port"`

	// A flag to indicate the primary redis node.
	IsPrimary *bool `mandatory:"true" json:"isPrimary"`

	// The subnet id of this redis node.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The availability domain of this redis node.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The hostname of the endpoint.
	Hostname *string `mandatory:"false" json:"hostname"`
}

func (m EndPoint) String() string {
	return common.PointerString(m)
}
