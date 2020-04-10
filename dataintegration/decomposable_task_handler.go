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

// DecomposableTaskHandler auto generated description
type DecomposableTaskHandler struct {
}

func (m DecomposableTaskHandler) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DecomposableTaskHandler) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDecomposableTaskHandler DecomposableTaskHandler
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDecomposableTaskHandler
	}{
		"DECOMPOSABLE_TASK_HANDLER",
		(MarshalTypeDecomposableTaskHandler)(m),
	}

	return json.Marshal(&s)
}
