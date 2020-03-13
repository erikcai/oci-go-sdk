// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CommonLogLevel Peer and Orderer Node log level
type CommonLogLevel struct {
	LogLevel CommonLogLevelLogLevelEnum `mandatory:"false" json:"logLevel,omitempty"`
}

func (m CommonLogLevel) String() string {
	return common.PointerString(m)
}

// CommonLogLevelLogLevelEnum Enum with underlying type: string
type CommonLogLevelLogLevelEnum string

// Set of constants representing the allowable values for CommonLogLevelLogLevelEnum
const (
	CommonLogLevelLogLevelPanic   CommonLogLevelLogLevelEnum = "PANIC"
	CommonLogLevelLogLevelError   CommonLogLevelLogLevelEnum = "ERROR"
	CommonLogLevelLogLevelWarning CommonLogLevelLogLevelEnum = "WARNING"
	CommonLogLevelLogLevelInfo    CommonLogLevelLogLevelEnum = "INFO"
	CommonLogLevelLogLevelNotice  CommonLogLevelLogLevelEnum = "NOTICE"
	CommonLogLevelLogLevelDebug   CommonLogLevelLogLevelEnum = "DEBUG"
)

var mappingCommonLogLevelLogLevel = map[string]CommonLogLevelLogLevelEnum{
	"PANIC":   CommonLogLevelLogLevelPanic,
	"ERROR":   CommonLogLevelLogLevelError,
	"WARNING": CommonLogLevelLogLevelWarning,
	"INFO":    CommonLogLevelLogLevelInfo,
	"NOTICE":  CommonLogLevelLogLevelNotice,
	"DEBUG":   CommonLogLevelLogLevelDebug,
}

// GetCommonLogLevelLogLevelEnumValues Enumerates the set of values for CommonLogLevelLogLevelEnum
func GetCommonLogLevelLogLevelEnumValues() []CommonLogLevelLogLevelEnum {
	values := make([]CommonLogLevelLogLevelEnum, 0)
	for _, v := range mappingCommonLogLevelLogLevel {
		values = append(values, v)
	}
	return values
}
