// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Orchestration
//
// Orchestration Maestro API
//

package orchestration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LogEntry Logs from executing a job
type LogEntry struct {

	// Type of log for which this is an entry
	Type LogEntryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The severity level of this entry
	Level LogEntryLevelEnum `mandatory:"false" json:"level,omitempty"`

	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	Message *string `mandatory:"false" json:"message"`
}

func (m LogEntry) String() string {
	return common.PointerString(m)
}

// LogEntryTypeEnum Enum with underlying type: string
type LogEntryTypeEnum string

// Set of constants representing the allowable values for LogEntryType
const (
	LogEntryTypeConsole LogEntryTypeEnum = "TERRAFORM_CONSOLE"
	LogEntryTypeLog     LogEntryTypeEnum = "TERRAFORM_LOG"
)

var mappingLogEntryType = map[string]LogEntryTypeEnum{
	"TERRAFORM_CONSOLE": LogEntryTypeConsole,
	"TERRAFORM_LOG":     LogEntryTypeLog,
}

// GetLogEntryTypeEnumValues Enumerates the set of values for LogEntryType
func GetLogEntryTypeEnumValues() []LogEntryTypeEnum {
	values := make([]LogEntryTypeEnum, 0)
	for _, v := range mappingLogEntryType {
		values = append(values, v)
	}
	return values
}

// LogEntryLevelEnum Enum with underlying type: string
type LogEntryLevelEnum string

// Set of constants representing the allowable values for LogEntryLevel
const (
	LogEntryLevelTrace LogEntryLevelEnum = "TRACE"
	LogEntryLevelDebug LogEntryLevelEnum = "DEBUG"
	LogEntryLevelInfo  LogEntryLevelEnum = "INFO"
	LogEntryLevelWarn  LogEntryLevelEnum = "WARN"
	LogEntryLevelError LogEntryLevelEnum = "ERROR"
	LogEntryLevelFatal LogEntryLevelEnum = "FATAL"
)

var mappingLogEntryLevel = map[string]LogEntryLevelEnum{
	"TRACE": LogEntryLevelTrace,
	"DEBUG": LogEntryLevelDebug,
	"INFO":  LogEntryLevelInfo,
	"WARN":  LogEntryLevelWarn,
	"ERROR": LogEntryLevelError,
	"FATAL": LogEntryLevelFatal,
}

// GetLogEntryLevelEnumValues Enumerates the set of values for LogEntryLevel
func GetLogEntryLevelEnumValues() []LogEntryLevelEnum {
	values := make([]LogEntryLevelEnum, 0)
	for _, v := range mappingLogEntryLevel {
		values = append(values, v)
	}
	return values
}
