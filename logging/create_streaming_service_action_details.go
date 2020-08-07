// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateStreamingServiceActionDetails Create an action that delivers to an Oracle Stream Service stream.
type CreateStreamingServiceActionDetails struct {

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream to which messages are delivered.
	StreamId *string `mandatory:"true" json:"streamId"`

	// The OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`
}

//GetId returns Id
func (m CreateStreamingServiceActionDetails) GetId() *string {
	return m.Id
}

//GetIsEnabled returns IsEnabled
func (m CreateStreamingServiceActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m CreateStreamingServiceActionDetails) GetDescription() *string {
	return m.Description
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
