// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastion

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v25/common"
)

// SessionSummary summary of session resource.
type SessionSummary struct {

	// OCID of session
	Id *string `mandatory:"true" json:"id"`

	// bastion display name that the session is created on
	BastionDisplayName *string `mandatory:"true" json:"bastionDisplayName"`

	// OCID of the bastion that the session is created on
	BastionId *string `mandatory:"true" json:"bastionId"`

	TargetResourceDetails TargetResourceDetails `mandatory:"true" json:"targetResourceDetails"`

	// The time the session was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the session.
	LifecycleState SessionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// TTL of the session.
	SessionTtlInSeconds *int `mandatory:"true" json:"sessionTtlInSeconds"`

	// The time the session was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m SessionSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *SessionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeUpdated           *common.SDKTime           `json:"timeUpdated"`
		LifecycleDetails      *string                   `json:"lifecycleDetails"`
		Id                    *string                   `json:"id"`
		BastionDisplayName    *string                   `json:"bastionDisplayName"`
		BastionId             *string                   `json:"bastionId"`
		TargetResourceDetails targetresourcedetails     `json:"targetResourceDetails"`
		TimeCreated           *common.SDKTime           `json:"timeCreated"`
		LifecycleState        SessionLifecycleStateEnum `json:"lifecycleState"`
		SessionTtlInSeconds   *int                      `json:"sessionTtlInSeconds"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.Id = model.Id

	m.BastionDisplayName = model.BastionDisplayName

	m.BastionId = model.BastionId

	nn, e = model.TargetResourceDetails.UnmarshalPolymorphicJSON(model.TargetResourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetResourceDetails = nn.(TargetResourceDetails)
	} else {
		m.TargetResourceDetails = nil
	}

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.SessionTtlInSeconds = model.SessionTtlInSeconds

	return
}
