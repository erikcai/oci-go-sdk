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

// CreateFlowLogConfigAttachmentDetails The representation of CreateFlowLogConfigAttachmentDetails
type CreateFlowLogConfigAttachmentDetails struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the resource to attach the
	// flow log configuration to. Attaching the configuration enables flow logs for the resource.
	TargetEntityId *string `mandatory:"true" json:"targetEntityId"`

	// The type of resource to attach the flow log configuration to.
	TargetEntityType CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum `mandatory:"true" json:"targetEntityType"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the flow log configuration to attach.
	FlowLogConfigId *string `mandatory:"true" json:"flowLogConfigId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateFlowLogConfigAttachmentDetails) String() string {
	return common.PointerString(m)
}

// CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum Enum with underlying type: string
type CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum string

// Set of constants representing the allowable values for CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum
const (
	CreateFlowLogConfigAttachmentDetailsTargetEntityTypeSubnet CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum = "SUBNET"
)

var mappingCreateFlowLogConfigAttachmentDetailsTargetEntityType = map[string]CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum{
	"SUBNET": CreateFlowLogConfigAttachmentDetailsTargetEntityTypeSubnet,
}

// GetCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumValues Enumerates the set of values for CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum
func GetCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumValues() []CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum {
	values := make([]CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum, 0)
	for _, v := range mappingCreateFlowLogConfigAttachmentDetailsTargetEntityType {
		values = append(values, v)
	}
	return values
}
