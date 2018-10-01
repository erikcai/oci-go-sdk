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

// UpdatePingMonitorDetails The request body used to update a Ping monitor.
type UpdatePingMonitorDetails struct {
	Targets []string `mandatory:"false" json:"targets"`

	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The probe port, defaults to the default port for the specified protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Legal values are: 10, 20, 30, and 60. The probe timeout
	// must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	Protocol UpdatePingMonitorDetailsProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	// A mutable name suitable for display in a user interface.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The monitor interval in seconds. Legal values are: 10, 30, and 60.
	IntervalInSeconds *int `mandatory:"false" json:"intervalInSeconds"`

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

func (m UpdatePingMonitorDetails) String() string {
	return common.PointerString(m)
}

// UpdatePingMonitorDetailsProtocolEnum Enum with underlying type: string
type UpdatePingMonitorDetailsProtocolEnum string

// Set of constants representing the allowable values for UpdatePingMonitorDetailsProtocolEnum
const (
	UpdatePingMonitorDetailsProtocolIcmp UpdatePingMonitorDetailsProtocolEnum = "ICMP"
	UpdatePingMonitorDetailsProtocolTcp  UpdatePingMonitorDetailsProtocolEnum = "TCP"
)

var mappingUpdatePingMonitorDetailsProtocol = map[string]UpdatePingMonitorDetailsProtocolEnum{
	"ICMP": UpdatePingMonitorDetailsProtocolIcmp,
	"TCP":  UpdatePingMonitorDetailsProtocolTcp,
}

// GetUpdatePingMonitorDetailsProtocolEnumValues Enumerates the set of values for UpdatePingMonitorDetailsProtocolEnum
func GetUpdatePingMonitorDetailsProtocolEnumValues() []UpdatePingMonitorDetailsProtocolEnum {
	values := make([]UpdatePingMonitorDetailsProtocolEnum, 0)
	for _, v := range mappingUpdatePingMonitorDetailsProtocol {
		values = append(values, v)
	}
	return values
}
