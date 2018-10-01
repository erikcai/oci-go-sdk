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

// CreateOnDemandPingProbeDetails The request body used to create an on-demand Ping probe.
type CreateOnDemandPingProbeDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Targets []string `mandatory:"true" json:"targets"`

	Protocol CreateOnDemandPingProbeDetailsProtocolEnum `mandatory:"true" json:"protocol"`

	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The probe port, defaults to the default port for the specified protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Legal values are: 10, 20, 30, and 60. The probe timeout
	// must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`
}

func (m CreateOnDemandPingProbeDetails) String() string {
	return common.PointerString(m)
}

// CreateOnDemandPingProbeDetailsProtocolEnum Enum with underlying type: string
type CreateOnDemandPingProbeDetailsProtocolEnum string

// Set of constants representing the allowable values for CreateOnDemandPingProbeDetailsProtocolEnum
const (
	CreateOnDemandPingProbeDetailsProtocolIcmp CreateOnDemandPingProbeDetailsProtocolEnum = "ICMP"
	CreateOnDemandPingProbeDetailsProtocolTcp  CreateOnDemandPingProbeDetailsProtocolEnum = "TCP"
)

var mappingCreateOnDemandPingProbeDetailsProtocol = map[string]CreateOnDemandPingProbeDetailsProtocolEnum{
	"ICMP": CreateOnDemandPingProbeDetailsProtocolIcmp,
	"TCP":  CreateOnDemandPingProbeDetailsProtocolTcp,
}

// GetCreateOnDemandPingProbeDetailsProtocolEnumValues Enumerates the set of values for CreateOnDemandPingProbeDetailsProtocolEnum
func GetCreateOnDemandPingProbeDetailsProtocolEnumValues() []CreateOnDemandPingProbeDetailsProtocolEnum {
	values := make([]CreateOnDemandPingProbeDetailsProtocolEnum, 0)
	for _, v := range mappingCreateOnDemandPingProbeDetailsProtocol {
		values = append(values, v)
	}
	return values
}
