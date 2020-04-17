// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateMetricsActionDetails Create an action that delivers metrics.
type CreateMetricsActionDetails struct {

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Compartment id of the resource.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// namespace of the metric
	MetricNamespace *string `mandatory:"false" json:"metricNamespace"`

	// name of the metric
	MetricName *string `mandatory:"false" json:"metricName"`
}

//GetId returns Id
func (m CreateMetricsActionDetails) GetId() *string {
	return m.Id
}

//GetIsEnabled returns IsEnabled
func (m CreateMetricsActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m CreateMetricsActionDetails) GetDescription() *string {
	return m.Description
}

func (m CreateMetricsActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateMetricsActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMetricsActionDetails CreateMetricsActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateMetricsActionDetails
	}{
		"METRICS",
		(MarshalTypeCreateMetricsActionDetails)(m),
	}

	return json.Marshal(&s)
}
