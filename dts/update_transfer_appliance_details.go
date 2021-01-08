// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// UpdateTransferApplianceDetails The representation of UpdateTransferApplianceDetails
type UpdateTransferApplianceDetails struct {
	LifecycleState UpdateTransferApplianceDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// Expected return date from customer for the device, time portion should be zero.
	ExpectedReturnDate *common.SDKTime `mandatory:"false" json:"expectedReturnDate"`

	// Start time for the window to pickup the device from customer.
	PickupWindowStartTime *common.SDKTime `mandatory:"false" json:"pickupWindowStartTime"`

	// End time for the window to pickup the device from customer.
	PickupWindowEndTime *common.SDKTime `mandatory:"false" json:"pickupWindowEndTime"`
}

func (m UpdateTransferApplianceDetails) String() string {
	return common.PointerString(m)
}

// UpdateTransferApplianceDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferApplianceDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferApplianceDetailsLifecycleStateEnum
const (
	UpdateTransferApplianceDetailsLifecycleStatePreparing             UpdateTransferApplianceDetailsLifecycleStateEnum = "PREPARING"
	UpdateTransferApplianceDetailsLifecycleStateFinalized             UpdateTransferApplianceDetailsLifecycleStateEnum = "FINALIZED"
	UpdateTransferApplianceDetailsLifecycleStateReturnLabelRequested  UpdateTransferApplianceDetailsLifecycleStateEnum = "RETURN_LABEL_REQUESTED"
	UpdateTransferApplianceDetailsLifecycleStateReturnLabelGenerating UpdateTransferApplianceDetailsLifecycleStateEnum = "RETURN_LABEL_GENERATING"
	UpdateTransferApplianceDetailsLifecycleStateReturnLabelAvailable  UpdateTransferApplianceDetailsLifecycleStateEnum = "RETURN_LABEL_AVAILABLE"
	UpdateTransferApplianceDetailsLifecycleStateDeleted               UpdateTransferApplianceDetailsLifecycleStateEnum = "DELETED"
	UpdateTransferApplianceDetailsLifecycleStateCustomerNeverReceived UpdateTransferApplianceDetailsLifecycleStateEnum = "CUSTOMER_NEVER_RECEIVED"
	UpdateTransferApplianceDetailsLifecycleStateCancelled             UpdateTransferApplianceDetailsLifecycleStateEnum = "CANCELLED"
)

var mappingUpdateTransferApplianceDetailsLifecycleState = map[string]UpdateTransferApplianceDetailsLifecycleStateEnum{
	"PREPARING":               UpdateTransferApplianceDetailsLifecycleStatePreparing,
	"FINALIZED":               UpdateTransferApplianceDetailsLifecycleStateFinalized,
	"RETURN_LABEL_REQUESTED":  UpdateTransferApplianceDetailsLifecycleStateReturnLabelRequested,
	"RETURN_LABEL_GENERATING": UpdateTransferApplianceDetailsLifecycleStateReturnLabelGenerating,
	"RETURN_LABEL_AVAILABLE":  UpdateTransferApplianceDetailsLifecycleStateReturnLabelAvailable,
	"DELETED":                 UpdateTransferApplianceDetailsLifecycleStateDeleted,
	"CUSTOMER_NEVER_RECEIVED": UpdateTransferApplianceDetailsLifecycleStateCustomerNeverReceived,
	"CANCELLED":               UpdateTransferApplianceDetailsLifecycleStateCancelled,
}

// GetUpdateTransferApplianceDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferApplianceDetailsLifecycleStateEnum
func GetUpdateTransferApplianceDetailsLifecycleStateEnumValues() []UpdateTransferApplianceDetailsLifecycleStateEnum {
	values := make([]UpdateTransferApplianceDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferApplianceDetailsLifecycleState {
		values = append(values, v)
	}
	return values
}
