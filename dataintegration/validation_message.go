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

// ValidationMessage The level, message key and validation message.
type ValidationMessage struct {

	// Total number of validation messages
	Level *string `mandatory:"false" json:"level"`

	// The key.
	MessageKey *string `mandatory:"false" json:"messageKey"`

	// The message itself.
	ValidationMessage *string `mandatory:"false" json:"validationMessage"`
}

func (m ValidationMessage) String() string {
	return common.PointerString(m)
}
