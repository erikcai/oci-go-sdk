// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// BackupCreationTypeEnum Enum with underlying type: string
type BackupCreationTypeEnum string

// Set of constants representing the allowable values for BackupCreationTypeEnum
const (
	BackupCreationTypeManual    BackupCreationTypeEnum = "MANUAL"
	BackupCreationTypeAutomatic BackupCreationTypeEnum = "AUTOMATIC"
)

var mappingBackupCreationType = map[string]BackupCreationTypeEnum{
	"MANUAL":    BackupCreationTypeManual,
	"AUTOMATIC": BackupCreationTypeAutomatic,
}

// GetBackupCreationTypeEnumValues Enumerates the set of values for BackupCreationTypeEnum
func GetBackupCreationTypeEnumValues() []BackupCreationTypeEnum {
	values := make([]BackupCreationTypeEnum, 0)
	for _, v := range mappingBackupCreationType {
		values = append(values, v)
	}
	return values
}
