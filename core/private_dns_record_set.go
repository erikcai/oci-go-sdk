// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PrivateDnsRecordSet A single Private DNS Resource Record Set.
// For more information, refer to RFC 7719 (https://tools.ietf.org/html/rfc7719).
type PrivateDnsRecordSet struct {

	// The OCID of the compartment containing the Private DNS Record Set.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Domain name of the PrivateDnsRecordSet, according to RFC 1035 (https://tools.ietf.org/html/rfc1035),
	// RFC2181 (https://tools.ietf.org/html/rfc2181), and RFC 7719 (https://tools.ietf.org/html/rfc7719).
	Domain *string `mandatory:"true" json:"domain"`

	// The OCID of the Private DNS Record Set.
	Id *string `mandatory:"true" json:"id"`

	// The Private DNS Record Set's current state.
	LifecycleState PrivateDnsRecordSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the DNS Zone the Private DNS Record Set belongs to.
	PrivateZoneId *string `mandatory:"true" json:"privateZoneId"`

	// List of RDATAs for the PrivateDnsRecordSet, in which individual RDATAs comply with RFC 1035 (https://tools.ietf.org/html/rfc1035).
	// -*A:* An IPv4 address.
	Rdatas []string `mandatory:"true" json:"rdatas"`

	// Type of PrivateDnsRecordSet according to RFC 1035 (https://tools.ietf.org/html/rfc1035) and RFC 7719 (https://tools.ietf.org/html/rfc7719).
	// Currently supported list of types are the following.
	// -*A:* Type 1, a host name to IPv4 address
	Type PrivateDnsRecordSetTypeEnum `mandatory:"true" json:"type"`

	// Time to live value for the PrivateDnsRecordSet, according to RFC 1035 (https://tools.ietf.org/html/rfc1035).
	// Defaults to 86400.
	Ttl *int `mandatory:"false" json:"ttl"`
}

func (m PrivateDnsRecordSet) String() string {
	return common.PointerString(m)
}

// PrivateDnsRecordSetLifecycleStateEnum Enum with underlying type: string
type PrivateDnsRecordSetLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateDnsRecordSetLifecycleState
const (
	PrivateDnsRecordSetLifecycleStateProvisioning PrivateDnsRecordSetLifecycleStateEnum = "PROVISIONING"
	PrivateDnsRecordSetLifecycleStateAvailable    PrivateDnsRecordSetLifecycleStateEnum = "AVAILABLE"
	PrivateDnsRecordSetLifecycleStateTerminating  PrivateDnsRecordSetLifecycleStateEnum = "TERMINATING"
	PrivateDnsRecordSetLifecycleStateTerminated   PrivateDnsRecordSetLifecycleStateEnum = "TERMINATED"
)

var mappingPrivateDnsRecordSetLifecycleState = map[string]PrivateDnsRecordSetLifecycleStateEnum{
	"PROVISIONING": PrivateDnsRecordSetLifecycleStateProvisioning,
	"AVAILABLE":    PrivateDnsRecordSetLifecycleStateAvailable,
	"TERMINATING":  PrivateDnsRecordSetLifecycleStateTerminating,
	"TERMINATED":   PrivateDnsRecordSetLifecycleStateTerminated,
}

// GetPrivateDnsRecordSetLifecycleStateEnumValues Enumerates the set of values for PrivateDnsRecordSetLifecycleState
func GetPrivateDnsRecordSetLifecycleStateEnumValues() []PrivateDnsRecordSetLifecycleStateEnum {
	values := make([]PrivateDnsRecordSetLifecycleStateEnum, 0)
	for _, v := range mappingPrivateDnsRecordSetLifecycleState {
		values = append(values, v)
	}
	return values
}

// PrivateDnsRecordSetTypeEnum Enum with underlying type: string
type PrivateDnsRecordSetTypeEnum string

// Set of constants representing the allowable values for PrivateDnsRecordSetType
const (
	PrivateDnsRecordSetTypeA PrivateDnsRecordSetTypeEnum = "A"
)

var mappingPrivateDnsRecordSetType = map[string]PrivateDnsRecordSetTypeEnum{
	"A": PrivateDnsRecordSetTypeA,
}

// GetPrivateDnsRecordSetTypeEnumValues Enumerates the set of values for PrivateDnsRecordSetType
func GetPrivateDnsRecordSetTypeEnumValues() []PrivateDnsRecordSetTypeEnum {
	values := make([]PrivateDnsRecordSetTypeEnum, 0)
	for _, v := range mappingPrivateDnsRecordSetType {
		values = append(values, v)
	}
	return values
}
