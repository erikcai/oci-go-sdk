// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
