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

// CreateSessionDetails The information about new session.
type CreateSessionDetails struct {

	// bastion instance ocid that the session is created on
	BastionId *string `mandatory:"true" json:"bastionId"`

	// name of the user to use on target compute instance
	TargetComputeInstanceUserName *string `mandatory:"true" json:"targetComputeInstanceUserName"`

	// The OCID of the target compute instance to connect to.
	TargetComputeInstanceId *string `mandatory:"true" json:"targetComputeInstanceId"`

	KeyDetails *PublicKeyDetails `mandatory:"true" json:"keyDetails"`

	// session Identifier, not unique
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the compartment
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The type of the session.
	SessionType SessionTypeEnum `mandatory:"false" json:"sessionType,omitempty"`

	// target compute instance ip address.
	TargetPrivateIpAddress *string `mandatory:"false" json:"targetPrivateIpAddress"`

	// target compute instance port to use.
	TargetComputeInstancePort *int `mandatory:"false" json:"targetComputeInstancePort"`

	// the type of the key, PUB refers to normal public key
	KeyType CreateSessionDetailsKeyTypeEnum `mandatory:"false" json:"keyType,omitempty"`

	// TTL of the session.
	SessionTtlInSeconds *int `mandatory:"false" json:"sessionTtlInSeconds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSessionDetails) String() string {
	return common.PointerString(m)
}

// CreateSessionDetailsKeyTypeEnum Enum with underlying type: string
type CreateSessionDetailsKeyTypeEnum string

// Set of constants representing the allowable values for CreateSessionDetailsKeyTypeEnum
const (
	CreateSessionDetailsKeyTypePub CreateSessionDetailsKeyTypeEnum = "PUB"
)

var mappingCreateSessionDetailsKeyType = map[string]CreateSessionDetailsKeyTypeEnum{
	"PUB": CreateSessionDetailsKeyTypePub,
}

// GetCreateSessionDetailsKeyTypeEnumValues Enumerates the set of values for CreateSessionDetailsKeyTypeEnum
func GetCreateSessionDetailsKeyTypeEnumValues() []CreateSessionDetailsKeyTypeEnum {
	values := make([]CreateSessionDetailsKeyTypeEnum, 0)
	for _, v := range mappingCreateSessionDetailsKeyType {
		values = append(values, v)
	}
	return values
}