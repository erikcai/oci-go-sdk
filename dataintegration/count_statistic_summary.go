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

// CountStatisticSummary Detail of object.
type CountStatisticSummary struct {

	// the type of object for the object count statistic.
	ObjectType CountStatisticSummaryObjectTypeEnum `mandatory:"false" json:"objectType,omitempty"`

	// the value for the object count statistic.
	ObjectCount *string `mandatory:"false" json:"objectCount"`
}

func (m CountStatisticSummary) String() string {
	return common.PointerString(m)
}

// CountStatisticSummaryObjectTypeEnum Enum with underlying type: string
type CountStatisticSummaryObjectTypeEnum string

// Set of constants representing the allowable values for CountStatisticSummaryObjectTypeEnum
const (
	CountStatisticSummaryObjectTypeProjects     CountStatisticSummaryObjectTypeEnum = "PROJECTS"
	CountStatisticSummaryObjectTypeFolders      CountStatisticSummaryObjectTypeEnum = "FOLDERS"
	CountStatisticSummaryObjectTypeDataflows    CountStatisticSummaryObjectTypeEnum = "DATAFLOWS"
	CountStatisticSummaryObjectTypeDataassets   CountStatisticSummaryObjectTypeEnum = "DATAASSETS"
	CountStatisticSummaryObjectTypeConnections  CountStatisticSummaryObjectTypeEnum = "CONNECTIONS"
	CountStatisticSummaryObjectTypeTasks        CountStatisticSummaryObjectTypeEnum = "TASKS"
	CountStatisticSummaryObjectTypeApplications CountStatisticSummaryObjectTypeEnum = "APPLICATIONS"
)

var mappingCountStatisticSummaryObjectType = map[string]CountStatisticSummaryObjectTypeEnum{
	"PROJECTS":     CountStatisticSummaryObjectTypeProjects,
	"FOLDERS":      CountStatisticSummaryObjectTypeFolders,
	"DATAFLOWS":    CountStatisticSummaryObjectTypeDataflows,
	"DATAASSETS":   CountStatisticSummaryObjectTypeDataassets,
	"CONNECTIONS":  CountStatisticSummaryObjectTypeConnections,
	"TASKS":        CountStatisticSummaryObjectTypeTasks,
	"APPLICATIONS": CountStatisticSummaryObjectTypeApplications,
}

// GetCountStatisticSummaryObjectTypeEnumValues Enumerates the set of values for CountStatisticSummaryObjectTypeEnum
func GetCountStatisticSummaryObjectTypeEnumValues() []CountStatisticSummaryObjectTypeEnum {
	values := make([]CountStatisticSummaryObjectTypeEnum, 0)
	for _, v := range mappingCountStatisticSummaryObjectType {
		values = append(values, v)
	}
	return values
}
