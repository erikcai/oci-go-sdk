// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DNS Service API
//
// API for managing DNS zones, records, and policies.
//

package dns

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PatchRrSetDetails The representation of PatchRrSetDetails
type PatchRrSetDetails struct {
	Items []RecordOperation `mandatory:"false" json:"items"`
}

func (m PatchRrSetDetails) String() string {
	return common.PointerString(m)
}
