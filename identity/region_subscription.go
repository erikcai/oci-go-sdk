// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RegionSubscription The representation of RegionSubscription
type RegionSubscription struct {

	// The key of the region such as PHX, IAD.
	RegionKey *string `mandatory:"true" json:"regionKey"`

	// The name of the region such as us-phoenix-1.
	RegionName *string `mandatory:"true" json:"regionName"`

	// The region subscription status such as Ready, InProgress
	Status RegionSubscriptionStatusEnum `mandatory:"true" json:"status"`

	// Indicates the region is home region or not.
	IsHomeRegion *bool `mandatory:"true" json:"isHomeRegion"`
}

func (m RegionSubscription) String() string {
	return common.PointerString(m)
}

// RegionSubscriptionStatusEnum Enum with underlying type: string
type RegionSubscriptionStatusEnum string

// Set of constants representing the allowable values for RegionSubscriptionStatus
const (
	RegionSubscriptionStatusReady      RegionSubscriptionStatusEnum = "READY"
	RegionSubscriptionStatusInProgress RegionSubscriptionStatusEnum = "IN_PROGRESS"
)

var mappingRegionSubscriptionStatus = map[string]RegionSubscriptionStatusEnum{
	"READY":       RegionSubscriptionStatusReady,
	"IN_PROGRESS": RegionSubscriptionStatusInProgress,
}

// GetRegionSubscriptionStatusEnumValues Enumerates the set of values for RegionSubscriptionStatus
func GetRegionSubscriptionStatusEnumValues() []RegionSubscriptionStatusEnum {
	values := make([]RegionSubscriptionStatusEnum, 0)
	for _, v := range mappingRegionSubscriptionStatus {
		values = append(values, v)
	}
	return values
}
