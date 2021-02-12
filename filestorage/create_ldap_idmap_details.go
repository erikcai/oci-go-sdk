// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// CreateLdapIdmapDetails Mount Target details about the LDAP ID Map configuration.
type CreateLdapIdmapDetails struct {

	// All LDAP searches are recursive starting at this Base Distinguished user name.
	BaseDistinguishedUserName *string `mandatory:"true" json:"baseDistinguishedUserName"`

	// All LDAP searches are recursive starting at this Base Distinguished group name.
	BaseDistinguishedGroupName *string `mandatory:"true" json:"baseDistinguishedGroupName"`

	// Schema type of LDAP account.
	SchemaType CreateLdapIdmapDetailsSchemaTypeEnum `mandatory:"false" json:"schemaType,omitempty"`

	// Integer for how often the mount target should recheck LDAP for updates.
	CacheRefreshIntervalSeconds *int `mandatory:"false" json:"cacheRefreshIntervalSeconds"`

	// Integer for how long cached entries may be used.
	CacheLifetimeSeconds *int `mandatory:"false" json:"cacheLifetimeSeconds"`

	// Integer for how long to cache if idmap information is missing.
	NegativeCacheLifetimeSeconds *int `mandatory:"false" json:"negativeCacheLifetimeSeconds"`

	// OCID of the first LDAP Account
	OutboundConnector1Id *string `mandatory:"false" json:"outboundConnector1Id"`

	// OCID of the second LDAP Account
	OutboundConnector2Id *string `mandatory:"false" json:"outboundConnector2Id"`
}

func (m CreateLdapIdmapDetails) String() string {
	return common.PointerString(m)
}

// CreateLdapIdmapDetailsSchemaTypeEnum Enum with underlying type: string
type CreateLdapIdmapDetailsSchemaTypeEnum string

// Set of constants representing the allowable values for CreateLdapIdmapDetailsSchemaTypeEnum
const (
	CreateLdapIdmapDetailsSchemaTypeRfc2307 CreateLdapIdmapDetailsSchemaTypeEnum = "RFC2307"
)

var mappingCreateLdapIdmapDetailsSchemaType = map[string]CreateLdapIdmapDetailsSchemaTypeEnum{
	"RFC2307": CreateLdapIdmapDetailsSchemaTypeRfc2307,
}

// GetCreateLdapIdmapDetailsSchemaTypeEnumValues Enumerates the set of values for CreateLdapIdmapDetailsSchemaTypeEnum
func GetCreateLdapIdmapDetailsSchemaTypeEnumValues() []CreateLdapIdmapDetailsSchemaTypeEnum {
	values := make([]CreateLdapIdmapDetailsSchemaTypeEnum, 0)
	for _, v := range mappingCreateLdapIdmapDetailsSchemaType {
		values = append(values, v)
	}
	return values
}
