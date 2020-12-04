// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// FleetMetricDefinition Database metric details.
type FleetMetricDefinition struct {

	// Name of the metric.
	MetricName *string `mandatory:"false" json:"metricName"`

	// Value of the metric.
	BaselineValue *float64 `mandatory:"false" json:"baselineValue"`

	// Value of the metric.
	TargetValue *float64 `mandatory:"false" json:"targetValue"`

	// Unit of the value.
	Unit *string `mandatory:"false" json:"unit"`

	// Data point date time in ISO UTC format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// Percentage change of value compare to baseline value.
	PercentageChange *float64 `mandatory:"false" json:"percentageChange"`

	// Dimensions of the base metric, this contains child metrics of the base metric.
	Dimensions []MetricDimensionDefinition `mandatory:"false" json:"dimensions"`
}

func (m FleetMetricDefinition) String() string {
	return common.PointerString(m)
}
