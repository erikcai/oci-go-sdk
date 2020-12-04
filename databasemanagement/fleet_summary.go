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

// FleetSummary This provides summary of metrics, which describes aggregated usage of cpu, storage so on so forth across fleet and databases inventory counts grouped by database type and subtype.
type FleetSummary struct {

	// List of databases present in the fleet and their usage metrics.
	AggregatedMetrics []FleetMetricSummaryDefinition `mandatory:"false" json:"aggregatedMetrics"`

	// Provides list of fleet database counts grouped by database type and sub type
	Inventory []FleetStatusByCategory `mandatory:"false" json:"inventory"`
}

func (m FleetSummary) String() string {
	return common.PointerString(m)
}
