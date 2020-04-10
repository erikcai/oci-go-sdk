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

// DbSystemSourceFromBackup Use the backupId to specify from which backup the operation shall proceed.
type DbSystemSourceFromBackup struct {

	// The OCID of the backup to be used as the source for the cloning
	// procedure.
	BackupId *string `mandatory:"true" json:"backupId"`
}

func (m DbSystemSourceFromBackup) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DbSystemSourceFromBackup) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbSystemSourceFromBackup DbSystemSourceFromBackup
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDbSystemSourceFromBackup
	}{
		"BACKUP",
		(MarshalTypeDbSystemSourceFromBackup)(m),
	}

	return json.Marshal(&s)
}
