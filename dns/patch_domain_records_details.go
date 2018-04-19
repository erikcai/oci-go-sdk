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

// PatchDomainRecordsDetails The representation of PatchDomainRecordsDetails
type PatchDomainRecordsDetails struct {
	Items []RecordOperation `mandatory:"false" json:"items"`
}

func (m PatchDomainRecordsDetails) String() string {
	return common.PointerString(m)
}
