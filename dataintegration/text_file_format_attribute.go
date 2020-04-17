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

// TextFileFormatAttribute The text file format attribute.
type TextFileFormatAttribute struct {

	// delimiter
	Delimiter *string `mandatory:"false" json:"delimiter"`

	// quote
	Quote *string `mandatory:"false" json:"quote"`

	// escapeCharacter
	EscapeCharacter *string `mandatory:"false" json:"escapeCharacter"`

	// encoding
	Encoding *string `mandatory:"false" json:"encoding"`
}

func (m TextFileFormatAttribute) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TextFileFormatAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTextFileFormatAttribute TextFileFormatAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTextFileFormatAttribute
	}{
		"TEXTFILEFORMAT",
		(MarshalTypeTextFileFormatAttribute)(m),
	}

	return json.Marshal(&s)
}
