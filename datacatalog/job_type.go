// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

// JobTypeEnum Enum with underlying type: string
type JobTypeEnum string

// Set of constants representing the allowable values for JobTypeEnum
const (
	JobTypeHarvest   JobTypeEnum = "HARVEST"
	JobTypeProfiling JobTypeEnum = "PROFILING"
	JobTypeSampling  JobTypeEnum = "SAMPLING"
	JobTypePreview   JobTypeEnum = "PREVIEW"
	JobTypeImport    JobTypeEnum = "IMPORT"
	JobTypeExport    JobTypeEnum = "EXPORT"
)

var mappingJobType = map[string]JobTypeEnum{
	"HARVEST":   JobTypeHarvest,
	"PROFILING": JobTypeProfiling,
	"SAMPLING":  JobTypeSampling,
	"PREVIEW":   JobTypePreview,
	"IMPORT":    JobTypeImport,
	"EXPORT":    JobTypeExport,
}

// GetJobTypeEnumValues Enumerates the set of values for JobTypeEnum
func GetJobTypeEnumValues() []JobTypeEnum {
	values := make([]JobTypeEnum, 0)
	for _, v := range mappingJobType {
		values = append(values, v)
	}
	return values
}
