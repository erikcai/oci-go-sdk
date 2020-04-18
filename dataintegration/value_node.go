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

// ValueNode auto generated description
type ValueNode struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// parsedString
	ParsedString *string `mandatory:"false" json:"parsedString"`

	// startPosition
	StartPosition *int `mandatory:"false" json:"startPosition"`

	// endPosition
	EndPosition *int `mandatory:"false" json:"endPosition"`

	ValueType *DataType `mandatory:"false" json:"valueType"`

	// auto generated description
	Value *interface{} `mandatory:"false" json:"value"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// type
	Type ValueNodeTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m ValueNode) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ValueNode) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeValueNode ValueNode
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeValueNode
	}{
		"VALUE_NODE",
		(MarshalTypeValueNode)(m),
	}

	return json.Marshal(&s)
}

// ValueNodeTypeEnum Enum with underlying type: string
type ValueNodeTypeEnum string

// Set of constants representing the allowable values for ValueNodeTypeEnum
const (
	ValueNodeTypeOther                    ValueNodeTypeEnum = "OTHER"
	ValueNodeTypeSelect                   ValueNodeTypeEnum = "SELECT"
	ValueNodeTypeJoin                     ValueNodeTypeEnum = "JOIN"
	ValueNodeTypeIdentifier               ValueNodeTypeEnum = "IDENTIFIER"
	ValueNodeTypeLiteral                  ValueNodeTypeEnum = "LITERAL"
	ValueNodeTypeOtherFunction            ValueNodeTypeEnum = "OTHER_FUNCTION"
	ValueNodeTypeExplain                  ValueNodeTypeEnum = "EXPLAIN"
	ValueNodeTypeDescribeSchema           ValueNodeTypeEnum = "DESCRIBE_SCHEMA"
	ValueNodeTypeDescribeTable            ValueNodeTypeEnum = "DESCRIBE_TABLE"
	ValueNodeTypeInsert                   ValueNodeTypeEnum = "INSERT"
	ValueNodeTypeDelete                   ValueNodeTypeEnum = "DELETE"
	ValueNodeTypeUpdate                   ValueNodeTypeEnum = "UPDATE"
	ValueNodeTypeSetOption                ValueNodeTypeEnum = "SET_OPTION"
	ValueNodeTypeDynamicParam             ValueNodeTypeEnum = "DYNAMIC_PARAM"
	ValueNodeTypeOrderBy                  ValueNodeTypeEnum = "ORDER_BY"
	ValueNodeTypeWith                     ValueNodeTypeEnum = "WITH"
	ValueNodeTypeWithItem                 ValueNodeTypeEnum = "WITH_ITEM"
	ValueNodeTypeUnion                    ValueNodeTypeEnum = "UNION"
	ValueNodeTypeExcept                   ValueNodeTypeEnum = "EXCEPT"
	ValueNodeTypeIntersect                ValueNodeTypeEnum = "INTERSECT"
	ValueNodeTypeAs                       ValueNodeTypeEnum = "AS"
	ValueNodeTypeArgumentAssignment       ValueNodeTypeEnum = "ARGUMENT_ASSIGNMENT"
	ValueNodeTypeDefault                  ValueNodeTypeEnum = "DEFAULT"
	ValueNodeTypeOver                     ValueNodeTypeEnum = "OVER"
	ValueNodeTypeFilter                   ValueNodeTypeEnum = "FILTER"
	ValueNodeTypeWindow                   ValueNodeTypeEnum = "WINDOW"
	ValueNodeTypeMerge                    ValueNodeTypeEnum = "MERGE"
	ValueNodeTypeTablesample              ValueNodeTypeEnum = "TABLESAMPLE"
	ValueNodeTypeMatchRecognize           ValueNodeTypeEnum = "MATCH_RECOGNIZE"
	ValueNodeTypeTimes                    ValueNodeTypeEnum = "TIMES"
	ValueNodeTypeDivide                   ValueNodeTypeEnum = "DIVIDE"
	ValueNodeTypeMod                      ValueNodeTypeEnum = "MOD"
	ValueNodeTypePlus                     ValueNodeTypeEnum = "PLUS"
	ValueNodeTypeMinus                    ValueNodeTypeEnum = "MINUS"
	ValueNodeTypePatternAlter             ValueNodeTypeEnum = "PATTERN_ALTER"
	ValueNodeTypePatternConcat            ValueNodeTypeEnum = "PATTERN_CONCAT"
	ValueNodeTypeIn                       ValueNodeTypeEnum = "IN"
	ValueNodeTypeNotIn                    ValueNodeTypeEnum = "NOT_IN"
	ValueNodeTypeLessThan                 ValueNodeTypeEnum = "LESS_THAN"
	ValueNodeTypeGreaterThan              ValueNodeTypeEnum = "GREATER_THAN"
	ValueNodeTypeLessThanOrEqual          ValueNodeTypeEnum = "LESS_THAN_OR_EQUAL"
	ValueNodeTypeGreaterThanOrEqual       ValueNodeTypeEnum = "GREATER_THAN_OR_EQUAL"
	ValueNodeTypeEquals                   ValueNodeTypeEnum = "EQUALS"
	ValueNodeTypeNotEquals                ValueNodeTypeEnum = "NOT_EQUALS"
	ValueNodeTypeIsDistinctFrom           ValueNodeTypeEnum = "IS_DISTINCT_FROM"
	ValueNodeTypeIsNotDistinctFrom        ValueNodeTypeEnum = "IS_NOT_DISTINCT_FROM"
	ValueNodeTypeOr                       ValueNodeTypeEnum = "OR"
	ValueNodeTypeAnd                      ValueNodeTypeEnum = "AND"
	ValueNodeTypeDot                      ValueNodeTypeEnum = "DOT"
	ValueNodeTypeOverlaps                 ValueNodeTypeEnum = "OVERLAPS"
	ValueNodeTypeContains                 ValueNodeTypeEnum = "CONTAINS"
	ValueNodeTypePrecedes                 ValueNodeTypeEnum = "PRECEDES"
	ValueNodeTypeImmediatelyPrecedes      ValueNodeTypeEnum = "IMMEDIATELY_PRECEDES"
	ValueNodeTypeSucceeds                 ValueNodeTypeEnum = "SUCCEEDS"
	ValueNodeTypeImmediatelySucceeds      ValueNodeTypeEnum = "IMMEDIATELY_SUCCEEDS"
	ValueNodeTypePeriodEquals             ValueNodeTypeEnum = "PERIOD_EQUALS"
	ValueNodeTypeLike                     ValueNodeTypeEnum = "LIKE"
	ValueNodeTypeSimilar                  ValueNodeTypeEnum = "SIMILAR"
	ValueNodeTypeBetween                  ValueNodeTypeEnum = "BETWEEN"
	ValueNodeTypeCase                     ValueNodeTypeEnum = "CASE"
	ValueNodeTypeNullif                   ValueNodeTypeEnum = "NULLIF"
	ValueNodeTypeCoalesce                 ValueNodeTypeEnum = "COALESCE"
	ValueNodeTypeDecode                   ValueNodeTypeEnum = "DECODE"
	ValueNodeTypeNvl                      ValueNodeTypeEnum = "NVL"
	ValueNodeTypeGreatest                 ValueNodeTypeEnum = "GREATEST"
	ValueNodeTypeLeast                    ValueNodeTypeEnum = "LEAST"
	ValueNodeTypeTimestampAdd             ValueNodeTypeEnum = "TIMESTAMP_ADD"
	ValueNodeTypeTimestampDiff            ValueNodeTypeEnum = "TIMESTAMP_DIFF"
	ValueNodeTypeNot                      ValueNodeTypeEnum = "NOT"
	ValueNodeTypePlusPrefix               ValueNodeTypeEnum = "PLUS_PREFIX"
	ValueNodeTypeMinusPrefix              ValueNodeTypeEnum = "MINUS_PREFIX"
	ValueNodeTypeExists                   ValueNodeTypeEnum = "EXISTS"
	ValueNodeTypeSome                     ValueNodeTypeEnum = "SOME"
	ValueNodeTypeAll                      ValueNodeTypeEnum = "ALL"
	ValueNodeTypeValues                   ValueNodeTypeEnum = "VALUES"
	ValueNodeTypeExplicitTable            ValueNodeTypeEnum = "EXPLICIT_TABLE"
	ValueNodeTypeScalarQuery              ValueNodeTypeEnum = "SCALAR_QUERY"
	ValueNodeTypeProcedureCall            ValueNodeTypeEnum = "PROCEDURE_CALL"
	ValueNodeTypeNewSpecification         ValueNodeTypeEnum = "NEW_SPECIFICATION"
	ValueNodeTypeFinal                    ValueNodeTypeEnum = "FINAL"
	ValueNodeTypeRunning                  ValueNodeTypeEnum = "RUNNING"
	ValueNodeTypePrev                     ValueNodeTypeEnum = "PREV"
	ValueNodeTypeNext                     ValueNodeTypeEnum = "NEXT"
	ValueNodeTypeFirst                    ValueNodeTypeEnum = "FIRST"
	ValueNodeTypeLast                     ValueNodeTypeEnum = "LAST"
	ValueNodeTypeClassifier               ValueNodeTypeEnum = "CLASSIFIER"
	ValueNodeTypeMatchNumber              ValueNodeTypeEnum = "MATCH_NUMBER"
	ValueNodeTypeSkipToFirst              ValueNodeTypeEnum = "SKIP_TO_FIRST"
	ValueNodeTypeSkipToLast               ValueNodeTypeEnum = "SKIP_TO_LAST"
	ValueNodeTypeDescending               ValueNodeTypeEnum = "DESCENDING"
	ValueNodeTypeNullsFirst               ValueNodeTypeEnum = "NULLS_FIRST"
	ValueNodeTypeNullsLast                ValueNodeTypeEnum = "NULLS_LAST"
	ValueNodeTypeIsTrue                   ValueNodeTypeEnum = "IS_TRUE"
	ValueNodeTypeIsFalse                  ValueNodeTypeEnum = "IS_FALSE"
	ValueNodeTypeIsNotTrue                ValueNodeTypeEnum = "IS_NOT_TRUE"
	ValueNodeTypeIsNotFalse               ValueNodeTypeEnum = "IS_NOT_FALSE"
	ValueNodeTypeIsUnknown                ValueNodeTypeEnum = "IS_UNKNOWN"
	ValueNodeTypeIsNull                   ValueNodeTypeEnum = "IS_NULL"
	ValueNodeTypeIsNotNull                ValueNodeTypeEnum = "IS_NOT_NULL"
	ValueNodeTypePreceding                ValueNodeTypeEnum = "PRECEDING"
	ValueNodeTypeFollowing                ValueNodeTypeEnum = "FOLLOWING"
	ValueNodeTypeFieldAccess              ValueNodeTypeEnum = "FIELD_ACCESS"
	ValueNodeTypeInputRef                 ValueNodeTypeEnum = "INPUT_REF"
	ValueNodeTypeTableInputRef            ValueNodeTypeEnum = "TABLE_INPUT_REF"
	ValueNodeTypePatternInputRef          ValueNodeTypeEnum = "PATTERN_INPUT_REF"
	ValueNodeTypeLocalRef                 ValueNodeTypeEnum = "LOCAL_REF"
	ValueNodeTypeCorrelVariable           ValueNodeTypeEnum = "CORREL_VARIABLE"
	ValueNodeTypePatternQuantifier        ValueNodeTypeEnum = "PATTERN_QUANTIFIER"
	ValueNodeTypeRow                      ValueNodeTypeEnum = "ROW"
	ValueNodeTypeColumnList               ValueNodeTypeEnum = "COLUMN_LIST"
	ValueNodeTypeCast                     ValueNodeTypeEnum = "CAST"
	ValueNodeTypeNextValue                ValueNodeTypeEnum = "NEXT_VALUE"
	ValueNodeTypeCurrentValue             ValueNodeTypeEnum = "CURRENT_VALUE"
	ValueNodeTypeFloor                    ValueNodeTypeEnum = "FLOOR"
	ValueNodeTypeCeil                     ValueNodeTypeEnum = "CEIL"
	ValueNodeTypeTrim                     ValueNodeTypeEnum = "TRIM"
	ValueNodeTypeLtrim                    ValueNodeTypeEnum = "LTRIM"
	ValueNodeTypeRtrim                    ValueNodeTypeEnum = "RTRIM"
	ValueNodeTypeExtract                  ValueNodeTypeEnum = "EXTRACT"
	ValueNodeTypeJdbcFn                   ValueNodeTypeEnum = "JDBC_FN"
	ValueNodeTypeMultisetValueConstructor ValueNodeTypeEnum = "MULTISET_VALUE_CONSTRUCTOR"
	ValueNodeTypeMultisetQueryConstructor ValueNodeTypeEnum = "MULTISET_QUERY_CONSTRUCTOR"
	ValueNodeTypeUnnest                   ValueNodeTypeEnum = "UNNEST"
	ValueNodeTypeLateral                  ValueNodeTypeEnum = "LATERAL"
	ValueNodeTypeCollectionTable          ValueNodeTypeEnum = "COLLECTION_TABLE"
	ValueNodeTypeArrayValueConstructor    ValueNodeTypeEnum = "ARRAY_VALUE_CONSTRUCTOR"
	ValueNodeTypeArrayQueryConstructor    ValueNodeTypeEnum = "ARRAY_QUERY_CONSTRUCTOR"
	ValueNodeTypeMapValueConstructor      ValueNodeTypeEnum = "MAP_VALUE_CONSTRUCTOR"
	ValueNodeTypeMapQueryConstructor      ValueNodeTypeEnum = "MAP_QUERY_CONSTRUCTOR"
	ValueNodeTypeCursor                   ValueNodeTypeEnum = "CURSOR"
	ValueNodeTypeLiteralChain             ValueNodeTypeEnum = "LITERAL_CHAIN"
	ValueNodeTypeEscape                   ValueNodeTypeEnum = "ESCAPE"
	ValueNodeTypeReinterpret              ValueNodeTypeEnum = "REINTERPRET"
	ValueNodeTypeExtend                   ValueNodeTypeEnum = "EXTEND"
	ValueNodeTypeCube                     ValueNodeTypeEnum = "CUBE"
	ValueNodeTypeRollup                   ValueNodeTypeEnum = "ROLLUP"
	ValueNodeTypeGroupingSets             ValueNodeTypeEnum = "GROUPING_SETS"
	ValueNodeTypeGrouping                 ValueNodeTypeEnum = "GROUPING"
	ValueNodeTypeGroupingId               ValueNodeTypeEnum = "GROUPING_ID"
	ValueNodeTypeGroupId                  ValueNodeTypeEnum = "GROUP_ID"
	ValueNodeTypePatternPermute           ValueNodeTypeEnum = "PATTERN_PERMUTE"
	ValueNodeTypePatternExcluded          ValueNodeTypeEnum = "PATTERN_EXCLUDED"
	ValueNodeTypeCount                    ValueNodeTypeEnum = "COUNT"
	ValueNodeTypeSum                      ValueNodeTypeEnum = "SUM"
	ValueNodeTypeSum0                     ValueNodeTypeEnum = "SUM0"
	ValueNodeTypeMin                      ValueNodeTypeEnum = "MIN"
	ValueNodeTypeMax                      ValueNodeTypeEnum = "MAX"
	ValueNodeTypeLead                     ValueNodeTypeEnum = "LEAD"
	ValueNodeTypeLag                      ValueNodeTypeEnum = "LAG"
	ValueNodeTypeFirstValue               ValueNodeTypeEnum = "FIRST_VALUE"
	ValueNodeTypeLastValue                ValueNodeTypeEnum = "LAST_VALUE"
	ValueNodeTypeCovarPop                 ValueNodeTypeEnum = "COVAR_POP"
	ValueNodeTypeCovarSamp                ValueNodeTypeEnum = "COVAR_SAMP"
	ValueNodeTypeRegrSxx                  ValueNodeTypeEnum = "REGR_SXX"
	ValueNodeTypeRegrSyy                  ValueNodeTypeEnum = "REGR_SYY"
	ValueNodeTypeAvg                      ValueNodeTypeEnum = "AVG"
	ValueNodeTypeStddevPop                ValueNodeTypeEnum = "STDDEV_POP"
	ValueNodeTypeStddevSamp               ValueNodeTypeEnum = "STDDEV_SAMP"
	ValueNodeTypeVarPop                   ValueNodeTypeEnum = "VAR_POP"
	ValueNodeTypeVarSamp                  ValueNodeTypeEnum = "VAR_SAMP"
	ValueNodeTypeNtile                    ValueNodeTypeEnum = "NTILE"
	ValueNodeTypeCollect                  ValueNodeTypeEnum = "COLLECT"
	ValueNodeTypeFusion                   ValueNodeTypeEnum = "FUSION"
	ValueNodeTypeSingleValue              ValueNodeTypeEnum = "SINGLE_VALUE"
	ValueNodeTypeRowNumber                ValueNodeTypeEnum = "ROW_NUMBER"
	ValueNodeTypeRank                     ValueNodeTypeEnum = "RANK"
	ValueNodeTypePercentRank              ValueNodeTypeEnum = "PERCENT_RANK"
	ValueNodeTypeDenseRank                ValueNodeTypeEnum = "DENSE_RANK"
	ValueNodeTypeCumeDist                 ValueNodeTypeEnum = "CUME_DIST"
	ValueNodeTypeTumble                   ValueNodeTypeEnum = "TUMBLE"
	ValueNodeTypeTumbleStart              ValueNodeTypeEnum = "TUMBLE_START"
	ValueNodeTypeTumbleEnd                ValueNodeTypeEnum = "TUMBLE_END"
	ValueNodeTypeHop                      ValueNodeTypeEnum = "HOP"
	ValueNodeTypeHopStart                 ValueNodeTypeEnum = "HOP_START"
	ValueNodeTypeHopEnd                   ValueNodeTypeEnum = "HOP_END"
	ValueNodeTypeSession                  ValueNodeTypeEnum = "SESSION"
	ValueNodeTypeSessionStart             ValueNodeTypeEnum = "SESSION_START"
	ValueNodeTypeSessionEnd               ValueNodeTypeEnum = "SESSION_END"
	ValueNodeTypeColumnDecl               ValueNodeTypeEnum = "COLUMN_DECL"
	ValueNodeTypeCheck                    ValueNodeTypeEnum = "CHECK"
	ValueNodeTypeUnique                   ValueNodeTypeEnum = "UNIQUE"
	ValueNodeTypePrimaryKey               ValueNodeTypeEnum = "PRIMARY_KEY"
	ValueNodeTypeForeignKey               ValueNodeTypeEnum = "FOREIGN_KEY"
	ValueNodeTypeCommit                   ValueNodeTypeEnum = "COMMIT"
	ValueNodeTypeRollback                 ValueNodeTypeEnum = "ROLLBACK"
	ValueNodeTypeAlterSession             ValueNodeTypeEnum = "ALTER_SESSION"
	ValueNodeTypeCreateSchema             ValueNodeTypeEnum = "CREATE_SCHEMA"
	ValueNodeTypeCreateForeignSchema      ValueNodeTypeEnum = "CREATE_FOREIGN_SCHEMA"
	ValueNodeTypeDropSchema               ValueNodeTypeEnum = "DROP_SCHEMA"
	ValueNodeTypeCreateTable              ValueNodeTypeEnum = "CREATE_TABLE"
	ValueNodeTypeAlterTable               ValueNodeTypeEnum = "ALTER_TABLE"
	ValueNodeTypeDropTable                ValueNodeTypeEnum = "DROP_TABLE"
	ValueNodeTypeCreateView               ValueNodeTypeEnum = "CREATE_VIEW"
	ValueNodeTypeAlterView                ValueNodeTypeEnum = "ALTER_VIEW"
	ValueNodeTypeDropView                 ValueNodeTypeEnum = "DROP_VIEW"
	ValueNodeTypeCreateMaterializedView   ValueNodeTypeEnum = "CREATE_MATERIALIZED_VIEW"
	ValueNodeTypeAlterMaterializedView    ValueNodeTypeEnum = "ALTER_MATERIALIZED_VIEW"
	ValueNodeTypeDropMaterializedView     ValueNodeTypeEnum = "DROP_MATERIALIZED_VIEW"
	ValueNodeTypeCreateSequence           ValueNodeTypeEnum = "CREATE_SEQUENCE"
	ValueNodeTypeAlterSequence            ValueNodeTypeEnum = "ALTER_SEQUENCE"
	ValueNodeTypeDropSequence             ValueNodeTypeEnum = "DROP_SEQUENCE"
	ValueNodeTypeCreateIndex              ValueNodeTypeEnum = "CREATE_INDEX"
	ValueNodeTypeAlterIndex               ValueNodeTypeEnum = "ALTER_INDEX"
	ValueNodeTypeDropIndex                ValueNodeTypeEnum = "DROP_INDEX"
	ValueNodeTypeOtherDdl                 ValueNodeTypeEnum = "OTHER_DDL"
)

var mappingValueNodeType = map[string]ValueNodeTypeEnum{
	"OTHER":           ValueNodeTypeOther,
	"SELECT":          ValueNodeTypeSelect,
	"JOIN":            ValueNodeTypeJoin,
	"IDENTIFIER":      ValueNodeTypeIdentifier,
	"LITERAL":         ValueNodeTypeLiteral,
	"OTHER_FUNCTION":  ValueNodeTypeOtherFunction,
	"EXPLAIN":         ValueNodeTypeExplain,
	"DESCRIBE_SCHEMA": ValueNodeTypeDescribeSchema,
	"DESCRIBE_TABLE":  ValueNodeTypeDescribeTable,
	"INSERT":          ValueNodeTypeInsert,
	"DELETE":          ValueNodeTypeDelete,
	"UPDATE":          ValueNodeTypeUpdate,
	"SET_OPTION":      ValueNodeTypeSetOption,
	"DYNAMIC_PARAM":   ValueNodeTypeDynamicParam,
	"ORDER_BY":        ValueNodeTypeOrderBy,
	"WITH":            ValueNodeTypeWith,
	"WITH_ITEM":       ValueNodeTypeWithItem,
	"UNION":           ValueNodeTypeUnion,
	"EXCEPT":          ValueNodeTypeExcept,
	"INTERSECT":       ValueNodeTypeIntersect,
	"AS":              ValueNodeTypeAs,
	"ARGUMENT_ASSIGNMENT":   ValueNodeTypeArgumentAssignment,
	"DEFAULT":               ValueNodeTypeDefault,
	"OVER":                  ValueNodeTypeOver,
	"FILTER":                ValueNodeTypeFilter,
	"WINDOW":                ValueNodeTypeWindow,
	"MERGE":                 ValueNodeTypeMerge,
	"TABLESAMPLE":           ValueNodeTypeTablesample,
	"MATCH_RECOGNIZE":       ValueNodeTypeMatchRecognize,
	"TIMES":                 ValueNodeTypeTimes,
	"DIVIDE":                ValueNodeTypeDivide,
	"MOD":                   ValueNodeTypeMod,
	"PLUS":                  ValueNodeTypePlus,
	"MINUS":                 ValueNodeTypeMinus,
	"PATTERN_ALTER":         ValueNodeTypePatternAlter,
	"PATTERN_CONCAT":        ValueNodeTypePatternConcat,
	"IN":                    ValueNodeTypeIn,
	"NOT_IN":                ValueNodeTypeNotIn,
	"LESS_THAN":             ValueNodeTypeLessThan,
	"GREATER_THAN":          ValueNodeTypeGreaterThan,
	"LESS_THAN_OR_EQUAL":    ValueNodeTypeLessThanOrEqual,
	"GREATER_THAN_OR_EQUAL": ValueNodeTypeGreaterThanOrEqual,
	"EQUALS":                ValueNodeTypeEquals,
	"NOT_EQUALS":            ValueNodeTypeNotEquals,
	"IS_DISTINCT_FROM":      ValueNodeTypeIsDistinctFrom,
	"IS_NOT_DISTINCT_FROM":  ValueNodeTypeIsNotDistinctFrom,
	"OR":                         ValueNodeTypeOr,
	"AND":                        ValueNodeTypeAnd,
	"DOT":                        ValueNodeTypeDot,
	"OVERLAPS":                   ValueNodeTypeOverlaps,
	"CONTAINS":                   ValueNodeTypeContains,
	"PRECEDES":                   ValueNodeTypePrecedes,
	"IMMEDIATELY_PRECEDES":       ValueNodeTypeImmediatelyPrecedes,
	"SUCCEEDS":                   ValueNodeTypeSucceeds,
	"IMMEDIATELY_SUCCEEDS":       ValueNodeTypeImmediatelySucceeds,
	"PERIOD_EQUALS":              ValueNodeTypePeriodEquals,
	"LIKE":                       ValueNodeTypeLike,
	"SIMILAR":                    ValueNodeTypeSimilar,
	"BETWEEN":                    ValueNodeTypeBetween,
	"CASE":                       ValueNodeTypeCase,
	"NULLIF":                     ValueNodeTypeNullif,
	"COALESCE":                   ValueNodeTypeCoalesce,
	"DECODE":                     ValueNodeTypeDecode,
	"NVL":                        ValueNodeTypeNvl,
	"GREATEST":                   ValueNodeTypeGreatest,
	"LEAST":                      ValueNodeTypeLeast,
	"TIMESTAMP_ADD":              ValueNodeTypeTimestampAdd,
	"TIMESTAMP_DIFF":             ValueNodeTypeTimestampDiff,
	"NOT":                        ValueNodeTypeNot,
	"PLUS_PREFIX":                ValueNodeTypePlusPrefix,
	"MINUS_PREFIX":               ValueNodeTypeMinusPrefix,
	"EXISTS":                     ValueNodeTypeExists,
	"SOME":                       ValueNodeTypeSome,
	"ALL":                        ValueNodeTypeAll,
	"VALUES":                     ValueNodeTypeValues,
	"EXPLICIT_TABLE":             ValueNodeTypeExplicitTable,
	"SCALAR_QUERY":               ValueNodeTypeScalarQuery,
	"PROCEDURE_CALL":             ValueNodeTypeProcedureCall,
	"NEW_SPECIFICATION":          ValueNodeTypeNewSpecification,
	"FINAL":                      ValueNodeTypeFinal,
	"RUNNING":                    ValueNodeTypeRunning,
	"PREV":                       ValueNodeTypePrev,
	"NEXT":                       ValueNodeTypeNext,
	"FIRST":                      ValueNodeTypeFirst,
	"LAST":                       ValueNodeTypeLast,
	"CLASSIFIER":                 ValueNodeTypeClassifier,
	"MATCH_NUMBER":               ValueNodeTypeMatchNumber,
	"SKIP_TO_FIRST":              ValueNodeTypeSkipToFirst,
	"SKIP_TO_LAST":               ValueNodeTypeSkipToLast,
	"DESCENDING":                 ValueNodeTypeDescending,
	"NULLS_FIRST":                ValueNodeTypeNullsFirst,
	"NULLS_LAST":                 ValueNodeTypeNullsLast,
	"IS_TRUE":                    ValueNodeTypeIsTrue,
	"IS_FALSE":                   ValueNodeTypeIsFalse,
	"IS_NOT_TRUE":                ValueNodeTypeIsNotTrue,
	"IS_NOT_FALSE":               ValueNodeTypeIsNotFalse,
	"IS_UNKNOWN":                 ValueNodeTypeIsUnknown,
	"IS_NULL":                    ValueNodeTypeIsNull,
	"IS_NOT_NULL":                ValueNodeTypeIsNotNull,
	"PRECEDING":                  ValueNodeTypePreceding,
	"FOLLOWING":                  ValueNodeTypeFollowing,
	"FIELD_ACCESS":               ValueNodeTypeFieldAccess,
	"INPUT_REF":                  ValueNodeTypeInputRef,
	"TABLE_INPUT_REF":            ValueNodeTypeTableInputRef,
	"PATTERN_INPUT_REF":          ValueNodeTypePatternInputRef,
	"LOCAL_REF":                  ValueNodeTypeLocalRef,
	"CORREL_VARIABLE":            ValueNodeTypeCorrelVariable,
	"PATTERN_QUANTIFIER":         ValueNodeTypePatternQuantifier,
	"ROW":                        ValueNodeTypeRow,
	"COLUMN_LIST":                ValueNodeTypeColumnList,
	"CAST":                       ValueNodeTypeCast,
	"NEXT_VALUE":                 ValueNodeTypeNextValue,
	"CURRENT_VALUE":              ValueNodeTypeCurrentValue,
	"FLOOR":                      ValueNodeTypeFloor,
	"CEIL":                       ValueNodeTypeCeil,
	"TRIM":                       ValueNodeTypeTrim,
	"LTRIM":                      ValueNodeTypeLtrim,
	"RTRIM":                      ValueNodeTypeRtrim,
	"EXTRACT":                    ValueNodeTypeExtract,
	"JDBC_FN":                    ValueNodeTypeJdbcFn,
	"MULTISET_VALUE_CONSTRUCTOR": ValueNodeTypeMultisetValueConstructor,
	"MULTISET_QUERY_CONSTRUCTOR": ValueNodeTypeMultisetQueryConstructor,
	"UNNEST":                     ValueNodeTypeUnnest,
	"LATERAL":                    ValueNodeTypeLateral,
	"COLLECTION_TABLE":           ValueNodeTypeCollectionTable,
	"ARRAY_VALUE_CONSTRUCTOR":    ValueNodeTypeArrayValueConstructor,
	"ARRAY_QUERY_CONSTRUCTOR":    ValueNodeTypeArrayQueryConstructor,
	"MAP_VALUE_CONSTRUCTOR":      ValueNodeTypeMapValueConstructor,
	"MAP_QUERY_CONSTRUCTOR":      ValueNodeTypeMapQueryConstructor,
	"CURSOR":                     ValueNodeTypeCursor,
	"LITERAL_CHAIN":              ValueNodeTypeLiteralChain,
	"ESCAPE":                     ValueNodeTypeEscape,
	"REINTERPRET":                ValueNodeTypeReinterpret,
	"EXTEND":                     ValueNodeTypeExtend,
	"CUBE":                       ValueNodeTypeCube,
	"ROLLUP":                     ValueNodeTypeRollup,
	"GROUPING_SETS":              ValueNodeTypeGroupingSets,
	"GROUPING":                   ValueNodeTypeGrouping,
	"GROUPING_ID":                ValueNodeTypeGroupingId,
	"GROUP_ID":                   ValueNodeTypeGroupId,
	"PATTERN_PERMUTE":            ValueNodeTypePatternPermute,
	"PATTERN_EXCLUDED":           ValueNodeTypePatternExcluded,
	"COUNT":                      ValueNodeTypeCount,
	"SUM":                        ValueNodeTypeSum,
	"SUM0":                       ValueNodeTypeSum0,
	"MIN":                        ValueNodeTypeMin,
	"MAX":                        ValueNodeTypeMax,
	"LEAD":                       ValueNodeTypeLead,
	"LAG":                        ValueNodeTypeLag,
	"FIRST_VALUE":                ValueNodeTypeFirstValue,
	"LAST_VALUE":                 ValueNodeTypeLastValue,
	"COVAR_POP":                  ValueNodeTypeCovarPop,
	"COVAR_SAMP":                 ValueNodeTypeCovarSamp,
	"REGR_SXX":                   ValueNodeTypeRegrSxx,
	"REGR_SYY":                   ValueNodeTypeRegrSyy,
	"AVG":                        ValueNodeTypeAvg,
	"STDDEV_POP":                 ValueNodeTypeStddevPop,
	"STDDEV_SAMP":                ValueNodeTypeStddevSamp,
	"VAR_POP":                    ValueNodeTypeVarPop,
	"VAR_SAMP":                   ValueNodeTypeVarSamp,
	"NTILE":                      ValueNodeTypeNtile,
	"COLLECT":                    ValueNodeTypeCollect,
	"FUSION":                     ValueNodeTypeFusion,
	"SINGLE_VALUE":               ValueNodeTypeSingleValue,
	"ROW_NUMBER":                 ValueNodeTypeRowNumber,
	"RANK":                       ValueNodeTypeRank,
	"PERCENT_RANK":               ValueNodeTypePercentRank,
	"DENSE_RANK":                 ValueNodeTypeDenseRank,
	"CUME_DIST":                  ValueNodeTypeCumeDist,
	"TUMBLE":                     ValueNodeTypeTumble,
	"TUMBLE_START":               ValueNodeTypeTumbleStart,
	"TUMBLE_END":                 ValueNodeTypeTumbleEnd,
	"HOP":                        ValueNodeTypeHop,
	"HOP_START":                  ValueNodeTypeHopStart,
	"HOP_END":                    ValueNodeTypeHopEnd,
	"SESSION":                    ValueNodeTypeSession,
	"SESSION_START":              ValueNodeTypeSessionStart,
	"SESSION_END":                ValueNodeTypeSessionEnd,
	"COLUMN_DECL":                ValueNodeTypeColumnDecl,
	"CHECK":                      ValueNodeTypeCheck,
	"UNIQUE":                     ValueNodeTypeUnique,
	"PRIMARY_KEY":                ValueNodeTypePrimaryKey,
	"FOREIGN_KEY":                ValueNodeTypeForeignKey,
	"COMMIT":                     ValueNodeTypeCommit,
	"ROLLBACK":                   ValueNodeTypeRollback,
	"ALTER_SESSION":              ValueNodeTypeAlterSession,
	"CREATE_SCHEMA":              ValueNodeTypeCreateSchema,
	"CREATE_FOREIGN_SCHEMA":      ValueNodeTypeCreateForeignSchema,
	"DROP_SCHEMA":                ValueNodeTypeDropSchema,
	"CREATE_TABLE":               ValueNodeTypeCreateTable,
	"ALTER_TABLE":                ValueNodeTypeAlterTable,
	"DROP_TABLE":                 ValueNodeTypeDropTable,
	"CREATE_VIEW":                ValueNodeTypeCreateView,
	"ALTER_VIEW":                 ValueNodeTypeAlterView,
	"DROP_VIEW":                  ValueNodeTypeDropView,
	"CREATE_MATERIALIZED_VIEW":   ValueNodeTypeCreateMaterializedView,
	"ALTER_MATERIALIZED_VIEW":    ValueNodeTypeAlterMaterializedView,
	"DROP_MATERIALIZED_VIEW":     ValueNodeTypeDropMaterializedView,
	"CREATE_SEQUENCE":            ValueNodeTypeCreateSequence,
	"ALTER_SEQUENCE":             ValueNodeTypeAlterSequence,
	"DROP_SEQUENCE":              ValueNodeTypeDropSequence,
	"CREATE_INDEX":               ValueNodeTypeCreateIndex,
	"ALTER_INDEX":                ValueNodeTypeAlterIndex,
	"DROP_INDEX":                 ValueNodeTypeDropIndex,
	"OTHER_DDL":                  ValueNodeTypeOtherDdl,
}

// GetValueNodeTypeEnumValues Enumerates the set of values for ValueNodeTypeEnum
func GetValueNodeTypeEnumValues() []ValueNodeTypeEnum {
	values := make([]ValueNodeTypeEnum, 0)
	for _, v := range mappingValueNodeType {
		values = append(values, v)
	}
	return values
}
