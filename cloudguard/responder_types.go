// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ResponderTypesEnum Enum with underlying type: string
type ResponderTypesEnum string

// Set of constants representing the allowable values for ResponderTypesEnum
const (
	ResponderTypesRemediation  ResponderTypesEnum = "REMEDIATION"
	ResponderTypesNotification ResponderTypesEnum = "NOTIFICATION"
)

var mappingResponderTypes = map[string]ResponderTypesEnum{
	"REMEDIATION":  ResponderTypesRemediation,
	"NOTIFICATION": ResponderTypesNotification,
}

// GetResponderTypesEnumValues Enumerates the set of values for ResponderTypesEnum
func GetResponderTypesEnumValues() []ResponderTypesEnum {
	values := make([]ResponderTypesEnum, 0)
	for _, v := range mappingResponderTypes {
		values = append(values, v)
	}
	return values
}
