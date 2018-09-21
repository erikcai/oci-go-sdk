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

// Dns DNS resolution results.  All durations are in milliseconds.
type Dns struct {

	// Total DNS resolution duration.  Calculated by domainLookupEnd - domainLookupStart.
	DomainLookupDuration *float64 `mandatory:"false" json:"domainLookupDuration"`

	// Addresses returned by DNS resolution.
	Addresses []string `mandatory:"false" json:"addresses"`
}

func (m Dns) String() string {
	return common.PointerString(m)
}
