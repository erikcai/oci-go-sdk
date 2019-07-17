// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DTS API
//
// A description of the DTS API
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateTransferPackageDetails The representation of UpdateTransferPackageDetails
type UpdateTransferPackageDetails struct {
	OriginalPackageDeliveryTrackingNumber *string `mandatory:"false" json:"originalPackageDeliveryTrackingNumber"`

	ReturnPackageDeliveryTrackingNumber *string `mandatory:"false" json:"returnPackageDeliveryTrackingNumber"`

	PackageDeliveryVendor *string `mandatory:"false" json:"packageDeliveryVendor"`

	LifecycleState UpdateTransferPackageDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m UpdateTransferPackageDetails) String() string {
	return common.PointerString(m)
}

// UpdateTransferPackageDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferPackageDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferPackageDetailsLifecycleStateEnum
const (
	UpdateTransferPackageDetailsLifecycleStateShipping  UpdateTransferPackageDetailsLifecycleStateEnum = "SHIPPING"
	UpdateTransferPackageDetailsLifecycleStateCancelled UpdateTransferPackageDetailsLifecycleStateEnum = "CANCELLED"
)

var mappingUpdateTransferPackageDetailsLifecycleState = map[string]UpdateTransferPackageDetailsLifecycleStateEnum{
	"SHIPPING":  UpdateTransferPackageDetailsLifecycleStateShipping,
	"CANCELLED": UpdateTransferPackageDetailsLifecycleStateCancelled,
}

// GetUpdateTransferPackageDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferPackageDetailsLifecycleStateEnum
func GetUpdateTransferPackageDetailsLifecycleStateEnumValues() []UpdateTransferPackageDetailsLifecycleStateEnum {
	values := make([]UpdateTransferPackageDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferPackageDetailsLifecycleState {
		values = append(values, v)
	}
	return values
}
