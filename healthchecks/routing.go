// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Health Checks Service API
//
// Health Checks Service API.  This API allows clients to configure and run probes (tests)
// that will be executed on one or more global vantage points to monitor OCI assets.  The API
// supports running on-demand probes as well as retrieving historical results.
//

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Routing Routing information for a vantage point.
type Routing struct {

	// Registry label for `asn`. May be null.
	AsLabel *string `mandatory:"false" json:"asLabel"`

	// Autonomous system number identifying the organization
	// responsible for routing packets to `prefix`.
	Asn *int `mandatory:"false" json:"asn"`

	// An IP prefix (CIDR syntax) that is less specific than
	// `address`, through which `address` is routed.
	Prefix *string `mandatory:"false" json:"prefix"`

	// An integer between 0 and 100 used to select between multiple
	// origin ASNs when routing to `prefix`. Most prefixes have
	// exactly one origin ASN, in which case `weight` will be 100.
	Weight *int `mandatory:"false" json:"weight"`
}

func (m Routing) String() string {
	return common.PointerString(m)
}
