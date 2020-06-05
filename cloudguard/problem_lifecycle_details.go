// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ProblemLifecycleDetailsEnum Enum with underlying type: string
type ProblemLifecycleDetailsEnum string

// Set of constants representing the allowable values for ProblemLifecycleDetailsEnum
const (
	ProblemLifecycleDetailsOpen      ProblemLifecycleDetailsEnum = "OPEN"
	ProblemLifecycleDetailsResolved  ProblemLifecycleDetailsEnum = "RESOLVED"
	ProblemLifecycleDetailsDismissed ProblemLifecycleDetailsEnum = "DISMISSED"
)

var mappingProblemLifecycleDetails = map[string]ProblemLifecycleDetailsEnum{
	"OPEN":      ProblemLifecycleDetailsOpen,
	"RESOLVED":  ProblemLifecycleDetailsResolved,
	"DISMISSED": ProblemLifecycleDetailsDismissed,
}

// GetProblemLifecycleDetailsEnumValues Enumerates the set of values for ProblemLifecycleDetailsEnum
func GetProblemLifecycleDetailsEnumValues() []ProblemLifecycleDetailsEnum {
	values := make([]ProblemLifecycleDetailsEnum, 0)
	for _, v := range mappingProblemLifecycleDetails {
		values = append(values, v)
	}
	return values
}
