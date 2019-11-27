// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
//
// API for managing and performing operations with keys and vaults.
//

package keymanagement

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// BackupVaultDetails The representation of BackupVaultDetails
type BackupVaultDetails struct {
	BackupLocation BackupLocation `mandatory:"false" json:"backupLocation"`

	IsIncludeKeys *bool `mandatory:"false" json:"isIncludeKeys"`
}

func (m BackupVaultDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *BackupVaultDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BackupLocation backuplocation `json:"backupLocation"`
		IsIncludeKeys  *bool          `json:"isIncludeKeys"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.BackupLocation.UnmarshalPolymorphicJSON(model.BackupLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.BackupLocation = nn.(BackupLocation)
	} else {
		m.BackupLocation = nil
	}

	m.IsIncludeKeys = model.IsIncludeKeys
	return
}
