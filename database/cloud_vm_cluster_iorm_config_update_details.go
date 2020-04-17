// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CloudVmClusterIormConfigUpdateDetails IORM Setting details for this cloud Vm Cluster to be updated
type CloudVmClusterIormConfigUpdateDetails struct {

	// Value for the IORM objective
	// Default is "Auto"
	Objective CloudVmClusterIormConfigUpdateDetailsObjectiveEnum `mandatory:"false" json:"objective,omitempty"`

	// Array of IORM Setting for all the database in
	// this Exadata DB System
	DbPlans []CloudVmClusterDbIormConfigUpdateDetail `mandatory:"false" json:"dbPlans"`
}

func (m CloudVmClusterIormConfigUpdateDetails) String() string {
	return common.PointerString(m)
}

// CloudVmClusterIormConfigUpdateDetailsObjectiveEnum Enum with underlying type: string
type CloudVmClusterIormConfigUpdateDetailsObjectiveEnum string

// Set of constants representing the allowable values for CloudVmClusterIormConfigUpdateDetailsObjectiveEnum
const (
	CloudVmClusterIormConfigUpdateDetailsObjectiveLowLatency     CloudVmClusterIormConfigUpdateDetailsObjectiveEnum = "LOW_LATENCY"
	CloudVmClusterIormConfigUpdateDetailsObjectiveHighThroughput CloudVmClusterIormConfigUpdateDetailsObjectiveEnum = "HIGH_THROUGHPUT"
	CloudVmClusterIormConfigUpdateDetailsObjectiveBalanced       CloudVmClusterIormConfigUpdateDetailsObjectiveEnum = "BALANCED"
	CloudVmClusterIormConfigUpdateDetailsObjectiveAuto           CloudVmClusterIormConfigUpdateDetailsObjectiveEnum = "AUTO"
	CloudVmClusterIormConfigUpdateDetailsObjectiveBasic          CloudVmClusterIormConfigUpdateDetailsObjectiveEnum = "BASIC"
)

var mappingCloudVmClusterIormConfigUpdateDetailsObjective = map[string]CloudVmClusterIormConfigUpdateDetailsObjectiveEnum{
	"LOW_LATENCY":     CloudVmClusterIormConfigUpdateDetailsObjectiveLowLatency,
	"HIGH_THROUGHPUT": CloudVmClusterIormConfigUpdateDetailsObjectiveHighThroughput,
	"BALANCED":        CloudVmClusterIormConfigUpdateDetailsObjectiveBalanced,
	"AUTO":            CloudVmClusterIormConfigUpdateDetailsObjectiveAuto,
	"BASIC":           CloudVmClusterIormConfigUpdateDetailsObjectiveBasic,
}

// GetCloudVmClusterIormConfigUpdateDetailsObjectiveEnumValues Enumerates the set of values for CloudVmClusterIormConfigUpdateDetailsObjectiveEnum
func GetCloudVmClusterIormConfigUpdateDetailsObjectiveEnumValues() []CloudVmClusterIormConfigUpdateDetailsObjectiveEnum {
	values := make([]CloudVmClusterIormConfigUpdateDetailsObjectiveEnum, 0)
	for _, v := range mappingCloudVmClusterIormConfigUpdateDetailsObjective {
		values = append(values, v)
	}
	return values
}
