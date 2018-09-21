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

// CreateHttpMonitorDetails The request body used to create an HTTP monitor.
type CreateHttpMonitorDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Targets []string `mandatory:"true" json:"targets"`

	Protocol CreateHttpMonitorDetailsProtocolEnum `mandatory:"true" json:"protocol"`

	// A mutable name suitable for display in a user interface.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The monitor interval in seconds. Legal values are: 10, 30, and 60.
	IntervalInSeconds *int `mandatory:"true" json:"intervalInSeconds"`

	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The probe port, defaults to the default port for the specified protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Legal values are: 10, 20, 30, and 60. The probe timeout
	// must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	Method CreateHttpMonitorDetailsMethodEnum `mandatory:"false" json:"method,omitempty"`

	// The optional URL path (including query parameters).
	Path *string `mandatory:"false" json:"path"`

	// A dictionary of HTTP request headers.
	Headers map[string]string `mandatory:"false" json:"headers"`

	// Controls execution of the monitor.  Set this property to start and stop the
	// monitor.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace.  For more information,
	// see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateHttpMonitorDetails) String() string {
	return common.PointerString(m)
}

// CreateHttpMonitorDetailsProtocolEnum Enum with underlying type: string
type CreateHttpMonitorDetailsProtocolEnum string

// Set of constants representing the allowable values for CreateHttpMonitorDetailsProtocol
const (
	CreateHttpMonitorDetailsProtocolHttp  CreateHttpMonitorDetailsProtocolEnum = "HTTP"
	CreateHttpMonitorDetailsProtocolHttps CreateHttpMonitorDetailsProtocolEnum = "HTTPS"
)

var mappingCreateHttpMonitorDetailsProtocol = map[string]CreateHttpMonitorDetailsProtocolEnum{
	"HTTP":  CreateHttpMonitorDetailsProtocolHttp,
	"HTTPS": CreateHttpMonitorDetailsProtocolHttps,
}

// GetCreateHttpMonitorDetailsProtocolEnumValues Enumerates the set of values for CreateHttpMonitorDetailsProtocol
func GetCreateHttpMonitorDetailsProtocolEnumValues() []CreateHttpMonitorDetailsProtocolEnum {
	values := make([]CreateHttpMonitorDetailsProtocolEnum, 0)
	for _, v := range mappingCreateHttpMonitorDetailsProtocol {
		values = append(values, v)
	}
	return values
}

// CreateHttpMonitorDetailsMethodEnum Enum with underlying type: string
type CreateHttpMonitorDetailsMethodEnum string

// Set of constants representing the allowable values for CreateHttpMonitorDetailsMethod
const (
	CreateHttpMonitorDetailsMethodGet  CreateHttpMonitorDetailsMethodEnum = "GET"
	CreateHttpMonitorDetailsMethodHead CreateHttpMonitorDetailsMethodEnum = "HEAD"
)

var mappingCreateHttpMonitorDetailsMethod = map[string]CreateHttpMonitorDetailsMethodEnum{
	"GET":  CreateHttpMonitorDetailsMethodGet,
	"HEAD": CreateHttpMonitorDetailsMethodHead,
}

// GetCreateHttpMonitorDetailsMethodEnumValues Enumerates the set of values for CreateHttpMonitorDetailsMethod
func GetCreateHttpMonitorDetailsMethodEnumValues() []CreateHttpMonitorDetailsMethodEnum {
	values := make([]CreateHttpMonitorDetailsMethodEnum, 0)
	for _, v := range mappingCreateHttpMonitorDetailsMethod {
		values = append(values, v)
	}
	return values
}
