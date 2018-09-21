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

// Geolocation Geographic information about a vantage point.
type Geolocation struct {

	// Opaque identifier for the geographical location of the vantage point.
	GeoKey *string `mandatory:"false" json:"geoKey"`

	// ISO 3166-2 code for the first-level administrative division (limited
	// to US states and Canadian provinces, both of which use 2 characters).
	AdminDivCode *string `mandatory:"false" json:"adminDivCode"`

	// Common English-language name for the city.
	CityName *string `mandatory:"false" json:"cityName"`

	// 2-letter ISO 3166-1 alpha-2 country code.
	CountryCode *string `mandatory:"false" json:"countryCode"`

	// Common English-language name for the country.
	CountryName *string `mandatory:"false" json:"countryName"`

	// Degrees north of the Equator.
	Latitude *float32 `mandatory:"false" json:"latitude"`

	// Degrees east of the prime meridian.
	Longitude *float32 `mandatory:"false" json:"longitude"`
}

func (m Geolocation) String() string {
	return common.PointerString(m)
}
