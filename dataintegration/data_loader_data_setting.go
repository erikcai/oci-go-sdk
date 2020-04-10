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

// DataLoaderDataSetting auto generated description
type DataLoaderDataSetting struct {

	// contentFormat
	ContentFormat *string `mandatory:"false" json:"contentFormat"`

	// contentOptions
	ContentOptions map[string]string `mandatory:"false" json:"contentOptions"`
}

func (m DataLoaderDataSetting) String() string {
	return common.PointerString(m)
}
