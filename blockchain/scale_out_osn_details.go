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

// ScaleOutOsnDetails The Ordering Service Node details to be added
type ScaleOutOsnDetails struct {

	// organization name
	Organization *string `mandatory:"false" json:"organization"`

	// Ordering Service Log level
	LogLevel CommonLogLevelLogLevelEnum `mandatory:"false" json:"logLevel,omitempty"`

	// port number
	Port *int `mandatory:"false" json:"port"`

	// Availability Domain to place new OSN
	Ad AvailabilityDomainAdsEnum `mandatory:"false" json:"ad,omitempty"`
}

func (m ScaleOutOsnDetails) String() string {
	return common.PointerString(m)
}