// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm),
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm), and
// Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as virtual cloud networks (VCNs),
// compute instances, block storage volumes, and container images.
//

package core

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// LdapConfigDetails The detail of LDAP's authentication configuration.
type LdapConfigDetails struct {

	// The IP address of primary LDAP server.
	PrimaryServer *string `mandatory:"false" json:"primaryServer"`

	// The IP address of secondary LDAP server.
	SecondaryServer *string `mandatory:"false" json:"secondaryServer"`

	// Option for LDAP SSL once customer enables SSL for certain ClientVpnEndpoint.
	// Allowed values:
	// * `NEVER`: Do not use SSL (the default setting).
	// * `ADAPTIVE`: Try using SSL, if that fails, use plain-text to get through authentication.
	// * `ALWAYS`: Always use SSL.
	UseSsl LdapConfigDetailsUseSslEnum `mandatory:"false" json:"useSsl,omitempty"`

	// Choose the authentication method once useSSL enabled.
	// Allowed values:
	// * `NEVER`: No peer certificate is required.
	// * `ALLOW`: Request a peer certificate, but session will not be aborted if certificate cannot be validated.
	// * `DEMAND`: A valid peer certificate is required, then session will be aborted if one is not provided.
	VerifySsl LdapConfigDetailsVerifySslEnum `mandatory:"false" json:"verifySsl,omitempty"`

	// Enable case-sensitivity or not in LDAP authentication.
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`

	// Whether to apply Anonymous bind or not.
	IsBindAnon *bool `mandatory:"false" json:"isBindAnon"`

	// The bind DN (Distinguished Name) includes the user and location of the
	// user in LDAP directory tree
	BindDN *string `mandatory:"false" json:"bindDN"`

	// The bind password is used to log in the LDAP server.
	BindPW *string `mandatory:"false" json:"bindPW"`

	// The starting point element helps LDAP service to navigate search scope.
	BaseDN *string `mandatory:"false" json:"baseDN"`

	// The username of client at attribute level.
	ClientUsername *string `mandatory:"false" json:"clientUsername"`

	// This additional requirement uses LDAP query syntax. E.g., to require that the user be a member of a particular LDAP group (specified by DN) use this filter:
	// memberOf=CN=VPN Users, CN=Users, DC=example, DC=net
	AdditionalRequirements *string `mandatory:"false" json:"additionalRequirements"`
}

func (m LdapConfigDetails) String() string {
	return common.PointerString(m)
}

// LdapConfigDetailsUseSslEnum Enum with underlying type: string
type LdapConfigDetailsUseSslEnum string

// Set of constants representing the allowable values for LdapConfigDetailsUseSslEnum
const (
	LdapConfigDetailsUseSslNever    LdapConfigDetailsUseSslEnum = "NEVER"
	LdapConfigDetailsUseSslAdaptive LdapConfigDetailsUseSslEnum = "ADAPTIVE"
	LdapConfigDetailsUseSslAlways   LdapConfigDetailsUseSslEnum = "ALWAYS"
)

var mappingLdapConfigDetailsUseSsl = map[string]LdapConfigDetailsUseSslEnum{
	"NEVER":    LdapConfigDetailsUseSslNever,
	"ADAPTIVE": LdapConfigDetailsUseSslAdaptive,
	"ALWAYS":   LdapConfigDetailsUseSslAlways,
}

// GetLdapConfigDetailsUseSslEnumValues Enumerates the set of values for LdapConfigDetailsUseSslEnum
func GetLdapConfigDetailsUseSslEnumValues() []LdapConfigDetailsUseSslEnum {
	values := make([]LdapConfigDetailsUseSslEnum, 0)
	for _, v := range mappingLdapConfigDetailsUseSsl {
		values = append(values, v)
	}
	return values
}

// LdapConfigDetailsVerifySslEnum Enum with underlying type: string
type LdapConfigDetailsVerifySslEnum string

// Set of constants representing the allowable values for LdapConfigDetailsVerifySslEnum
const (
	LdapConfigDetailsVerifySslNever  LdapConfigDetailsVerifySslEnum = "NEVER"
	LdapConfigDetailsVerifySslAllow  LdapConfigDetailsVerifySslEnum = "ALLOW"
	LdapConfigDetailsVerifySslDemand LdapConfigDetailsVerifySslEnum = "DEMAND"
)

var mappingLdapConfigDetailsVerifySsl = map[string]LdapConfigDetailsVerifySslEnum{
	"NEVER":  LdapConfigDetailsVerifySslNever,
	"ALLOW":  LdapConfigDetailsVerifySslAllow,
	"DEMAND": LdapConfigDetailsVerifySslDemand,
}

// GetLdapConfigDetailsVerifySslEnumValues Enumerates the set of values for LdapConfigDetailsVerifySslEnum
func GetLdapConfigDetailsVerifySslEnumValues() []LdapConfigDetailsVerifySslEnum {
	values := make([]LdapConfigDetailsVerifySslEnum, 0)
	for _, v := range mappingLdapConfigDetailsVerifySsl {
		values = append(values, v)
	}
	return values
}
