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

// ReferenceNode auto generated description
type ReferenceNode struct {

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

	// referenceKind
	ReferenceKind *string `mandatory:"false" json:"referenceKind"`

	// referenceName
	ReferenceName *string `mandatory:"false" json:"referenceName"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// type
	Type ReferenceNodeTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m ReferenceNode) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ReferenceNode) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeReferenceNode ReferenceNode
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeReferenceNode
	}{
		"REFERENCE_NODE",
		(MarshalTypeReferenceNode)(m),
	}

	return json.Marshal(&s)
}

// ReferenceNodeTypeEnum Enum with underlying type: string
type ReferenceNodeTypeEnum string

// Set of constants representing the allowable values for ReferenceNodeTypeEnum
const (
	ReferenceNodeTypeOther                    ReferenceNodeTypeEnum = "OTHER"
	ReferenceNodeTypeSelect                   ReferenceNodeTypeEnum = "SELECT"
	ReferenceNodeTypeJoin                     ReferenceNodeTypeEnum = "JOIN"
	ReferenceNodeTypeIdentifier               ReferenceNodeTypeEnum = "IDENTIFIER"
	ReferenceNodeTypeLiteral                  ReferenceNodeTypeEnum = "LITERAL"
	ReferenceNodeTypeOtherFunction            ReferenceNodeTypeEnum = "OTHER_FUNCTION"
	ReferenceNodeTypeExplain                  ReferenceNodeTypeEnum = "EXPLAIN"
	ReferenceNodeTypeDescribeSchema           ReferenceNodeTypeEnum = "DESCRIBE_SCHEMA"
	ReferenceNodeTypeDescribeTable            ReferenceNodeTypeEnum = "DESCRIBE_TABLE"
	ReferenceNodeTypeInsert                   ReferenceNodeTypeEnum = "INSERT"
	ReferenceNodeTypeDelete                   ReferenceNodeTypeEnum = "DELETE"
	ReferenceNodeTypeUpdate                   ReferenceNodeTypeEnum = "UPDATE"
	ReferenceNodeTypeSetOption                ReferenceNodeTypeEnum = "SET_OPTION"
	ReferenceNodeTypeDynamicParam             ReferenceNodeTypeEnum = "DYNAMIC_PARAM"
	ReferenceNodeTypeOrderBy                  ReferenceNodeTypeEnum = "ORDER_BY"
	ReferenceNodeTypeWith                     ReferenceNodeTypeEnum = "WITH"
	ReferenceNodeTypeWithItem                 ReferenceNodeTypeEnum = "WITH_ITEM"
	ReferenceNodeTypeUnion                    ReferenceNodeTypeEnum = "UNION"
	ReferenceNodeTypeExcept                   ReferenceNodeTypeEnum = "EXCEPT"
	ReferenceNodeTypeIntersect                ReferenceNodeTypeEnum = "INTERSECT"
	ReferenceNodeTypeAs                       ReferenceNodeTypeEnum = "AS"
	ReferenceNodeTypeArgumentAssignment       ReferenceNodeTypeEnum = "ARGUMENT_ASSIGNMENT"
	ReferenceNodeTypeDefault                  ReferenceNodeTypeEnum = "DEFAULT"
	ReferenceNodeTypeOver                     ReferenceNodeTypeEnum = "OVER"
	ReferenceNodeTypeFilter                   ReferenceNodeTypeEnum = "FILTER"
	ReferenceNodeTypeWindow                   ReferenceNodeTypeEnum = "WINDOW"
	ReferenceNodeTypeMerge                    ReferenceNodeTypeEnum = "MERGE"
	ReferenceNodeTypeTablesample              ReferenceNodeTypeEnum = "TABLESAMPLE"
	ReferenceNodeTypeMatchRecognize           ReferenceNodeTypeEnum = "MATCH_RECOGNIZE"
	ReferenceNodeTypeTimes                    ReferenceNodeTypeEnum = "TIMES"
	ReferenceNodeTypeDivide                   ReferenceNodeTypeEnum = "DIVIDE"
	ReferenceNodeTypeMod                      ReferenceNodeTypeEnum = "MOD"
	ReferenceNodeTypePlus                     ReferenceNodeTypeEnum = "PLUS"
	ReferenceNodeTypeMinus                    ReferenceNodeTypeEnum = "MINUS"
	ReferenceNodeTypePatternAlter             ReferenceNodeTypeEnum = "PATTERN_ALTER"
	ReferenceNodeTypePatternConcat            ReferenceNodeTypeEnum = "PATTERN_CONCAT"
	ReferenceNodeTypeIn                       ReferenceNodeTypeEnum = "IN"
	ReferenceNodeTypeNotIn                    ReferenceNodeTypeEnum = "NOT_IN"
	ReferenceNodeTypeLessThan                 ReferenceNodeTypeEnum = "LESS_THAN"
	ReferenceNodeTypeGreaterThan              ReferenceNodeTypeEnum = "GREATER_THAN"
	ReferenceNodeTypeLessThanOrEqual          ReferenceNodeTypeEnum = "LESS_THAN_OR_EQUAL"
	ReferenceNodeTypeGreaterThanOrEqual       ReferenceNodeTypeEnum = "GREATER_THAN_OR_EQUAL"
	ReferenceNodeTypeEquals                   ReferenceNodeTypeEnum = "EQUALS"
	ReferenceNodeTypeNotEquals                ReferenceNodeTypeEnum = "NOT_EQUALS"
	ReferenceNodeTypeIsDistinctFrom           ReferenceNodeTypeEnum = "IS_DISTINCT_FROM"
	ReferenceNodeTypeIsNotDistinctFrom        ReferenceNodeTypeEnum = "IS_NOT_DISTINCT_FROM"
	ReferenceNodeTypeOr                       ReferenceNodeTypeEnum = "OR"
	ReferenceNodeTypeAnd                      ReferenceNodeTypeEnum = "AND"
	ReferenceNodeTypeDot                      ReferenceNodeTypeEnum = "DOT"
	ReferenceNodeTypeOverlaps                 ReferenceNodeTypeEnum = "OVERLAPS"
	ReferenceNodeTypeContains                 ReferenceNodeTypeEnum = "CONTAINS"
	ReferenceNodeTypePrecedes                 ReferenceNodeTypeEnum = "PRECEDES"
	ReferenceNodeTypeImmediatelyPrecedes      ReferenceNodeTypeEnum = "IMMEDIATELY_PRECEDES"
	ReferenceNodeTypeSucceeds                 ReferenceNodeTypeEnum = "SUCCEEDS"
	ReferenceNodeTypeImmediatelySucceeds      ReferenceNodeTypeEnum = "IMMEDIATELY_SUCCEEDS"
	ReferenceNodeTypePeriodEquals             ReferenceNodeTypeEnum = "PERIOD_EQUALS"
	ReferenceNodeTypeLike                     ReferenceNodeTypeEnum = "LIKE"
	ReferenceNodeTypeSimilar                  ReferenceNodeTypeEnum = "SIMILAR"
	ReferenceNodeTypeBetween                  ReferenceNodeTypeEnum = "BETWEEN"
	ReferenceNodeTypeCase                     ReferenceNodeTypeEnum = "CASE"
	ReferenceNodeTypeNullif                   ReferenceNodeTypeEnum = "NULLIF"
	ReferenceNodeTypeCoalesce                 ReferenceNodeTypeEnum = "COALESCE"
	ReferenceNodeTypeDecode                   ReferenceNodeTypeEnum = "DECODE"
	ReferenceNodeTypeNvl                      ReferenceNodeTypeEnum = "NVL"
	ReferenceNodeTypeGreatest                 ReferenceNodeTypeEnum = "GREATEST"
	ReferenceNodeTypeLeast                    ReferenceNodeTypeEnum = "LEAST"
	ReferenceNodeTypeTimestampAdd             ReferenceNodeTypeEnum = "TIMESTAMP_ADD"
	ReferenceNodeTypeTimestampDiff            ReferenceNodeTypeEnum = "TIMESTAMP_DIFF"
	ReferenceNodeTypeNot                      ReferenceNodeTypeEnum = "NOT"
	ReferenceNodeTypePlusPrefix               ReferenceNodeTypeEnum = "PLUS_PREFIX"
	ReferenceNodeTypeMinusPrefix              ReferenceNodeTypeEnum = "MINUS_PREFIX"
	ReferenceNodeTypeExists                   ReferenceNodeTypeEnum = "EXISTS"
	ReferenceNodeTypeSome                     ReferenceNodeTypeEnum = "SOME"
	ReferenceNodeTypeAll                      ReferenceNodeTypeEnum = "ALL"
	ReferenceNodeTypeValues                   ReferenceNodeTypeEnum = "VALUES"
	ReferenceNodeTypeExplicitTable            ReferenceNodeTypeEnum = "EXPLICIT_TABLE"
	ReferenceNodeTypeScalarQuery              ReferenceNodeTypeEnum = "SCALAR_QUERY"
	ReferenceNodeTypeProcedureCall            ReferenceNodeTypeEnum = "PROCEDURE_CALL"
	ReferenceNodeTypeNewSpecification         ReferenceNodeTypeEnum = "NEW_SPECIFICATION"
	ReferenceNodeTypeFinal                    ReferenceNodeTypeEnum = "FINAL"
	ReferenceNodeTypeRunning                  ReferenceNodeTypeEnum = "RUNNING"
	ReferenceNodeTypePrev                     ReferenceNodeTypeEnum = "PREV"
	ReferenceNodeTypeNext                     ReferenceNodeTypeEnum = "NEXT"
	ReferenceNodeTypeFirst                    ReferenceNodeTypeEnum = "FIRST"
	ReferenceNodeTypeLast                     ReferenceNodeTypeEnum = "LAST"
	ReferenceNodeTypeClassifier               ReferenceNodeTypeEnum = "CLASSIFIER"
	ReferenceNodeTypeMatchNumber              ReferenceNodeTypeEnum = "MATCH_NUMBER"
	ReferenceNodeTypeSkipToFirst              ReferenceNodeTypeEnum = "SKIP_TO_FIRST"
	ReferenceNodeTypeSkipToLast               ReferenceNodeTypeEnum = "SKIP_TO_LAST"
	ReferenceNodeTypeDescending               ReferenceNodeTypeEnum = "DESCENDING"
	ReferenceNodeTypeNullsFirst               ReferenceNodeTypeEnum = "NULLS_FIRST"
	ReferenceNodeTypeNullsLast                ReferenceNodeTypeEnum = "NULLS_LAST"
	ReferenceNodeTypeIsTrue                   ReferenceNodeTypeEnum = "IS_TRUE"
	ReferenceNodeTypeIsFalse                  ReferenceNodeTypeEnum = "IS_FALSE"
	ReferenceNodeTypeIsNotTrue                ReferenceNodeTypeEnum = "IS_NOT_TRUE"
	ReferenceNodeTypeIsNotFalse               ReferenceNodeTypeEnum = "IS_NOT_FALSE"
	ReferenceNodeTypeIsUnknown                ReferenceNodeTypeEnum = "IS_UNKNOWN"
	ReferenceNodeTypeIsNull                   ReferenceNodeTypeEnum = "IS_NULL"
	ReferenceNodeTypeIsNotNull                ReferenceNodeTypeEnum = "IS_NOT_NULL"
	ReferenceNodeTypePreceding                ReferenceNodeTypeEnum = "PRECEDING"
	ReferenceNodeTypeFollowing                ReferenceNodeTypeEnum = "FOLLOWING"
	ReferenceNodeTypeFieldAccess              ReferenceNodeTypeEnum = "FIELD_ACCESS"
	ReferenceNodeTypeInputRef                 ReferenceNodeTypeEnum = "INPUT_REF"
	ReferenceNodeTypeTableInputRef            ReferenceNodeTypeEnum = "TABLE_INPUT_REF"
	ReferenceNodeTypePatternInputRef          ReferenceNodeTypeEnum = "PATTERN_INPUT_REF"
	ReferenceNodeTypeLocalRef                 ReferenceNodeTypeEnum = "LOCAL_REF"
	ReferenceNodeTypeCorrelVariable           ReferenceNodeTypeEnum = "CORREL_VARIABLE"
	ReferenceNodeTypePatternQuantifier        ReferenceNodeTypeEnum = "PATTERN_QUANTIFIER"
	ReferenceNodeTypeRow                      ReferenceNodeTypeEnum = "ROW"
	ReferenceNodeTypeColumnList               ReferenceNodeTypeEnum = "COLUMN_LIST"
	ReferenceNodeTypeCast                     ReferenceNodeTypeEnum = "CAST"
	ReferenceNodeTypeNextValue                ReferenceNodeTypeEnum = "NEXT_VALUE"
	ReferenceNodeTypeCurrentValue             ReferenceNodeTypeEnum = "CURRENT_VALUE"
	ReferenceNodeTypeFloor                    ReferenceNodeTypeEnum = "FLOOR"
	ReferenceNodeTypeCeil                     ReferenceNodeTypeEnum = "CEIL"
	ReferenceNodeTypeTrim                     ReferenceNodeTypeEnum = "TRIM"
	ReferenceNodeTypeLtrim                    ReferenceNodeTypeEnum = "LTRIM"
	ReferenceNodeTypeRtrim                    ReferenceNodeTypeEnum = "RTRIM"
	ReferenceNodeTypeExtract                  ReferenceNodeTypeEnum = "EXTRACT"
	ReferenceNodeTypeJdbcFn                   ReferenceNodeTypeEnum = "JDBC_FN"
	ReferenceNodeTypeMultisetValueConstructor ReferenceNodeTypeEnum = "MULTISET_VALUE_CONSTRUCTOR"
	ReferenceNodeTypeMultisetQueryConstructor ReferenceNodeTypeEnum = "MULTISET_QUERY_CONSTRUCTOR"
	ReferenceNodeTypeUnnest                   ReferenceNodeTypeEnum = "UNNEST"
	ReferenceNodeTypeLateral                  ReferenceNodeTypeEnum = "LATERAL"
	ReferenceNodeTypeCollectionTable          ReferenceNodeTypeEnum = "COLLECTION_TABLE"
	ReferenceNodeTypeArrayValueConstructor    ReferenceNodeTypeEnum = "ARRAY_VALUE_CONSTRUCTOR"
	ReferenceNodeTypeArrayQueryConstructor    ReferenceNodeTypeEnum = "ARRAY_QUERY_CONSTRUCTOR"
	ReferenceNodeTypeMapValueConstructor      ReferenceNodeTypeEnum = "MAP_VALUE_CONSTRUCTOR"
	ReferenceNodeTypeMapQueryConstructor      ReferenceNodeTypeEnum = "MAP_QUERY_CONSTRUCTOR"
	ReferenceNodeTypeCursor                   ReferenceNodeTypeEnum = "CURSOR"
	ReferenceNodeTypeLiteralChain             ReferenceNodeTypeEnum = "LITERAL_CHAIN"
	ReferenceNodeTypeEscape                   ReferenceNodeTypeEnum = "ESCAPE"
	ReferenceNodeTypeReinterpret              ReferenceNodeTypeEnum = "REINTERPRET"
	ReferenceNodeTypeExtend                   ReferenceNodeTypeEnum = "EXTEND"
	ReferenceNodeTypeCube                     ReferenceNodeTypeEnum = "CUBE"
	ReferenceNodeTypeRollup                   ReferenceNodeTypeEnum = "ROLLUP"
	ReferenceNodeTypeGroupingSets             ReferenceNodeTypeEnum = "GROUPING_SETS"
	ReferenceNodeTypeGrouping                 ReferenceNodeTypeEnum = "GROUPING"
	ReferenceNodeTypeGroupingId               ReferenceNodeTypeEnum = "GROUPING_ID"
	ReferenceNodeTypeGroupId                  ReferenceNodeTypeEnum = "GROUP_ID"
	ReferenceNodeTypePatternPermute           ReferenceNodeTypeEnum = "PATTERN_PERMUTE"
	ReferenceNodeTypePatternExcluded          ReferenceNodeTypeEnum = "PATTERN_EXCLUDED"
	ReferenceNodeTypeCount                    ReferenceNodeTypeEnum = "COUNT"
	ReferenceNodeTypeSum                      ReferenceNodeTypeEnum = "SUM"
	ReferenceNodeTypeSum0                     ReferenceNodeTypeEnum = "SUM0"
	ReferenceNodeTypeMin                      ReferenceNodeTypeEnum = "MIN"
	ReferenceNodeTypeMax                      ReferenceNodeTypeEnum = "MAX"
	ReferenceNodeTypeLead                     ReferenceNodeTypeEnum = "LEAD"
	ReferenceNodeTypeLag                      ReferenceNodeTypeEnum = "LAG"
	ReferenceNodeTypeFirstValue               ReferenceNodeTypeEnum = "FIRST_VALUE"
	ReferenceNodeTypeLastValue                ReferenceNodeTypeEnum = "LAST_VALUE"
	ReferenceNodeTypeCovarPop                 ReferenceNodeTypeEnum = "COVAR_POP"
	ReferenceNodeTypeCovarSamp                ReferenceNodeTypeEnum = "COVAR_SAMP"
	ReferenceNodeTypeRegrSxx                  ReferenceNodeTypeEnum = "REGR_SXX"
	ReferenceNodeTypeRegrSyy                  ReferenceNodeTypeEnum = "REGR_SYY"
	ReferenceNodeTypeAvg                      ReferenceNodeTypeEnum = "AVG"
	ReferenceNodeTypeStddevPop                ReferenceNodeTypeEnum = "STDDEV_POP"
	ReferenceNodeTypeStddevSamp               ReferenceNodeTypeEnum = "STDDEV_SAMP"
	ReferenceNodeTypeVarPop                   ReferenceNodeTypeEnum = "VAR_POP"
	ReferenceNodeTypeVarSamp                  ReferenceNodeTypeEnum = "VAR_SAMP"
	ReferenceNodeTypeNtile                    ReferenceNodeTypeEnum = "NTILE"
	ReferenceNodeTypeCollect                  ReferenceNodeTypeEnum = "COLLECT"
	ReferenceNodeTypeFusion                   ReferenceNodeTypeEnum = "FUSION"
	ReferenceNodeTypeSingleValue              ReferenceNodeTypeEnum = "SINGLE_VALUE"
	ReferenceNodeTypeRowNumber                ReferenceNodeTypeEnum = "ROW_NUMBER"
	ReferenceNodeTypeRank                     ReferenceNodeTypeEnum = "RANK"
	ReferenceNodeTypePercentRank              ReferenceNodeTypeEnum = "PERCENT_RANK"
	ReferenceNodeTypeDenseRank                ReferenceNodeTypeEnum = "DENSE_RANK"
	ReferenceNodeTypeCumeDist                 ReferenceNodeTypeEnum = "CUME_DIST"
	ReferenceNodeTypeTumble                   ReferenceNodeTypeEnum = "TUMBLE"
	ReferenceNodeTypeTumbleStart              ReferenceNodeTypeEnum = "TUMBLE_START"
	ReferenceNodeTypeTumbleEnd                ReferenceNodeTypeEnum = "TUMBLE_END"
	ReferenceNodeTypeHop                      ReferenceNodeTypeEnum = "HOP"
	ReferenceNodeTypeHopStart                 ReferenceNodeTypeEnum = "HOP_START"
	ReferenceNodeTypeHopEnd                   ReferenceNodeTypeEnum = "HOP_END"
	ReferenceNodeTypeSession                  ReferenceNodeTypeEnum = "SESSION"
	ReferenceNodeTypeSessionStart             ReferenceNodeTypeEnum = "SESSION_START"
	ReferenceNodeTypeSessionEnd               ReferenceNodeTypeEnum = "SESSION_END"
	ReferenceNodeTypeColumnDecl               ReferenceNodeTypeEnum = "COLUMN_DECL"
	ReferenceNodeTypeCheck                    ReferenceNodeTypeEnum = "CHECK"
	ReferenceNodeTypeUnique                   ReferenceNodeTypeEnum = "UNIQUE"
	ReferenceNodeTypePrimaryKey               ReferenceNodeTypeEnum = "PRIMARY_KEY"
	ReferenceNodeTypeForeignKey               ReferenceNodeTypeEnum = "FOREIGN_KEY"
	ReferenceNodeTypeCommit                   ReferenceNodeTypeEnum = "COMMIT"
	ReferenceNodeTypeRollback                 ReferenceNodeTypeEnum = "ROLLBACK"
	ReferenceNodeTypeAlterSession             ReferenceNodeTypeEnum = "ALTER_SESSION"
	ReferenceNodeTypeCreateSchema             ReferenceNodeTypeEnum = "CREATE_SCHEMA"
	ReferenceNodeTypeCreateForeignSchema      ReferenceNodeTypeEnum = "CREATE_FOREIGN_SCHEMA"
	ReferenceNodeTypeDropSchema               ReferenceNodeTypeEnum = "DROP_SCHEMA"
	ReferenceNodeTypeCreateTable              ReferenceNodeTypeEnum = "CREATE_TABLE"
	ReferenceNodeTypeAlterTable               ReferenceNodeTypeEnum = "ALTER_TABLE"
	ReferenceNodeTypeDropTable                ReferenceNodeTypeEnum = "DROP_TABLE"
	ReferenceNodeTypeCreateView               ReferenceNodeTypeEnum = "CREATE_VIEW"
	ReferenceNodeTypeAlterView                ReferenceNodeTypeEnum = "ALTER_VIEW"
	ReferenceNodeTypeDropView                 ReferenceNodeTypeEnum = "DROP_VIEW"
	ReferenceNodeTypeCreateMaterializedView   ReferenceNodeTypeEnum = "CREATE_MATERIALIZED_VIEW"
	ReferenceNodeTypeAlterMaterializedView    ReferenceNodeTypeEnum = "ALTER_MATERIALIZED_VIEW"
	ReferenceNodeTypeDropMaterializedView     ReferenceNodeTypeEnum = "DROP_MATERIALIZED_VIEW"
	ReferenceNodeTypeCreateSequence           ReferenceNodeTypeEnum = "CREATE_SEQUENCE"
	ReferenceNodeTypeAlterSequence            ReferenceNodeTypeEnum = "ALTER_SEQUENCE"
	ReferenceNodeTypeDropSequence             ReferenceNodeTypeEnum = "DROP_SEQUENCE"
	ReferenceNodeTypeCreateIndex              ReferenceNodeTypeEnum = "CREATE_INDEX"
	ReferenceNodeTypeAlterIndex               ReferenceNodeTypeEnum = "ALTER_INDEX"
	ReferenceNodeTypeDropIndex                ReferenceNodeTypeEnum = "DROP_INDEX"
	ReferenceNodeTypeOtherDdl                 ReferenceNodeTypeEnum = "OTHER_DDL"
)

var mappingReferenceNodeType = map[string]ReferenceNodeTypeEnum{
	"OTHER":                      ReferenceNodeTypeOther,
	"SELECT":                     ReferenceNodeTypeSelect,
	"JOIN":                       ReferenceNodeTypeJoin,
	"IDENTIFIER":                 ReferenceNodeTypeIdentifier,
	"LITERAL":                    ReferenceNodeTypeLiteral,
	"OTHER_FUNCTION":             ReferenceNodeTypeOtherFunction,
	"EXPLAIN":                    ReferenceNodeTypeExplain,
	"DESCRIBE_SCHEMA":            ReferenceNodeTypeDescribeSchema,
	"DESCRIBE_TABLE":             ReferenceNodeTypeDescribeTable,
	"INSERT":                     ReferenceNodeTypeInsert,
	"DELETE":                     ReferenceNodeTypeDelete,
	"UPDATE":                     ReferenceNodeTypeUpdate,
	"SET_OPTION":                 ReferenceNodeTypeSetOption,
	"DYNAMIC_PARAM":              ReferenceNodeTypeDynamicParam,
	"ORDER_BY":                   ReferenceNodeTypeOrderBy,
	"WITH":                       ReferenceNodeTypeWith,
	"WITH_ITEM":                  ReferenceNodeTypeWithItem,
	"UNION":                      ReferenceNodeTypeUnion,
	"EXCEPT":                     ReferenceNodeTypeExcept,
	"INTERSECT":                  ReferenceNodeTypeIntersect,
	"AS":                         ReferenceNodeTypeAs,
	"ARGUMENT_ASSIGNMENT":        ReferenceNodeTypeArgumentAssignment,
	"DEFAULT":                    ReferenceNodeTypeDefault,
	"OVER":                       ReferenceNodeTypeOver,
	"FILTER":                     ReferenceNodeTypeFilter,
	"WINDOW":                     ReferenceNodeTypeWindow,
	"MERGE":                      ReferenceNodeTypeMerge,
	"TABLESAMPLE":                ReferenceNodeTypeTablesample,
	"MATCH_RECOGNIZE":            ReferenceNodeTypeMatchRecognize,
	"TIMES":                      ReferenceNodeTypeTimes,
	"DIVIDE":                     ReferenceNodeTypeDivide,
	"MOD":                        ReferenceNodeTypeMod,
	"PLUS":                       ReferenceNodeTypePlus,
	"MINUS":                      ReferenceNodeTypeMinus,
	"PATTERN_ALTER":              ReferenceNodeTypePatternAlter,
	"PATTERN_CONCAT":             ReferenceNodeTypePatternConcat,
	"IN":                         ReferenceNodeTypeIn,
	"NOT_IN":                     ReferenceNodeTypeNotIn,
	"LESS_THAN":                  ReferenceNodeTypeLessThan,
	"GREATER_THAN":               ReferenceNodeTypeGreaterThan,
	"LESS_THAN_OR_EQUAL":         ReferenceNodeTypeLessThanOrEqual,
	"GREATER_THAN_OR_EQUAL":      ReferenceNodeTypeGreaterThanOrEqual,
	"EQUALS":                     ReferenceNodeTypeEquals,
	"NOT_EQUALS":                 ReferenceNodeTypeNotEquals,
	"IS_DISTINCT_FROM":           ReferenceNodeTypeIsDistinctFrom,
	"IS_NOT_DISTINCT_FROM":       ReferenceNodeTypeIsNotDistinctFrom,
	"OR":                         ReferenceNodeTypeOr,
	"AND":                        ReferenceNodeTypeAnd,
	"DOT":                        ReferenceNodeTypeDot,
	"OVERLAPS":                   ReferenceNodeTypeOverlaps,
	"CONTAINS":                   ReferenceNodeTypeContains,
	"PRECEDES":                   ReferenceNodeTypePrecedes,
	"IMMEDIATELY_PRECEDES":       ReferenceNodeTypeImmediatelyPrecedes,
	"SUCCEEDS":                   ReferenceNodeTypeSucceeds,
	"IMMEDIATELY_SUCCEEDS":       ReferenceNodeTypeImmediatelySucceeds,
	"PERIOD_EQUALS":              ReferenceNodeTypePeriodEquals,
	"LIKE":                       ReferenceNodeTypeLike,
	"SIMILAR":                    ReferenceNodeTypeSimilar,
	"BETWEEN":                    ReferenceNodeTypeBetween,
	"CASE":                       ReferenceNodeTypeCase,
	"NULLIF":                     ReferenceNodeTypeNullif,
	"COALESCE":                   ReferenceNodeTypeCoalesce,
	"DECODE":                     ReferenceNodeTypeDecode,
	"NVL":                        ReferenceNodeTypeNvl,
	"GREATEST":                   ReferenceNodeTypeGreatest,
	"LEAST":                      ReferenceNodeTypeLeast,
	"TIMESTAMP_ADD":              ReferenceNodeTypeTimestampAdd,
	"TIMESTAMP_DIFF":             ReferenceNodeTypeTimestampDiff,
	"NOT":                        ReferenceNodeTypeNot,
	"PLUS_PREFIX":                ReferenceNodeTypePlusPrefix,
	"MINUS_PREFIX":               ReferenceNodeTypeMinusPrefix,
	"EXISTS":                     ReferenceNodeTypeExists,
	"SOME":                       ReferenceNodeTypeSome,
	"ALL":                        ReferenceNodeTypeAll,
	"VALUES":                     ReferenceNodeTypeValues,
	"EXPLICIT_TABLE":             ReferenceNodeTypeExplicitTable,
	"SCALAR_QUERY":               ReferenceNodeTypeScalarQuery,
	"PROCEDURE_CALL":             ReferenceNodeTypeProcedureCall,
	"NEW_SPECIFICATION":          ReferenceNodeTypeNewSpecification,
	"FINAL":                      ReferenceNodeTypeFinal,
	"RUNNING":                    ReferenceNodeTypeRunning,
	"PREV":                       ReferenceNodeTypePrev,
	"NEXT":                       ReferenceNodeTypeNext,
	"FIRST":                      ReferenceNodeTypeFirst,
	"LAST":                       ReferenceNodeTypeLast,
	"CLASSIFIER":                 ReferenceNodeTypeClassifier,
	"MATCH_NUMBER":               ReferenceNodeTypeMatchNumber,
	"SKIP_TO_FIRST":              ReferenceNodeTypeSkipToFirst,
	"SKIP_TO_LAST":               ReferenceNodeTypeSkipToLast,
	"DESCENDING":                 ReferenceNodeTypeDescending,
	"NULLS_FIRST":                ReferenceNodeTypeNullsFirst,
	"NULLS_LAST":                 ReferenceNodeTypeNullsLast,
	"IS_TRUE":                    ReferenceNodeTypeIsTrue,
	"IS_FALSE":                   ReferenceNodeTypeIsFalse,
	"IS_NOT_TRUE":                ReferenceNodeTypeIsNotTrue,
	"IS_NOT_FALSE":               ReferenceNodeTypeIsNotFalse,
	"IS_UNKNOWN":                 ReferenceNodeTypeIsUnknown,
	"IS_NULL":                    ReferenceNodeTypeIsNull,
	"IS_NOT_NULL":                ReferenceNodeTypeIsNotNull,
	"PRECEDING":                  ReferenceNodeTypePreceding,
	"FOLLOWING":                  ReferenceNodeTypeFollowing,
	"FIELD_ACCESS":               ReferenceNodeTypeFieldAccess,
	"INPUT_REF":                  ReferenceNodeTypeInputRef,
	"TABLE_INPUT_REF":            ReferenceNodeTypeTableInputRef,
	"PATTERN_INPUT_REF":          ReferenceNodeTypePatternInputRef,
	"LOCAL_REF":                  ReferenceNodeTypeLocalRef,
	"CORREL_VARIABLE":            ReferenceNodeTypeCorrelVariable,
	"PATTERN_QUANTIFIER":         ReferenceNodeTypePatternQuantifier,
	"ROW":                        ReferenceNodeTypeRow,
	"COLUMN_LIST":                ReferenceNodeTypeColumnList,
	"CAST":                       ReferenceNodeTypeCast,
	"NEXT_VALUE":                 ReferenceNodeTypeNextValue,
	"CURRENT_VALUE":              ReferenceNodeTypeCurrentValue,
	"FLOOR":                      ReferenceNodeTypeFloor,
	"CEIL":                       ReferenceNodeTypeCeil,
	"TRIM":                       ReferenceNodeTypeTrim,
	"LTRIM":                      ReferenceNodeTypeLtrim,
	"RTRIM":                      ReferenceNodeTypeRtrim,
	"EXTRACT":                    ReferenceNodeTypeExtract,
	"JDBC_FN":                    ReferenceNodeTypeJdbcFn,
	"MULTISET_VALUE_CONSTRUCTOR": ReferenceNodeTypeMultisetValueConstructor,
	"MULTISET_QUERY_CONSTRUCTOR": ReferenceNodeTypeMultisetQueryConstructor,
	"UNNEST":                     ReferenceNodeTypeUnnest,
	"LATERAL":                    ReferenceNodeTypeLateral,
	"COLLECTION_TABLE":           ReferenceNodeTypeCollectionTable,
	"ARRAY_VALUE_CONSTRUCTOR":    ReferenceNodeTypeArrayValueConstructor,
	"ARRAY_QUERY_CONSTRUCTOR":    ReferenceNodeTypeArrayQueryConstructor,
	"MAP_VALUE_CONSTRUCTOR":      ReferenceNodeTypeMapValueConstructor,
	"MAP_QUERY_CONSTRUCTOR":      ReferenceNodeTypeMapQueryConstructor,
	"CURSOR":                     ReferenceNodeTypeCursor,
	"LITERAL_CHAIN":              ReferenceNodeTypeLiteralChain,
	"ESCAPE":                     ReferenceNodeTypeEscape,
	"REINTERPRET":                ReferenceNodeTypeReinterpret,
	"EXTEND":                     ReferenceNodeTypeExtend,
	"CUBE":                       ReferenceNodeTypeCube,
	"ROLLUP":                     ReferenceNodeTypeRollup,
	"GROUPING_SETS":              ReferenceNodeTypeGroupingSets,
	"GROUPING":                   ReferenceNodeTypeGrouping,
	"GROUPING_ID":                ReferenceNodeTypeGroupingId,
	"GROUP_ID":                   ReferenceNodeTypeGroupId,
	"PATTERN_PERMUTE":            ReferenceNodeTypePatternPermute,
	"PATTERN_EXCLUDED":           ReferenceNodeTypePatternExcluded,
	"COUNT":                      ReferenceNodeTypeCount,
	"SUM":                        ReferenceNodeTypeSum,
	"SUM0":                       ReferenceNodeTypeSum0,
	"MIN":                        ReferenceNodeTypeMin,
	"MAX":                        ReferenceNodeTypeMax,
	"LEAD":                       ReferenceNodeTypeLead,
	"LAG":                        ReferenceNodeTypeLag,
	"FIRST_VALUE":                ReferenceNodeTypeFirstValue,
	"LAST_VALUE":                 ReferenceNodeTypeLastValue,
	"COVAR_POP":                  ReferenceNodeTypeCovarPop,
	"COVAR_SAMP":                 ReferenceNodeTypeCovarSamp,
	"REGR_SXX":                   ReferenceNodeTypeRegrSxx,
	"REGR_SYY":                   ReferenceNodeTypeRegrSyy,
	"AVG":                        ReferenceNodeTypeAvg,
	"STDDEV_POP":                 ReferenceNodeTypeStddevPop,
	"STDDEV_SAMP":                ReferenceNodeTypeStddevSamp,
	"VAR_POP":                    ReferenceNodeTypeVarPop,
	"VAR_SAMP":                   ReferenceNodeTypeVarSamp,
	"NTILE":                      ReferenceNodeTypeNtile,
	"COLLECT":                    ReferenceNodeTypeCollect,
	"FUSION":                     ReferenceNodeTypeFusion,
	"SINGLE_VALUE":               ReferenceNodeTypeSingleValue,
	"ROW_NUMBER":                 ReferenceNodeTypeRowNumber,
	"RANK":                       ReferenceNodeTypeRank,
	"PERCENT_RANK":               ReferenceNodeTypePercentRank,
	"DENSE_RANK":                 ReferenceNodeTypeDenseRank,
	"CUME_DIST":                  ReferenceNodeTypeCumeDist,
	"TUMBLE":                     ReferenceNodeTypeTumble,
	"TUMBLE_START":               ReferenceNodeTypeTumbleStart,
	"TUMBLE_END":                 ReferenceNodeTypeTumbleEnd,
	"HOP":                        ReferenceNodeTypeHop,
	"HOP_START":                  ReferenceNodeTypeHopStart,
	"HOP_END":                    ReferenceNodeTypeHopEnd,
	"SESSION":                    ReferenceNodeTypeSession,
	"SESSION_START":              ReferenceNodeTypeSessionStart,
	"SESSION_END":                ReferenceNodeTypeSessionEnd,
	"COLUMN_DECL":                ReferenceNodeTypeColumnDecl,
	"CHECK":                      ReferenceNodeTypeCheck,
	"UNIQUE":                     ReferenceNodeTypeUnique,
	"PRIMARY_KEY":                ReferenceNodeTypePrimaryKey,
	"FOREIGN_KEY":                ReferenceNodeTypeForeignKey,
	"COMMIT":                     ReferenceNodeTypeCommit,
	"ROLLBACK":                   ReferenceNodeTypeRollback,
	"ALTER_SESSION":              ReferenceNodeTypeAlterSession,
	"CREATE_SCHEMA":              ReferenceNodeTypeCreateSchema,
	"CREATE_FOREIGN_SCHEMA":      ReferenceNodeTypeCreateForeignSchema,
	"DROP_SCHEMA":                ReferenceNodeTypeDropSchema,
	"CREATE_TABLE":               ReferenceNodeTypeCreateTable,
	"ALTER_TABLE":                ReferenceNodeTypeAlterTable,
	"DROP_TABLE":                 ReferenceNodeTypeDropTable,
	"CREATE_VIEW":                ReferenceNodeTypeCreateView,
	"ALTER_VIEW":                 ReferenceNodeTypeAlterView,
	"DROP_VIEW":                  ReferenceNodeTypeDropView,
	"CREATE_MATERIALIZED_VIEW":   ReferenceNodeTypeCreateMaterializedView,
	"ALTER_MATERIALIZED_VIEW":    ReferenceNodeTypeAlterMaterializedView,
	"DROP_MATERIALIZED_VIEW":     ReferenceNodeTypeDropMaterializedView,
	"CREATE_SEQUENCE":            ReferenceNodeTypeCreateSequence,
	"ALTER_SEQUENCE":             ReferenceNodeTypeAlterSequence,
	"DROP_SEQUENCE":              ReferenceNodeTypeDropSequence,
	"CREATE_INDEX":               ReferenceNodeTypeCreateIndex,
	"ALTER_INDEX":                ReferenceNodeTypeAlterIndex,
	"DROP_INDEX":                 ReferenceNodeTypeDropIndex,
	"OTHER_DDL":                  ReferenceNodeTypeOtherDdl,
}

// GetReferenceNodeTypeEnumValues Enumerates the set of values for ReferenceNodeTypeEnum
func GetReferenceNodeTypeEnumValues() []ReferenceNodeTypeEnum {
	values := make([]ReferenceNodeTypeEnum, 0)
	for _, v := range mappingReferenceNodeType {
		values = append(values, v)
	}
	return values
}
