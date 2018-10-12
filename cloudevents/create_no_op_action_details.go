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

// CreateNoOpActionDetails Create an action that writes to a log.
type CreateNoOpActionDetails struct {

	// Whether or not this action is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// This string is meant for internal testing.
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
