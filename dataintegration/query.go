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

// Query auto generated description
type Query struct {

	// query
	Query *string `mandatory:"false" json:"query"`
}

func (m Query) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Query) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeQuery Query
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeQuery
	}{
		"QUERY",
		(MarshalTypeQuery)(m),
	}

	return json.Marshal(&s)
}
