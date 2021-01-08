// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastion

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v31/common"
)

// Session the session resource.
type Session struct {

	// OCID of session
	Id *string `mandatory:"true" json:"id"`

	// OCID of the bastion that the session is created on
	BastionId *string `mandatory:"true" json:"bastionId"`

	// bastion display name that the session is created on
	BastionDisplayName *string `mandatory:"true" json:"bastionDisplayName"`

	TargetResourceDetails TargetResourceDetails `mandatory:"true" json:"targetResourceDetails"`

	KeyDetails *PublicKeyDetails `mandatory:"true" json:"keyDetails"`

	// The time the session was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the session.
	LifecycleState SessionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// TTL of the session.
	SessionTtlInSeconds *int `mandatory:"true" json:"sessionTtlInSeconds"`

	// session Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// name of the user to use on bastion
	BastionUserName *string `mandatory:"false" json:"bastionUserName"`

	// bastion ip address to connect from.
	BastionIp *string `mandatory:"false" json:"bastionIp"`

	// bastion port to connect from.
	BastionPort *int `mandatory:"false" json:"bastionPort"`

	// the connection message for the session.
	SshMetadata map[string]string `mandatory:"false" json:"sshMetadata"`

	// the type of the key, PUB refers to normal public key
	KeyType SessionKeyTypeEnum `mandatory:"false" json:"keyType,omitempty"`

	// bastion host Key for customer to make sure they connect the correct bastion
	BastionPublicHostKeyInfo *string `mandatory:"false" json:"bastionPublicHostKeyInfo"`

	// The time the session was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Session) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Session) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                           `json:"displayName"`
		BastionUserName          *string                           `json:"bastionUserName"`
		BastionIp                *string                           `json:"bastionIp"`
		BastionPort              *int                              `json:"bastionPort"`
		SshMetadata              map[string]string                 `json:"sshMetadata"`
		KeyType                  SessionKeyTypeEnum                `json:"keyType"`
		BastionPublicHostKeyInfo *string                           `json:"bastionPublicHostKeyInfo"`
		TimeUpdated              *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails         *string                           `json:"lifecycleDetails"`
		FreeformTags             map[string]string                 `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{} `json:"definedTags"`
		Id                       *string                           `json:"id"`
		BastionId                *string                           `json:"bastionId"`
		BastionDisplayName       *string                           `json:"bastionDisplayName"`
		TargetResourceDetails    targetresourcedetails             `json:"targetResourceDetails"`
		KeyDetails               *PublicKeyDetails                 `json:"keyDetails"`
		TimeCreated              *common.SDKTime                   `json:"timeCreated"`
		LifecycleState           SessionLifecycleStateEnum         `json:"lifecycleState"`
		SessionTtlInSeconds      *int                              `json:"sessionTtlInSeconds"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.BastionUserName = model.BastionUserName

	m.BastionIp = model.BastionIp

	m.BastionPort = model.BastionPort

	m.SshMetadata = model.SshMetadata

	m.KeyType = model.KeyType

	m.BastionPublicHostKeyInfo = model.BastionPublicHostKeyInfo

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.BastionId = model.BastionId

	m.BastionDisplayName = model.BastionDisplayName

	nn, e = model.TargetResourceDetails.UnmarshalPolymorphicJSON(model.TargetResourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetResourceDetails = nn.(TargetResourceDetails)
	} else {
		m.TargetResourceDetails = nil
	}

	m.KeyDetails = model.KeyDetails

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.SessionTtlInSeconds = model.SessionTtlInSeconds

	return
}

// SessionKeyTypeEnum Enum with underlying type: string
type SessionKeyTypeEnum string

// Set of constants representing the allowable values for SessionKeyTypeEnum
const (
	SessionKeyTypePub SessionKeyTypeEnum = "PUB"
)

var mappingSessionKeyType = map[string]SessionKeyTypeEnum{
	"PUB": SessionKeyTypePub,
}

// GetSessionKeyTypeEnumValues Enumerates the set of values for SessionKeyTypeEnum
func GetSessionKeyTypeEnumValues() []SessionKeyTypeEnum {
	values := make([]SessionKeyTypeEnum, 0)
	for _, v := range mappingSessionKeyType {
		values = append(values, v)
	}
	return values
}
