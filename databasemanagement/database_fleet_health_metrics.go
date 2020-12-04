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

// DatabaseFleetHealthMetrics Fleet health metrics
type DatabaseFleetHealthMetrics struct {

	// Baseline datetime in ISO 8601 format which is "yyyy-mm-dd'T'hh:mm:ss'Z'".
	// This is datetime against which percentage change is calculated
	CompareBaselineTime *string `mandatory:"true" json:"compareBaselineTime"`

	// Datetime format in ISO 8601 which is "yyyy-mm-dd'T'hh:mm:ss'Z'". All metrics are returned for this datetime and
	// percentage change is calculated against baseline datetime
	CompareTargetTime *string `mandatory:"true" json:"compareTargetTime"`

	// List of databases present in the fleet and their usage metrics.
	FleetDatabases []DatabaseUsageMetrics `mandatory:"true" json:"fleetDatabases"`

	// The current state of the Managed Database Group.
	CompareType CompareTypeEnum `mandatory:"false" json:"compareType,omitempty"`

	FleetSummary *FleetSummary `mandatory:"false" json:"fleetSummary"`
}

func (m DatabaseFleetHealthMetrics) String() string {
	return common.PointerString(m)
}
