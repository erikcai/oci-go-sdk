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

// PrivateIpAddress A private IP address is an IP address allocated from the CIDR block of IP addresses for a subnet.
// An IP address is automatically allocated and mapped to a VNIC when the VNIC is created
// to provide network connectivity to the compute instance to which the VNIC is attached.
// This IP address is called the "primary" IP address for the VNIC. "Secondary" IP
// addresses may be allocated by Oracle cloud services and mapped and moved between VNICs.
// The "mapping" subresource indicates the VNIC to which the private
// IP address is currently mapped. If not mapped, the "mapping" subresource is not present.
type PrivateIpAddress struct {

	// The IP address within the subnet.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The IP address's current state.
	LifecycleState PrivateIpAddressLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// OCID of the subnet the IP address is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Creation time of the IP address.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The Availability Domain of the IP address definition.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// A user-friendly name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The hostname DNS label used in the FQDN to identify the IP address.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// OCID of the IP address if it is secondary.
	Id *string `mandatory:"false" json:"id"`

	// The IP address mapping if mapped or in the process of mapping or unmapping.
	Mapping *PrivateIpAddressMapping `mandatory:"false" json:"mapping"`
}

func (m PrivateIpAddress) String() string {
	return common.PointerString(m)
}

// PrivateIpAddressLifecycleStateEnum Enum with underlying type: string
type PrivateIpAddressLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateIpAddressLifecycleState
const (
	PrivateIpAddressLifecycleStateProvisioning PrivateIpAddressLifecycleStateEnum = "PROVISIONING"
	PrivateIpAddressLifecycleStateAvailable    PrivateIpAddressLifecycleStateEnum = "AVAILABLE"
	PrivateIpAddressLifecycleStateTerminating  PrivateIpAddressLifecycleStateEnum = "TERMINATING"
	PrivateIpAddressLifecycleStateTerminated   PrivateIpAddressLifecycleStateEnum = "TERMINATED"
)

var mappingPrivateIpAddressLifecycleState = map[string]PrivateIpAddressLifecycleStateEnum{
	"PROVISIONING": PrivateIpAddressLifecycleStateProvisioning,
	"AVAILABLE":    PrivateIpAddressLifecycleStateAvailable,
	"TERMINATING":  PrivateIpAddressLifecycleStateTerminating,
	"TERMINATED":   PrivateIpAddressLifecycleStateTerminated,
}

// GetPrivateIpAddressLifecycleStateEnumValues Enumerates the set of values for PrivateIpAddressLifecycleState
func GetPrivateIpAddressLifecycleStateEnumValues() []PrivateIpAddressLifecycleStateEnum {
	values := make([]PrivateIpAddressLifecycleStateEnum, 0)
	for _, v := range mappingPrivateIpAddressLifecycleState {
		values = append(values, v)
	}
	return values
}
