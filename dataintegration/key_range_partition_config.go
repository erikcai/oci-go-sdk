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

// KeyRangePartitionConfig auto generated description
type KeyRangePartitionConfig struct {

	// partitionNum
	PartitionNum *int `mandatory:"false" json:"partitionNum"`

	KeyRange *KeyRange `mandatory:"false" json:"keyRange"`
}

func (m KeyRangePartitionConfig) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m KeyRangePartitionConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKeyRangePartitionConfig KeyRangePartitionConfig
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeKeyRangePartitionConfig
	}{
		"KEYRANGEPARTITIONCONFIG",
		(MarshalTypeKeyRangePartitionConfig)(m),
	}

	return json.Marshal(&s)
}
