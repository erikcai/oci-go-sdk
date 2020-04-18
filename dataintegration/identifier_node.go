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

// IdentifierNode auto generated description
type IdentifierNode struct {

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

	// nameList
	NameList []string `mandatory:"false" json:"nameList"`

	// Reference to a typed object
	ReferencedObject *string `mandatory:"false" json:"referencedObject"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// type
	Type IdentifierNodeTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m IdentifierNode) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m IdentifierNode) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIdentifierNode IdentifierNode
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeIdentifierNode
	}{
		"IDENTIFIER_NODE",
		(MarshalTypeIdentifierNode)(m),
	}

	return json.Marshal(&s)
}

// IdentifierNodeTypeEnum Enum with underlying type: string
type IdentifierNodeTypeEnum string

// Set of constants representing the allowable values for IdentifierNodeTypeEnum
const (
	IdentifierNodeTypeOther                    IdentifierNodeTypeEnum = "OTHER"
	IdentifierNodeTypeSelect                   IdentifierNodeTypeEnum = "SELECT"
	IdentifierNodeTypeJoin                     IdentifierNodeTypeEnum = "JOIN"
	IdentifierNodeTypeIdentifier               IdentifierNodeTypeEnum = "IDENTIFIER"
	IdentifierNodeTypeLiteral                  IdentifierNodeTypeEnum = "LITERAL"
	IdentifierNodeTypeOtherFunction            IdentifierNodeTypeEnum = "OTHER_FUNCTION"
	IdentifierNodeTypeExplain                  IdentifierNodeTypeEnum = "EXPLAIN"
	IdentifierNodeTypeDescribeSchema           IdentifierNodeTypeEnum = "DESCRIBE_SCHEMA"
	IdentifierNodeTypeDescribeTable            IdentifierNodeTypeEnum = "DESCRIBE_TABLE"
	IdentifierNodeTypeInsert                   IdentifierNodeTypeEnum = "INSERT"
	IdentifierNodeTypeDelete                   IdentifierNodeTypeEnum = "DELETE"
	IdentifierNodeTypeUpdate                   IdentifierNodeTypeEnum = "UPDATE"
	IdentifierNodeTypeSetOption                IdentifierNodeTypeEnum = "SET_OPTION"
	IdentifierNodeTypeDynamicParam             IdentifierNodeTypeEnum = "DYNAMIC_PARAM"
	IdentifierNodeTypeOrderBy                  IdentifierNodeTypeEnum = "ORDER_BY"
	IdentifierNodeTypeWith                     IdentifierNodeTypeEnum = "WITH"
	IdentifierNodeTypeWithItem                 IdentifierNodeTypeEnum = "WITH_ITEM"
	IdentifierNodeTypeUnion                    IdentifierNodeTypeEnum = "UNION"
	IdentifierNodeTypeExcept                   IdentifierNodeTypeEnum = "EXCEPT"
	IdentifierNodeTypeIntersect                IdentifierNodeTypeEnum = "INTERSECT"
	IdentifierNodeTypeAs                       IdentifierNodeTypeEnum = "AS"
	IdentifierNodeTypeArgumentAssignment       IdentifierNodeTypeEnum = "ARGUMENT_ASSIGNMENT"
	IdentifierNodeTypeDefault                  IdentifierNodeTypeEnum = "DEFAULT"
	IdentifierNodeTypeOver                     IdentifierNodeTypeEnum = "OVER"
	IdentifierNodeTypeFilter                   IdentifierNodeTypeEnum = "FILTER"
	IdentifierNodeTypeWindow                   IdentifierNodeTypeEnum = "WINDOW"
	IdentifierNodeTypeMerge                    IdentifierNodeTypeEnum = "MERGE"
	IdentifierNodeTypeTablesample              IdentifierNodeTypeEnum = "TABLESAMPLE"
	IdentifierNodeTypeMatchRecognize           IdentifierNodeTypeEnum = "MATCH_RECOGNIZE"
	IdentifierNodeTypeTimes                    IdentifierNodeTypeEnum = "TIMES"
	IdentifierNodeTypeDivide                   IdentifierNodeTypeEnum = "DIVIDE"
	IdentifierNodeTypeMod                      IdentifierNodeTypeEnum = "MOD"
	IdentifierNodeTypePlus                     IdentifierNodeTypeEnum = "PLUS"
	IdentifierNodeTypeMinus                    IdentifierNodeTypeEnum = "MINUS"
	IdentifierNodeTypePatternAlter             IdentifierNodeTypeEnum = "PATTERN_ALTER"
	IdentifierNodeTypePatternConcat            IdentifierNodeTypeEnum = "PATTERN_CONCAT"
	IdentifierNodeTypeIn                       IdentifierNodeTypeEnum = "IN"
	IdentifierNodeTypeNotIn                    IdentifierNodeTypeEnum = "NOT_IN"
	IdentifierNodeTypeLessThan                 IdentifierNodeTypeEnum = "LESS_THAN"
	IdentifierNodeTypeGreaterThan              IdentifierNodeTypeEnum = "GREATER_THAN"
	IdentifierNodeTypeLessThanOrEqual          IdentifierNodeTypeEnum = "LESS_THAN_OR_EQUAL"
	IdentifierNodeTypeGreaterThanOrEqual       IdentifierNodeTypeEnum = "GREATER_THAN_OR_EQUAL"
	IdentifierNodeTypeEquals                   IdentifierNodeTypeEnum = "EQUALS"
	IdentifierNodeTypeNotEquals                IdentifierNodeTypeEnum = "NOT_EQUALS"
	IdentifierNodeTypeIsDistinctFrom           IdentifierNodeTypeEnum = "IS_DISTINCT_FROM"
	IdentifierNodeTypeIsNotDistinctFrom        IdentifierNodeTypeEnum = "IS_NOT_DISTINCT_FROM"
	IdentifierNodeTypeOr                       IdentifierNodeTypeEnum = "OR"
	IdentifierNodeTypeAnd                      IdentifierNodeTypeEnum = "AND"
	IdentifierNodeTypeDot                      IdentifierNodeTypeEnum = "DOT"
	IdentifierNodeTypeOverlaps                 IdentifierNodeTypeEnum = "OVERLAPS"
	IdentifierNodeTypeContains                 IdentifierNodeTypeEnum = "CONTAINS"
	IdentifierNodeTypePrecedes                 IdentifierNodeTypeEnum = "PRECEDES"
	IdentifierNodeTypeImmediatelyPrecedes      IdentifierNodeTypeEnum = "IMMEDIATELY_PRECEDES"
	IdentifierNodeTypeSucceeds                 IdentifierNodeTypeEnum = "SUCCEEDS"
	IdentifierNodeTypeImmediatelySucceeds      IdentifierNodeTypeEnum = "IMMEDIATELY_SUCCEEDS"
	IdentifierNodeTypePeriodEquals             IdentifierNodeTypeEnum = "PERIOD_EQUALS"
	IdentifierNodeTypeLike                     IdentifierNodeTypeEnum = "LIKE"
	IdentifierNodeTypeSimilar                  IdentifierNodeTypeEnum = "SIMILAR"
	IdentifierNodeTypeBetween                  IdentifierNodeTypeEnum = "BETWEEN"
	IdentifierNodeTypeCase                     IdentifierNodeTypeEnum = "CASE"
	IdentifierNodeTypeNullif                   IdentifierNodeTypeEnum = "NULLIF"
	IdentifierNodeTypeCoalesce                 IdentifierNodeTypeEnum = "COALESCE"
	IdentifierNodeTypeDecode                   IdentifierNodeTypeEnum = "DECODE"
	IdentifierNodeTypeNvl                      IdentifierNodeTypeEnum = "NVL"
	IdentifierNodeTypeGreatest                 IdentifierNodeTypeEnum = "GREATEST"
	IdentifierNodeTypeLeast                    IdentifierNodeTypeEnum = "LEAST"
	IdentifierNodeTypeTimestampAdd             IdentifierNodeTypeEnum = "TIMESTAMP_ADD"
	IdentifierNodeTypeTimestampDiff            IdentifierNodeTypeEnum = "TIMESTAMP_DIFF"
	IdentifierNodeTypeNot                      IdentifierNodeTypeEnum = "NOT"
	IdentifierNodeTypePlusPrefix               IdentifierNodeTypeEnum = "PLUS_PREFIX"
	IdentifierNodeTypeMinusPrefix              IdentifierNodeTypeEnum = "MINUS_PREFIX"
	IdentifierNodeTypeExists                   IdentifierNodeTypeEnum = "EXISTS"
	IdentifierNodeTypeSome                     IdentifierNodeTypeEnum = "SOME"
	IdentifierNodeTypeAll                      IdentifierNodeTypeEnum = "ALL"
	IdentifierNodeTypeValues                   IdentifierNodeTypeEnum = "VALUES"
	IdentifierNodeTypeExplicitTable            IdentifierNodeTypeEnum = "EXPLICIT_TABLE"
	IdentifierNodeTypeScalarQuery              IdentifierNodeTypeEnum = "SCALAR_QUERY"
	IdentifierNodeTypeProcedureCall            IdentifierNodeTypeEnum = "PROCEDURE_CALL"
	IdentifierNodeTypeNewSpecification         IdentifierNodeTypeEnum = "NEW_SPECIFICATION"
	IdentifierNodeTypeFinal                    IdentifierNodeTypeEnum = "FINAL"
	IdentifierNodeTypeRunning                  IdentifierNodeTypeEnum = "RUNNING"
	IdentifierNodeTypePrev                     IdentifierNodeTypeEnum = "PREV"
	IdentifierNodeTypeNext                     IdentifierNodeTypeEnum = "NEXT"
	IdentifierNodeTypeFirst                    IdentifierNodeTypeEnum = "FIRST"
	IdentifierNodeTypeLast                     IdentifierNodeTypeEnum = "LAST"
	IdentifierNodeTypeClassifier               IdentifierNodeTypeEnum = "CLASSIFIER"
	IdentifierNodeTypeMatchNumber              IdentifierNodeTypeEnum = "MATCH_NUMBER"
	IdentifierNodeTypeSkipToFirst              IdentifierNodeTypeEnum = "SKIP_TO_FIRST"
	IdentifierNodeTypeSkipToLast               IdentifierNodeTypeEnum = "SKIP_TO_LAST"
	IdentifierNodeTypeDescending               IdentifierNodeTypeEnum = "DESCENDING"
	IdentifierNodeTypeNullsFirst               IdentifierNodeTypeEnum = "NULLS_FIRST"
	IdentifierNodeTypeNullsLast                IdentifierNodeTypeEnum = "NULLS_LAST"
	IdentifierNodeTypeIsTrue                   IdentifierNodeTypeEnum = "IS_TRUE"
	IdentifierNodeTypeIsFalse                  IdentifierNodeTypeEnum = "IS_FALSE"
	IdentifierNodeTypeIsNotTrue                IdentifierNodeTypeEnum = "IS_NOT_TRUE"
	IdentifierNodeTypeIsNotFalse               IdentifierNodeTypeEnum = "IS_NOT_FALSE"
	IdentifierNodeTypeIsUnknown                IdentifierNodeTypeEnum = "IS_UNKNOWN"
	IdentifierNodeTypeIsNull                   IdentifierNodeTypeEnum = "IS_NULL"
	IdentifierNodeTypeIsNotNull                IdentifierNodeTypeEnum = "IS_NOT_NULL"
	IdentifierNodeTypePreceding                IdentifierNodeTypeEnum = "PRECEDING"
	IdentifierNodeTypeFollowing                IdentifierNodeTypeEnum = "FOLLOWING"
	IdentifierNodeTypeFieldAccess              IdentifierNodeTypeEnum = "FIELD_ACCESS"
	IdentifierNodeTypeInputRef                 IdentifierNodeTypeEnum = "INPUT_REF"
	IdentifierNodeTypeTableInputRef            IdentifierNodeTypeEnum = "TABLE_INPUT_REF"
	IdentifierNodeTypePatternInputRef          IdentifierNodeTypeEnum = "PATTERN_INPUT_REF"
	IdentifierNodeTypeLocalRef                 IdentifierNodeTypeEnum = "LOCAL_REF"
	IdentifierNodeTypeCorrelVariable           IdentifierNodeTypeEnum = "CORREL_VARIABLE"
	IdentifierNodeTypePatternQuantifier        IdentifierNodeTypeEnum = "PATTERN_QUANTIFIER"
	IdentifierNodeTypeRow                      IdentifierNodeTypeEnum = "ROW"
	IdentifierNodeTypeColumnList               IdentifierNodeTypeEnum = "COLUMN_LIST"
	IdentifierNodeTypeCast                     IdentifierNodeTypeEnum = "CAST"
	IdentifierNodeTypeNextValue                IdentifierNodeTypeEnum = "NEXT_VALUE"
	IdentifierNodeTypeCurrentValue             IdentifierNodeTypeEnum = "CURRENT_VALUE"
	IdentifierNodeTypeFloor                    IdentifierNodeTypeEnum = "FLOOR"
	IdentifierNodeTypeCeil                     IdentifierNodeTypeEnum = "CEIL"
	IdentifierNodeTypeTrim                     IdentifierNodeTypeEnum = "TRIM"
	IdentifierNodeTypeLtrim                    IdentifierNodeTypeEnum = "LTRIM"
	IdentifierNodeTypeRtrim                    IdentifierNodeTypeEnum = "RTRIM"
	IdentifierNodeTypeExtract                  IdentifierNodeTypeEnum = "EXTRACT"
	IdentifierNodeTypeJdbcFn                   IdentifierNodeTypeEnum = "JDBC_FN"
	IdentifierNodeTypeMultisetValueConstructor IdentifierNodeTypeEnum = "MULTISET_VALUE_CONSTRUCTOR"
	IdentifierNodeTypeMultisetQueryConstructor IdentifierNodeTypeEnum = "MULTISET_QUERY_CONSTRUCTOR"
	IdentifierNodeTypeUnnest                   IdentifierNodeTypeEnum = "UNNEST"
	IdentifierNodeTypeLateral                  IdentifierNodeTypeEnum = "LATERAL"
	IdentifierNodeTypeCollectionTable          IdentifierNodeTypeEnum = "COLLECTION_TABLE"
	IdentifierNodeTypeArrayValueConstructor    IdentifierNodeTypeEnum = "ARRAY_VALUE_CONSTRUCTOR"
	IdentifierNodeTypeArrayQueryConstructor    IdentifierNodeTypeEnum = "ARRAY_QUERY_CONSTRUCTOR"
	IdentifierNodeTypeMapValueConstructor      IdentifierNodeTypeEnum = "MAP_VALUE_CONSTRUCTOR"
	IdentifierNodeTypeMapQueryConstructor      IdentifierNodeTypeEnum = "MAP_QUERY_CONSTRUCTOR"
	IdentifierNodeTypeCursor                   IdentifierNodeTypeEnum = "CURSOR"
	IdentifierNodeTypeLiteralChain             IdentifierNodeTypeEnum = "LITERAL_CHAIN"
	IdentifierNodeTypeEscape                   IdentifierNodeTypeEnum = "ESCAPE"
	IdentifierNodeTypeReinterpret              IdentifierNodeTypeEnum = "REINTERPRET"
	IdentifierNodeTypeExtend                   IdentifierNodeTypeEnum = "EXTEND"
	IdentifierNodeTypeCube                     IdentifierNodeTypeEnum = "CUBE"
	IdentifierNodeTypeRollup                   IdentifierNodeTypeEnum = "ROLLUP"
	IdentifierNodeTypeGroupingSets             IdentifierNodeTypeEnum = "GROUPING_SETS"
	IdentifierNodeTypeGrouping                 IdentifierNodeTypeEnum = "GROUPING"
	IdentifierNodeTypeGroupingId               IdentifierNodeTypeEnum = "GROUPING_ID"
	IdentifierNodeTypeGroupId                  IdentifierNodeTypeEnum = "GROUP_ID"
	IdentifierNodeTypePatternPermute           IdentifierNodeTypeEnum = "PATTERN_PERMUTE"
	IdentifierNodeTypePatternExcluded          IdentifierNodeTypeEnum = "PATTERN_EXCLUDED"
	IdentifierNodeTypeCount                    IdentifierNodeTypeEnum = "COUNT"
	IdentifierNodeTypeSum                      IdentifierNodeTypeEnum = "SUM"
	IdentifierNodeTypeSum0                     IdentifierNodeTypeEnum = "SUM0"
	IdentifierNodeTypeMin                      IdentifierNodeTypeEnum = "MIN"
	IdentifierNodeTypeMax                      IdentifierNodeTypeEnum = "MAX"
	IdentifierNodeTypeLead                     IdentifierNodeTypeEnum = "LEAD"
	IdentifierNodeTypeLag                      IdentifierNodeTypeEnum = "LAG"
	IdentifierNodeTypeFirstValue               IdentifierNodeTypeEnum = "FIRST_VALUE"
	IdentifierNodeTypeLastValue                IdentifierNodeTypeEnum = "LAST_VALUE"
	IdentifierNodeTypeCovarPop                 IdentifierNodeTypeEnum = "COVAR_POP"
	IdentifierNodeTypeCovarSamp                IdentifierNodeTypeEnum = "COVAR_SAMP"
	IdentifierNodeTypeRegrSxx                  IdentifierNodeTypeEnum = "REGR_SXX"
	IdentifierNodeTypeRegrSyy                  IdentifierNodeTypeEnum = "REGR_SYY"
	IdentifierNodeTypeAvg                      IdentifierNodeTypeEnum = "AVG"
	IdentifierNodeTypeStddevPop                IdentifierNodeTypeEnum = "STDDEV_POP"
	IdentifierNodeTypeStddevSamp               IdentifierNodeTypeEnum = "STDDEV_SAMP"
	IdentifierNodeTypeVarPop                   IdentifierNodeTypeEnum = "VAR_POP"
	IdentifierNodeTypeVarSamp                  IdentifierNodeTypeEnum = "VAR_SAMP"
	IdentifierNodeTypeNtile                    IdentifierNodeTypeEnum = "NTILE"
	IdentifierNodeTypeCollect                  IdentifierNodeTypeEnum = "COLLECT"
	IdentifierNodeTypeFusion                   IdentifierNodeTypeEnum = "FUSION"
	IdentifierNodeTypeSingleValue              IdentifierNodeTypeEnum = "SINGLE_VALUE"
	IdentifierNodeTypeRowNumber                IdentifierNodeTypeEnum = "ROW_NUMBER"
	IdentifierNodeTypeRank                     IdentifierNodeTypeEnum = "RANK"
	IdentifierNodeTypePercentRank              IdentifierNodeTypeEnum = "PERCENT_RANK"
	IdentifierNodeTypeDenseRank                IdentifierNodeTypeEnum = "DENSE_RANK"
	IdentifierNodeTypeCumeDist                 IdentifierNodeTypeEnum = "CUME_DIST"
	IdentifierNodeTypeTumble                   IdentifierNodeTypeEnum = "TUMBLE"
	IdentifierNodeTypeTumbleStart              IdentifierNodeTypeEnum = "TUMBLE_START"
	IdentifierNodeTypeTumbleEnd                IdentifierNodeTypeEnum = "TUMBLE_END"
	IdentifierNodeTypeHop                      IdentifierNodeTypeEnum = "HOP"
	IdentifierNodeTypeHopStart                 IdentifierNodeTypeEnum = "HOP_START"
	IdentifierNodeTypeHopEnd                   IdentifierNodeTypeEnum = "HOP_END"
	IdentifierNodeTypeSession                  IdentifierNodeTypeEnum = "SESSION"
	IdentifierNodeTypeSessionStart             IdentifierNodeTypeEnum = "SESSION_START"
	IdentifierNodeTypeSessionEnd               IdentifierNodeTypeEnum = "SESSION_END"
	IdentifierNodeTypeColumnDecl               IdentifierNodeTypeEnum = "COLUMN_DECL"
	IdentifierNodeTypeCheck                    IdentifierNodeTypeEnum = "CHECK"
	IdentifierNodeTypeUnique                   IdentifierNodeTypeEnum = "UNIQUE"
	IdentifierNodeTypePrimaryKey               IdentifierNodeTypeEnum = "PRIMARY_KEY"
	IdentifierNodeTypeForeignKey               IdentifierNodeTypeEnum = "FOREIGN_KEY"
	IdentifierNodeTypeCommit                   IdentifierNodeTypeEnum = "COMMIT"
	IdentifierNodeTypeRollback                 IdentifierNodeTypeEnum = "ROLLBACK"
	IdentifierNodeTypeAlterSession             IdentifierNodeTypeEnum = "ALTER_SESSION"
	IdentifierNodeTypeCreateSchema             IdentifierNodeTypeEnum = "CREATE_SCHEMA"
	IdentifierNodeTypeCreateForeignSchema      IdentifierNodeTypeEnum = "CREATE_FOREIGN_SCHEMA"
	IdentifierNodeTypeDropSchema               IdentifierNodeTypeEnum = "DROP_SCHEMA"
	IdentifierNodeTypeCreateTable              IdentifierNodeTypeEnum = "CREATE_TABLE"
	IdentifierNodeTypeAlterTable               IdentifierNodeTypeEnum = "ALTER_TABLE"
	IdentifierNodeTypeDropTable                IdentifierNodeTypeEnum = "DROP_TABLE"
	IdentifierNodeTypeCreateView               IdentifierNodeTypeEnum = "CREATE_VIEW"
	IdentifierNodeTypeAlterView                IdentifierNodeTypeEnum = "ALTER_VIEW"
	IdentifierNodeTypeDropView                 IdentifierNodeTypeEnum = "DROP_VIEW"
	IdentifierNodeTypeCreateMaterializedView   IdentifierNodeTypeEnum = "CREATE_MATERIALIZED_VIEW"
	IdentifierNodeTypeAlterMaterializedView    IdentifierNodeTypeEnum = "ALTER_MATERIALIZED_VIEW"
	IdentifierNodeTypeDropMaterializedView     IdentifierNodeTypeEnum = "DROP_MATERIALIZED_VIEW"
	IdentifierNodeTypeCreateSequence           IdentifierNodeTypeEnum = "CREATE_SEQUENCE"
	IdentifierNodeTypeAlterSequence            IdentifierNodeTypeEnum = "ALTER_SEQUENCE"
	IdentifierNodeTypeDropSequence             IdentifierNodeTypeEnum = "DROP_SEQUENCE"
	IdentifierNodeTypeCreateIndex              IdentifierNodeTypeEnum = "CREATE_INDEX"
	IdentifierNodeTypeAlterIndex               IdentifierNodeTypeEnum = "ALTER_INDEX"
	IdentifierNodeTypeDropIndex                IdentifierNodeTypeEnum = "DROP_INDEX"
	IdentifierNodeTypeOtherDdl                 IdentifierNodeTypeEnum = "OTHER_DDL"
)

var mappingIdentifierNodeType = map[string]IdentifierNodeTypeEnum{
	"OTHER":           IdentifierNodeTypeOther,
	"SELECT":          IdentifierNodeTypeSelect,
	"JOIN":            IdentifierNodeTypeJoin,
	"IDENTIFIER":      IdentifierNodeTypeIdentifier,
	"LITERAL":         IdentifierNodeTypeLiteral,
	"OTHER_FUNCTION":  IdentifierNodeTypeOtherFunction,
	"EXPLAIN":         IdentifierNodeTypeExplain,
	"DESCRIBE_SCHEMA": IdentifierNodeTypeDescribeSchema,
	"DESCRIBE_TABLE":  IdentifierNodeTypeDescribeTable,
	"INSERT":          IdentifierNodeTypeInsert,
	"DELETE":          IdentifierNodeTypeDelete,
	"UPDATE":          IdentifierNodeTypeUpdate,
	"SET_OPTION":      IdentifierNodeTypeSetOption,
	"DYNAMIC_PARAM":   IdentifierNodeTypeDynamicParam,
	"ORDER_BY":        IdentifierNodeTypeOrderBy,
	"WITH":            IdentifierNodeTypeWith,
	"WITH_ITEM":       IdentifierNodeTypeWithItem,
	"UNION":           IdentifierNodeTypeUnion,
	"EXCEPT":          IdentifierNodeTypeExcept,
	"INTERSECT":       IdentifierNodeTypeIntersect,
	"AS":              IdentifierNodeTypeAs,
	"ARGUMENT_ASSIGNMENT":   IdentifierNodeTypeArgumentAssignment,
	"DEFAULT":               IdentifierNodeTypeDefault,
	"OVER":                  IdentifierNodeTypeOver,
	"FILTER":                IdentifierNodeTypeFilter,
	"WINDOW":                IdentifierNodeTypeWindow,
	"MERGE":                 IdentifierNodeTypeMerge,
	"TABLESAMPLE":           IdentifierNodeTypeTablesample,
	"MATCH_RECOGNIZE":       IdentifierNodeTypeMatchRecognize,
	"TIMES":                 IdentifierNodeTypeTimes,
	"DIVIDE":                IdentifierNodeTypeDivide,
	"MOD":                   IdentifierNodeTypeMod,
	"PLUS":                  IdentifierNodeTypePlus,
	"MINUS":                 IdentifierNodeTypeMinus,
	"PATTERN_ALTER":         IdentifierNodeTypePatternAlter,
	"PATTERN_CONCAT":        IdentifierNodeTypePatternConcat,
	"IN":                    IdentifierNodeTypeIn,
	"NOT_IN":                IdentifierNodeTypeNotIn,
	"LESS_THAN":             IdentifierNodeTypeLessThan,
	"GREATER_THAN":          IdentifierNodeTypeGreaterThan,
	"LESS_THAN_OR_EQUAL":    IdentifierNodeTypeLessThanOrEqual,
	"GREATER_THAN_OR_EQUAL": IdentifierNodeTypeGreaterThanOrEqual,
	"EQUALS":                IdentifierNodeTypeEquals,
	"NOT_EQUALS":            IdentifierNodeTypeNotEquals,
	"IS_DISTINCT_FROM":      IdentifierNodeTypeIsDistinctFrom,
	"IS_NOT_DISTINCT_FROM":  IdentifierNodeTypeIsNotDistinctFrom,
	"OR":                         IdentifierNodeTypeOr,
	"AND":                        IdentifierNodeTypeAnd,
	"DOT":                        IdentifierNodeTypeDot,
	"OVERLAPS":                   IdentifierNodeTypeOverlaps,
	"CONTAINS":                   IdentifierNodeTypeContains,
	"PRECEDES":                   IdentifierNodeTypePrecedes,
	"IMMEDIATELY_PRECEDES":       IdentifierNodeTypeImmediatelyPrecedes,
	"SUCCEEDS":                   IdentifierNodeTypeSucceeds,
	"IMMEDIATELY_SUCCEEDS":       IdentifierNodeTypeImmediatelySucceeds,
	"PERIOD_EQUALS":              IdentifierNodeTypePeriodEquals,
	"LIKE":                       IdentifierNodeTypeLike,
	"SIMILAR":                    IdentifierNodeTypeSimilar,
	"BETWEEN":                    IdentifierNodeTypeBetween,
	"CASE":                       IdentifierNodeTypeCase,
	"NULLIF":                     IdentifierNodeTypeNullif,
	"COALESCE":                   IdentifierNodeTypeCoalesce,
	"DECODE":                     IdentifierNodeTypeDecode,
	"NVL":                        IdentifierNodeTypeNvl,
	"GREATEST":                   IdentifierNodeTypeGreatest,
	"LEAST":                      IdentifierNodeTypeLeast,
	"TIMESTAMP_ADD":              IdentifierNodeTypeTimestampAdd,
	"TIMESTAMP_DIFF":             IdentifierNodeTypeTimestampDiff,
	"NOT":                        IdentifierNodeTypeNot,
	"PLUS_PREFIX":                IdentifierNodeTypePlusPrefix,
	"MINUS_PREFIX":               IdentifierNodeTypeMinusPrefix,
	"EXISTS":                     IdentifierNodeTypeExists,
	"SOME":                       IdentifierNodeTypeSome,
	"ALL":                        IdentifierNodeTypeAll,
	"VALUES":                     IdentifierNodeTypeValues,
	"EXPLICIT_TABLE":             IdentifierNodeTypeExplicitTable,
	"SCALAR_QUERY":               IdentifierNodeTypeScalarQuery,
	"PROCEDURE_CALL":             IdentifierNodeTypeProcedureCall,
	"NEW_SPECIFICATION":          IdentifierNodeTypeNewSpecification,
	"FINAL":                      IdentifierNodeTypeFinal,
	"RUNNING":                    IdentifierNodeTypeRunning,
	"PREV":                       IdentifierNodeTypePrev,
	"NEXT":                       IdentifierNodeTypeNext,
	"FIRST":                      IdentifierNodeTypeFirst,
	"LAST":                       IdentifierNodeTypeLast,
	"CLASSIFIER":                 IdentifierNodeTypeClassifier,
	"MATCH_NUMBER":               IdentifierNodeTypeMatchNumber,
	"SKIP_TO_FIRST":              IdentifierNodeTypeSkipToFirst,
	"SKIP_TO_LAST":               IdentifierNodeTypeSkipToLast,
	"DESCENDING":                 IdentifierNodeTypeDescending,
	"NULLS_FIRST":                IdentifierNodeTypeNullsFirst,
	"NULLS_LAST":                 IdentifierNodeTypeNullsLast,
	"IS_TRUE":                    IdentifierNodeTypeIsTrue,
	"IS_FALSE":                   IdentifierNodeTypeIsFalse,
	"IS_NOT_TRUE":                IdentifierNodeTypeIsNotTrue,
	"IS_NOT_FALSE":               IdentifierNodeTypeIsNotFalse,
	"IS_UNKNOWN":                 IdentifierNodeTypeIsUnknown,
	"IS_NULL":                    IdentifierNodeTypeIsNull,
	"IS_NOT_NULL":                IdentifierNodeTypeIsNotNull,
	"PRECEDING":                  IdentifierNodeTypePreceding,
	"FOLLOWING":                  IdentifierNodeTypeFollowing,
	"FIELD_ACCESS":               IdentifierNodeTypeFieldAccess,
	"INPUT_REF":                  IdentifierNodeTypeInputRef,
	"TABLE_INPUT_REF":            IdentifierNodeTypeTableInputRef,
	"PATTERN_INPUT_REF":          IdentifierNodeTypePatternInputRef,
	"LOCAL_REF":                  IdentifierNodeTypeLocalRef,
	"CORREL_VARIABLE":            IdentifierNodeTypeCorrelVariable,
	"PATTERN_QUANTIFIER":         IdentifierNodeTypePatternQuantifier,
	"ROW":                        IdentifierNodeTypeRow,
	"COLUMN_LIST":                IdentifierNodeTypeColumnList,
	"CAST":                       IdentifierNodeTypeCast,
	"NEXT_VALUE":                 IdentifierNodeTypeNextValue,
	"CURRENT_VALUE":              IdentifierNodeTypeCurrentValue,
	"FLOOR":                      IdentifierNodeTypeFloor,
	"CEIL":                       IdentifierNodeTypeCeil,
	"TRIM":                       IdentifierNodeTypeTrim,
	"LTRIM":                      IdentifierNodeTypeLtrim,
	"RTRIM":                      IdentifierNodeTypeRtrim,
	"EXTRACT":                    IdentifierNodeTypeExtract,
	"JDBC_FN":                    IdentifierNodeTypeJdbcFn,
	"MULTISET_VALUE_CONSTRUCTOR": IdentifierNodeTypeMultisetValueConstructor,
	"MULTISET_QUERY_CONSTRUCTOR": IdentifierNodeTypeMultisetQueryConstructor,
	"UNNEST":                     IdentifierNodeTypeUnnest,
	"LATERAL":                    IdentifierNodeTypeLateral,
	"COLLECTION_TABLE":           IdentifierNodeTypeCollectionTable,
	"ARRAY_VALUE_CONSTRUCTOR":    IdentifierNodeTypeArrayValueConstructor,
	"ARRAY_QUERY_CONSTRUCTOR":    IdentifierNodeTypeArrayQueryConstructor,
	"MAP_VALUE_CONSTRUCTOR":      IdentifierNodeTypeMapValueConstructor,
	"MAP_QUERY_CONSTRUCTOR":      IdentifierNodeTypeMapQueryConstructor,
	"CURSOR":                     IdentifierNodeTypeCursor,
	"LITERAL_CHAIN":              IdentifierNodeTypeLiteralChain,
	"ESCAPE":                     IdentifierNodeTypeEscape,
	"REINTERPRET":                IdentifierNodeTypeReinterpret,
	"EXTEND":                     IdentifierNodeTypeExtend,
	"CUBE":                       IdentifierNodeTypeCube,
	"ROLLUP":                     IdentifierNodeTypeRollup,
	"GROUPING_SETS":              IdentifierNodeTypeGroupingSets,
	"GROUPING":                   IdentifierNodeTypeGrouping,
	"GROUPING_ID":                IdentifierNodeTypeGroupingId,
	"GROUP_ID":                   IdentifierNodeTypeGroupId,
	"PATTERN_PERMUTE":            IdentifierNodeTypePatternPermute,
	"PATTERN_EXCLUDED":           IdentifierNodeTypePatternExcluded,
	"COUNT":                      IdentifierNodeTypeCount,
	"SUM":                        IdentifierNodeTypeSum,
	"SUM0":                       IdentifierNodeTypeSum0,
	"MIN":                        IdentifierNodeTypeMin,
	"MAX":                        IdentifierNodeTypeMax,
	"LEAD":                       IdentifierNodeTypeLead,
	"LAG":                        IdentifierNodeTypeLag,
	"FIRST_VALUE":                IdentifierNodeTypeFirstValue,
	"LAST_VALUE":                 IdentifierNodeTypeLastValue,
	"COVAR_POP":                  IdentifierNodeTypeCovarPop,
	"COVAR_SAMP":                 IdentifierNodeTypeCovarSamp,
	"REGR_SXX":                   IdentifierNodeTypeRegrSxx,
	"REGR_SYY":                   IdentifierNodeTypeRegrSyy,
	"AVG":                        IdentifierNodeTypeAvg,
	"STDDEV_POP":                 IdentifierNodeTypeStddevPop,
	"STDDEV_SAMP":                IdentifierNodeTypeStddevSamp,
	"VAR_POP":                    IdentifierNodeTypeVarPop,
	"VAR_SAMP":                   IdentifierNodeTypeVarSamp,
	"NTILE":                      IdentifierNodeTypeNtile,
	"COLLECT":                    IdentifierNodeTypeCollect,
	"FUSION":                     IdentifierNodeTypeFusion,
	"SINGLE_VALUE":               IdentifierNodeTypeSingleValue,
	"ROW_NUMBER":                 IdentifierNodeTypeRowNumber,
	"RANK":                       IdentifierNodeTypeRank,
	"PERCENT_RANK":               IdentifierNodeTypePercentRank,
	"DENSE_RANK":                 IdentifierNodeTypeDenseRank,
	"CUME_DIST":                  IdentifierNodeTypeCumeDist,
	"TUMBLE":                     IdentifierNodeTypeTumble,
	"TUMBLE_START":               IdentifierNodeTypeTumbleStart,
	"TUMBLE_END":                 IdentifierNodeTypeTumbleEnd,
	"HOP":                        IdentifierNodeTypeHop,
	"HOP_START":                  IdentifierNodeTypeHopStart,
	"HOP_END":                    IdentifierNodeTypeHopEnd,
	"SESSION":                    IdentifierNodeTypeSession,
	"SESSION_START":              IdentifierNodeTypeSessionStart,
	"SESSION_END":                IdentifierNodeTypeSessionEnd,
	"COLUMN_DECL":                IdentifierNodeTypeColumnDecl,
	"CHECK":                      IdentifierNodeTypeCheck,
	"UNIQUE":                     IdentifierNodeTypeUnique,
	"PRIMARY_KEY":                IdentifierNodeTypePrimaryKey,
	"FOREIGN_KEY":                IdentifierNodeTypeForeignKey,
	"COMMIT":                     IdentifierNodeTypeCommit,
	"ROLLBACK":                   IdentifierNodeTypeRollback,
	"ALTER_SESSION":              IdentifierNodeTypeAlterSession,
	"CREATE_SCHEMA":              IdentifierNodeTypeCreateSchema,
	"CREATE_FOREIGN_SCHEMA":      IdentifierNodeTypeCreateForeignSchema,
	"DROP_SCHEMA":                IdentifierNodeTypeDropSchema,
	"CREATE_TABLE":               IdentifierNodeTypeCreateTable,
	"ALTER_TABLE":                IdentifierNodeTypeAlterTable,
	"DROP_TABLE":                 IdentifierNodeTypeDropTable,
	"CREATE_VIEW":                IdentifierNodeTypeCreateView,
	"ALTER_VIEW":                 IdentifierNodeTypeAlterView,
	"DROP_VIEW":                  IdentifierNodeTypeDropView,
	"CREATE_MATERIALIZED_VIEW":   IdentifierNodeTypeCreateMaterializedView,
	"ALTER_MATERIALIZED_VIEW":    IdentifierNodeTypeAlterMaterializedView,
	"DROP_MATERIALIZED_VIEW":     IdentifierNodeTypeDropMaterializedView,
	"CREATE_SEQUENCE":            IdentifierNodeTypeCreateSequence,
	"ALTER_SEQUENCE":             IdentifierNodeTypeAlterSequence,
	"DROP_SEQUENCE":              IdentifierNodeTypeDropSequence,
	"CREATE_INDEX":               IdentifierNodeTypeCreateIndex,
	"ALTER_INDEX":                IdentifierNodeTypeAlterIndex,
	"DROP_INDEX":                 IdentifierNodeTypeDropIndex,
	"OTHER_DDL":                  IdentifierNodeTypeOtherDdl,
}

// GetIdentifierNodeTypeEnumValues Enumerates the set of values for IdentifierNodeTypeEnum
func GetIdentifierNodeTypeEnumValues() []IdentifierNodeTypeEnum {
	values := make([]IdentifierNodeTypeEnum, 0)
	for _, v := range mappingIdentifierNodeType {
		values = append(values, v)
	}
	return values
}
