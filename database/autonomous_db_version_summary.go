// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDbVersionSummary The supported Autonomous Database version.
type AutonomousDbVersionSummary struct {

	// A valid Oracle Database version for Autonomous Database.
	Version *string `mandatory:"true" json:"version"`

	// The Autonomous Database workload type. OLTP indicates an Autonomous Transaction Processing database and DW indicates an Autonomous Data Warehouse database.
	DbWorkload AutonomousDbVersionSummaryDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// True if the database uses dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// A URL that points to a detailed description of the Autonomous Database version.
	Details *string `mandatory:"false" json:"details"`
}

func (m AutonomousDbVersionSummary) String() string {
	return common.PointerString(m)
}

// AutonomousDbVersionSummaryDbWorkloadEnum Enum with underlying type: string
type AutonomousDbVersionSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDbVersionSummaryDbWorkloadEnum
const (
	AutonomousDbVersionSummaryDbWorkloadOltp AutonomousDbVersionSummaryDbWorkloadEnum = "OLTP"
	AutonomousDbVersionSummaryDbWorkloadDw   AutonomousDbVersionSummaryDbWorkloadEnum = "DW"
)

var mappingAutonomousDbVersionSummaryDbWorkload = map[string]AutonomousDbVersionSummaryDbWorkloadEnum{
	"OLTP": AutonomousDbVersionSummaryDbWorkloadOltp,
	"DW":   AutonomousDbVersionSummaryDbWorkloadDw,
}

// GetAutonomousDbVersionSummaryDbWorkloadEnumValues Enumerates the set of values for AutonomousDbVersionSummaryDbWorkloadEnum
func GetAutonomousDbVersionSummaryDbWorkloadEnumValues() []AutonomousDbVersionSummaryDbWorkloadEnum {
	values := make([]AutonomousDbVersionSummaryDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDbVersionSummaryDbWorkload {
		values = append(values, v)
	}
	return values
}
