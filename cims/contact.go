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

// Contact Contact object
type Contact struct {
	Name *string `mandatory:"false" json:"name"`

	Email *string `mandatory:"false" json:"email"`

	Phone *string `mandatory:"false" json:"phone"`

	ContactType ContactContactTypeEnum `mandatory:"false" json:"contactType,omitempty"`
}

func (m Contact) String() string {
	return common.PointerString(m)
}

// ContactContactTypeEnum Enum with underlying type: string
type ContactContactTypeEnum string

// Set of constants representing the allowable values for ContactContactTypeEnum
const (
	ContactContactTypePrimary   ContactContactTypeEnum = "PRIMARY"
	ContactContactTypeAlternate ContactContactTypeEnum = "ALTERNATE"
	ContactContactTypeAdmin     ContactContactTypeEnum = "ADMIN"
)

var mappingContactContactType = map[string]ContactContactTypeEnum{
	"PRIMARY":   ContactContactTypePrimary,
	"ALTERNATE": ContactContactTypeAlternate,
	"ADMIN":     ContactContactTypeAdmin,
}

// GetContactContactTypeEnumValues Enumerates the set of values for ContactContactTypeEnum
func GetContactContactTypeEnumValues() []ContactContactTypeEnum {
	values := make([]ContactContactTypeEnum, 0)
	for _, v := range mappingContactContactType {
		values = append(values, v)
	}
	return values
}
