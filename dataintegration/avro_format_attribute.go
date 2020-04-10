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

// AvroFormatAttribute Th avro format attribute.
type AvroFormatAttribute struct {

	// compression
	Compression *string `mandatory:"false" json:"compression"`
}

func (m AvroFormatAttribute) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AvroFormatAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAvroFormatAttribute AvroFormatAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeAvroFormatAttribute
	}{
		"AVROFORMAT",
		(MarshalTypeAvroFormatAttribute)(m),
	}

	return json.Marshal(&s)
}
