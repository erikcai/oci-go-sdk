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

// HttpProbeResultSummary Results returned by running an HTTP probe.  All times and durations are in milliseconds.
// All start and end times are relative to the same _origin_ as `startTime`.  Time
// properties conform to the W3C Resource Timing
// PerformanceResourceTiming (https://w3c.github.io/resource-timing/#sec-resource-timing)
// interface.
type HttpProbeResultSummary struct {

	// A value uniquely identifying this probe result within the currently-persisted results
	// of the probe configuration responsible for creating it.
	Key *string `mandatory:"false" json:"key"`

	// The OCID of the monitor or on-demand probe responsible for creating this result.
	ProbeConfigurationId *string `mandatory:"false" json:"probeConfigurationId"`

	// The time that that probe began execution. expressed in milliseconds since the
	// POSIX epoch.  Since this API conforms to the W3C Resource Timing
	// PerformanceResourceTiming (https://w3c.github.io/resource-timing/#sec-resource-timing)
	// interface, this field is used instead of an RFC 3339-formatted `timeStarted`
	// field.
	StartTime *float64 `mandatory:"false" json:"startTime"`

	// The hostname or IP address target of the probe.
	Target *string `mandatory:"false" json:"target"`

	// The name of the vantage point that executed the probe.
	VantagePointName *string `mandatory:"false" json:"vantagePointName"`

	// True if the probe did not complete before the configured time-out value.
	IsTimedOut *bool `mandatory:"false" json:"isTimedOut"`

	// True if the probe result is determined to be healthy based on probe type-specific
	// criteria.  For HTTP probes, a probe result is considered healthy if the HTTP
	// response code is greater than or equal to 200 and less than 300.
	IsHealthy *bool `mandatory:"false" json:"isHealthy"`

	// The category of error, if an error occurs executing the probe.  The `errorMessage`
	// field provides a message with the error details.
	// * NONE - No error
	// * DNS - DNS errors
	// * TRANSPORT - Transport-related errors, for example a "TLS certificate expired" error.
	// * NETWORK - Network-related errors, for example a "network unreachable" error.
	// * SYSTEM - Internal system errors.
	ErrorCategory HttpProbeResultSummaryErrorCategoryEnum `mandatory:"false" json:"errorCategory,omitempty"`

	// Error information indicating why a probe execution failed.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	Protocol HttpProbeResultSummaryProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	Connection *TcpConnection `mandatory:"false" json:"connection"`

	Dns *Dns `mandatory:"false" json:"dns"`

	// The HTTP response status code.
	StatusCode *int `mandatory:"false" json:"statusCode"`

	// The time immediately before the vantage point starts the domain name lookup for
	// the resource.
	DomainLookupStart *float64 `mandatory:"false" json:"domainLookupStart"`

	// The time immediately before the vantage point finishes the domain name lookup for
	// the resource.
	DomainLookupEnd *float64 `mandatory:"false" json:"domainLookupEnd"`

	// The time immediately before the vantage point starts establishing the connection
	// to the server to retrieve the resource.
	ConnectStart *float64 `mandatory:"false" json:"connectStart"`

	// The time immediately before the vantage point starts the handshake process to
	// secure the current connection.
	SecureConnectionStart *float64 `mandatory:"false" json:"secureConnectionStart"`

	// The time immediately after the vantage point finishes establishing the connection
	// to the server to retrieve the resource.
	ConnectEnd *float64 `mandatory:"false" json:"connectEnd"`

	// The time immediately before the vantage point starts to fetch the resource.
	FetchStart *float64 `mandatory:"false" json:"fetchStart"`

	// The time immediately before the vantage point starts requesting the resource from
	// the server.
	RequestStart *float64 `mandatory:"false" json:"requestStart"`

	// The time immediately after the vantage point's HTTP parser receives the first byte
	// of the response.
	ResponseStart *float64 `mandatory:"false" json:"responseStart"`

	// The time immediately after the vantage point receives the last byte of the response
	// or immediately before the transport connection is closed, whichever comes first.
	ResponseEnd *float64 `mandatory:"false" json:"responseEnd"`

	// Total duration from start of request until response is fully consumed or the
	// connection is closed.
	Duration *float64 `mandatory:"false" json:"duration"`

	// The size, in octets, of the payload body prior to removing any applied
	// content-codings.
	EncodedBodySize *int `mandatory:"false" json:"encodedBodySize"`
}

func (m HttpProbeResultSummary) String() string {
	return common.PointerString(m)
}

// HttpProbeResultSummaryErrorCategoryEnum Enum with underlying type: string
type HttpProbeResultSummaryErrorCategoryEnum string

// Set of constants representing the allowable values for HttpProbeResultSummaryErrorCategory
const (
	HttpProbeResultSummaryErrorCategoryNone      HttpProbeResultSummaryErrorCategoryEnum = "NONE"
	HttpProbeResultSummaryErrorCategoryDns       HttpProbeResultSummaryErrorCategoryEnum = "DNS"
	HttpProbeResultSummaryErrorCategoryTransport HttpProbeResultSummaryErrorCategoryEnum = "TRANSPORT"
	HttpProbeResultSummaryErrorCategoryNetwork   HttpProbeResultSummaryErrorCategoryEnum = "NETWORK"
	HttpProbeResultSummaryErrorCategorySystem    HttpProbeResultSummaryErrorCategoryEnum = "SYSTEM"
)

var mappingHttpProbeResultSummaryErrorCategory = map[string]HttpProbeResultSummaryErrorCategoryEnum{
	"NONE":      HttpProbeResultSummaryErrorCategoryNone,
	"DNS":       HttpProbeResultSummaryErrorCategoryDns,
	"TRANSPORT": HttpProbeResultSummaryErrorCategoryTransport,
	"NETWORK":   HttpProbeResultSummaryErrorCategoryNetwork,
	"SYSTEM":    HttpProbeResultSummaryErrorCategorySystem,
}

// GetHttpProbeResultSummaryErrorCategoryEnumValues Enumerates the set of values for HttpProbeResultSummaryErrorCategory
func GetHttpProbeResultSummaryErrorCategoryEnumValues() []HttpProbeResultSummaryErrorCategoryEnum {
	values := make([]HttpProbeResultSummaryErrorCategoryEnum, 0)
	for _, v := range mappingHttpProbeResultSummaryErrorCategory {
		values = append(values, v)
	}
	return values
}

// HttpProbeResultSummaryProtocolEnum Enum with underlying type: string
type HttpProbeResultSummaryProtocolEnum string

// Set of constants representing the allowable values for HttpProbeResultSummaryProtocol
const (
	HttpProbeResultSummaryProtocolHttp  HttpProbeResultSummaryProtocolEnum = "HTTP"
	HttpProbeResultSummaryProtocolHttps HttpProbeResultSummaryProtocolEnum = "HTTPS"
)

var mappingHttpProbeResultSummaryProtocol = map[string]HttpProbeResultSummaryProtocolEnum{
	"HTTP":  HttpProbeResultSummaryProtocolHttp,
	"HTTPS": HttpProbeResultSummaryProtocolHttps,
}

// GetHttpProbeResultSummaryProtocolEnumValues Enumerates the set of values for HttpProbeResultSummaryProtocol
func GetHttpProbeResultSummaryProtocolEnumValues() []HttpProbeResultSummaryProtocolEnum {
	values := make([]HttpProbeResultSummaryProtocolEnum, 0)
	for _, v := range mappingHttpProbeResultSummaryProtocol {
		values = append(values, v)
	}
	return values
}
