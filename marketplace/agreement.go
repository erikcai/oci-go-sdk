// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Agreement The model for an end user license agreement.
type Agreement struct {

	// The unique identifier of the agreement.
	Id *string `mandatory:"true" json:"id"`

	// The content URL of the agreement.
	ContentUrl *string `mandatory:"true" json:"contentUrl"`

	// A time-based signature that can be used to accept an agreement or cancel a
	// previous acceptance.
	Signature *string `mandatory:"true" json:"signature"`

	// The unique identifier of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Who authored the agreement.
	Author AgreementAuthorEnum `mandatory:"false" json:"author,omitempty"`

	// Call to action to read and accept the agreement.
	Prompt *string `mandatory:"false" json:"prompt"`
}

func (m Agreement) String() string {
	return common.PointerString(m)
}

// AgreementAuthorEnum Enum with underlying type: string
type AgreementAuthorEnum string

// Set of constants representing the allowable values for AgreementAuthorEnum
const (
	AgreementAuthorOracle  AgreementAuthorEnum = "ORACLE"
	AgreementAuthorPartner AgreementAuthorEnum = "PARTNER"
)

var mappingAgreementAuthor = map[string]AgreementAuthorEnum{
	"ORACLE":  AgreementAuthorOracle,
	"PARTNER": AgreementAuthorPartner,
}

// GetAgreementAuthorEnumValues Enumerates the set of values for AgreementAuthorEnum
func GetAgreementAuthorEnumValues() []AgreementAuthorEnum {
	values := make([]AgreementAuthorEnum, 0)
	for _, v := range mappingAgreementAuthor {
		values = append(values, v)
	}
	return values
}
