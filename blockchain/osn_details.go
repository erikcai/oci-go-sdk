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

// OsnDetails An Ordering Service Node details
type OsnDetails struct {

	// Ordering Service Log level
	LogLevel CommonLogLevelLogLevelEnum `mandatory:"false" json:"logLevel,omitempty"`

	// OSN identifier
	OrdererName *string `mandatory:"false" json:"ordererName"`

	// organization name
	Organization *string `mandatory:"false" json:"organization"`

	// port numbers
	Port []int `mandatory:"false" json:"port"`

	// Availability Domain of OSN
	Ad AvailabilityDomainAdsEnum `mandatory:"false" json:"ad,omitempty"`
}

func (m OsnDetails) String() string {
	return common.PointerString(m)
}
