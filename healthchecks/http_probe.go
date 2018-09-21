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

// HttpProbe This model contains all of the mutable and immutable properties for an HTTP probe.
type HttpProbe struct {

	// The OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// Location for fetching results.
	ResultsUrl *string `mandatory:"false" json:"resultsUrl"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	Targets []string `mandatory:"false" json:"targets"`

	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The probe port, defaults to the default port for the specified protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Legal values are: 10, 20, 30, and 60. The probe timeout
	// must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	Protocol HttpProbeProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	Method HttpProbeMethodEnum `mandatory:"false" json:"method,omitempty"`

	// The optional URL path (including query parameters).
	Path *string `mandatory:"false" json:"path"`

	// A dictionary of HTTP request headers.
	Headers map[string]string `mandatory:"false" json:"headers"`
}

func (m HttpProbe) String() string {
	return common.PointerString(m)
}

// HttpProbeProtocolEnum Enum with underlying type: string
type HttpProbeProtocolEnum string

// Set of constants representing the allowable values for HttpProbeProtocol
const (
	HttpProbeProtocolHttp  HttpProbeProtocolEnum = "HTTP"
	HttpProbeProtocolHttps HttpProbeProtocolEnum = "HTTPS"
)

var mappingHttpProbeProtocol = map[string]HttpProbeProtocolEnum{
	"HTTP":  HttpProbeProtocolHttp,
	"HTTPS": HttpProbeProtocolHttps,
}

// GetHttpProbeProtocolEnumValues Enumerates the set of values for HttpProbeProtocol
func GetHttpProbeProtocolEnumValues() []HttpProbeProtocolEnum {
	values := make([]HttpProbeProtocolEnum, 0)
	for _, v := range mappingHttpProbeProtocol {
		values = append(values, v)
	}
	return values
}

// HttpProbeMethodEnum Enum with underlying type: string
type HttpProbeMethodEnum string

// Set of constants representing the allowable values for HttpProbeMethod
const (
	HttpProbeMethodGet  HttpProbeMethodEnum = "GET"
	HttpProbeMethodHead HttpProbeMethodEnum = "HEAD"
)

var mappingHttpProbeMethod = map[string]HttpProbeMethodEnum{
	"GET":  HttpProbeMethodGet,
	"HEAD": HttpProbeMethodHead,
}

// GetHttpProbeMethodEnumValues Enumerates the set of values for HttpProbeMethod
func GetHttpProbeMethodEnumValues() []HttpProbeMethodEnum {
	values := make([]HttpProbeMethodEnum, 0)
	for _, v := range mappingHttpProbeMethod {
		values = append(values, v)
	}
	return values
}
