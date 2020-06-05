// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ProblemLifecycleStatesEnum Enum with underlying type: string
type ProblemLifecycleStatesEnum string

// Set of constants representing the allowable values for ProblemLifecycleStatesEnum
const (
	ProblemLifecycleStatesActive   ProblemLifecycleStatesEnum = "ACTIVE"
	ProblemLifecycleStatesInactive ProblemLifecycleStatesEnum = "INACTIVE"
)

var mappingProblemLifecycleStates = map[string]ProblemLifecycleStatesEnum{
	"ACTIVE":   ProblemLifecycleStatesActive,
	"INACTIVE": ProblemLifecycleStatesInactive,
}

// GetProblemLifecycleStatesEnumValues Enumerates the set of values for ProblemLifecycleStatesEnum
func GetProblemLifecycleStatesEnumValues() []ProblemLifecycleStatesEnum {
	values := make([]ProblemLifecycleStatesEnum, 0)
	for _, v := range mappingProblemLifecycleStates {
		values = append(values, v)
	}
	return values
}
