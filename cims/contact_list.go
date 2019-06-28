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

// ContactList Contact list
type ContactList struct {
	ContactList []Contact `mandatory:"false" json:"contactList"`
}

func (m ContactList) String() string {
	return common.PointerString(m)
}
