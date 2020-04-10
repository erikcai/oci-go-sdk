// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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

// OracleReadAttribute The Oracle read attribute
type OracleReadAttribute struct {

	// fetchSize
	FetchSize *int `mandatory:"false" json:"fetchSize"`
}

func (m OracleReadAttribute) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OracleReadAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleReadAttribute OracleReadAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOracleReadAttribute
	}{
		"ORACLEREADATTRIBUTE",
		(MarshalTypeOracleReadAttribute)(m),
	}

	return json.Marshal(&s)
}
