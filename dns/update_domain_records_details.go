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

// UpdateDomainRecordsDetails The representation of UpdateDomainRecordsDetails
type UpdateDomainRecordsDetails struct {
	Items []RecordDetails `mandatory:"false" json:"items"`
}

func (m UpdateDomainRecordsDetails) String() string {
	return common.PointerString(m)
}
