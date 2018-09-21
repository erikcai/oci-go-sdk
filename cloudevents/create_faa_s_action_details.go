// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// EventsControlService API
//
// This service exposes APIs to create, update and delete Rules. Rules are used to tap into the Events stream.
//

package cloudevents

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateFaaSActionDetails Deliver to a Function-as-a-service endpoint
type CreateFaaSActionDetails struct {

	// whether or not this aciton is currently enabled
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// OCID of the function to call
	FunctionId *string `mandatory:"false" json:"functionId"`
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
