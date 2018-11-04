// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CloudEvents API
//
// API for the CloudEvents Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package cloudevents

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// FaaSAction An action that delivers to an Oracle Functions endpoint.
type FaaSAction struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the action.
	Id *string `mandatory:"true" json:"id"`

	// A message generated by the CloudEvents service about the current state of this rule.
	LifecycleMessage *string `mandatory:"true" json:"lifecycleMessage"`

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The unique URI for a function hosted by Oracle Functions Service.
	FaasFunctionUri *string `mandatory:"false" json:"faasFunctionUri"`

	// The compartment OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) which hosts the
	// Oracle Functions Service application. Applications in Functions hold the functions and triggers.
	FaasAppCompartmentId *string `mandatory:"false" json:"faasAppCompartmentId"`

	// The current state of the rule.
	LifecycleState ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m FaaSAction) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m FaaSAction) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m FaaSAction) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

//GetDescription returns Description
func (m FaaSAction) GetDescription() *string {
	return m.Description
}

func (m FaaSAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FaaSAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFaaSAction FaaSAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeFaaSAction
	}{
		"FAAS",
		(MarshalTypeFaaSAction)(m),
	}

	return json.Marshal(&s)
}
