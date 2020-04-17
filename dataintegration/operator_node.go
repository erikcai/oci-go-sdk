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

// OperatorNode auto generated description
type OperatorNode struct {

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

	// operatorName
	OperatorName *string `mandatory:"false" json:"operatorName"`

	// isBinaryOperator
	IsBinaryOperator *bool `mandatory:"false" json:"isBinaryOperator"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// operands
	Operands []ExpressionNode `mandatory:"false" json:"operands"`

	// type
	Type OperatorNodeTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m OperatorNode) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OperatorNode) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOperatorNode OperatorNode
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOperatorNode
	}{
		"OPERATOR_NODE",
		(MarshalTypeOperatorNode)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OperatorNode) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key              *string              `json:"key"`
		ModelVersion     *string              `json:"modelVersion"`
		ParentRef        *ParentReference     `json:"parentRef"`
		Type             OperatorNodeTypeEnum `json:"type"`
		ParsedString     *string              `json:"parsedString"`
		StartPosition    *int                 `json:"startPosition"`
		EndPosition      *int                 `json:"endPosition"`
		OperatorName     *string              `json:"operatorName"`
		IsBinaryOperator *bool                `json:"isBinaryOperator"`
		ObjectStatus     *int                 `json:"objectStatus"`
		Operands         []expressionnode     `json:"operands"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Type = model.Type

	m.ParsedString = model.ParsedString

	m.StartPosition = model.StartPosition

	m.EndPosition = model.EndPosition

	m.OperatorName = model.OperatorName

	m.IsBinaryOperator = model.IsBinaryOperator

	m.ObjectStatus = model.ObjectStatus

	m.Operands = make([]ExpressionNode, len(model.Operands))
	for i, n := range model.Operands {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Operands[i] = nn.(ExpressionNode)
		} else {
			m.Operands[i] = nil
		}
	}
	return
}

// OperatorNodeTypeEnum Enum with underlying type: string
type OperatorNodeTypeEnum string

// Set of constants representing the allowable values for OperatorNodeTypeEnum
const (
	OperatorNodeTypeOther                    OperatorNodeTypeEnum = "OTHER"
	OperatorNodeTypeSelect                   OperatorNodeTypeEnum = "SELECT"
	OperatorNodeTypeJoin                     OperatorNodeTypeEnum = "JOIN"
	OperatorNodeTypeIdentifier               OperatorNodeTypeEnum = "IDENTIFIER"
	OperatorNodeTypeLiteral                  OperatorNodeTypeEnum = "LITERAL"
	OperatorNodeTypeOtherFunction            OperatorNodeTypeEnum = "OTHER_FUNCTION"
	OperatorNodeTypeExplain                  OperatorNodeTypeEnum = "EXPLAIN"
	OperatorNodeTypeDescribeSchema           OperatorNodeTypeEnum = "DESCRIBE_SCHEMA"
	OperatorNodeTypeDescribeTable            OperatorNodeTypeEnum = "DESCRIBE_TABLE"
	OperatorNodeTypeInsert                   OperatorNodeTypeEnum = "INSERT"
	OperatorNodeTypeDelete                   OperatorNodeTypeEnum = "DELETE"
	OperatorNodeTypeUpdate                   OperatorNodeTypeEnum = "UPDATE"
	OperatorNodeTypeSetOption                OperatorNodeTypeEnum = "SET_OPTION"
	OperatorNodeTypeDynamicParam             OperatorNodeTypeEnum = "DYNAMIC_PARAM"
	OperatorNodeTypeOrderBy                  OperatorNodeTypeEnum = "ORDER_BY"
	OperatorNodeTypeWith                     OperatorNodeTypeEnum = "WITH"
	OperatorNodeTypeWithItem                 OperatorNodeTypeEnum = "WITH_ITEM"
	OperatorNodeTypeUnion                    OperatorNodeTypeEnum = "UNION"
	OperatorNodeTypeExcept                   OperatorNodeTypeEnum = "EXCEPT"
	OperatorNodeTypeIntersect                OperatorNodeTypeEnum = "INTERSECT"
	OperatorNodeTypeAs                       OperatorNodeTypeEnum = "AS"
	OperatorNodeTypeArgumentAssignment       OperatorNodeTypeEnum = "ARGUMENT_ASSIGNMENT"
	OperatorNodeTypeDefault                  OperatorNodeTypeEnum = "DEFAULT"
	OperatorNodeTypeOver                     OperatorNodeTypeEnum = "OVER"
	OperatorNodeTypeFilter                   OperatorNodeTypeEnum = "FILTER"
	OperatorNodeTypeWindow                   OperatorNodeTypeEnum = "WINDOW"
	OperatorNodeTypeMerge                    OperatorNodeTypeEnum = "MERGE"
	OperatorNodeTypeTablesample              OperatorNodeTypeEnum = "TABLESAMPLE"
	OperatorNodeTypeMatchRecognize           OperatorNodeTypeEnum = "MATCH_RECOGNIZE"
	OperatorNodeTypeTimes                    OperatorNodeTypeEnum = "TIMES"
	OperatorNodeTypeDivide                   OperatorNodeTypeEnum = "DIVIDE"
	OperatorNodeTypeMod                      OperatorNodeTypeEnum = "MOD"
	OperatorNodeTypePlus                     OperatorNodeTypeEnum = "PLUS"
	OperatorNodeTypeMinus                    OperatorNodeTypeEnum = "MINUS"
	OperatorNodeTypePatternAlter             OperatorNodeTypeEnum = "PATTERN_ALTER"
	OperatorNodeTypePatternConcat            OperatorNodeTypeEnum = "PATTERN_CONCAT"
	OperatorNodeTypeIn                       OperatorNodeTypeEnum = "IN"
	OperatorNodeTypeNotIn                    OperatorNodeTypeEnum = "NOT_IN"
	OperatorNodeTypeLessThan                 OperatorNodeTypeEnum = "LESS_THAN"
	OperatorNodeTypeGreaterThan              OperatorNodeTypeEnum = "GREATER_THAN"
	OperatorNodeTypeLessThanOrEqual          OperatorNodeTypeEnum = "LESS_THAN_OR_EQUAL"
	OperatorNodeTypeGreaterThanOrEqual       OperatorNodeTypeEnum = "GREATER_THAN_OR_EQUAL"
	OperatorNodeTypeEquals                   OperatorNodeTypeEnum = "EQUALS"
	OperatorNodeTypeNotEquals                OperatorNodeTypeEnum = "NOT_EQUALS"
	OperatorNodeTypeIsDistinctFrom           OperatorNodeTypeEnum = "IS_DISTINCT_FROM"
	OperatorNodeTypeIsNotDistinctFrom        OperatorNodeTypeEnum = "IS_NOT_DISTINCT_FROM"
	OperatorNodeTypeOr                       OperatorNodeTypeEnum = "OR"
	OperatorNodeTypeAnd                      OperatorNodeTypeEnum = "AND"
	OperatorNodeTypeDot                      OperatorNodeTypeEnum = "DOT"
	OperatorNodeTypeOverlaps                 OperatorNodeTypeEnum = "OVERLAPS"
	OperatorNodeTypeContains                 OperatorNodeTypeEnum = "CONTAINS"
	OperatorNodeTypePrecedes                 OperatorNodeTypeEnum = "PRECEDES"
	OperatorNodeTypeImmediatelyPrecedes      OperatorNodeTypeEnum = "IMMEDIATELY_PRECEDES"
	OperatorNodeTypeSucceeds                 OperatorNodeTypeEnum = "SUCCEEDS"
	OperatorNodeTypeImmediatelySucceeds      OperatorNodeTypeEnum = "IMMEDIATELY_SUCCEEDS"
	OperatorNodeTypePeriodEquals             OperatorNodeTypeEnum = "PERIOD_EQUALS"
	OperatorNodeTypeLike                     OperatorNodeTypeEnum = "LIKE"
	OperatorNodeTypeSimilar                  OperatorNodeTypeEnum = "SIMILAR"
	OperatorNodeTypeBetween                  OperatorNodeTypeEnum = "BETWEEN"
	OperatorNodeTypeCase                     OperatorNodeTypeEnum = "CASE"
	OperatorNodeTypeNullif                   OperatorNodeTypeEnum = "NULLIF"
	OperatorNodeTypeCoalesce                 OperatorNodeTypeEnum = "COALESCE"
	OperatorNodeTypeDecode                   OperatorNodeTypeEnum = "DECODE"
	OperatorNodeTypeNvl                      OperatorNodeTypeEnum = "NVL"
	OperatorNodeTypeGreatest                 OperatorNodeTypeEnum = "GREATEST"
	OperatorNodeTypeLeast                    OperatorNodeTypeEnum = "LEAST"
	OperatorNodeTypeTimestampAdd             OperatorNodeTypeEnum = "TIMESTAMP_ADD"
	OperatorNodeTypeTimestampDiff            OperatorNodeTypeEnum = "TIMESTAMP_DIFF"
	OperatorNodeTypeNot                      OperatorNodeTypeEnum = "NOT"
	OperatorNodeTypePlusPrefix               OperatorNodeTypeEnum = "PLUS_PREFIX"
	OperatorNodeTypeMinusPrefix              OperatorNodeTypeEnum = "MINUS_PREFIX"
	OperatorNodeTypeExists                   OperatorNodeTypeEnum = "EXISTS"
	OperatorNodeTypeSome                     OperatorNodeTypeEnum = "SOME"
	OperatorNodeTypeAll                      OperatorNodeTypeEnum = "ALL"
	OperatorNodeTypeValues                   OperatorNodeTypeEnum = "VALUES"
	OperatorNodeTypeExplicitTable            OperatorNodeTypeEnum = "EXPLICIT_TABLE"
	OperatorNodeTypeScalarQuery              OperatorNodeTypeEnum = "SCALAR_QUERY"
	OperatorNodeTypeProcedureCall            OperatorNodeTypeEnum = "PROCEDURE_CALL"
	OperatorNodeTypeNewSpecification         OperatorNodeTypeEnum = "NEW_SPECIFICATION"
	OperatorNodeTypeFinal                    OperatorNodeTypeEnum = "FINAL"
	OperatorNodeTypeRunning                  OperatorNodeTypeEnum = "RUNNING"
	OperatorNodeTypePrev                     OperatorNodeTypeEnum = "PREV"
	OperatorNodeTypeNext                     OperatorNodeTypeEnum = "NEXT"
	OperatorNodeTypeFirst                    OperatorNodeTypeEnum = "FIRST"
	OperatorNodeTypeLast                     OperatorNodeTypeEnum = "LAST"
	OperatorNodeTypeClassifier               OperatorNodeTypeEnum = "CLASSIFIER"
	OperatorNodeTypeMatchNumber              OperatorNodeTypeEnum = "MATCH_NUMBER"
	OperatorNodeTypeSkipToFirst              OperatorNodeTypeEnum = "SKIP_TO_FIRST"
	OperatorNodeTypeSkipToLast               OperatorNodeTypeEnum = "SKIP_TO_LAST"
	OperatorNodeTypeDescending               OperatorNodeTypeEnum = "DESCENDING"
	OperatorNodeTypeNullsFirst               OperatorNodeTypeEnum = "NULLS_FIRST"
	OperatorNodeTypeNullsLast                OperatorNodeTypeEnum = "NULLS_LAST"
	OperatorNodeTypeIsTrue                   OperatorNodeTypeEnum = "IS_TRUE"
	OperatorNodeTypeIsFalse                  OperatorNodeTypeEnum = "IS_FALSE"
	OperatorNodeTypeIsNotTrue                OperatorNodeTypeEnum = "IS_NOT_TRUE"
	OperatorNodeTypeIsNotFalse               OperatorNodeTypeEnum = "IS_NOT_FALSE"
	OperatorNodeTypeIsUnknown                OperatorNodeTypeEnum = "IS_UNKNOWN"
	OperatorNodeTypeIsNull                   OperatorNodeTypeEnum = "IS_NULL"
	OperatorNodeTypeIsNotNull                OperatorNodeTypeEnum = "IS_NOT_NULL"
	OperatorNodeTypePreceding                OperatorNodeTypeEnum = "PRECEDING"
	OperatorNodeTypeFollowing                OperatorNodeTypeEnum = "FOLLOWING"
	OperatorNodeTypeFieldAccess              OperatorNodeTypeEnum = "FIELD_ACCESS"
	OperatorNodeTypeInputRef                 OperatorNodeTypeEnum = "INPUT_REF"
	OperatorNodeTypeTableInputRef            OperatorNodeTypeEnum = "TABLE_INPUT_REF"
	OperatorNodeTypePatternInputRef          OperatorNodeTypeEnum = "PATTERN_INPUT_REF"
	OperatorNodeTypeLocalRef                 OperatorNodeTypeEnum = "LOCAL_REF"
	OperatorNodeTypeCorrelVariable           OperatorNodeTypeEnum = "CORREL_VARIABLE"
	OperatorNodeTypePatternQuantifier        OperatorNodeTypeEnum = "PATTERN_QUANTIFIER"
	OperatorNodeTypeRow                      OperatorNodeTypeEnum = "ROW"
	OperatorNodeTypeColumnList               OperatorNodeTypeEnum = "COLUMN_LIST"
	OperatorNodeTypeCast                     OperatorNodeTypeEnum = "CAST"
	OperatorNodeTypeNextValue                OperatorNodeTypeEnum = "NEXT_VALUE"
	OperatorNodeTypeCurrentValue             OperatorNodeTypeEnum = "CURRENT_VALUE"
	OperatorNodeTypeFloor                    OperatorNodeTypeEnum = "FLOOR"
	OperatorNodeTypeCeil                     OperatorNodeTypeEnum = "CEIL"
	OperatorNodeTypeTrim                     OperatorNodeTypeEnum = "TRIM"
	OperatorNodeTypeLtrim                    OperatorNodeTypeEnum = "LTRIM"
	OperatorNodeTypeRtrim                    OperatorNodeTypeEnum = "RTRIM"
	OperatorNodeTypeExtract                  OperatorNodeTypeEnum = "EXTRACT"
	OperatorNodeTypeJdbcFn                   OperatorNodeTypeEnum = "JDBC_FN"
	OperatorNodeTypeMultisetValueConstructor OperatorNodeTypeEnum = "MULTISET_VALUE_CONSTRUCTOR"
	OperatorNodeTypeMultisetQueryConstructor OperatorNodeTypeEnum = "MULTISET_QUERY_CONSTRUCTOR"
	OperatorNodeTypeUnnest                   OperatorNodeTypeEnum = "UNNEST"
	OperatorNodeTypeLateral                  OperatorNodeTypeEnum = "LATERAL"
	OperatorNodeTypeCollectionTable          OperatorNodeTypeEnum = "COLLECTION_TABLE"
	OperatorNodeTypeArrayValueConstructor    OperatorNodeTypeEnum = "ARRAY_VALUE_CONSTRUCTOR"
	OperatorNodeTypeArrayQueryConstructor    OperatorNodeTypeEnum = "ARRAY_QUERY_CONSTRUCTOR"
	OperatorNodeTypeMapValueConstructor      OperatorNodeTypeEnum = "MAP_VALUE_CONSTRUCTOR"
	OperatorNodeTypeMapQueryConstructor      OperatorNodeTypeEnum = "MAP_QUERY_CONSTRUCTOR"
	OperatorNodeTypeCursor                   OperatorNodeTypeEnum = "CURSOR"
	OperatorNodeTypeLiteralChain             OperatorNodeTypeEnum = "LITERAL_CHAIN"
	OperatorNodeTypeEscape                   OperatorNodeTypeEnum = "ESCAPE"
	OperatorNodeTypeReinterpret              OperatorNodeTypeEnum = "REINTERPRET"
	OperatorNodeTypeExtend                   OperatorNodeTypeEnum = "EXTEND"
	OperatorNodeTypeCube                     OperatorNodeTypeEnum = "CUBE"
	OperatorNodeTypeRollup                   OperatorNodeTypeEnum = "ROLLUP"
	OperatorNodeTypeGroupingSets             OperatorNodeTypeEnum = "GROUPING_SETS"
	OperatorNodeTypeGrouping                 OperatorNodeTypeEnum = "GROUPING"
	OperatorNodeTypeGroupingId               OperatorNodeTypeEnum = "GROUPING_ID"
	OperatorNodeTypeGroupId                  OperatorNodeTypeEnum = "GROUP_ID"
	OperatorNodeTypePatternPermute           OperatorNodeTypeEnum = "PATTERN_PERMUTE"
	OperatorNodeTypePatternExcluded          OperatorNodeTypeEnum = "PATTERN_EXCLUDED"
	OperatorNodeTypeCount                    OperatorNodeTypeEnum = "COUNT"
	OperatorNodeTypeSum                      OperatorNodeTypeEnum = "SUM"
	OperatorNodeTypeSum0                     OperatorNodeTypeEnum = "SUM0"
	OperatorNodeTypeMin                      OperatorNodeTypeEnum = "MIN"
	OperatorNodeTypeMax                      OperatorNodeTypeEnum = "MAX"
	OperatorNodeTypeLead                     OperatorNodeTypeEnum = "LEAD"
	OperatorNodeTypeLag                      OperatorNodeTypeEnum = "LAG"
	OperatorNodeTypeFirstValue               OperatorNodeTypeEnum = "FIRST_VALUE"
	OperatorNodeTypeLastValue                OperatorNodeTypeEnum = "LAST_VALUE"
	OperatorNodeTypeCovarPop                 OperatorNodeTypeEnum = "COVAR_POP"
	OperatorNodeTypeCovarSamp                OperatorNodeTypeEnum = "COVAR_SAMP"
	OperatorNodeTypeRegrSxx                  OperatorNodeTypeEnum = "REGR_SXX"
	OperatorNodeTypeRegrSyy                  OperatorNodeTypeEnum = "REGR_SYY"
	OperatorNodeTypeAvg                      OperatorNodeTypeEnum = "AVG"
	OperatorNodeTypeStddevPop                OperatorNodeTypeEnum = "STDDEV_POP"
	OperatorNodeTypeStddevSamp               OperatorNodeTypeEnum = "STDDEV_SAMP"
	OperatorNodeTypeVarPop                   OperatorNodeTypeEnum = "VAR_POP"
	OperatorNodeTypeVarSamp                  OperatorNodeTypeEnum = "VAR_SAMP"
	OperatorNodeTypeNtile                    OperatorNodeTypeEnum = "NTILE"
	OperatorNodeTypeCollect                  OperatorNodeTypeEnum = "COLLECT"
	OperatorNodeTypeFusion                   OperatorNodeTypeEnum = "FUSION"
	OperatorNodeTypeSingleValue              OperatorNodeTypeEnum = "SINGLE_VALUE"
	OperatorNodeTypeRowNumber                OperatorNodeTypeEnum = "ROW_NUMBER"
	OperatorNodeTypeRank                     OperatorNodeTypeEnum = "RANK"
	OperatorNodeTypePercentRank              OperatorNodeTypeEnum = "PERCENT_RANK"
	OperatorNodeTypeDenseRank                OperatorNodeTypeEnum = "DENSE_RANK"
	OperatorNodeTypeCumeDist                 OperatorNodeTypeEnum = "CUME_DIST"
	OperatorNodeTypeTumble                   OperatorNodeTypeEnum = "TUMBLE"
	OperatorNodeTypeTumbleStart              OperatorNodeTypeEnum = "TUMBLE_START"
	OperatorNodeTypeTumbleEnd                OperatorNodeTypeEnum = "TUMBLE_END"
	OperatorNodeTypeHop                      OperatorNodeTypeEnum = "HOP"
	OperatorNodeTypeHopStart                 OperatorNodeTypeEnum = "HOP_START"
	OperatorNodeTypeHopEnd                   OperatorNodeTypeEnum = "HOP_END"
	OperatorNodeTypeSession                  OperatorNodeTypeEnum = "SESSION"
	OperatorNodeTypeSessionStart             OperatorNodeTypeEnum = "SESSION_START"
	OperatorNodeTypeSessionEnd               OperatorNodeTypeEnum = "SESSION_END"
	OperatorNodeTypeColumnDecl               OperatorNodeTypeEnum = "COLUMN_DECL"
	OperatorNodeTypeCheck                    OperatorNodeTypeEnum = "CHECK"
	OperatorNodeTypeUnique                   OperatorNodeTypeEnum = "UNIQUE"
	OperatorNodeTypePrimaryKey               OperatorNodeTypeEnum = "PRIMARY_KEY"
	OperatorNodeTypeForeignKey               OperatorNodeTypeEnum = "FOREIGN_KEY"
	OperatorNodeTypeCommit                   OperatorNodeTypeEnum = "COMMIT"
	OperatorNodeTypeRollback                 OperatorNodeTypeEnum = "ROLLBACK"
	OperatorNodeTypeAlterSession             OperatorNodeTypeEnum = "ALTER_SESSION"
	OperatorNodeTypeCreateSchema             OperatorNodeTypeEnum = "CREATE_SCHEMA"
	OperatorNodeTypeCreateForeignSchema      OperatorNodeTypeEnum = "CREATE_FOREIGN_SCHEMA"
	OperatorNodeTypeDropSchema               OperatorNodeTypeEnum = "DROP_SCHEMA"
	OperatorNodeTypeCreateTable              OperatorNodeTypeEnum = "CREATE_TABLE"
	OperatorNodeTypeAlterTable               OperatorNodeTypeEnum = "ALTER_TABLE"
	OperatorNodeTypeDropTable                OperatorNodeTypeEnum = "DROP_TABLE"
	OperatorNodeTypeCreateView               OperatorNodeTypeEnum = "CREATE_VIEW"
	OperatorNodeTypeAlterView                OperatorNodeTypeEnum = "ALTER_VIEW"
	OperatorNodeTypeDropView                 OperatorNodeTypeEnum = "DROP_VIEW"
	OperatorNodeTypeCreateMaterializedView   OperatorNodeTypeEnum = "CREATE_MATERIALIZED_VIEW"
	OperatorNodeTypeAlterMaterializedView    OperatorNodeTypeEnum = "ALTER_MATERIALIZED_VIEW"
	OperatorNodeTypeDropMaterializedView     OperatorNodeTypeEnum = "DROP_MATERIALIZED_VIEW"
	OperatorNodeTypeCreateSequence           OperatorNodeTypeEnum = "CREATE_SEQUENCE"
	OperatorNodeTypeAlterSequence            OperatorNodeTypeEnum = "ALTER_SEQUENCE"
	OperatorNodeTypeDropSequence             OperatorNodeTypeEnum = "DROP_SEQUENCE"
	OperatorNodeTypeCreateIndex              OperatorNodeTypeEnum = "CREATE_INDEX"
	OperatorNodeTypeAlterIndex               OperatorNodeTypeEnum = "ALTER_INDEX"
	OperatorNodeTypeDropIndex                OperatorNodeTypeEnum = "DROP_INDEX"
	OperatorNodeTypeOtherDdl                 OperatorNodeTypeEnum = "OTHER_DDL"
)

var mappingOperatorNodeType = map[string]OperatorNodeTypeEnum{
	"OTHER":                      OperatorNodeTypeOther,
	"SELECT":                     OperatorNodeTypeSelect,
	"JOIN":                       OperatorNodeTypeJoin,
	"IDENTIFIER":                 OperatorNodeTypeIdentifier,
	"LITERAL":                    OperatorNodeTypeLiteral,
	"OTHER_FUNCTION":             OperatorNodeTypeOtherFunction,
	"EXPLAIN":                    OperatorNodeTypeExplain,
	"DESCRIBE_SCHEMA":            OperatorNodeTypeDescribeSchema,
	"DESCRIBE_TABLE":             OperatorNodeTypeDescribeTable,
	"INSERT":                     OperatorNodeTypeInsert,
	"DELETE":                     OperatorNodeTypeDelete,
	"UPDATE":                     OperatorNodeTypeUpdate,
	"SET_OPTION":                 OperatorNodeTypeSetOption,
	"DYNAMIC_PARAM":              OperatorNodeTypeDynamicParam,
	"ORDER_BY":                   OperatorNodeTypeOrderBy,
	"WITH":                       OperatorNodeTypeWith,
	"WITH_ITEM":                  OperatorNodeTypeWithItem,
	"UNION":                      OperatorNodeTypeUnion,
	"EXCEPT":                     OperatorNodeTypeExcept,
	"INTERSECT":                  OperatorNodeTypeIntersect,
	"AS":                         OperatorNodeTypeAs,
	"ARGUMENT_ASSIGNMENT":        OperatorNodeTypeArgumentAssignment,
	"DEFAULT":                    OperatorNodeTypeDefault,
	"OVER":                       OperatorNodeTypeOver,
	"FILTER":                     OperatorNodeTypeFilter,
	"WINDOW":                     OperatorNodeTypeWindow,
	"MERGE":                      OperatorNodeTypeMerge,
	"TABLESAMPLE":                OperatorNodeTypeTablesample,
	"MATCH_RECOGNIZE":            OperatorNodeTypeMatchRecognize,
	"TIMES":                      OperatorNodeTypeTimes,
	"DIVIDE":                     OperatorNodeTypeDivide,
	"MOD":                        OperatorNodeTypeMod,
	"PLUS":                       OperatorNodeTypePlus,
	"MINUS":                      OperatorNodeTypeMinus,
	"PATTERN_ALTER":              OperatorNodeTypePatternAlter,
	"PATTERN_CONCAT":             OperatorNodeTypePatternConcat,
	"IN":                         OperatorNodeTypeIn,
	"NOT_IN":                     OperatorNodeTypeNotIn,
	"LESS_THAN":                  OperatorNodeTypeLessThan,
	"GREATER_THAN":               OperatorNodeTypeGreaterThan,
	"LESS_THAN_OR_EQUAL":         OperatorNodeTypeLessThanOrEqual,
	"GREATER_THAN_OR_EQUAL":      OperatorNodeTypeGreaterThanOrEqual,
	"EQUALS":                     OperatorNodeTypeEquals,
	"NOT_EQUALS":                 OperatorNodeTypeNotEquals,
	"IS_DISTINCT_FROM":           OperatorNodeTypeIsDistinctFrom,
	"IS_NOT_DISTINCT_FROM":       OperatorNodeTypeIsNotDistinctFrom,
	"OR":                         OperatorNodeTypeOr,
	"AND":                        OperatorNodeTypeAnd,
	"DOT":                        OperatorNodeTypeDot,
	"OVERLAPS":                   OperatorNodeTypeOverlaps,
	"CONTAINS":                   OperatorNodeTypeContains,
	"PRECEDES":                   OperatorNodeTypePrecedes,
	"IMMEDIATELY_PRECEDES":       OperatorNodeTypeImmediatelyPrecedes,
	"SUCCEEDS":                   OperatorNodeTypeSucceeds,
	"IMMEDIATELY_SUCCEEDS":       OperatorNodeTypeImmediatelySucceeds,
	"PERIOD_EQUALS":              OperatorNodeTypePeriodEquals,
	"LIKE":                       OperatorNodeTypeLike,
	"SIMILAR":                    OperatorNodeTypeSimilar,
	"BETWEEN":                    OperatorNodeTypeBetween,
	"CASE":                       OperatorNodeTypeCase,
	"NULLIF":                     OperatorNodeTypeNullif,
	"COALESCE":                   OperatorNodeTypeCoalesce,
	"DECODE":                     OperatorNodeTypeDecode,
	"NVL":                        OperatorNodeTypeNvl,
	"GREATEST":                   OperatorNodeTypeGreatest,
	"LEAST":                      OperatorNodeTypeLeast,
	"TIMESTAMP_ADD":              OperatorNodeTypeTimestampAdd,
	"TIMESTAMP_DIFF":             OperatorNodeTypeTimestampDiff,
	"NOT":                        OperatorNodeTypeNot,
	"PLUS_PREFIX":                OperatorNodeTypePlusPrefix,
	"MINUS_PREFIX":               OperatorNodeTypeMinusPrefix,
	"EXISTS":                     OperatorNodeTypeExists,
	"SOME":                       OperatorNodeTypeSome,
	"ALL":                        OperatorNodeTypeAll,
	"VALUES":                     OperatorNodeTypeValues,
	"EXPLICIT_TABLE":             OperatorNodeTypeExplicitTable,
	"SCALAR_QUERY":               OperatorNodeTypeScalarQuery,
	"PROCEDURE_CALL":             OperatorNodeTypeProcedureCall,
	"NEW_SPECIFICATION":          OperatorNodeTypeNewSpecification,
	"FINAL":                      OperatorNodeTypeFinal,
	"RUNNING":                    OperatorNodeTypeRunning,
	"PREV":                       OperatorNodeTypePrev,
	"NEXT":                       OperatorNodeTypeNext,
	"FIRST":                      OperatorNodeTypeFirst,
	"LAST":                       OperatorNodeTypeLast,
	"CLASSIFIER":                 OperatorNodeTypeClassifier,
	"MATCH_NUMBER":               OperatorNodeTypeMatchNumber,
	"SKIP_TO_FIRST":              OperatorNodeTypeSkipToFirst,
	"SKIP_TO_LAST":               OperatorNodeTypeSkipToLast,
	"DESCENDING":                 OperatorNodeTypeDescending,
	"NULLS_FIRST":                OperatorNodeTypeNullsFirst,
	"NULLS_LAST":                 OperatorNodeTypeNullsLast,
	"IS_TRUE":                    OperatorNodeTypeIsTrue,
	"IS_FALSE":                   OperatorNodeTypeIsFalse,
	"IS_NOT_TRUE":                OperatorNodeTypeIsNotTrue,
	"IS_NOT_FALSE":               OperatorNodeTypeIsNotFalse,
	"IS_UNKNOWN":                 OperatorNodeTypeIsUnknown,
	"IS_NULL":                    OperatorNodeTypeIsNull,
	"IS_NOT_NULL":                OperatorNodeTypeIsNotNull,
	"PRECEDING":                  OperatorNodeTypePreceding,
	"FOLLOWING":                  OperatorNodeTypeFollowing,
	"FIELD_ACCESS":               OperatorNodeTypeFieldAccess,
	"INPUT_REF":                  OperatorNodeTypeInputRef,
	"TABLE_INPUT_REF":            OperatorNodeTypeTableInputRef,
	"PATTERN_INPUT_REF":          OperatorNodeTypePatternInputRef,
	"LOCAL_REF":                  OperatorNodeTypeLocalRef,
	"CORREL_VARIABLE":            OperatorNodeTypeCorrelVariable,
	"PATTERN_QUANTIFIER":         OperatorNodeTypePatternQuantifier,
	"ROW":                        OperatorNodeTypeRow,
	"COLUMN_LIST":                OperatorNodeTypeColumnList,
	"CAST":                       OperatorNodeTypeCast,
	"NEXT_VALUE":                 OperatorNodeTypeNextValue,
	"CURRENT_VALUE":              OperatorNodeTypeCurrentValue,
	"FLOOR":                      OperatorNodeTypeFloor,
	"CEIL":                       OperatorNodeTypeCeil,
	"TRIM":                       OperatorNodeTypeTrim,
	"LTRIM":                      OperatorNodeTypeLtrim,
	"RTRIM":                      OperatorNodeTypeRtrim,
	"EXTRACT":                    OperatorNodeTypeExtract,
	"JDBC_FN":                    OperatorNodeTypeJdbcFn,
	"MULTISET_VALUE_CONSTRUCTOR": OperatorNodeTypeMultisetValueConstructor,
	"MULTISET_QUERY_CONSTRUCTOR": OperatorNodeTypeMultisetQueryConstructor,
	"UNNEST":                     OperatorNodeTypeUnnest,
	"LATERAL":                    OperatorNodeTypeLateral,
	"COLLECTION_TABLE":           OperatorNodeTypeCollectionTable,
	"ARRAY_VALUE_CONSTRUCTOR":    OperatorNodeTypeArrayValueConstructor,
	"ARRAY_QUERY_CONSTRUCTOR":    OperatorNodeTypeArrayQueryConstructor,
	"MAP_VALUE_CONSTRUCTOR":      OperatorNodeTypeMapValueConstructor,
	"MAP_QUERY_CONSTRUCTOR":      OperatorNodeTypeMapQueryConstructor,
	"CURSOR":                     OperatorNodeTypeCursor,
	"LITERAL_CHAIN":              OperatorNodeTypeLiteralChain,
	"ESCAPE":                     OperatorNodeTypeEscape,
	"REINTERPRET":                OperatorNodeTypeReinterpret,
	"EXTEND":                     OperatorNodeTypeExtend,
	"CUBE":                       OperatorNodeTypeCube,
	"ROLLUP":                     OperatorNodeTypeRollup,
	"GROUPING_SETS":              OperatorNodeTypeGroupingSets,
	"GROUPING":                   OperatorNodeTypeGrouping,
	"GROUPING_ID":                OperatorNodeTypeGroupingId,
	"GROUP_ID":                   OperatorNodeTypeGroupId,
	"PATTERN_PERMUTE":            OperatorNodeTypePatternPermute,
	"PATTERN_EXCLUDED":           OperatorNodeTypePatternExcluded,
	"COUNT":                      OperatorNodeTypeCount,
	"SUM":                        OperatorNodeTypeSum,
	"SUM0":                       OperatorNodeTypeSum0,
	"MIN":                        OperatorNodeTypeMin,
	"MAX":                        OperatorNodeTypeMax,
	"LEAD":                       OperatorNodeTypeLead,
	"LAG":                        OperatorNodeTypeLag,
	"FIRST_VALUE":                OperatorNodeTypeFirstValue,
	"LAST_VALUE":                 OperatorNodeTypeLastValue,
	"COVAR_POP":                  OperatorNodeTypeCovarPop,
	"COVAR_SAMP":                 OperatorNodeTypeCovarSamp,
	"REGR_SXX":                   OperatorNodeTypeRegrSxx,
	"REGR_SYY":                   OperatorNodeTypeRegrSyy,
	"AVG":                        OperatorNodeTypeAvg,
	"STDDEV_POP":                 OperatorNodeTypeStddevPop,
	"STDDEV_SAMP":                OperatorNodeTypeStddevSamp,
	"VAR_POP":                    OperatorNodeTypeVarPop,
	"VAR_SAMP":                   OperatorNodeTypeVarSamp,
	"NTILE":                      OperatorNodeTypeNtile,
	"COLLECT":                    OperatorNodeTypeCollect,
	"FUSION":                     OperatorNodeTypeFusion,
	"SINGLE_VALUE":               OperatorNodeTypeSingleValue,
	"ROW_NUMBER":                 OperatorNodeTypeRowNumber,
	"RANK":                       OperatorNodeTypeRank,
	"PERCENT_RANK":               OperatorNodeTypePercentRank,
	"DENSE_RANK":                 OperatorNodeTypeDenseRank,
	"CUME_DIST":                  OperatorNodeTypeCumeDist,
	"TUMBLE":                     OperatorNodeTypeTumble,
	"TUMBLE_START":               OperatorNodeTypeTumbleStart,
	"TUMBLE_END":                 OperatorNodeTypeTumbleEnd,
	"HOP":                        OperatorNodeTypeHop,
	"HOP_START":                  OperatorNodeTypeHopStart,
	"HOP_END":                    OperatorNodeTypeHopEnd,
	"SESSION":                    OperatorNodeTypeSession,
	"SESSION_START":              OperatorNodeTypeSessionStart,
	"SESSION_END":                OperatorNodeTypeSessionEnd,
	"COLUMN_DECL":                OperatorNodeTypeColumnDecl,
	"CHECK":                      OperatorNodeTypeCheck,
	"UNIQUE":                     OperatorNodeTypeUnique,
	"PRIMARY_KEY":                OperatorNodeTypePrimaryKey,
	"FOREIGN_KEY":                OperatorNodeTypeForeignKey,
	"COMMIT":                     OperatorNodeTypeCommit,
	"ROLLBACK":                   OperatorNodeTypeRollback,
	"ALTER_SESSION":              OperatorNodeTypeAlterSession,
	"CREATE_SCHEMA":              OperatorNodeTypeCreateSchema,
	"CREATE_FOREIGN_SCHEMA":      OperatorNodeTypeCreateForeignSchema,
	"DROP_SCHEMA":                OperatorNodeTypeDropSchema,
	"CREATE_TABLE":               OperatorNodeTypeCreateTable,
	"ALTER_TABLE":                OperatorNodeTypeAlterTable,
	"DROP_TABLE":                 OperatorNodeTypeDropTable,
	"CREATE_VIEW":                OperatorNodeTypeCreateView,
	"ALTER_VIEW":                 OperatorNodeTypeAlterView,
	"DROP_VIEW":                  OperatorNodeTypeDropView,
	"CREATE_MATERIALIZED_VIEW":   OperatorNodeTypeCreateMaterializedView,
	"ALTER_MATERIALIZED_VIEW":    OperatorNodeTypeAlterMaterializedView,
	"DROP_MATERIALIZED_VIEW":     OperatorNodeTypeDropMaterializedView,
	"CREATE_SEQUENCE":            OperatorNodeTypeCreateSequence,
	"ALTER_SEQUENCE":             OperatorNodeTypeAlterSequence,
	"DROP_SEQUENCE":              OperatorNodeTypeDropSequence,
	"CREATE_INDEX":               OperatorNodeTypeCreateIndex,
	"ALTER_INDEX":                OperatorNodeTypeAlterIndex,
	"DROP_INDEX":                 OperatorNodeTypeDropIndex,
	"OTHER_DDL":                  OperatorNodeTypeOtherDdl,
}

// GetOperatorNodeTypeEnumValues Enumerates the set of values for OperatorNodeTypeEnum
func GetOperatorNodeTypeEnumValues() []OperatorNodeTypeEnum {
	values := make([]OperatorNodeTypeEnum, 0)
	for _, v := range mappingOperatorNodeType {
		values = append(values, v)
	}
	return values
}
