// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// DecomposableTaskHandler The decomposable task handler.
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
