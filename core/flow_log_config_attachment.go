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

// FlowLogConfigAttachment Represents an attachment between a flow log configuration and a resource such as a subnet. By
// creating a `FlowLogConfigAttachment`, you turn on flow logs for the attached resource. See
// CreateFlowLogConfigAttachment.
// **Warning:** Oracle recommends that you avoid using any confidential information when you
// supply string values using the API.
type FlowLogConfigAttachment struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the
	// flow log configuration attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The flow log configuration attachment's OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The flow log configuration attachment's current state.
	LifecycleState FlowLogConfigAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the resource that the flow log
	// configuration is attached to.
	TargetEntityId *string `mandatory:"true" json:"targetEntityId"`

	// The type of resource that the flow log configuration is attached to.
	TargetEntityType FlowLogConfigAttachmentTargetEntityTypeEnum `mandatory:"true" json:"targetEntityType"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the flow log configuration that
	// is attached to the resource.
	FlowLogConfigId *string `mandatory:"true" json:"flowLogConfigId"`

	// The date and time the flow log configuration attachment was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m FlowLogConfigAttachment) String() string {
	return common.PointerString(m)
}

// FlowLogConfigAttachmentLifecycleStateEnum Enum with underlying type: string
type FlowLogConfigAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for FlowLogConfigAttachmentLifecycleStateEnum
const (
	FlowLogConfigAttachmentLifecycleStateProvisioning FlowLogConfigAttachmentLifecycleStateEnum = "PROVISIONING"
	FlowLogConfigAttachmentLifecycleStateAvailable    FlowLogConfigAttachmentLifecycleStateEnum = "AVAILABLE"
	FlowLogConfigAttachmentLifecycleStateTerminating  FlowLogConfigAttachmentLifecycleStateEnum = "TERMINATING"
	FlowLogConfigAttachmentLifecycleStateTerminated   FlowLogConfigAttachmentLifecycleStateEnum = "TERMINATED"
)

var mappingFlowLogConfigAttachmentLifecycleState = map[string]FlowLogConfigAttachmentLifecycleStateEnum{
	"PROVISIONING": FlowLogConfigAttachmentLifecycleStateProvisioning,
	"AVAILABLE":    FlowLogConfigAttachmentLifecycleStateAvailable,
	"TERMINATING":  FlowLogConfigAttachmentLifecycleStateTerminating,
	"TERMINATED":   FlowLogConfigAttachmentLifecycleStateTerminated,
}

// GetFlowLogConfigAttachmentLifecycleStateEnumValues Enumerates the set of values for FlowLogConfigAttachmentLifecycleStateEnum
func GetFlowLogConfigAttachmentLifecycleStateEnumValues() []FlowLogConfigAttachmentLifecycleStateEnum {
	values := make([]FlowLogConfigAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingFlowLogConfigAttachmentLifecycleState {
		values = append(values, v)
	}
	return values
}

// FlowLogConfigAttachmentTargetEntityTypeEnum Enum with underlying type: string
type FlowLogConfigAttachmentTargetEntityTypeEnum string

// Set of constants representing the allowable values for FlowLogConfigAttachmentTargetEntityTypeEnum
const (
	FlowLogConfigAttachmentTargetEntityTypeSubnet FlowLogConfigAttachmentTargetEntityTypeEnum = "SUBNET"
)

var mappingFlowLogConfigAttachmentTargetEntityType = map[string]FlowLogConfigAttachmentTargetEntityTypeEnum{
	"SUBNET": FlowLogConfigAttachmentTargetEntityTypeSubnet,
}

// GetFlowLogConfigAttachmentTargetEntityTypeEnumValues Enumerates the set of values for FlowLogConfigAttachmentTargetEntityTypeEnum
func GetFlowLogConfigAttachmentTargetEntityTypeEnumValues() []FlowLogConfigAttachmentTargetEntityTypeEnum {
	values := make([]FlowLogConfigAttachmentTargetEntityTypeEnum, 0)
	for _, v := range mappingFlowLogConfigAttachmentTargetEntityType {
		values = append(values, v)
	}
	return values
}
