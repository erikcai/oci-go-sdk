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

// DbSystemSourceImportFromUrl An Object Storage PAR from which to import the DB System initial data.
type DbSystemSourceImportFromUrl struct {
}

func (m DbSystemSourceImportFromUrl) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DbSystemSourceImportFromUrl) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbSystemSourceImportFromUrl DbSystemSourceImportFromUrl
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDbSystemSourceImportFromUrl
	}{
		"IMPORTURL",
		(MarshalTypeDbSystemSourceImportFromUrl)(m),
	}

	return json.Marshal(&s)
}
