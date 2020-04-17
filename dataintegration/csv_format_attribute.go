// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CsvFormatAttribute The CSV format attribute.
type CsvFormatAttribute struct {

	// encoding
	Encoding *string `mandatory:"false" json:"encoding"`

	// escapeChar
	EscapeChar *string `mandatory:"false" json:"escapeChar"`

	// delimiter
	Delimiter *string `mandatory:"false" json:"delimiter"`

	// quoteCharacter
	QuoteCharacter *string `mandatory:"false" json:"quoteCharacter"`

	// hasHeader
	IsHasHeader *bool `mandatory:"false" json:"isHasHeader"`
}

func (m CsvFormatAttribute) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CsvFormatAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCsvFormatAttribute CsvFormatAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCsvFormatAttribute
	}{
		"CSVFORMAT",
		(MarshalTypeCsvFormatAttribute)(m),
	}

	return json.Marshal(&s)
}
