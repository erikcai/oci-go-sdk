// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDbSystemSourceFromBackupDetails Use the backupId to specify from which backup the operation shall proceed.
type CreateDbSystemSourceFromBackupDetails struct {

	// The OCID of the backup to be used as the source for the cloning
	// procedure.
	BackupId *string `mandatory:"true" json:"backupId"`
}

func (m CreateDbSystemSourceFromBackupDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbSystemSourceFromBackupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbSystemSourceFromBackupDetails CreateDbSystemSourceFromBackupDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateDbSystemSourceFromBackupDetails
	}{
		"BACKUP",
		(MarshalTypeCreateDbSystemSourceFromBackupDetails)(m),
	}

	return json.Marshal(&s)
}
