// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastions

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Session the session resource.
type Session struct {

	// OCID of session
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of the session.
	SessionType SessionTypeEnum `mandatory:"true" json:"sessionType"`

	// OCID of the bastion that the session is created on
	BastionId *string `mandatory:"true" json:"bastionId"`

	// bastion display name that the session is created on
	BastionDisplayName *string `mandatory:"true" json:"bastionDisplayName"`

	// display name of the target compute instance
	TargetComputeInstanceDisplayName *string `mandatory:"true" json:"targetComputeInstanceDisplayName"`

	// name of the user to use on end instance
	TargetComputeInstanceUserName *string `mandatory:"true" json:"targetComputeInstanceUserName"`

	// The OCID of the target compute instance to connect to.
	TargetComputeInstanceId *string `mandatory:"true" json:"targetComputeInstanceId"`

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

	// target compute instance ip address.
	TargetPrivateIpAddress *string `mandatory:"false" json:"targetPrivateIpAddress"`

	// target compute instance port to use.
	TargetComputeInstancePort *int `mandatory:"false" json:"targetComputeInstancePort"`

	// bastion ip address to connect from.
	BastionIp *string `mandatory:"false" json:"bastionIp"`

	// bastion port to connect from.
	BastionPort *int `mandatory:"false" json:"bastionPort"`

	// the connection message for the session.
	SshMetadata map[string]string `mandatory:"false" json:"sshMetadata"`

	// the type of the key, PUB refers to normal public key
	KeyType SessionKeyTypeEnum `mandatory:"false" json:"keyType,omitempty"`

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
