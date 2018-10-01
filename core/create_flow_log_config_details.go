// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateFlowLogConfigDetails The representation of CreateFlowLogConfigDetails
type CreateFlowLogConfigDetails struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// flow log configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type or types of flow logs to store. `ALL` includes records for both accepted traffic and
	// rejected traffic.
	FlowLogType CreateFlowLogConfigDetailsFlowLogTypeEnum `mandatory:"true" json:"flowLogType"`

	// Where to store the flow logs.
	Destination FlowLogDestination `mandatory:"true" json:"destination"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateFlowLogConfigDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateFlowLogConfigDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags   map[string]map[string]interface{}         `json:"definedTags"`
		DisplayName   *string                                   `json:"displayName"`
		FreeformTags  map[string]string                         `json:"freeformTags"`
		CompartmentId *string                                   `json:"compartmentId"`
		FlowLogType   CreateFlowLogConfigDetailsFlowLogTypeEnum `json:"flowLogType"`
		Destination   flowlogdestination                        `json:"destination"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DefinedTags = model.DefinedTags
	m.DisplayName = model.DisplayName
	m.FreeformTags = model.FreeformTags
	m.CompartmentId = model.CompartmentId
	m.FlowLogType = model.FlowLogType
	nn, e := model.Destination.UnmarshalPolymorphicJSON(model.Destination.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Destination = nn.(FlowLogDestination)
	} else {
		m.Destination = nil
	}
	return
}

// CreateFlowLogConfigDetailsFlowLogTypeEnum Enum with underlying type: string
type CreateFlowLogConfigDetailsFlowLogTypeEnum string

// Set of constants representing the allowable values for CreateFlowLogConfigDetailsFlowLogTypeEnum
const (
	CreateFlowLogConfigDetailsFlowLogTypeAll CreateFlowLogConfigDetailsFlowLogTypeEnum = "ALL"
)

var mappingCreateFlowLogConfigDetailsFlowLogType = map[string]CreateFlowLogConfigDetailsFlowLogTypeEnum{
	"ALL": CreateFlowLogConfigDetailsFlowLogTypeAll,
}

// GetCreateFlowLogConfigDetailsFlowLogTypeEnumValues Enumerates the set of values for CreateFlowLogConfigDetailsFlowLogTypeEnum
func GetCreateFlowLogConfigDetailsFlowLogTypeEnumValues() []CreateFlowLogConfigDetailsFlowLogTypeEnum {
	values := make([]CreateFlowLogConfigDetailsFlowLogTypeEnum, 0)
	for _, v := range mappingCreateFlowLogConfigDetailsFlowLogType {
		values = append(values, v)
	}
	return values
}
