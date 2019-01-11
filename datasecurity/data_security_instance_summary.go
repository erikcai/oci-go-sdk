// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Security Control Plane API
//
// The API to manage data security instance creation and deletion
//

package datasecurity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DataSecurityInstanceSummary Summary of the data security instance.
type DataSecurityInstanceSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Data security instance name, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Service URL of the data security instance
	Url *string `mandatory:"false" json:"url"`

	// Description of the data security instance
	Description *string `mandatory:"false" json:"description"`

	// The time the the data security instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the data security instance.
	LifecycleState DataSecurityInstanceSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m DataSecurityInstanceSummary) String() string {
	return common.PointerString(m)
}

// DataSecurityInstanceSummaryLifecycleStateEnum Enum with underlying type: string
type DataSecurityInstanceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DataSecurityInstanceSummaryLifecycleStateEnum
const (
	DataSecurityInstanceSummaryLifecycleStateCreating DataSecurityInstanceSummaryLifecycleStateEnum = "CREATING"
	DataSecurityInstanceSummaryLifecycleStateUpdating DataSecurityInstanceSummaryLifecycleStateEnum = "UPDATING"
	DataSecurityInstanceSummaryLifecycleStateActive   DataSecurityInstanceSummaryLifecycleStateEnum = "ACTIVE"
	DataSecurityInstanceSummaryLifecycleStateDeleting DataSecurityInstanceSummaryLifecycleStateEnum = "DELETING"
	DataSecurityInstanceSummaryLifecycleStateDeleted  DataSecurityInstanceSummaryLifecycleStateEnum = "DELETED"
	DataSecurityInstanceSummaryLifecycleStateFailed   DataSecurityInstanceSummaryLifecycleStateEnum = "FAILED"
)

var mappingDataSecurityInstanceSummaryLifecycleState = map[string]DataSecurityInstanceSummaryLifecycleStateEnum{
	"CREATING": DataSecurityInstanceSummaryLifecycleStateCreating,
	"UPDATING": DataSecurityInstanceSummaryLifecycleStateUpdating,
	"ACTIVE":   DataSecurityInstanceSummaryLifecycleStateActive,
	"DELETING": DataSecurityInstanceSummaryLifecycleStateDeleting,
	"DELETED":  DataSecurityInstanceSummaryLifecycleStateDeleted,
	"FAILED":   DataSecurityInstanceSummaryLifecycleStateFailed,
}

// GetDataSecurityInstanceSummaryLifecycleStateEnumValues Enumerates the set of values for DataSecurityInstanceSummaryLifecycleStateEnum
func GetDataSecurityInstanceSummaryLifecycleStateEnumValues() []DataSecurityInstanceSummaryLifecycleStateEnum {
	values := make([]DataSecurityInstanceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDataSecurityInstanceSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
