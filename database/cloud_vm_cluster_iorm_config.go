// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CloudVmClusterIormConfig Response details which has IORM Settings for this cloud VM cluster
type CloudVmClusterIormConfig struct {

	// The current config state of IORM settings for this cloud VM cluster.
	LifecycleState CloudVmClusterIormConfigLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Value for the IORM objective
	// Default is "Auto"
	Objective CloudVmClusterIormConfigObjectiveEnum `mandatory:"false" json:"objective,omitempty"`

	// Array of IORM Setting for all the database in
	// this cloud Vm Cluster
	DbPlans []CloudVmClusterDbIormConfig `mandatory:"false" json:"dbPlans"`
}

func (m CloudVmClusterIormConfig) String() string {
	return common.PointerString(m)
}

// CloudVmClusterIormConfigLifecycleStateEnum Enum with underlying type: string
type CloudVmClusterIormConfigLifecycleStateEnum string

// Set of constants representing the allowable values for CloudVmClusterIormConfigLifecycleStateEnum
const (
	CloudVmClusterIormConfigLifecycleStateBootstrapping CloudVmClusterIormConfigLifecycleStateEnum = "BOOTSTRAPPING"
	CloudVmClusterIormConfigLifecycleStateEnabled       CloudVmClusterIormConfigLifecycleStateEnum = "ENABLED"
	CloudVmClusterIormConfigLifecycleStateDisabled      CloudVmClusterIormConfigLifecycleStateEnum = "DISABLED"
	CloudVmClusterIormConfigLifecycleStateUpdating      CloudVmClusterIormConfigLifecycleStateEnum = "UPDATING"
	CloudVmClusterIormConfigLifecycleStateFailed        CloudVmClusterIormConfigLifecycleStateEnum = "FAILED"
)

var mappingCloudVmClusterIormConfigLifecycleState = map[string]CloudVmClusterIormConfigLifecycleStateEnum{
	"BOOTSTRAPPING": CloudVmClusterIormConfigLifecycleStateBootstrapping,
	"ENABLED":       CloudVmClusterIormConfigLifecycleStateEnabled,
	"DISABLED":      CloudVmClusterIormConfigLifecycleStateDisabled,
	"UPDATING":      CloudVmClusterIormConfigLifecycleStateUpdating,
	"FAILED":        CloudVmClusterIormConfigLifecycleStateFailed,
}

// GetCloudVmClusterIormConfigLifecycleStateEnumValues Enumerates the set of values for CloudVmClusterIormConfigLifecycleStateEnum
func GetCloudVmClusterIormConfigLifecycleStateEnumValues() []CloudVmClusterIormConfigLifecycleStateEnum {
	values := make([]CloudVmClusterIormConfigLifecycleStateEnum, 0)
	for _, v := range mappingCloudVmClusterIormConfigLifecycleState {
		values = append(values, v)
	}
	return values
}

// CloudVmClusterIormConfigObjectiveEnum Enum with underlying type: string
type CloudVmClusterIormConfigObjectiveEnum string

// Set of constants representing the allowable values for CloudVmClusterIormConfigObjectiveEnum
const (
	CloudVmClusterIormConfigObjectiveLowLatency     CloudVmClusterIormConfigObjectiveEnum = "LOW_LATENCY"
	CloudVmClusterIormConfigObjectiveHighThroughput CloudVmClusterIormConfigObjectiveEnum = "HIGH_THROUGHPUT"
	CloudVmClusterIormConfigObjectiveBalanced       CloudVmClusterIormConfigObjectiveEnum = "BALANCED"
	CloudVmClusterIormConfigObjectiveAuto           CloudVmClusterIormConfigObjectiveEnum = "AUTO"
	CloudVmClusterIormConfigObjectiveBasic          CloudVmClusterIormConfigObjectiveEnum = "BASIC"
)

var mappingCloudVmClusterIormConfigObjective = map[string]CloudVmClusterIormConfigObjectiveEnum{
	"LOW_LATENCY":     CloudVmClusterIormConfigObjectiveLowLatency,
	"HIGH_THROUGHPUT": CloudVmClusterIormConfigObjectiveHighThroughput,
	"BALANCED":        CloudVmClusterIormConfigObjectiveBalanced,
	"AUTO":            CloudVmClusterIormConfigObjectiveAuto,
	"BASIC":           CloudVmClusterIormConfigObjectiveBasic,
}

// GetCloudVmClusterIormConfigObjectiveEnumValues Enumerates the set of values for CloudVmClusterIormConfigObjectiveEnum
func GetCloudVmClusterIormConfigObjectiveEnumValues() []CloudVmClusterIormConfigObjectiveEnum {
	values := make([]CloudVmClusterIormConfigObjectiveEnum, 0)
	for _, v := range mappingCloudVmClusterIormConfigObjective {
		values = append(values, v)
	}
	return values
}
