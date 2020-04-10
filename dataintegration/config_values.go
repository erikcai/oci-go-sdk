// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ConfigValues Configuration values can be string, objects or parameters.
type ConfigValues struct {

	// configParamValues
	ConfigParamValues map[string]ConfigParamValue `mandatory:"false" json:"configParamValues"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The configuration definition reference.
	ConfigDef *string `mandatory:"false" json:"configDef"`

	// The key of the config owner for this value.
	ConfigOwner *string `mandatory:"false" json:"configOwner"`

	// configParamNames
	ConfigParamNames []string `mandatory:"false" json:"configParamNames"`
}

func (m ConfigValues) String() string {
	return common.PointerString(m)
}
