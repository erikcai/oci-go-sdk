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

// PingProbeResultSummary Results returned by running a Ping probe.  All times and durations are in milliseconds.
// All start and end times are relative to the same _origin_ as `startTime`.
type PingProbeResultSummary struct {

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
	ErrorCategory PingProbeResultSummaryErrorCategoryEnum `mandatory:"false" json:"errorCategory,omitempty"`

	// Error information indicating why a probe execution failed.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	Protocol PingProbeResultSummaryProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	Connection *Connection `mandatory:"false" json:"connection"`

	Dns *Dns `mandatory:"false" json:"dns"`

	// The time immediately before the vantage point starts the domain name lookup for
	// the resource.
	DomainLookupStart *float64 `mandatory:"false" json:"domainLookupStart"`

	// The time immediately before the vantage point finishes the domain name lookup for
	// the resource.
	DomainLookupEnd *float64 `mandatory:"false" json:"domainLookupEnd"`

	// The number of attempted measurements.
	Count *int `mandatory:"false" json:"count"`

	// The minimum latency in milliseconds across measurements in the test run.
	MinimumLatencyInMs *float64 `mandatory:"false" json:"minimumLatencyInMs"`

	// The maximum latency in milliseconds across measurements in the test run.
	MaximumLatencyInMs *float64 `mandatory:"false" json:"maximumLatencyInMs"`

	// The mean latency in milliseconds across measurements in the test run.
	MeanLatencyInMs *float64 `mandatory:"false" json:"meanLatencyInMs"`

	// The median latency in milliseconds across measurements in the test run.
	MedianLatencyInMs *float64 `mandatory:"false" json:"medianLatencyInMs"`

	// The packet loss percentage expressed as a number between 0 and 100 inclusive.
	PacketLoss *float64 `mandatory:"false" json:"packetLoss"`
}

func (m PingProbeResultSummary) String() string {
	return common.PointerString(m)
}

// PingProbeResultSummaryErrorCategoryEnum Enum with underlying type: string
type PingProbeResultSummaryErrorCategoryEnum string

// Set of constants representing the allowable values for PingProbeResultSummaryErrorCategoryEnum
const (
	PingProbeResultSummaryErrorCategoryNone      PingProbeResultSummaryErrorCategoryEnum = "NONE"
	PingProbeResultSummaryErrorCategoryDns       PingProbeResultSummaryErrorCategoryEnum = "DNS"
	PingProbeResultSummaryErrorCategoryTransport PingProbeResultSummaryErrorCategoryEnum = "TRANSPORT"
	PingProbeResultSummaryErrorCategoryNetwork   PingProbeResultSummaryErrorCategoryEnum = "NETWORK"
	PingProbeResultSummaryErrorCategorySystem    PingProbeResultSummaryErrorCategoryEnum = "SYSTEM"
)

var mappingPingProbeResultSummaryErrorCategory = map[string]PingProbeResultSummaryErrorCategoryEnum{
	"NONE":      PingProbeResultSummaryErrorCategoryNone,
	"DNS":       PingProbeResultSummaryErrorCategoryDns,
	"TRANSPORT": PingProbeResultSummaryErrorCategoryTransport,
	"NETWORK":   PingProbeResultSummaryErrorCategoryNetwork,
	"SYSTEM":    PingProbeResultSummaryErrorCategorySystem,
}

// GetPingProbeResultSummaryErrorCategoryEnumValues Enumerates the set of values for PingProbeResultSummaryErrorCategoryEnum
func GetPingProbeResultSummaryErrorCategoryEnumValues() []PingProbeResultSummaryErrorCategoryEnum {
	values := make([]PingProbeResultSummaryErrorCategoryEnum, 0)
	for _, v := range mappingPingProbeResultSummaryErrorCategory {
		values = append(values, v)
	}
	return values
}

// PingProbeResultSummaryProtocolEnum Enum with underlying type: string
type PingProbeResultSummaryProtocolEnum string

// Set of constants representing the allowable values for PingProbeResultSummaryProtocolEnum
const (
	PingProbeResultSummaryProtocolIcmp PingProbeResultSummaryProtocolEnum = "ICMP"
	PingProbeResultSummaryProtocolTcp  PingProbeResultSummaryProtocolEnum = "TCP"
)

var mappingPingProbeResultSummaryProtocol = map[string]PingProbeResultSummaryProtocolEnum{
	"ICMP": PingProbeResultSummaryProtocolIcmp,
	"TCP":  PingProbeResultSummaryProtocolTcp,
}

// GetPingProbeResultSummaryProtocolEnumValues Enumerates the set of values for PingProbeResultSummaryProtocolEnum
func GetPingProbeResultSummaryProtocolEnumValues() []PingProbeResultSummaryProtocolEnum {
	values := make([]PingProbeResultSummaryProtocolEnum, 0)
	for _, v := range mappingPingProbeResultSummaryProtocol {
		values = append(values, v)
	}
	return values
}
