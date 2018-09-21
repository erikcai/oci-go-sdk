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

// HealthChecksVantagePointSummary Information about a vantage point.
type HealthChecksVantagePointSummary struct {

	// The display name for the vantagepoint.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The organization on whose infrastructure this vantage point runs.
	// Typically not unique, as we usually have many vantage points on each
	// major provider.
	ProviderName *string `mandatory:"false" json:"providerName"`

	// The unique permanent name for the vantage point.
	Name *string `mandatory:"false" json:"name"`

	Geo *Geolocation `mandatory:"false" json:"geo"`

	// Array of objects that describe how the vantage point is routed, i.e.
	// what prefixes and ASNs connect it to the rest of the Internet. Sorted
	// from most-specific to least-specific prefix. When a prefix has
	// multiple origin ASNs (MOAS routing), they are sorted by weight
	// (highest to lowest).
	// Will be null if routing information is unknown (e.g. if `address` is
	// in an RFC 1918 prefix).
	Routing []Routing `mandatory:"false" json:"routing"`
}

func (m HealthChecksVantagePointSummary) String() string {
	return common.PointerString(m)
}
