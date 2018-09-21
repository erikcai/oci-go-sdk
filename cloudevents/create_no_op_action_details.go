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

// CreateNoOpActionDetails delivery to logs
type CreateNoOpActionDetails struct {

	// whether or not this aciton is currently enabled
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// name of the noop action; will appear in the logs
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetIsEnabled returns IsEnabled
func (m CreateNoOpActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m CreateNoOpActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateNoOpActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateNoOpActionDetails CreateNoOpActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateNoOpActionDetails
	}{
		"NOOP",
		(MarshalTypeCreateNoOpActionDetails)(m),
	}

	return json.Marshal(&s)
}
