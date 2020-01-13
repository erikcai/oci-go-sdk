// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OptionMetadataValue The representation of OptionMetadataValue
type OptionMetadataValue struct {

	// The literal value being described.
	Value *string `mandatory:"false" json:"value"`

	// The semantics of the value being described.
	Description *string `mandatory:"false" json:"description"`
}

func (m OptionMetadataValue) String() string {
	return common.PointerString(m)
}
