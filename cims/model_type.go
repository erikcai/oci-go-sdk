// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
//
// Cloud Incident Management Service
//

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ModelType The representation of ModelType
type ModelType struct {
	TypeName *string `mandatory:"false" json:"typeName"`
}

func (m ModelType) String() string {
	return common.PointerString(m)
}
