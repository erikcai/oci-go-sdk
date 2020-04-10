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

// Message Message details
type Message struct {

	// Type of message (error, warning, or info)
	Type MessageTypeEnum `mandatory:"true" json:"type"`

	// The message code
	Code *string `mandatory:"true" json:"code"`

	// The message text
	Message *string `mandatory:"true" json:"message"`
}

func (m Message) String() string {
	return common.PointerString(m)
}

// MessageTypeEnum Enum with underlying type: string
type MessageTypeEnum string

// Set of constants representing the allowable values for MessageTypeEnum
const (
	MessageTypeError   MessageTypeEnum = "ERROR"
	MessageTypeWarning MessageTypeEnum = "WARNING"
	MessageTypeInfo    MessageTypeEnum = "INFO"
)

var mappingMessageType = map[string]MessageTypeEnum{
	"ERROR":   MessageTypeError,
	"WARNING": MessageTypeWarning,
	"INFO":    MessageTypeInfo,
}

// GetMessageTypeEnumValues Enumerates the set of values for MessageTypeEnum
func GetMessageTypeEnumValues() []MessageTypeEnum {
	values := make([]MessageTypeEnum, 0)
	for _, v := range mappingMessageType {
		values = append(values, v)
	}
	return values
}
