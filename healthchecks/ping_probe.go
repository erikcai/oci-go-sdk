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

// PingProbe This model contains all of the mutable and immutable properties for a Ping probe.
type PingProbe struct {

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

	Protocol PingProbeProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

func (m PingProbe) String() string {
	return common.PointerString(m)
}

// PingProbeProtocolEnum Enum with underlying type: string
type PingProbeProtocolEnum string

// Set of constants representing the allowable values for PingProbeProtocolEnum
const (
	PingProbeProtocolIcmp PingProbeProtocolEnum = "ICMP"
	PingProbeProtocolTcp  PingProbeProtocolEnum = "TCP"
)

var mappingPingProbeProtocol = map[string]PingProbeProtocolEnum{
	"ICMP": PingProbeProtocolIcmp,
	"TCP":  PingProbeProtocolTcp,
}

// GetPingProbeProtocolEnumValues Enumerates the set of values for PingProbeProtocolEnum
func GetPingProbeProtocolEnumValues() []PingProbeProtocolEnum {
	values := make([]PingProbeProtocolEnum, 0)
	for _, v := range mappingPingProbeProtocol {
		values = append(values, v)
	}
	return values
}
