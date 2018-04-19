// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DNS Service API
//
// API for managing DNS zones, records, and policies.
//

package dns

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Nameserver A server that has been set up to answer DNS queries for a zone.
type Nameserver struct {

	// The hostname of the nameserver.
	Hostname *string `mandatory:"true" json:"hostname"`
}

func (m Nameserver) String() string {
	return common.PointerString(m)
}
