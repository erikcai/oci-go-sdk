// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// TenantManager API
//
// A description of the TenantManager API
//

package tenantmanagercontrolplane

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateCreating   LifecycleStateEnum = "CREATING"
	LifecycleStateActive     LifecycleStateEnum = "ACTIVE"
	LifecycleStateInactive   LifecycleStateEnum = "INACTIVE"
	LifecycleStateUpdating   LifecycleStateEnum = "UPDATING"
	LifecycleStateFailed     LifecycleStateEnum = "FAILED"
	LifecycleStateTerminated LifecycleStateEnum = "TERMINATED"
)

var mappingLifecycleState = map[string]LifecycleStateEnum{
	"CREATING":   LifecycleStateCreating,
	"ACTIVE":     LifecycleStateActive,
	"INACTIVE":   LifecycleStateInactive,
	"UPDATING":   LifecycleStateUpdating,
	"FAILED":     LifecycleStateFailed,
	"TERMINATED": LifecycleStateTerminated,
}

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
	values := make([]LifecycleStateEnum, 0)
	for _, v := range mappingLifecycleState {
		values = append(values, v)
	}
	return values
}
