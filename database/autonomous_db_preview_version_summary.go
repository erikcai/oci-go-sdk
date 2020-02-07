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

// AutonomousDbPreviewVersionSummary The Autonomous Database preview version. Note that preview version software is only available for databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI).
type AutonomousDbPreviewVersionSummary struct {

	// A valid Autonomous Database preview version.
	Version *string `mandatory:"true" json:"version"`

	// The date and time when the preview version availability begins.
	TimePreviewBegin *common.SDKTime `mandatory:"false" json:"timePreviewBegin"`

	// The date and time when the preview version availability ends.
	TimePreviewEnd *common.SDKTime `mandatory:"false" json:"timePreviewEnd"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	DbWorkload AutonomousDbPreviewVersionSummaryDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// A URL that points to a detailed description of the preview version.
	Details *string `mandatory:"false" json:"details"`
}

func (m AutonomousDbPreviewVersionSummary) String() string {
	return common.PointerString(m)
}

// AutonomousDbPreviewVersionSummaryDbWorkloadEnum Enum with underlying type: string
type AutonomousDbPreviewVersionSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDbPreviewVersionSummaryDbWorkloadEnum
const (
	AutonomousDbPreviewVersionSummaryDbWorkloadOltp AutonomousDbPreviewVersionSummaryDbWorkloadEnum = "OLTP"
	AutonomousDbPreviewVersionSummaryDbWorkloadDw   AutonomousDbPreviewVersionSummaryDbWorkloadEnum = "DW"
)

var mappingAutonomousDbPreviewVersionSummaryDbWorkload = map[string]AutonomousDbPreviewVersionSummaryDbWorkloadEnum{
	"OLTP": AutonomousDbPreviewVersionSummaryDbWorkloadOltp,
	"DW":   AutonomousDbPreviewVersionSummaryDbWorkloadDw,
}

// GetAutonomousDbPreviewVersionSummaryDbWorkloadEnumValues Enumerates the set of values for AutonomousDbPreviewVersionSummaryDbWorkloadEnum
func GetAutonomousDbPreviewVersionSummaryDbWorkloadEnumValues() []AutonomousDbPreviewVersionSummaryDbWorkloadEnum {
	values := make([]AutonomousDbPreviewVersionSummaryDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDbPreviewVersionSummaryDbWorkload {
		values = append(values, v)
	}
	return values
}
