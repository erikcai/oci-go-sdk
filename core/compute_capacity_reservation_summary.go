// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm),
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm), and
// Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as virtual cloud networks (VCNs),
// compute instances, block storage volumes, and container images.
//

package core

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// ComputeCapacityReservationSummary Summary information for a compute capacity reservation.
type ComputeCapacityReservationSummary struct {

	// The OCID of the instance reservation configuration.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the capacity reservation was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name for the capacity reservation. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My Reservation`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the capacity reservation.
	LifecycleState ComputeCapacityReservationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The availability domain of the capacity reservation.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The number of instances for which capacity will be held in this
	// compute capacity reservation. This number is the sum of the values of the `reservedCount` fields
	// for all of the instance reservation configurations under this reservation.
	// The purpose of this field is to calculate the percentage usage of the reservation.
	ReservedInstanceCount *int64 `mandatory:"false" json:"reservedInstanceCount"`

	// The total number of instances currently consuming space in
	// this compute capacity reservation. This number is the sum of the values of the `usedCount` fields
	// for all of the instance reservation configurations under this reservation.
	// The purpose of this field is to calculate the percentage usage of the reservation.
	UsedInstanceCount *int64 `mandatory:"false" json:"usedInstanceCount"`

	// Whether this capacity reservation is the default.
	// For more information, see Capacity Reservations (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
	IsDefaultReservation *bool `mandatory:"false" json:"isDefaultReservation"`
}

func (m ComputeCapacityReservationSummary) String() string {
	return common.PointerString(m)
}
