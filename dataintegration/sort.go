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

// Sort auto generated description
type Sort struct {

	// sortClauses
	SortClauses []SortClause `mandatory:"false" json:"sortClauses"`
}

func (m Sort) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Sort) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSort Sort
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeSort
	}{
		"SORT",
		(MarshalTypeSort)(m),
	}

	return json.Marshal(&s)
}
