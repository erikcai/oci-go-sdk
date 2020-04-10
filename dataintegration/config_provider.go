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

// ConfigProvider auto generated description
type ConfigProvider struct {

	// bindings
	Bindings map[string]ParameterValue `mandatory:"false" json:"bindings"`

	// childProviders
	ChildProviders map[string]ConfigProvider `mandatory:"false" json:"childProviders"`
}

func (m ConfigProvider) String() string {
	return common.PointerString(m)
}
