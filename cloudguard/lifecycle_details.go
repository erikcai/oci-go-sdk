// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// LifecycleDetailsEnum Enum with underlying type: string
type LifecycleDetailsEnum string

// Set of constants representing the allowable values for LifecycleDetailsEnum
const (
	LifecycleDetailsOpen      LifecycleDetailsEnum = "OPEN"
	LifecycleDetailsResolved  LifecycleDetailsEnum = "RESOLVED"
	LifecycleDetailsDismissed LifecycleDetailsEnum = "DISMISSED"
)

var mappingLifecycleDetails = map[string]LifecycleDetailsEnum{
	"OPEN":      LifecycleDetailsOpen,
	"RESOLVED":  LifecycleDetailsResolved,
	"DISMISSED": LifecycleDetailsDismissed,
}

// GetLifecycleDetailsEnumValues Enumerates the set of values for LifecycleDetailsEnum
func GetLifecycleDetailsEnumValues() []LifecycleDetailsEnum {
	values := make([]LifecycleDetailsEnum, 0)
	for _, v := range mappingLifecycleDetails {
		values = append(values, v)
	}
	return values
}
