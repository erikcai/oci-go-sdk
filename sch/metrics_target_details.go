// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connectors API
//
// A description of the Connectors API
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// MetricsTargetDetails The metrics target.
type MetricsTargetDetails struct {

	// The compartment OCID of the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The namespace of the metric.
	MetricNamespace *string `mandatory:"true" json:"metricNamespace"`

	// The name of the metric.
	MetricName *string `mandatory:"true" json:"metricName"`
}

func (m MetricsTargetDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m MetricsTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMetricsTargetDetails MetricsTargetDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeMetricsTargetDetails
	}{
		"metrics",
		(MarshalTypeMetricsTargetDetails)(m),
	}

	return json.Marshal(&s)
}
