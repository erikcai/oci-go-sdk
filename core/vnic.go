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

// Vnic A virtual network interface card. Each instance automatically has a VNIC attached to it,
// the "primary" VNIC, and the VNIC connects the instance to the subnet. Additional VNICs
// can be created, "secondary" VNICs, and attached to an instance provide a connection to other
// subnets. For more information, see
// Overview of the Compute Service (https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Concepts/computeoverview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
type Vnic struct {

	// The VNIC's Availability Domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment containing the VNIC.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VNIC.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the VNIC.
	LifecycleState VnicLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The private IP address of the VNIC. The address is within the subnet's CIDR
	// and is accessible within the VCN.
	PrivateIp *string `mandatory:"true" json:"privateIp"`

	// The OCID of the subnet the VNIC is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time the VNIC was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name. Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The hostname for the VNIC that is created during instance launch.
	// Used for DNS. The value is the hostname portion of the instance's
	// fully qualified domain name (FQDN) (e.g., `bminstance-1` in FQDN
	// `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// RFC 952 (https://tools.ietf.org/html/rfc952) and
	// RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// The value cannot be changed.
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).
	// Example: `bminstance-1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// Whether the VNIC is a primary VNIC.
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// The MAC address of the VNIC.
	// Example: `00:00:17:B6:4D:DD`
	MacAddress *string `mandatory:"false" json:"macAddress"`

	// The public IP address of the VNIC, if one is assigned.
	PublicIp *string `mandatory:"false" json:"publicIp"`

	// Indicates whether Source/Destination check is disabled on the VNIC.
	// Defaults to `false`, in which case we enable Source/Destination check on the VNIC.
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`
}

func (m Vnic) String() string {
	return common.PointerString(m)
}

// VnicLifecycleStateEnum Enum with underlying type: string
type VnicLifecycleStateEnum string

// Set of constants representing the allowable values for VnicLifecycleState
const (
	VnicLifecycleStateProvisioning VnicLifecycleStateEnum = "PROVISIONING"
	VnicLifecycleStateAvailable    VnicLifecycleStateEnum = "AVAILABLE"
	VnicLifecycleStateTerminating  VnicLifecycleStateEnum = "TERMINATING"
	VnicLifecycleStateTerminated   VnicLifecycleStateEnum = "TERMINATED"
)

var mappingVnicLifecycleState = map[string]VnicLifecycleStateEnum{
	"PROVISIONING": VnicLifecycleStateProvisioning,
	"AVAILABLE":    VnicLifecycleStateAvailable,
	"TERMINATING":  VnicLifecycleStateTerminating,
	"TERMINATED":   VnicLifecycleStateTerminated,
}

// GetVnicLifecycleStateEnumValues Enumerates the set of values for VnicLifecycleState
func GetVnicLifecycleStateEnumValues() []VnicLifecycleStateEnum {
	values := make([]VnicLifecycleStateEnum, 0)
	for _, v := range mappingVnicLifecycleState {
		values = append(values, v)
	}
	return values
}
