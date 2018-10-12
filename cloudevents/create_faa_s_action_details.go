// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Event Control Service API
//
// API for managing event rules and actions.
// For more information, see Overview of Events (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package cloudevents

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateFaaSActionDetails Create an action that delivers to an Oracle Functions endpoint.
type CreateFaaSActionDetails struct {

	// Whether or not this action is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The unique URI for a function hosted by Functions-as-a-Service.
	FaasFunctionUri *string `mandatory:"false" json:"faasFunctionUri"`

	// The compartment OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) which hosts the FaaS application. Applications in FaaS hold the functions and triggers.
	FaasAppCompartmentId *string `mandatory:"false" json:"faasAppCompartmentId"`
}

//GetIsEnabled returns IsEnabled
func (m CreateFaaSActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
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
