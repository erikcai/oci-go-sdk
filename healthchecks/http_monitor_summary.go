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

// HttpMonitorSummary This model contains all of the mutable and immutable summary properties for an HTTP monitor.
type HttpMonitorSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// Location for fetching results.
	ResultsUrl *string `mandatory:"false" json:"resultsUrl"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// A mutable name suitable for display in a user interface.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The monitor interval in seconds. Legal values are: 10, 30, and 60.
	IntervalInSeconds *int `mandatory:"false" json:"intervalInSeconds"`

	// Controls execution of the monitor.  Set this property to start and stop the
	// monitor.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	Protocol HttpMonitorSummaryProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

func (m HttpMonitorSummary) String() string {
	return common.PointerString(m)
}

// HttpMonitorSummaryProtocolEnum Enum with underlying type: string
type HttpMonitorSummaryProtocolEnum string

// Set of constants representing the allowable values for HttpMonitorSummaryProtocolEnum
const (
	HttpMonitorSummaryProtocolHttp  HttpMonitorSummaryProtocolEnum = "HTTP"
	HttpMonitorSummaryProtocolHttps HttpMonitorSummaryProtocolEnum = "HTTPS"
)

var mappingHttpMonitorSummaryProtocol = map[string]HttpMonitorSummaryProtocolEnum{
	"HTTP":  HttpMonitorSummaryProtocolHttp,
	"HTTPS": HttpMonitorSummaryProtocolHttps,
}

// GetHttpMonitorSummaryProtocolEnumValues Enumerates the set of values for HttpMonitorSummaryProtocolEnum
func GetHttpMonitorSummaryProtocolEnumValues() []HttpMonitorSummaryProtocolEnum {
	values := make([]HttpMonitorSummaryProtocolEnum, 0)
	for _, v := range mappingHttpMonitorSummaryProtocol {
		values = append(values, v)
	}
	return values
}
