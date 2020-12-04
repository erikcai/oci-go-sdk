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

// DatabaseUsageMetrics Database metric aggregated values
type DatabaseUsageMetrics struct {

	// ocid of the database asset.
	DbId *string `mandatory:"false" json:"dbId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The type of Oracle database installation.
	DatabaseType DatabaseTypeEnum `mandatory:"false" json:"databaseType,omitempty"`

	// Indicates whether an Oracle database is a Container/Pluggable/Non-Container database.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	// Display name of the database.
	DatabaseName *string `mandatory:"false" json:"databaseName"`

	// If the database is a Pluggable database then this value indicates its managed Container database ocid
	DatabaseContainerId *string `mandatory:"false" json:"databaseContainerId"`

	// List of fleet databases health metrics like cpu, storage and memory so on so forth.
	Metrics []FleetMetricDefinition `mandatory:"false" json:"metrics"`
}

func (m DatabaseUsageMetrics) String() string {
	return common.PointerString(m)
}
