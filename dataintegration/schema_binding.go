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

// SchemaBinding The schema binding object.
type SchemaBinding struct {
	DataAsset *DataAsset `mandatory:"false" json:"dataAsset"`

	// resourceName
	ResourceName *string `mandatory:"false" json:"resourceName"`
}

func (m SchemaBinding) String() string {
	return common.PointerString(m)
}
