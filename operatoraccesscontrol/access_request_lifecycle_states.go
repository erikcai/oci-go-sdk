// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

// AccessRequestLifecycleStatesEnum Enum with underlying type: string
type AccessRequestLifecycleStatesEnum string

// Set of constants representing the allowable values for AccessRequestLifecycleStatesEnum
const (
	AccessRequestLifecycleStatesCreated         AccessRequestLifecycleStatesEnum = "CREATED"
	AccessRequestLifecycleStatesApprovalwaiting AccessRequestLifecycleStatesEnum = "APPROVALWAITING"
	AccessRequestLifecycleStatesPreapproved     AccessRequestLifecycleStatesEnum = "PREAPPROVED"
	AccessRequestLifecycleStatesApproved        AccessRequestLifecycleStatesEnum = "APPROVED"
	AccessRequestLifecycleStatesRejected        AccessRequestLifecycleStatesEnum = "REJECTED"
	AccessRequestLifecycleStatesDeployed        AccessRequestLifecycleStatesEnum = "DEPLOYED"
	AccessRequestLifecycleStatesDeployfailed    AccessRequestLifecycleStatesEnum = "DEPLOYFAILED"
	AccessRequestLifecycleStatesUndeployed      AccessRequestLifecycleStatesEnum = "UNDEPLOYED"
	AccessRequestLifecycleStatesUndeployfailed  AccessRequestLifecycleStatesEnum = "UNDEPLOYFAILED"
	AccessRequestLifecycleStatesRevoking        AccessRequestLifecycleStatesEnum = "REVOKING"
	AccessRequestLifecycleStatesRevoked         AccessRequestLifecycleStatesEnum = "REVOKED"
	AccessRequestLifecycleStatesExtending       AccessRequestLifecycleStatesEnum = "EXTENDING"
	AccessRequestLifecycleStatesExtended        AccessRequestLifecycleStatesEnum = "EXTENDED"
	AccessRequestLifecycleStatesCompleting      AccessRequestLifecycleStatesEnum = "COMPLETING"
	AccessRequestLifecycleStatesCompleted       AccessRequestLifecycleStatesEnum = "COMPLETED"
	AccessRequestLifecycleStatesExpired         AccessRequestLifecycleStatesEnum = "EXPIRED"
)

var mappingAccessRequestLifecycleStates = map[string]AccessRequestLifecycleStatesEnum{
	"CREATED":         AccessRequestLifecycleStatesCreated,
	"APPROVALWAITING": AccessRequestLifecycleStatesApprovalwaiting,
	"PREAPPROVED":     AccessRequestLifecycleStatesPreapproved,
	"APPROVED":        AccessRequestLifecycleStatesApproved,
	"REJECTED":        AccessRequestLifecycleStatesRejected,
	"DEPLOYED":        AccessRequestLifecycleStatesDeployed,
	"DEPLOYFAILED":    AccessRequestLifecycleStatesDeployfailed,
	"UNDEPLOYED":      AccessRequestLifecycleStatesUndeployed,
	"UNDEPLOYFAILED":  AccessRequestLifecycleStatesUndeployfailed,
	"REVOKING":        AccessRequestLifecycleStatesRevoking,
	"REVOKED":         AccessRequestLifecycleStatesRevoked,
	"EXTENDING":       AccessRequestLifecycleStatesExtending,
	"EXTENDED":        AccessRequestLifecycleStatesExtended,
	"COMPLETING":      AccessRequestLifecycleStatesCompleting,
	"COMPLETED":       AccessRequestLifecycleStatesCompleted,
	"EXPIRED":         AccessRequestLifecycleStatesExpired,
}

// GetAccessRequestLifecycleStatesEnumValues Enumerates the set of values for AccessRequestLifecycleStatesEnum
func GetAccessRequestLifecycleStatesEnumValues() []AccessRequestLifecycleStatesEnum {
	values := make([]AccessRequestLifecycleStatesEnum, 0)
	for _, v := range mappingAccessRequestLifecycleStates {
		values = append(values, v)
	}
	return values
}
