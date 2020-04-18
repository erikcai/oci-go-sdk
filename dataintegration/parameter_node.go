// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ParameterNode auto generated description
type ParameterNode struct {

	// auto generated description
	ModelType ParameterNodeModelTypeEnum `mandatory:"true" json:"modelType"`

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// type
	Type ParameterNodeTypeEnum `mandatory:"false" json:"type,omitempty"`

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

	// paramName
	ParamName *string `mandatory:"false" json:"paramName"`
}

func (m ParameterNode) String() string {
	return common.PointerString(m)
}

// ParameterNodeModelTypeEnum Enum with underlying type: string
type ParameterNodeModelTypeEnum string

// Set of constants representing the allowable values for ParameterNodeModelTypeEnum
const (
	ParameterNodeModelTypeReferenceNode  ParameterNodeModelTypeEnum = "REFERENCE_NODE"
	ParameterNodeModelTypeOperatorNode   ParameterNodeModelTypeEnum = "OPERATOR_NODE"
	ParameterNodeModelTypeRootNode       ParameterNodeModelTypeEnum = "ROOT_NODE"
	ParameterNodeModelTypeValueNode      ParameterNodeModelTypeEnum = "VALUE_NODE"
	ParameterNodeModelTypeIdentifierNode ParameterNodeModelTypeEnum = "IDENTIFIER_NODE"
	ParameterNodeModelTypeParameterNode  ParameterNodeModelTypeEnum = "PARAMETER_NODE"
	ParameterNodeModelTypeFunctionNode   ParameterNodeModelTypeEnum = "FUNCTION_NODE"
)

var mappingParameterNodeModelType = map[string]ParameterNodeModelTypeEnum{
	"REFERENCE_NODE":  ParameterNodeModelTypeReferenceNode,
	"OPERATOR_NODE":   ParameterNodeModelTypeOperatorNode,
	"ROOT_NODE":       ParameterNodeModelTypeRootNode,
	"VALUE_NODE":      ParameterNodeModelTypeValueNode,
	"IDENTIFIER_NODE": ParameterNodeModelTypeIdentifierNode,
	"PARAMETER_NODE":  ParameterNodeModelTypeParameterNode,
	"FUNCTION_NODE":   ParameterNodeModelTypeFunctionNode,
}

// GetParameterNodeModelTypeEnumValues Enumerates the set of values for ParameterNodeModelTypeEnum
func GetParameterNodeModelTypeEnumValues() []ParameterNodeModelTypeEnum {
	values := make([]ParameterNodeModelTypeEnum, 0)
	for _, v := range mappingParameterNodeModelType {
		values = append(values, v)
	}
	return values
}

// ParameterNodeTypeEnum Enum with underlying type: string
type ParameterNodeTypeEnum string

// Set of constants representing the allowable values for ParameterNodeTypeEnum
const (
	ParameterNodeTypeOther                    ParameterNodeTypeEnum = "OTHER"
	ParameterNodeTypeSelect                   ParameterNodeTypeEnum = "SELECT"
	ParameterNodeTypeJoin                     ParameterNodeTypeEnum = "JOIN"
	ParameterNodeTypeIdentifier               ParameterNodeTypeEnum = "IDENTIFIER"
	ParameterNodeTypeLiteral                  ParameterNodeTypeEnum = "LITERAL"
	ParameterNodeTypeOtherFunction            ParameterNodeTypeEnum = "OTHER_FUNCTION"
	ParameterNodeTypeExplain                  ParameterNodeTypeEnum = "EXPLAIN"
	ParameterNodeTypeDescribeSchema           ParameterNodeTypeEnum = "DESCRIBE_SCHEMA"
	ParameterNodeTypeDescribeTable            ParameterNodeTypeEnum = "DESCRIBE_TABLE"
	ParameterNodeTypeInsert                   ParameterNodeTypeEnum = "INSERT"
	ParameterNodeTypeDelete                   ParameterNodeTypeEnum = "DELETE"
	ParameterNodeTypeUpdate                   ParameterNodeTypeEnum = "UPDATE"
	ParameterNodeTypeSetOption                ParameterNodeTypeEnum = "SET_OPTION"
	ParameterNodeTypeDynamicParam             ParameterNodeTypeEnum = "DYNAMIC_PARAM"
	ParameterNodeTypeOrderBy                  ParameterNodeTypeEnum = "ORDER_BY"
	ParameterNodeTypeWith                     ParameterNodeTypeEnum = "WITH"
	ParameterNodeTypeWithItem                 ParameterNodeTypeEnum = "WITH_ITEM"
	ParameterNodeTypeUnion                    ParameterNodeTypeEnum = "UNION"
	ParameterNodeTypeExcept                   ParameterNodeTypeEnum = "EXCEPT"
	ParameterNodeTypeIntersect                ParameterNodeTypeEnum = "INTERSECT"
	ParameterNodeTypeAs                       ParameterNodeTypeEnum = "AS"
	ParameterNodeTypeArgumentAssignment       ParameterNodeTypeEnum = "ARGUMENT_ASSIGNMENT"
	ParameterNodeTypeDefault                  ParameterNodeTypeEnum = "DEFAULT"
	ParameterNodeTypeOver                     ParameterNodeTypeEnum = "OVER"
	ParameterNodeTypeFilter                   ParameterNodeTypeEnum = "FILTER"
	ParameterNodeTypeWindow                   ParameterNodeTypeEnum = "WINDOW"
	ParameterNodeTypeMerge                    ParameterNodeTypeEnum = "MERGE"
	ParameterNodeTypeTablesample              ParameterNodeTypeEnum = "TABLESAMPLE"
	ParameterNodeTypeMatchRecognize           ParameterNodeTypeEnum = "MATCH_RECOGNIZE"
	ParameterNodeTypeTimes                    ParameterNodeTypeEnum = "TIMES"
	ParameterNodeTypeDivide                   ParameterNodeTypeEnum = "DIVIDE"
	ParameterNodeTypeMod                      ParameterNodeTypeEnum = "MOD"
	ParameterNodeTypePlus                     ParameterNodeTypeEnum = "PLUS"
	ParameterNodeTypeMinus                    ParameterNodeTypeEnum = "MINUS"
	ParameterNodeTypePatternAlter             ParameterNodeTypeEnum = "PATTERN_ALTER"
	ParameterNodeTypePatternConcat            ParameterNodeTypeEnum = "PATTERN_CONCAT"
	ParameterNodeTypeIn                       ParameterNodeTypeEnum = "IN"
	ParameterNodeTypeNotIn                    ParameterNodeTypeEnum = "NOT_IN"
	ParameterNodeTypeLessThan                 ParameterNodeTypeEnum = "LESS_THAN"
	ParameterNodeTypeGreaterThan              ParameterNodeTypeEnum = "GREATER_THAN"
	ParameterNodeTypeLessThanOrEqual          ParameterNodeTypeEnum = "LESS_THAN_OR_EQUAL"
	ParameterNodeTypeGreaterThanOrEqual       ParameterNodeTypeEnum = "GREATER_THAN_OR_EQUAL"
	ParameterNodeTypeEquals                   ParameterNodeTypeEnum = "EQUALS"
	ParameterNodeTypeNotEquals                ParameterNodeTypeEnum = "NOT_EQUALS"
	ParameterNodeTypeIsDistinctFrom           ParameterNodeTypeEnum = "IS_DISTINCT_FROM"
	ParameterNodeTypeIsNotDistinctFrom        ParameterNodeTypeEnum = "IS_NOT_DISTINCT_FROM"
	ParameterNodeTypeOr                       ParameterNodeTypeEnum = "OR"
	ParameterNodeTypeAnd                      ParameterNodeTypeEnum = "AND"
	ParameterNodeTypeDot                      ParameterNodeTypeEnum = "DOT"
	ParameterNodeTypeOverlaps                 ParameterNodeTypeEnum = "OVERLAPS"
	ParameterNodeTypeContains                 ParameterNodeTypeEnum = "CONTAINS"
	ParameterNodeTypePrecedes                 ParameterNodeTypeEnum = "PRECEDES"
	ParameterNodeTypeImmediatelyPrecedes      ParameterNodeTypeEnum = "IMMEDIATELY_PRECEDES"
	ParameterNodeTypeSucceeds                 ParameterNodeTypeEnum = "SUCCEEDS"
	ParameterNodeTypeImmediatelySucceeds      ParameterNodeTypeEnum = "IMMEDIATELY_SUCCEEDS"
	ParameterNodeTypePeriodEquals             ParameterNodeTypeEnum = "PERIOD_EQUALS"
	ParameterNodeTypeLike                     ParameterNodeTypeEnum = "LIKE"
	ParameterNodeTypeSimilar                  ParameterNodeTypeEnum = "SIMILAR"
	ParameterNodeTypeBetween                  ParameterNodeTypeEnum = "BETWEEN"
	ParameterNodeTypeCase                     ParameterNodeTypeEnum = "CASE"
	ParameterNodeTypeNullif                   ParameterNodeTypeEnum = "NULLIF"
	ParameterNodeTypeCoalesce                 ParameterNodeTypeEnum = "COALESCE"
	ParameterNodeTypeDecode                   ParameterNodeTypeEnum = "DECODE"
	ParameterNodeTypeNvl                      ParameterNodeTypeEnum = "NVL"
	ParameterNodeTypeGreatest                 ParameterNodeTypeEnum = "GREATEST"
	ParameterNodeTypeLeast                    ParameterNodeTypeEnum = "LEAST"
	ParameterNodeTypeTimestampAdd             ParameterNodeTypeEnum = "TIMESTAMP_ADD"
	ParameterNodeTypeTimestampDiff            ParameterNodeTypeEnum = "TIMESTAMP_DIFF"
	ParameterNodeTypeNot                      ParameterNodeTypeEnum = "NOT"
	ParameterNodeTypePlusPrefix               ParameterNodeTypeEnum = "PLUS_PREFIX"
	ParameterNodeTypeMinusPrefix              ParameterNodeTypeEnum = "MINUS_PREFIX"
	ParameterNodeTypeExists                   ParameterNodeTypeEnum = "EXISTS"
	ParameterNodeTypeSome                     ParameterNodeTypeEnum = "SOME"
	ParameterNodeTypeAll                      ParameterNodeTypeEnum = "ALL"
	ParameterNodeTypeValues                   ParameterNodeTypeEnum = "VALUES"
	ParameterNodeTypeExplicitTable            ParameterNodeTypeEnum = "EXPLICIT_TABLE"
	ParameterNodeTypeScalarQuery              ParameterNodeTypeEnum = "SCALAR_QUERY"
	ParameterNodeTypeProcedureCall            ParameterNodeTypeEnum = "PROCEDURE_CALL"
	ParameterNodeTypeNewSpecification         ParameterNodeTypeEnum = "NEW_SPECIFICATION"
	ParameterNodeTypeFinal                    ParameterNodeTypeEnum = "FINAL"
	ParameterNodeTypeRunning                  ParameterNodeTypeEnum = "RUNNING"
	ParameterNodeTypePrev                     ParameterNodeTypeEnum = "PREV"
	ParameterNodeTypeNext                     ParameterNodeTypeEnum = "NEXT"
	ParameterNodeTypeFirst                    ParameterNodeTypeEnum = "FIRST"
	ParameterNodeTypeLast                     ParameterNodeTypeEnum = "LAST"
	ParameterNodeTypeClassifier               ParameterNodeTypeEnum = "CLASSIFIER"
	ParameterNodeTypeMatchNumber              ParameterNodeTypeEnum = "MATCH_NUMBER"
	ParameterNodeTypeSkipToFirst              ParameterNodeTypeEnum = "SKIP_TO_FIRST"
	ParameterNodeTypeSkipToLast               ParameterNodeTypeEnum = "SKIP_TO_LAST"
	ParameterNodeTypeDescending               ParameterNodeTypeEnum = "DESCENDING"
	ParameterNodeTypeNullsFirst               ParameterNodeTypeEnum = "NULLS_FIRST"
	ParameterNodeTypeNullsLast                ParameterNodeTypeEnum = "NULLS_LAST"
	ParameterNodeTypeIsTrue                   ParameterNodeTypeEnum = "IS_TRUE"
	ParameterNodeTypeIsFalse                  ParameterNodeTypeEnum = "IS_FALSE"
	ParameterNodeTypeIsNotTrue                ParameterNodeTypeEnum = "IS_NOT_TRUE"
	ParameterNodeTypeIsNotFalse               ParameterNodeTypeEnum = "IS_NOT_FALSE"
	ParameterNodeTypeIsUnknown                ParameterNodeTypeEnum = "IS_UNKNOWN"
	ParameterNodeTypeIsNull                   ParameterNodeTypeEnum = "IS_NULL"
	ParameterNodeTypeIsNotNull                ParameterNodeTypeEnum = "IS_NOT_NULL"
	ParameterNodeTypePreceding                ParameterNodeTypeEnum = "PRECEDING"
	ParameterNodeTypeFollowing                ParameterNodeTypeEnum = "FOLLOWING"
	ParameterNodeTypeFieldAccess              ParameterNodeTypeEnum = "FIELD_ACCESS"
	ParameterNodeTypeInputRef                 ParameterNodeTypeEnum = "INPUT_REF"
	ParameterNodeTypeTableInputRef            ParameterNodeTypeEnum = "TABLE_INPUT_REF"
	ParameterNodeTypePatternInputRef          ParameterNodeTypeEnum = "PATTERN_INPUT_REF"
	ParameterNodeTypeLocalRef                 ParameterNodeTypeEnum = "LOCAL_REF"
	ParameterNodeTypeCorrelVariable           ParameterNodeTypeEnum = "CORREL_VARIABLE"
	ParameterNodeTypePatternQuantifier        ParameterNodeTypeEnum = "PATTERN_QUANTIFIER"
	ParameterNodeTypeRow                      ParameterNodeTypeEnum = "ROW"
	ParameterNodeTypeColumnList               ParameterNodeTypeEnum = "COLUMN_LIST"
	ParameterNodeTypeCast                     ParameterNodeTypeEnum = "CAST"
	ParameterNodeTypeNextValue                ParameterNodeTypeEnum = "NEXT_VALUE"
	ParameterNodeTypeCurrentValue             ParameterNodeTypeEnum = "CURRENT_VALUE"
	ParameterNodeTypeFloor                    ParameterNodeTypeEnum = "FLOOR"
	ParameterNodeTypeCeil                     ParameterNodeTypeEnum = "CEIL"
	ParameterNodeTypeTrim                     ParameterNodeTypeEnum = "TRIM"
	ParameterNodeTypeLtrim                    ParameterNodeTypeEnum = "LTRIM"
	ParameterNodeTypeRtrim                    ParameterNodeTypeEnum = "RTRIM"
	ParameterNodeTypeExtract                  ParameterNodeTypeEnum = "EXTRACT"
	ParameterNodeTypeJdbcFn                   ParameterNodeTypeEnum = "JDBC_FN"
	ParameterNodeTypeMultisetValueConstructor ParameterNodeTypeEnum = "MULTISET_VALUE_CONSTRUCTOR"
	ParameterNodeTypeMultisetQueryConstructor ParameterNodeTypeEnum = "MULTISET_QUERY_CONSTRUCTOR"
	ParameterNodeTypeUnnest                   ParameterNodeTypeEnum = "UNNEST"
	ParameterNodeTypeLateral                  ParameterNodeTypeEnum = "LATERAL"
	ParameterNodeTypeCollectionTable          ParameterNodeTypeEnum = "COLLECTION_TABLE"
	ParameterNodeTypeArrayValueConstructor    ParameterNodeTypeEnum = "ARRAY_VALUE_CONSTRUCTOR"
	ParameterNodeTypeArrayQueryConstructor    ParameterNodeTypeEnum = "ARRAY_QUERY_CONSTRUCTOR"
	ParameterNodeTypeMapValueConstructor      ParameterNodeTypeEnum = "MAP_VALUE_CONSTRUCTOR"
	ParameterNodeTypeMapQueryConstructor      ParameterNodeTypeEnum = "MAP_QUERY_CONSTRUCTOR"
	ParameterNodeTypeCursor                   ParameterNodeTypeEnum = "CURSOR"
	ParameterNodeTypeLiteralChain             ParameterNodeTypeEnum = "LITERAL_CHAIN"
	ParameterNodeTypeEscape                   ParameterNodeTypeEnum = "ESCAPE"
	ParameterNodeTypeReinterpret              ParameterNodeTypeEnum = "REINTERPRET"
	ParameterNodeTypeExtend                   ParameterNodeTypeEnum = "EXTEND"
	ParameterNodeTypeCube                     ParameterNodeTypeEnum = "CUBE"
	ParameterNodeTypeRollup                   ParameterNodeTypeEnum = "ROLLUP"
	ParameterNodeTypeGroupingSets             ParameterNodeTypeEnum = "GROUPING_SETS"
	ParameterNodeTypeGrouping                 ParameterNodeTypeEnum = "GROUPING"
	ParameterNodeTypeGroupingId               ParameterNodeTypeEnum = "GROUPING_ID"
	ParameterNodeTypeGroupId                  ParameterNodeTypeEnum = "GROUP_ID"
	ParameterNodeTypePatternPermute           ParameterNodeTypeEnum = "PATTERN_PERMUTE"
	ParameterNodeTypePatternExcluded          ParameterNodeTypeEnum = "PATTERN_EXCLUDED"
	ParameterNodeTypeCount                    ParameterNodeTypeEnum = "COUNT"
	ParameterNodeTypeSum                      ParameterNodeTypeEnum = "SUM"
	ParameterNodeTypeSum0                     ParameterNodeTypeEnum = "SUM0"
	ParameterNodeTypeMin                      ParameterNodeTypeEnum = "MIN"
	ParameterNodeTypeMax                      ParameterNodeTypeEnum = "MAX"
	ParameterNodeTypeLead                     ParameterNodeTypeEnum = "LEAD"
	ParameterNodeTypeLag                      ParameterNodeTypeEnum = "LAG"
	ParameterNodeTypeFirstValue               ParameterNodeTypeEnum = "FIRST_VALUE"
	ParameterNodeTypeLastValue                ParameterNodeTypeEnum = "LAST_VALUE"
	ParameterNodeTypeCovarPop                 ParameterNodeTypeEnum = "COVAR_POP"
	ParameterNodeTypeCovarSamp                ParameterNodeTypeEnum = "COVAR_SAMP"
	ParameterNodeTypeRegrSxx                  ParameterNodeTypeEnum = "REGR_SXX"
	ParameterNodeTypeRegrSyy                  ParameterNodeTypeEnum = "REGR_SYY"
	ParameterNodeTypeAvg                      ParameterNodeTypeEnum = "AVG"
	ParameterNodeTypeStddevPop                ParameterNodeTypeEnum = "STDDEV_POP"
	ParameterNodeTypeStddevSamp               ParameterNodeTypeEnum = "STDDEV_SAMP"
	ParameterNodeTypeVarPop                   ParameterNodeTypeEnum = "VAR_POP"
	ParameterNodeTypeVarSamp                  ParameterNodeTypeEnum = "VAR_SAMP"
	ParameterNodeTypeNtile                    ParameterNodeTypeEnum = "NTILE"
	ParameterNodeTypeCollect                  ParameterNodeTypeEnum = "COLLECT"
	ParameterNodeTypeFusion                   ParameterNodeTypeEnum = "FUSION"
	ParameterNodeTypeSingleValue              ParameterNodeTypeEnum = "SINGLE_VALUE"
	ParameterNodeTypeRowNumber                ParameterNodeTypeEnum = "ROW_NUMBER"
	ParameterNodeTypeRank                     ParameterNodeTypeEnum = "RANK"
	ParameterNodeTypePercentRank              ParameterNodeTypeEnum = "PERCENT_RANK"
	ParameterNodeTypeDenseRank                ParameterNodeTypeEnum = "DENSE_RANK"
	ParameterNodeTypeCumeDist                 ParameterNodeTypeEnum = "CUME_DIST"
	ParameterNodeTypeTumble                   ParameterNodeTypeEnum = "TUMBLE"
	ParameterNodeTypeTumbleStart              ParameterNodeTypeEnum = "TUMBLE_START"
	ParameterNodeTypeTumbleEnd                ParameterNodeTypeEnum = "TUMBLE_END"
	ParameterNodeTypeHop                      ParameterNodeTypeEnum = "HOP"
	ParameterNodeTypeHopStart                 ParameterNodeTypeEnum = "HOP_START"
	ParameterNodeTypeHopEnd                   ParameterNodeTypeEnum = "HOP_END"
	ParameterNodeTypeSession                  ParameterNodeTypeEnum = "SESSION"
	ParameterNodeTypeSessionStart             ParameterNodeTypeEnum = "SESSION_START"
	ParameterNodeTypeSessionEnd               ParameterNodeTypeEnum = "SESSION_END"
	ParameterNodeTypeColumnDecl               ParameterNodeTypeEnum = "COLUMN_DECL"
	ParameterNodeTypeCheck                    ParameterNodeTypeEnum = "CHECK"
	ParameterNodeTypeUnique                   ParameterNodeTypeEnum = "UNIQUE"
	ParameterNodeTypePrimaryKey               ParameterNodeTypeEnum = "PRIMARY_KEY"
	ParameterNodeTypeForeignKey               ParameterNodeTypeEnum = "FOREIGN_KEY"
	ParameterNodeTypeCommit                   ParameterNodeTypeEnum = "COMMIT"
	ParameterNodeTypeRollback                 ParameterNodeTypeEnum = "ROLLBACK"
	ParameterNodeTypeAlterSession             ParameterNodeTypeEnum = "ALTER_SESSION"
	ParameterNodeTypeCreateSchema             ParameterNodeTypeEnum = "CREATE_SCHEMA"
	ParameterNodeTypeCreateForeignSchema      ParameterNodeTypeEnum = "CREATE_FOREIGN_SCHEMA"
	ParameterNodeTypeDropSchema               ParameterNodeTypeEnum = "DROP_SCHEMA"
	ParameterNodeTypeCreateTable              ParameterNodeTypeEnum = "CREATE_TABLE"
	ParameterNodeTypeAlterTable               ParameterNodeTypeEnum = "ALTER_TABLE"
	ParameterNodeTypeDropTable                ParameterNodeTypeEnum = "DROP_TABLE"
	ParameterNodeTypeCreateView               ParameterNodeTypeEnum = "CREATE_VIEW"
	ParameterNodeTypeAlterView                ParameterNodeTypeEnum = "ALTER_VIEW"
	ParameterNodeTypeDropView                 ParameterNodeTypeEnum = "DROP_VIEW"
	ParameterNodeTypeCreateMaterializedView   ParameterNodeTypeEnum = "CREATE_MATERIALIZED_VIEW"
	ParameterNodeTypeAlterMaterializedView    ParameterNodeTypeEnum = "ALTER_MATERIALIZED_VIEW"
	ParameterNodeTypeDropMaterializedView     ParameterNodeTypeEnum = "DROP_MATERIALIZED_VIEW"
	ParameterNodeTypeCreateSequence           ParameterNodeTypeEnum = "CREATE_SEQUENCE"
	ParameterNodeTypeAlterSequence            ParameterNodeTypeEnum = "ALTER_SEQUENCE"
	ParameterNodeTypeDropSequence             ParameterNodeTypeEnum = "DROP_SEQUENCE"
	ParameterNodeTypeCreateIndex              ParameterNodeTypeEnum = "CREATE_INDEX"
	ParameterNodeTypeAlterIndex               ParameterNodeTypeEnum = "ALTER_INDEX"
	ParameterNodeTypeDropIndex                ParameterNodeTypeEnum = "DROP_INDEX"
	ParameterNodeTypeOtherDdl                 ParameterNodeTypeEnum = "OTHER_DDL"
)

var mappingParameterNodeType = map[string]ParameterNodeTypeEnum{
	"OTHER":           ParameterNodeTypeOther,
	"SELECT":          ParameterNodeTypeSelect,
	"JOIN":            ParameterNodeTypeJoin,
	"IDENTIFIER":      ParameterNodeTypeIdentifier,
	"LITERAL":         ParameterNodeTypeLiteral,
	"OTHER_FUNCTION":  ParameterNodeTypeOtherFunction,
	"EXPLAIN":         ParameterNodeTypeExplain,
	"DESCRIBE_SCHEMA": ParameterNodeTypeDescribeSchema,
	"DESCRIBE_TABLE":  ParameterNodeTypeDescribeTable,
	"INSERT":          ParameterNodeTypeInsert,
	"DELETE":          ParameterNodeTypeDelete,
	"UPDATE":          ParameterNodeTypeUpdate,
	"SET_OPTION":      ParameterNodeTypeSetOption,
	"DYNAMIC_PARAM":   ParameterNodeTypeDynamicParam,
	"ORDER_BY":        ParameterNodeTypeOrderBy,
	"WITH":            ParameterNodeTypeWith,
	"WITH_ITEM":       ParameterNodeTypeWithItem,
	"UNION":           ParameterNodeTypeUnion,
	"EXCEPT":          ParameterNodeTypeExcept,
	"INTERSECT":       ParameterNodeTypeIntersect,
	"AS":              ParameterNodeTypeAs,
	"ARGUMENT_ASSIGNMENT":   ParameterNodeTypeArgumentAssignment,
	"DEFAULT":               ParameterNodeTypeDefault,
	"OVER":                  ParameterNodeTypeOver,
	"FILTER":                ParameterNodeTypeFilter,
	"WINDOW":                ParameterNodeTypeWindow,
	"MERGE":                 ParameterNodeTypeMerge,
	"TABLESAMPLE":           ParameterNodeTypeTablesample,
	"MATCH_RECOGNIZE":       ParameterNodeTypeMatchRecognize,
	"TIMES":                 ParameterNodeTypeTimes,
	"DIVIDE":                ParameterNodeTypeDivide,
	"MOD":                   ParameterNodeTypeMod,
	"PLUS":                  ParameterNodeTypePlus,
	"MINUS":                 ParameterNodeTypeMinus,
	"PATTERN_ALTER":         ParameterNodeTypePatternAlter,
	"PATTERN_CONCAT":        ParameterNodeTypePatternConcat,
	"IN":                    ParameterNodeTypeIn,
	"NOT_IN":                ParameterNodeTypeNotIn,
	"LESS_THAN":             ParameterNodeTypeLessThan,
	"GREATER_THAN":          ParameterNodeTypeGreaterThan,
	"LESS_THAN_OR_EQUAL":    ParameterNodeTypeLessThanOrEqual,
	"GREATER_THAN_OR_EQUAL": ParameterNodeTypeGreaterThanOrEqual,
	"EQUALS":                ParameterNodeTypeEquals,
	"NOT_EQUALS":            ParameterNodeTypeNotEquals,
	"IS_DISTINCT_FROM":      ParameterNodeTypeIsDistinctFrom,
	"IS_NOT_DISTINCT_FROM":  ParameterNodeTypeIsNotDistinctFrom,
	"OR":                         ParameterNodeTypeOr,
	"AND":                        ParameterNodeTypeAnd,
	"DOT":                        ParameterNodeTypeDot,
	"OVERLAPS":                   ParameterNodeTypeOverlaps,
	"CONTAINS":                   ParameterNodeTypeContains,
	"PRECEDES":                   ParameterNodeTypePrecedes,
	"IMMEDIATELY_PRECEDES":       ParameterNodeTypeImmediatelyPrecedes,
	"SUCCEEDS":                   ParameterNodeTypeSucceeds,
	"IMMEDIATELY_SUCCEEDS":       ParameterNodeTypeImmediatelySucceeds,
	"PERIOD_EQUALS":              ParameterNodeTypePeriodEquals,
	"LIKE":                       ParameterNodeTypeLike,
	"SIMILAR":                    ParameterNodeTypeSimilar,
	"BETWEEN":                    ParameterNodeTypeBetween,
	"CASE":                       ParameterNodeTypeCase,
	"NULLIF":                     ParameterNodeTypeNullif,
	"COALESCE":                   ParameterNodeTypeCoalesce,
	"DECODE":                     ParameterNodeTypeDecode,
	"NVL":                        ParameterNodeTypeNvl,
	"GREATEST":                   ParameterNodeTypeGreatest,
	"LEAST":                      ParameterNodeTypeLeast,
	"TIMESTAMP_ADD":              ParameterNodeTypeTimestampAdd,
	"TIMESTAMP_DIFF":             ParameterNodeTypeTimestampDiff,
	"NOT":                        ParameterNodeTypeNot,
	"PLUS_PREFIX":                ParameterNodeTypePlusPrefix,
	"MINUS_PREFIX":               ParameterNodeTypeMinusPrefix,
	"EXISTS":                     ParameterNodeTypeExists,
	"SOME":                       ParameterNodeTypeSome,
	"ALL":                        ParameterNodeTypeAll,
	"VALUES":                     ParameterNodeTypeValues,
	"EXPLICIT_TABLE":             ParameterNodeTypeExplicitTable,
	"SCALAR_QUERY":               ParameterNodeTypeScalarQuery,
	"PROCEDURE_CALL":             ParameterNodeTypeProcedureCall,
	"NEW_SPECIFICATION":          ParameterNodeTypeNewSpecification,
	"FINAL":                      ParameterNodeTypeFinal,
	"RUNNING":                    ParameterNodeTypeRunning,
	"PREV":                       ParameterNodeTypePrev,
	"NEXT":                       ParameterNodeTypeNext,
	"FIRST":                      ParameterNodeTypeFirst,
	"LAST":                       ParameterNodeTypeLast,
	"CLASSIFIER":                 ParameterNodeTypeClassifier,
	"MATCH_NUMBER":               ParameterNodeTypeMatchNumber,
	"SKIP_TO_FIRST":              ParameterNodeTypeSkipToFirst,
	"SKIP_TO_LAST":               ParameterNodeTypeSkipToLast,
	"DESCENDING":                 ParameterNodeTypeDescending,
	"NULLS_FIRST":                ParameterNodeTypeNullsFirst,
	"NULLS_LAST":                 ParameterNodeTypeNullsLast,
	"IS_TRUE":                    ParameterNodeTypeIsTrue,
	"IS_FALSE":                   ParameterNodeTypeIsFalse,
	"IS_NOT_TRUE":                ParameterNodeTypeIsNotTrue,
	"IS_NOT_FALSE":               ParameterNodeTypeIsNotFalse,
	"IS_UNKNOWN":                 ParameterNodeTypeIsUnknown,
	"IS_NULL":                    ParameterNodeTypeIsNull,
	"IS_NOT_NULL":                ParameterNodeTypeIsNotNull,
	"PRECEDING":                  ParameterNodeTypePreceding,
	"FOLLOWING":                  ParameterNodeTypeFollowing,
	"FIELD_ACCESS":               ParameterNodeTypeFieldAccess,
	"INPUT_REF":                  ParameterNodeTypeInputRef,
	"TABLE_INPUT_REF":            ParameterNodeTypeTableInputRef,
	"PATTERN_INPUT_REF":          ParameterNodeTypePatternInputRef,
	"LOCAL_REF":                  ParameterNodeTypeLocalRef,
	"CORREL_VARIABLE":            ParameterNodeTypeCorrelVariable,
	"PATTERN_QUANTIFIER":         ParameterNodeTypePatternQuantifier,
	"ROW":                        ParameterNodeTypeRow,
	"COLUMN_LIST":                ParameterNodeTypeColumnList,
	"CAST":                       ParameterNodeTypeCast,
	"NEXT_VALUE":                 ParameterNodeTypeNextValue,
	"CURRENT_VALUE":              ParameterNodeTypeCurrentValue,
	"FLOOR":                      ParameterNodeTypeFloor,
	"CEIL":                       ParameterNodeTypeCeil,
	"TRIM":                       ParameterNodeTypeTrim,
	"LTRIM":                      ParameterNodeTypeLtrim,
	"RTRIM":                      ParameterNodeTypeRtrim,
	"EXTRACT":                    ParameterNodeTypeExtract,
	"JDBC_FN":                    ParameterNodeTypeJdbcFn,
	"MULTISET_VALUE_CONSTRUCTOR": ParameterNodeTypeMultisetValueConstructor,
	"MULTISET_QUERY_CONSTRUCTOR": ParameterNodeTypeMultisetQueryConstructor,
	"UNNEST":                     ParameterNodeTypeUnnest,
	"LATERAL":                    ParameterNodeTypeLateral,
	"COLLECTION_TABLE":           ParameterNodeTypeCollectionTable,
	"ARRAY_VALUE_CONSTRUCTOR":    ParameterNodeTypeArrayValueConstructor,
	"ARRAY_QUERY_CONSTRUCTOR":    ParameterNodeTypeArrayQueryConstructor,
	"MAP_VALUE_CONSTRUCTOR":      ParameterNodeTypeMapValueConstructor,
	"MAP_QUERY_CONSTRUCTOR":      ParameterNodeTypeMapQueryConstructor,
	"CURSOR":                     ParameterNodeTypeCursor,
	"LITERAL_CHAIN":              ParameterNodeTypeLiteralChain,
	"ESCAPE":                     ParameterNodeTypeEscape,
	"REINTERPRET":                ParameterNodeTypeReinterpret,
	"EXTEND":                     ParameterNodeTypeExtend,
	"CUBE":                       ParameterNodeTypeCube,
	"ROLLUP":                     ParameterNodeTypeRollup,
	"GROUPING_SETS":              ParameterNodeTypeGroupingSets,
	"GROUPING":                   ParameterNodeTypeGrouping,
	"GROUPING_ID":                ParameterNodeTypeGroupingId,
	"GROUP_ID":                   ParameterNodeTypeGroupId,
	"PATTERN_PERMUTE":            ParameterNodeTypePatternPermute,
	"PATTERN_EXCLUDED":           ParameterNodeTypePatternExcluded,
	"COUNT":                      ParameterNodeTypeCount,
	"SUM":                        ParameterNodeTypeSum,
	"SUM0":                       ParameterNodeTypeSum0,
	"MIN":                        ParameterNodeTypeMin,
	"MAX":                        ParameterNodeTypeMax,
	"LEAD":                       ParameterNodeTypeLead,
	"LAG":                        ParameterNodeTypeLag,
	"FIRST_VALUE":                ParameterNodeTypeFirstValue,
	"LAST_VALUE":                 ParameterNodeTypeLastValue,
	"COVAR_POP":                  ParameterNodeTypeCovarPop,
	"COVAR_SAMP":                 ParameterNodeTypeCovarSamp,
	"REGR_SXX":                   ParameterNodeTypeRegrSxx,
	"REGR_SYY":                   ParameterNodeTypeRegrSyy,
	"AVG":                        ParameterNodeTypeAvg,
	"STDDEV_POP":                 ParameterNodeTypeStddevPop,
	"STDDEV_SAMP":                ParameterNodeTypeStddevSamp,
	"VAR_POP":                    ParameterNodeTypeVarPop,
	"VAR_SAMP":                   ParameterNodeTypeVarSamp,
	"NTILE":                      ParameterNodeTypeNtile,
	"COLLECT":                    ParameterNodeTypeCollect,
	"FUSION":                     ParameterNodeTypeFusion,
	"SINGLE_VALUE":               ParameterNodeTypeSingleValue,
	"ROW_NUMBER":                 ParameterNodeTypeRowNumber,
	"RANK":                       ParameterNodeTypeRank,
	"PERCENT_RANK":               ParameterNodeTypePercentRank,
	"DENSE_RANK":                 ParameterNodeTypeDenseRank,
	"CUME_DIST":                  ParameterNodeTypeCumeDist,
	"TUMBLE":                     ParameterNodeTypeTumble,
	"TUMBLE_START":               ParameterNodeTypeTumbleStart,
	"TUMBLE_END":                 ParameterNodeTypeTumbleEnd,
	"HOP":                        ParameterNodeTypeHop,
	"HOP_START":                  ParameterNodeTypeHopStart,
	"HOP_END":                    ParameterNodeTypeHopEnd,
	"SESSION":                    ParameterNodeTypeSession,
	"SESSION_START":              ParameterNodeTypeSessionStart,
	"SESSION_END":                ParameterNodeTypeSessionEnd,
	"COLUMN_DECL":                ParameterNodeTypeColumnDecl,
	"CHECK":                      ParameterNodeTypeCheck,
	"UNIQUE":                     ParameterNodeTypeUnique,
	"PRIMARY_KEY":                ParameterNodeTypePrimaryKey,
	"FOREIGN_KEY":                ParameterNodeTypeForeignKey,
	"COMMIT":                     ParameterNodeTypeCommit,
	"ROLLBACK":                   ParameterNodeTypeRollback,
	"ALTER_SESSION":              ParameterNodeTypeAlterSession,
	"CREATE_SCHEMA":              ParameterNodeTypeCreateSchema,
	"CREATE_FOREIGN_SCHEMA":      ParameterNodeTypeCreateForeignSchema,
	"DROP_SCHEMA":                ParameterNodeTypeDropSchema,
	"CREATE_TABLE":               ParameterNodeTypeCreateTable,
	"ALTER_TABLE":                ParameterNodeTypeAlterTable,
	"DROP_TABLE":                 ParameterNodeTypeDropTable,
	"CREATE_VIEW":                ParameterNodeTypeCreateView,
	"ALTER_VIEW":                 ParameterNodeTypeAlterView,
	"DROP_VIEW":                  ParameterNodeTypeDropView,
	"CREATE_MATERIALIZED_VIEW":   ParameterNodeTypeCreateMaterializedView,
	"ALTER_MATERIALIZED_VIEW":    ParameterNodeTypeAlterMaterializedView,
	"DROP_MATERIALIZED_VIEW":     ParameterNodeTypeDropMaterializedView,
	"CREATE_SEQUENCE":            ParameterNodeTypeCreateSequence,
	"ALTER_SEQUENCE":             ParameterNodeTypeAlterSequence,
	"DROP_SEQUENCE":              ParameterNodeTypeDropSequence,
	"CREATE_INDEX":               ParameterNodeTypeCreateIndex,
	"ALTER_INDEX":                ParameterNodeTypeAlterIndex,
	"DROP_INDEX":                 ParameterNodeTypeDropIndex,
	"OTHER_DDL":                  ParameterNodeTypeOtherDdl,
}

// GetParameterNodeTypeEnumValues Enumerates the set of values for ParameterNodeTypeEnum
func GetParameterNodeTypeEnumValues() []ParameterNodeTypeEnum {
	values := make([]ParameterNodeTypeEnum, 0)
	for _, v := range mappingParameterNodeType {
		values = append(values, v)
	}
	return values
}
