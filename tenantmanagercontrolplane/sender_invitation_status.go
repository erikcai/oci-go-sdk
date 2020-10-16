// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// TenantManager API
//
// A description of the TenantManager API
//

package tenantmanagercontrolplane

// SenderInvitationStatusEnum Enum with underlying type: string
type SenderInvitationStatusEnum string

// Set of constants representing the allowable values for SenderInvitationStatusEnum
const (
	SenderInvitationStatusPending  SenderInvitationStatusEnum = "PENDING"
	SenderInvitationStatusCanceled SenderInvitationStatusEnum = "CANCELED"
	SenderInvitationStatusAccepted SenderInvitationStatusEnum = "ACCEPTED"
	SenderInvitationStatusExpired  SenderInvitationStatusEnum = "EXPIRED"
	SenderInvitationStatusFailed   SenderInvitationStatusEnum = "FAILED"
)

var mappingSenderInvitationStatus = map[string]SenderInvitationStatusEnum{
	"PENDING":  SenderInvitationStatusPending,
	"CANCELED": SenderInvitationStatusCanceled,
	"ACCEPTED": SenderInvitationStatusAccepted,
	"EXPIRED":  SenderInvitationStatusExpired,
	"FAILED":   SenderInvitationStatusFailed,
}

// GetSenderInvitationStatusEnumValues Enumerates the set of values for SenderInvitationStatusEnum
func GetSenderInvitationStatusEnumValues() []SenderInvitationStatusEnum {
	values := make([]SenderInvitationStatusEnum, 0)
	for _, v := range mappingSenderInvitationStatus {
		values = append(values, v)
	}
	return values
}
