// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// OracleWriteAttribute The Oracle write attribute
type OracleWriteAttribute struct {

	// batchSize
	BatchSize *int `mandatory:"false" json:"batchSize"`

	// truncate
	IsTruncate *bool `mandatory:"false" json:"isTruncate"`

	// isolationLevel
	IsolationLevel *string `mandatory:"false" json:"isolationLevel"`
}

func (m OracleWriteAttribute) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OracleWriteAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleWriteAttribute OracleWriteAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOracleWriteAttribute
	}{
		"ORACLEWRITEATTRIBUTE",
		(MarshalTypeOracleWriteAttribute)(m),
	}

	return json.Marshal(&s)
}
