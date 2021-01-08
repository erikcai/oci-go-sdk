// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastion

// BastionLifecycleStateEnum Enum with underlying type: string
type BastionLifecycleStateEnum string

// Set of constants representing the allowable values for BastionLifecycleStateEnum
const (
	BastionLifecycleStateCreating BastionLifecycleStateEnum = "CREATING"
	BastionLifecycleStateUpdating BastionLifecycleStateEnum = "UPDATING"
	BastionLifecycleStateActive   BastionLifecycleStateEnum = "ACTIVE"
	BastionLifecycleStateDeleting BastionLifecycleStateEnum = "DELETING"
	BastionLifecycleStateDeleted  BastionLifecycleStateEnum = "DELETED"
	BastionLifecycleStateFailed   BastionLifecycleStateEnum = "FAILED"
)

var mappingBastionLifecycleState = map[string]BastionLifecycleStateEnum{
	"CREATING": BastionLifecycleStateCreating,
	"UPDATING": BastionLifecycleStateUpdating,
	"ACTIVE":   BastionLifecycleStateActive,
	"DELETING": BastionLifecycleStateDeleting,
	"DELETED":  BastionLifecycleStateDeleted,
	"FAILED":   BastionLifecycleStateFailed,
}

// GetBastionLifecycleStateEnumValues Enumerates the set of values for BastionLifecycleStateEnum
func GetBastionLifecycleStateEnumValues() []BastionLifecycleStateEnum {
	values := make([]BastionLifecycleStateEnum, 0)
	for _, v := range mappingBastionLifecycleState {
		values = append(values, v)
	}
	return values
}
