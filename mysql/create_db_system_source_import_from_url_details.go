// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDbSystemSourceImportFromUrlDetails An Object Storage PAR from which to import the DB System initial data.
type CreateDbSystemSourceImportFromUrlDetails struct {

	// The Pre-Authenticated Request (PAR) URL of the file you want to import from Object Storage.
	SourceUrl *string `mandatory:"true" json:"sourceUrl"`
}

func (m CreateDbSystemSourceImportFromUrlDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbSystemSourceImportFromUrlDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbSystemSourceImportFromUrlDetails CreateDbSystemSourceImportFromUrlDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateDbSystemSourceImportFromUrlDetails
	}{
		"IMPORTURL",
		(MarshalTypeCreateDbSystemSourceImportFromUrlDetails)(m),
	}

	return json.Marshal(&s)
}
