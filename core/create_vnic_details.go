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

// CreateVnicDetails The representation of CreateVnicDetails
type CreateVnicDetails struct {

	// The OCID of the subnet to create the VNIC in. Use this `subnetId` instead
	// of the deprecated `subnetId` in
	// LaunchInstanceDetails.
	// At least one of them is required; if you provide both, the values must match.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Whether the VNIC should be assigned a public IP address. Defaults to whether
	// the subnet is public or private. If not set and the VNIC is being created
	// in a private subnet (i.e., where `prohibitPublicIpOnVnic`=true in the
	// Subnet), then no public IP address is assigned.
	// If not set and the subnet is public (`prohibitPublicIpOnVnic`=false), then
	// a public IP address is assigned. If set to true and
	// `prohibitPublicIpOnVnic`=true, an error is returned.
	// Example: `false`
	AssignPublicIp *bool `mandatory:"false" json:"assignPublicIp"`

	// A user-friendly name for the VNIC. Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The hostname for the VNIC that is created during instance launch.
	// Used for DNS. The value is the hostname portion of the instance's
	// fully qualified domain name (FQDN) (e.g., `bminstance-1` in FQDN
	// `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// RFC 952 (https://tools.ietf.org/html/rfc952) and
	// RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// The value cannot be changed, and it can be retrieved from the
	// Vnic.
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).
	// Use this `hostnameLabel` instead
	// of the deprecated `hostnameLabel` in
	// LaunchInstanceDetails.
	// If you provide both, the values must match.
	// Example: `bminstance-1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// A private IP address of your choice to assign to the VNIC. Must be an
	// available IP address within the subnet's CIDR. If no value is specified,
	// a private IP address from the subnet will be automatically assigned.
	// Example: `10.0.3.1`
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// Indicates whether Source/Destination check is disabled on the VNIC.
	// Defaults to `false`, in which case we enable Source/Destination check on the VNIC.
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`
}

func (m CreateVnicDetails) String() string {
	return common.PointerString(m)
}
