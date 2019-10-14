// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDbHomeWithDbSystemIdFromBackupDetails Note that a valid `dbSystemId` value must be supplied for the `CreateDbHomeWithDbSystemIdFromBackup` API operation to successfully complete.
type CreateDbHomeWithDbSystemIdFromBackupDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	Database *CreateDatabaseFromBackupDetails `mandatory:"true" json:"database"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`
}

//GetDisplayName returns DisplayName
func (m CreateDbHomeWithDbSystemIdFromBackupDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetKmsKeyId returns KmsKeyId
func (m CreateDbHomeWithDbSystemIdFromBackupDetails) GetKmsKeyId() *string {
	return m.KmsKeyId
}

//GetKmsKeyVersionId returns KmsKeyVersionId
func (m CreateDbHomeWithDbSystemIdFromBackupDetails) GetKmsKeyVersionId() *string {
	return m.KmsKeyVersionId
}

func (m CreateDbHomeWithDbSystemIdFromBackupDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbHomeWithDbSystemIdFromBackupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithDbSystemIdFromBackupDetails CreateDbHomeWithDbSystemIdFromBackupDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithDbSystemIdFromBackupDetails
	}{
		"DB_BACKUP",
		(MarshalTypeCreateDbHomeWithDbSystemIdFromBackupDetails)(m),
	}

	return json.Marshal(&s)
}
