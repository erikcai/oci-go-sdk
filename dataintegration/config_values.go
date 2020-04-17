// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
