// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// VnicAttachment Represents an attachment between a VNIC and an instance. For more information, see
// Overview of the Compute Service (https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Concepts/computeoverview.htm).
type VnicAttachment struct {

	// The Availability Domain of an instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VNIC attachment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The current state of the VNIC attachment.
	LifecycleState VnicAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the subnet of the VNIC.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time the VNIC attachment was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name. Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Identifies the physical network port on which the VNIC is attached.
	NicIndex *int `mandatory:"false" json:"nicIndex"`

	// The VLAN tag of the attached VNIC.
	VlanTag *int `mandatory:"false" json:"vlanTag"`

	// The OCID of the VNIC once it is attached.
	VnicId *string `mandatory:"false" json:"vnicId"`
}

func (m VnicAttachment) String() string {
	return common.PointerString(m)
}

// VnicAttachmentLifecycleStateEnum Enum with underlying type: string
type VnicAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for VnicAttachmentLifecycleState
const (
	VnicAttachmentLifecycleStateAttaching VnicAttachmentLifecycleStateEnum = "ATTACHING"
	VnicAttachmentLifecycleStateAttached  VnicAttachmentLifecycleStateEnum = "ATTACHED"
	VnicAttachmentLifecycleStateDetaching VnicAttachmentLifecycleStateEnum = "DETACHING"
	VnicAttachmentLifecycleStateDetached  VnicAttachmentLifecycleStateEnum = "DETACHED"
)

var mappingVnicAttachmentLifecycleState = map[string]VnicAttachmentLifecycleStateEnum{
	"ATTACHING": VnicAttachmentLifecycleStateAttaching,
	"ATTACHED":  VnicAttachmentLifecycleStateAttached,
	"DETACHING": VnicAttachmentLifecycleStateDetaching,
	"DETACHED":  VnicAttachmentLifecycleStateDetached,
}

// GetVnicAttachmentLifecycleStateEnumValues Enumerates the set of values for VnicAttachmentLifecycleState
func GetVnicAttachmentLifecycleStateEnumValues() []VnicAttachmentLifecycleStateEnum {
	values := make([]VnicAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingVnicAttachmentLifecycleState {
		values = append(values, v)
	}
	return values
}
