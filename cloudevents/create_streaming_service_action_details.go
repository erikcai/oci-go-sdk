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

// CreateStreamingServiceActionDetails Create an action that delivers to an Oracle Stream Service stream.
type CreateStreamingServiceActionDetails struct {

	// Whether or not this action is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the stream to which messages are delivered.
	StreamId *string `mandatory:"true" json:"streamId"`
}

//GetIsEnabled returns IsEnabled
func (m CreateStreamingServiceActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m CreateStreamingServiceActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateStreamingServiceActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateStreamingServiceActionDetails CreateStreamingServiceActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateStreamingServiceActionDetails
	}{
		"OSS",
		(MarshalTypeCreateStreamingServiceActionDetails)(m),
	}

	return json.Marshal(&s)
}
