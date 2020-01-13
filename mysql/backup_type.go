// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// BackupTypeEnum Enum with underlying type: string
type BackupTypeEnum string

// Set of constants representing the allowable values for BackupTypeEnum
const (
	BackupTypeFull        BackupTypeEnum = "FULL"
	BackupTypeIncremental BackupTypeEnum = "INCREMENTAL"
)

var mappingBackupType = map[string]BackupTypeEnum{
	"FULL":        BackupTypeFull,
	"INCREMENTAL": BackupTypeIncremental,
}

// GetBackupTypeEnumValues Enumerates the set of values for BackupTypeEnum
func GetBackupTypeEnumValues() []BackupTypeEnum {
	values := make([]BackupTypeEnum, 0)
	for _, v := range mappingBackupType {
		values = append(values, v)
	}
	return values
}
