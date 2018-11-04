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

// CreateFaaSActionDetails Create an action that delivers to an Oracle Functions Service endpoint.
type CreateFaaSActionDetails struct {

	// Whether or not this action is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The unique URI for a function hosted by Oracle Functions Service.
	FaasFunctionUri *string `mandatory:"false" json:"faasFunctionUri"`

	// The compartment OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) which hosts the
	// Oracle Functions Service application. Applications in Functions hold the functions and triggers.
	FaasAppCompartmentId *string `mandatory:"false" json:"faasAppCompartmentId"`
}

//GetIsEnabled returns IsEnabled
func (m CreateFaaSActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m CreateFaaSActionDetails) GetDescription() *string {
	return m.Description
}

func (m CreateFaaSActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateFaaSActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateFaaSActionDetails CreateFaaSActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateFaaSActionDetails
	}{
		"FAAS",
		(MarshalTypeCreateFaaSActionDetails)(m),
	}

	return json.Marshal(&s)
}
