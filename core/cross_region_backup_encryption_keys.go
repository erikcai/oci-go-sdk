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

// CrossRegionBackupEncryptionKeys The representation of CrossRegionBackupEncryptionKeys
type CrossRegionBackupEncryptionKeys struct {

	// volume encryption keys
	Keys []string `mandatory:"true" json:"keys"`

	// the encryption type used to encrypt the keys.
	EncryptionType CrossRegionBackupEncryptionKeysEncryptionTypeEnum `mandatory:"true" json:"encryptionType"`

	// version of the master cross region key used for encryption
	MasterKeyVersion *string `mandatory:"false" json:"masterKeyVersion"`
}

func (m CrossRegionBackupEncryptionKeys) String() string {
	return common.PointerString(m)
}

// CrossRegionBackupEncryptionKeysEncryptionTypeEnum Enum with underlying type: string
type CrossRegionBackupEncryptionKeysEncryptionTypeEnum string

// Set of constants representing the allowable values for CrossRegionBackupEncryptionKeysEncryptionType
const (
	CrossRegionBackupEncryptionKeysEncryptionTypeUnencrypted                 CrossRegionBackupEncryptionKeysEncryptionTypeEnum = "UNENCRYPTED"
	CrossRegionBackupEncryptionKeysEncryptionTypeMasterCrossRegionKeyEncrypt CrossRegionBackupEncryptionKeysEncryptionTypeEnum = "MASTER_CROSS_REGION_KEY_ENCRYPT"
)

var mappingCrossRegionBackupEncryptionKeysEncryptionType = map[string]CrossRegionBackupEncryptionKeysEncryptionTypeEnum{
	"UNENCRYPTED":                     CrossRegionBackupEncryptionKeysEncryptionTypeUnencrypted,
	"MASTER_CROSS_REGION_KEY_ENCRYPT": CrossRegionBackupEncryptionKeysEncryptionTypeMasterCrossRegionKeyEncrypt,
}

// GetCrossRegionBackupEncryptionKeysEncryptionTypeEnumValues Enumerates the set of values for CrossRegionBackupEncryptionKeysEncryptionType
func GetCrossRegionBackupEncryptionKeysEncryptionTypeEnumValues() []CrossRegionBackupEncryptionKeysEncryptionTypeEnum {
	values := make([]CrossRegionBackupEncryptionKeysEncryptionTypeEnum, 0)
	for _, v := range mappingCrossRegionBackupEncryptionKeysEncryptionType {
		values = append(values, v)
	}
	return values
}
