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

// AdFd A combination of AvailabiltyDomain and FaultDomain
type AdFd struct {

	// As name.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// As name.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`
}

func (m AdFd) String() string {
	return common.PointerString(m)
}
