// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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

// RootNode auto generated description
type RootNode struct {

	// auto generated description
	ModelType RootNodeModelTypeEnum `mandatory:"true" json:"modelType"`

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// type
	Type RootNodeTypeEnum `mandatory:"false" json:"type,omitempty"`

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

	// exprString
	ExprString *string `mandatory:"false" json:"exprString"`
}

func (m RootNode) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *RootNode) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key              *string               `json:"key"`
		ModelVersion     *string               `json:"modelVersion"`
		ParentRef        *ParentReference      `json:"parentRef"`
		Type             RootNodeTypeEnum      `json:"type"`
		ParsedString     *string               `json:"parsedString"`
		StartPosition    *int                  `json:"startPosition"`
		EndPosition      *int                  `json:"endPosition"`
		OperatorName     *string               `json:"operatorName"`
		IsBinaryOperator *bool                 `json:"isBinaryOperator"`
		ObjectStatus     *int                  `json:"objectStatus"`
		Operands         []expressionnode      `json:"operands"`
		ExprString       *string               `json:"exprString"`
		ModelType        RootNodeModelTypeEnum `json:"modelType"`
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

	m.ExprString = model.ExprString

	m.ModelType = model.ModelType
	return
}

// RootNodeModelTypeEnum Enum with underlying type: string
type RootNodeModelTypeEnum string

// Set of constants representing the allowable values for RootNodeModelTypeEnum
const (
	RootNodeModelTypeReferenceNode  RootNodeModelTypeEnum = "REFERENCE_NODE"
	RootNodeModelTypeOperatorNode   RootNodeModelTypeEnum = "OPERATOR_NODE"
	RootNodeModelTypeRootNode       RootNodeModelTypeEnum = "ROOT_NODE"
	RootNodeModelTypeValueNode      RootNodeModelTypeEnum = "VALUE_NODE"
	RootNodeModelTypeIdentifierNode RootNodeModelTypeEnum = "IDENTIFIER_NODE"
	RootNodeModelTypeParameterNode  RootNodeModelTypeEnum = "PARAMETER_NODE"
	RootNodeModelTypeFunctionNode   RootNodeModelTypeEnum = "FUNCTION_NODE"
)

var mappingRootNodeModelType = map[string]RootNodeModelTypeEnum{
	"REFERENCE_NODE":  RootNodeModelTypeReferenceNode,
	"OPERATOR_NODE":   RootNodeModelTypeOperatorNode,
	"ROOT_NODE":       RootNodeModelTypeRootNode,
	"VALUE_NODE":      RootNodeModelTypeValueNode,
	"IDENTIFIER_NODE": RootNodeModelTypeIdentifierNode,
	"PARAMETER_NODE":  RootNodeModelTypeParameterNode,
	"FUNCTION_NODE":   RootNodeModelTypeFunctionNode,
}

// GetRootNodeModelTypeEnumValues Enumerates the set of values for RootNodeModelTypeEnum
func GetRootNodeModelTypeEnumValues() []RootNodeModelTypeEnum {
	values := make([]RootNodeModelTypeEnum, 0)
	for _, v := range mappingRootNodeModelType {
		values = append(values, v)
	}
	return values
}

// RootNodeTypeEnum Enum with underlying type: string
type RootNodeTypeEnum string

// Set of constants representing the allowable values for RootNodeTypeEnum
const (
	RootNodeTypeOther                    RootNodeTypeEnum = "OTHER"
	RootNodeTypeSelect                   RootNodeTypeEnum = "SELECT"
	RootNodeTypeJoin                     RootNodeTypeEnum = "JOIN"
	RootNodeTypeIdentifier               RootNodeTypeEnum = "IDENTIFIER"
	RootNodeTypeLiteral                  RootNodeTypeEnum = "LITERAL"
	RootNodeTypeOtherFunction            RootNodeTypeEnum = "OTHER_FUNCTION"
	RootNodeTypeExplain                  RootNodeTypeEnum = "EXPLAIN"
	RootNodeTypeDescribeSchema           RootNodeTypeEnum = "DESCRIBE_SCHEMA"
	RootNodeTypeDescribeTable            RootNodeTypeEnum = "DESCRIBE_TABLE"
	RootNodeTypeInsert                   RootNodeTypeEnum = "INSERT"
	RootNodeTypeDelete                   RootNodeTypeEnum = "DELETE"
	RootNodeTypeUpdate                   RootNodeTypeEnum = "UPDATE"
	RootNodeTypeSetOption                RootNodeTypeEnum = "SET_OPTION"
	RootNodeTypeDynamicParam             RootNodeTypeEnum = "DYNAMIC_PARAM"
	RootNodeTypeOrderBy                  RootNodeTypeEnum = "ORDER_BY"
	RootNodeTypeWith                     RootNodeTypeEnum = "WITH"
	RootNodeTypeWithItem                 RootNodeTypeEnum = "WITH_ITEM"
	RootNodeTypeUnion                    RootNodeTypeEnum = "UNION"
	RootNodeTypeExcept                   RootNodeTypeEnum = "EXCEPT"
	RootNodeTypeIntersect                RootNodeTypeEnum = "INTERSECT"
	RootNodeTypeAs                       RootNodeTypeEnum = "AS"
	RootNodeTypeArgumentAssignment       RootNodeTypeEnum = "ARGUMENT_ASSIGNMENT"
	RootNodeTypeDefault                  RootNodeTypeEnum = "DEFAULT"
	RootNodeTypeOver                     RootNodeTypeEnum = "OVER"
	RootNodeTypeFilter                   RootNodeTypeEnum = "FILTER"
	RootNodeTypeWindow                   RootNodeTypeEnum = "WINDOW"
	RootNodeTypeMerge                    RootNodeTypeEnum = "MERGE"
	RootNodeTypeTablesample              RootNodeTypeEnum = "TABLESAMPLE"
	RootNodeTypeMatchRecognize           RootNodeTypeEnum = "MATCH_RECOGNIZE"
	RootNodeTypeTimes                    RootNodeTypeEnum = "TIMES"
	RootNodeTypeDivide                   RootNodeTypeEnum = "DIVIDE"
	RootNodeTypeMod                      RootNodeTypeEnum = "MOD"
	RootNodeTypePlus                     RootNodeTypeEnum = "PLUS"
	RootNodeTypeMinus                    RootNodeTypeEnum = "MINUS"
	RootNodeTypePatternAlter             RootNodeTypeEnum = "PATTERN_ALTER"
	RootNodeTypePatternConcat            RootNodeTypeEnum = "PATTERN_CONCAT"
	RootNodeTypeIn                       RootNodeTypeEnum = "IN"
	RootNodeTypeNotIn                    RootNodeTypeEnum = "NOT_IN"
	RootNodeTypeLessThan                 RootNodeTypeEnum = "LESS_THAN"
	RootNodeTypeGreaterThan              RootNodeTypeEnum = "GREATER_THAN"
	RootNodeTypeLessThanOrEqual          RootNodeTypeEnum = "LESS_THAN_OR_EQUAL"
	RootNodeTypeGreaterThanOrEqual       RootNodeTypeEnum = "GREATER_THAN_OR_EQUAL"
	RootNodeTypeEquals                   RootNodeTypeEnum = "EQUALS"
	RootNodeTypeNotEquals                RootNodeTypeEnum = "NOT_EQUALS"
	RootNodeTypeIsDistinctFrom           RootNodeTypeEnum = "IS_DISTINCT_FROM"
	RootNodeTypeIsNotDistinctFrom        RootNodeTypeEnum = "IS_NOT_DISTINCT_FROM"
	RootNodeTypeOr                       RootNodeTypeEnum = "OR"
	RootNodeTypeAnd                      RootNodeTypeEnum = "AND"
	RootNodeTypeDot                      RootNodeTypeEnum = "DOT"
	RootNodeTypeOverlaps                 RootNodeTypeEnum = "OVERLAPS"
	RootNodeTypeContains                 RootNodeTypeEnum = "CONTAINS"
	RootNodeTypePrecedes                 RootNodeTypeEnum = "PRECEDES"
	RootNodeTypeImmediatelyPrecedes      RootNodeTypeEnum = "IMMEDIATELY_PRECEDES"
	RootNodeTypeSucceeds                 RootNodeTypeEnum = "SUCCEEDS"
	RootNodeTypeImmediatelySucceeds      RootNodeTypeEnum = "IMMEDIATELY_SUCCEEDS"
	RootNodeTypePeriodEquals             RootNodeTypeEnum = "PERIOD_EQUALS"
	RootNodeTypeLike                     RootNodeTypeEnum = "LIKE"
	RootNodeTypeSimilar                  RootNodeTypeEnum = "SIMILAR"
	RootNodeTypeBetween                  RootNodeTypeEnum = "BETWEEN"
	RootNodeTypeCase                     RootNodeTypeEnum = "CASE"
	RootNodeTypeNullif                   RootNodeTypeEnum = "NULLIF"
	RootNodeTypeCoalesce                 RootNodeTypeEnum = "COALESCE"
	RootNodeTypeDecode                   RootNodeTypeEnum = "DECODE"
	RootNodeTypeNvl                      RootNodeTypeEnum = "NVL"
	RootNodeTypeGreatest                 RootNodeTypeEnum = "GREATEST"
	RootNodeTypeLeast                    RootNodeTypeEnum = "LEAST"
	RootNodeTypeTimestampAdd             RootNodeTypeEnum = "TIMESTAMP_ADD"
	RootNodeTypeTimestampDiff            RootNodeTypeEnum = "TIMESTAMP_DIFF"
	RootNodeTypeNot                      RootNodeTypeEnum = "NOT"
	RootNodeTypePlusPrefix               RootNodeTypeEnum = "PLUS_PREFIX"
	RootNodeTypeMinusPrefix              RootNodeTypeEnum = "MINUS_PREFIX"
	RootNodeTypeExists                   RootNodeTypeEnum = "EXISTS"
	RootNodeTypeSome                     RootNodeTypeEnum = "SOME"
	RootNodeTypeAll                      RootNodeTypeEnum = "ALL"
	RootNodeTypeValues                   RootNodeTypeEnum = "VALUES"
	RootNodeTypeExplicitTable            RootNodeTypeEnum = "EXPLICIT_TABLE"
	RootNodeTypeScalarQuery              RootNodeTypeEnum = "SCALAR_QUERY"
	RootNodeTypeProcedureCall            RootNodeTypeEnum = "PROCEDURE_CALL"
	RootNodeTypeNewSpecification         RootNodeTypeEnum = "NEW_SPECIFICATION"
	RootNodeTypeFinal                    RootNodeTypeEnum = "FINAL"
	RootNodeTypeRunning                  RootNodeTypeEnum = "RUNNING"
	RootNodeTypePrev                     RootNodeTypeEnum = "PREV"
	RootNodeTypeNext                     RootNodeTypeEnum = "NEXT"
	RootNodeTypeFirst                    RootNodeTypeEnum = "FIRST"
	RootNodeTypeLast                     RootNodeTypeEnum = "LAST"
	RootNodeTypeClassifier               RootNodeTypeEnum = "CLASSIFIER"
	RootNodeTypeMatchNumber              RootNodeTypeEnum = "MATCH_NUMBER"
	RootNodeTypeSkipToFirst              RootNodeTypeEnum = "SKIP_TO_FIRST"
	RootNodeTypeSkipToLast               RootNodeTypeEnum = "SKIP_TO_LAST"
	RootNodeTypeDescending               RootNodeTypeEnum = "DESCENDING"
	RootNodeTypeNullsFirst               RootNodeTypeEnum = "NULLS_FIRST"
	RootNodeTypeNullsLast                RootNodeTypeEnum = "NULLS_LAST"
	RootNodeTypeIsTrue                   RootNodeTypeEnum = "IS_TRUE"
	RootNodeTypeIsFalse                  RootNodeTypeEnum = "IS_FALSE"
	RootNodeTypeIsNotTrue                RootNodeTypeEnum = "IS_NOT_TRUE"
	RootNodeTypeIsNotFalse               RootNodeTypeEnum = "IS_NOT_FALSE"
	RootNodeTypeIsUnknown                RootNodeTypeEnum = "IS_UNKNOWN"
	RootNodeTypeIsNull                   RootNodeTypeEnum = "IS_NULL"
	RootNodeTypeIsNotNull                RootNodeTypeEnum = "IS_NOT_NULL"
	RootNodeTypePreceding                RootNodeTypeEnum = "PRECEDING"
	RootNodeTypeFollowing                RootNodeTypeEnum = "FOLLOWING"
	RootNodeTypeFieldAccess              RootNodeTypeEnum = "FIELD_ACCESS"
	RootNodeTypeInputRef                 RootNodeTypeEnum = "INPUT_REF"
	RootNodeTypeTableInputRef            RootNodeTypeEnum = "TABLE_INPUT_REF"
	RootNodeTypePatternInputRef          RootNodeTypeEnum = "PATTERN_INPUT_REF"
	RootNodeTypeLocalRef                 RootNodeTypeEnum = "LOCAL_REF"
	RootNodeTypeCorrelVariable           RootNodeTypeEnum = "CORREL_VARIABLE"
	RootNodeTypePatternQuantifier        RootNodeTypeEnum = "PATTERN_QUANTIFIER"
	RootNodeTypeRow                      RootNodeTypeEnum = "ROW"
	RootNodeTypeColumnList               RootNodeTypeEnum = "COLUMN_LIST"
	RootNodeTypeCast                     RootNodeTypeEnum = "CAST"
	RootNodeTypeNextValue                RootNodeTypeEnum = "NEXT_VALUE"
	RootNodeTypeCurrentValue             RootNodeTypeEnum = "CURRENT_VALUE"
	RootNodeTypeFloor                    RootNodeTypeEnum = "FLOOR"
	RootNodeTypeCeil                     RootNodeTypeEnum = "CEIL"
	RootNodeTypeTrim                     RootNodeTypeEnum = "TRIM"
	RootNodeTypeLtrim                    RootNodeTypeEnum = "LTRIM"
	RootNodeTypeRtrim                    RootNodeTypeEnum = "RTRIM"
	RootNodeTypeExtract                  RootNodeTypeEnum = "EXTRACT"
	RootNodeTypeJdbcFn                   RootNodeTypeEnum = "JDBC_FN"
	RootNodeTypeMultisetValueConstructor RootNodeTypeEnum = "MULTISET_VALUE_CONSTRUCTOR"
	RootNodeTypeMultisetQueryConstructor RootNodeTypeEnum = "MULTISET_QUERY_CONSTRUCTOR"
	RootNodeTypeUnnest                   RootNodeTypeEnum = "UNNEST"
	RootNodeTypeLateral                  RootNodeTypeEnum = "LATERAL"
	RootNodeTypeCollectionTable          RootNodeTypeEnum = "COLLECTION_TABLE"
	RootNodeTypeArrayValueConstructor    RootNodeTypeEnum = "ARRAY_VALUE_CONSTRUCTOR"
	RootNodeTypeArrayQueryConstructor    RootNodeTypeEnum = "ARRAY_QUERY_CONSTRUCTOR"
	RootNodeTypeMapValueConstructor      RootNodeTypeEnum = "MAP_VALUE_CONSTRUCTOR"
	RootNodeTypeMapQueryConstructor      RootNodeTypeEnum = "MAP_QUERY_CONSTRUCTOR"
	RootNodeTypeCursor                   RootNodeTypeEnum = "CURSOR"
	RootNodeTypeLiteralChain             RootNodeTypeEnum = "LITERAL_CHAIN"
	RootNodeTypeEscape                   RootNodeTypeEnum = "ESCAPE"
	RootNodeTypeReinterpret              RootNodeTypeEnum = "REINTERPRET"
	RootNodeTypeExtend                   RootNodeTypeEnum = "EXTEND"
	RootNodeTypeCube                     RootNodeTypeEnum = "CUBE"
	RootNodeTypeRollup                   RootNodeTypeEnum = "ROLLUP"
	RootNodeTypeGroupingSets             RootNodeTypeEnum = "GROUPING_SETS"
	RootNodeTypeGrouping                 RootNodeTypeEnum = "GROUPING"
	RootNodeTypeGroupingId               RootNodeTypeEnum = "GROUPING_ID"
	RootNodeTypeGroupId                  RootNodeTypeEnum = "GROUP_ID"
	RootNodeTypePatternPermute           RootNodeTypeEnum = "PATTERN_PERMUTE"
	RootNodeTypePatternExcluded          RootNodeTypeEnum = "PATTERN_EXCLUDED"
	RootNodeTypeCount                    RootNodeTypeEnum = "COUNT"
	RootNodeTypeSum                      RootNodeTypeEnum = "SUM"
	RootNodeTypeSum0                     RootNodeTypeEnum = "SUM0"
	RootNodeTypeMin                      RootNodeTypeEnum = "MIN"
	RootNodeTypeMax                      RootNodeTypeEnum = "MAX"
	RootNodeTypeLead                     RootNodeTypeEnum = "LEAD"
	RootNodeTypeLag                      RootNodeTypeEnum = "LAG"
	RootNodeTypeFirstValue               RootNodeTypeEnum = "FIRST_VALUE"
	RootNodeTypeLastValue                RootNodeTypeEnum = "LAST_VALUE"
	RootNodeTypeCovarPop                 RootNodeTypeEnum = "COVAR_POP"
	RootNodeTypeCovarSamp                RootNodeTypeEnum = "COVAR_SAMP"
	RootNodeTypeRegrSxx                  RootNodeTypeEnum = "REGR_SXX"
	RootNodeTypeRegrSyy                  RootNodeTypeEnum = "REGR_SYY"
	RootNodeTypeAvg                      RootNodeTypeEnum = "AVG"
	RootNodeTypeStddevPop                RootNodeTypeEnum = "STDDEV_POP"
	RootNodeTypeStddevSamp               RootNodeTypeEnum = "STDDEV_SAMP"
	RootNodeTypeVarPop                   RootNodeTypeEnum = "VAR_POP"
	RootNodeTypeVarSamp                  RootNodeTypeEnum = "VAR_SAMP"
	RootNodeTypeNtile                    RootNodeTypeEnum = "NTILE"
	RootNodeTypeCollect                  RootNodeTypeEnum = "COLLECT"
	RootNodeTypeFusion                   RootNodeTypeEnum = "FUSION"
	RootNodeTypeSingleValue              RootNodeTypeEnum = "SINGLE_VALUE"
	RootNodeTypeRowNumber                RootNodeTypeEnum = "ROW_NUMBER"
	RootNodeTypeRank                     RootNodeTypeEnum = "RANK"
	RootNodeTypePercentRank              RootNodeTypeEnum = "PERCENT_RANK"
	RootNodeTypeDenseRank                RootNodeTypeEnum = "DENSE_RANK"
	RootNodeTypeCumeDist                 RootNodeTypeEnum = "CUME_DIST"
	RootNodeTypeTumble                   RootNodeTypeEnum = "TUMBLE"
	RootNodeTypeTumbleStart              RootNodeTypeEnum = "TUMBLE_START"
	RootNodeTypeTumbleEnd                RootNodeTypeEnum = "TUMBLE_END"
	RootNodeTypeHop                      RootNodeTypeEnum = "HOP"
	RootNodeTypeHopStart                 RootNodeTypeEnum = "HOP_START"
	RootNodeTypeHopEnd                   RootNodeTypeEnum = "HOP_END"
	RootNodeTypeSession                  RootNodeTypeEnum = "SESSION"
	RootNodeTypeSessionStart             RootNodeTypeEnum = "SESSION_START"
	RootNodeTypeSessionEnd               RootNodeTypeEnum = "SESSION_END"
	RootNodeTypeColumnDecl               RootNodeTypeEnum = "COLUMN_DECL"
	RootNodeTypeCheck                    RootNodeTypeEnum = "CHECK"
	RootNodeTypeUnique                   RootNodeTypeEnum = "UNIQUE"
	RootNodeTypePrimaryKey               RootNodeTypeEnum = "PRIMARY_KEY"
	RootNodeTypeForeignKey               RootNodeTypeEnum = "FOREIGN_KEY"
	RootNodeTypeCommit                   RootNodeTypeEnum = "COMMIT"
	RootNodeTypeRollback                 RootNodeTypeEnum = "ROLLBACK"
	RootNodeTypeAlterSession             RootNodeTypeEnum = "ALTER_SESSION"
	RootNodeTypeCreateSchema             RootNodeTypeEnum = "CREATE_SCHEMA"
	RootNodeTypeCreateForeignSchema      RootNodeTypeEnum = "CREATE_FOREIGN_SCHEMA"
	RootNodeTypeDropSchema               RootNodeTypeEnum = "DROP_SCHEMA"
	RootNodeTypeCreateTable              RootNodeTypeEnum = "CREATE_TABLE"
	RootNodeTypeAlterTable               RootNodeTypeEnum = "ALTER_TABLE"
	RootNodeTypeDropTable                RootNodeTypeEnum = "DROP_TABLE"
	RootNodeTypeCreateView               RootNodeTypeEnum = "CREATE_VIEW"
	RootNodeTypeAlterView                RootNodeTypeEnum = "ALTER_VIEW"
	RootNodeTypeDropView                 RootNodeTypeEnum = "DROP_VIEW"
	RootNodeTypeCreateMaterializedView   RootNodeTypeEnum = "CREATE_MATERIALIZED_VIEW"
	RootNodeTypeAlterMaterializedView    RootNodeTypeEnum = "ALTER_MATERIALIZED_VIEW"
	RootNodeTypeDropMaterializedView     RootNodeTypeEnum = "DROP_MATERIALIZED_VIEW"
	RootNodeTypeCreateSequence           RootNodeTypeEnum = "CREATE_SEQUENCE"
	RootNodeTypeAlterSequence            RootNodeTypeEnum = "ALTER_SEQUENCE"
	RootNodeTypeDropSequence             RootNodeTypeEnum = "DROP_SEQUENCE"
	RootNodeTypeCreateIndex              RootNodeTypeEnum = "CREATE_INDEX"
	RootNodeTypeAlterIndex               RootNodeTypeEnum = "ALTER_INDEX"
	RootNodeTypeDropIndex                RootNodeTypeEnum = "DROP_INDEX"
	RootNodeTypeOtherDdl                 RootNodeTypeEnum = "OTHER_DDL"
)

var mappingRootNodeType = map[string]RootNodeTypeEnum{
	"OTHER":           RootNodeTypeOther,
	"SELECT":          RootNodeTypeSelect,
	"JOIN":            RootNodeTypeJoin,
	"IDENTIFIER":      RootNodeTypeIdentifier,
	"LITERAL":         RootNodeTypeLiteral,
	"OTHER_FUNCTION":  RootNodeTypeOtherFunction,
	"EXPLAIN":         RootNodeTypeExplain,
	"DESCRIBE_SCHEMA": RootNodeTypeDescribeSchema,
	"DESCRIBE_TABLE":  RootNodeTypeDescribeTable,
	"INSERT":          RootNodeTypeInsert,
	"DELETE":          RootNodeTypeDelete,
	"UPDATE":          RootNodeTypeUpdate,
	"SET_OPTION":      RootNodeTypeSetOption,
	"DYNAMIC_PARAM":   RootNodeTypeDynamicParam,
	"ORDER_BY":        RootNodeTypeOrderBy,
	"WITH":            RootNodeTypeWith,
	"WITH_ITEM":       RootNodeTypeWithItem,
	"UNION":           RootNodeTypeUnion,
	"EXCEPT":          RootNodeTypeExcept,
	"INTERSECT":       RootNodeTypeIntersect,
	"AS":              RootNodeTypeAs,
	"ARGUMENT_ASSIGNMENT":   RootNodeTypeArgumentAssignment,
	"DEFAULT":               RootNodeTypeDefault,
	"OVER":                  RootNodeTypeOver,
	"FILTER":                RootNodeTypeFilter,
	"WINDOW":                RootNodeTypeWindow,
	"MERGE":                 RootNodeTypeMerge,
	"TABLESAMPLE":           RootNodeTypeTablesample,
	"MATCH_RECOGNIZE":       RootNodeTypeMatchRecognize,
	"TIMES":                 RootNodeTypeTimes,
	"DIVIDE":                RootNodeTypeDivide,
	"MOD":                   RootNodeTypeMod,
	"PLUS":                  RootNodeTypePlus,
	"MINUS":                 RootNodeTypeMinus,
	"PATTERN_ALTER":         RootNodeTypePatternAlter,
	"PATTERN_CONCAT":        RootNodeTypePatternConcat,
	"IN":                    RootNodeTypeIn,
	"NOT_IN":                RootNodeTypeNotIn,
	"LESS_THAN":             RootNodeTypeLessThan,
	"GREATER_THAN":          RootNodeTypeGreaterThan,
	"LESS_THAN_OR_EQUAL":    RootNodeTypeLessThanOrEqual,
	"GREATER_THAN_OR_EQUAL": RootNodeTypeGreaterThanOrEqual,
	"EQUALS":                RootNodeTypeEquals,
	"NOT_EQUALS":            RootNodeTypeNotEquals,
	"IS_DISTINCT_FROM":      RootNodeTypeIsDistinctFrom,
	"IS_NOT_DISTINCT_FROM":  RootNodeTypeIsNotDistinctFrom,
	"OR":                         RootNodeTypeOr,
	"AND":                        RootNodeTypeAnd,
	"DOT":                        RootNodeTypeDot,
	"OVERLAPS":                   RootNodeTypeOverlaps,
	"CONTAINS":                   RootNodeTypeContains,
	"PRECEDES":                   RootNodeTypePrecedes,
	"IMMEDIATELY_PRECEDES":       RootNodeTypeImmediatelyPrecedes,
	"SUCCEEDS":                   RootNodeTypeSucceeds,
	"IMMEDIATELY_SUCCEEDS":       RootNodeTypeImmediatelySucceeds,
	"PERIOD_EQUALS":              RootNodeTypePeriodEquals,
	"LIKE":                       RootNodeTypeLike,
	"SIMILAR":                    RootNodeTypeSimilar,
	"BETWEEN":                    RootNodeTypeBetween,
	"CASE":                       RootNodeTypeCase,
	"NULLIF":                     RootNodeTypeNullif,
	"COALESCE":                   RootNodeTypeCoalesce,
	"DECODE":                     RootNodeTypeDecode,
	"NVL":                        RootNodeTypeNvl,
	"GREATEST":                   RootNodeTypeGreatest,
	"LEAST":                      RootNodeTypeLeast,
	"TIMESTAMP_ADD":              RootNodeTypeTimestampAdd,
	"TIMESTAMP_DIFF":             RootNodeTypeTimestampDiff,
	"NOT":                        RootNodeTypeNot,
	"PLUS_PREFIX":                RootNodeTypePlusPrefix,
	"MINUS_PREFIX":               RootNodeTypeMinusPrefix,
	"EXISTS":                     RootNodeTypeExists,
	"SOME":                       RootNodeTypeSome,
	"ALL":                        RootNodeTypeAll,
	"VALUES":                     RootNodeTypeValues,
	"EXPLICIT_TABLE":             RootNodeTypeExplicitTable,
	"SCALAR_QUERY":               RootNodeTypeScalarQuery,
	"PROCEDURE_CALL":             RootNodeTypeProcedureCall,
	"NEW_SPECIFICATION":          RootNodeTypeNewSpecification,
	"FINAL":                      RootNodeTypeFinal,
	"RUNNING":                    RootNodeTypeRunning,
	"PREV":                       RootNodeTypePrev,
	"NEXT":                       RootNodeTypeNext,
	"FIRST":                      RootNodeTypeFirst,
	"LAST":                       RootNodeTypeLast,
	"CLASSIFIER":                 RootNodeTypeClassifier,
	"MATCH_NUMBER":               RootNodeTypeMatchNumber,
	"SKIP_TO_FIRST":              RootNodeTypeSkipToFirst,
	"SKIP_TO_LAST":               RootNodeTypeSkipToLast,
	"DESCENDING":                 RootNodeTypeDescending,
	"NULLS_FIRST":                RootNodeTypeNullsFirst,
	"NULLS_LAST":                 RootNodeTypeNullsLast,
	"IS_TRUE":                    RootNodeTypeIsTrue,
	"IS_FALSE":                   RootNodeTypeIsFalse,
	"IS_NOT_TRUE":                RootNodeTypeIsNotTrue,
	"IS_NOT_FALSE":               RootNodeTypeIsNotFalse,
	"IS_UNKNOWN":                 RootNodeTypeIsUnknown,
	"IS_NULL":                    RootNodeTypeIsNull,
	"IS_NOT_NULL":                RootNodeTypeIsNotNull,
	"PRECEDING":                  RootNodeTypePreceding,
	"FOLLOWING":                  RootNodeTypeFollowing,
	"FIELD_ACCESS":               RootNodeTypeFieldAccess,
	"INPUT_REF":                  RootNodeTypeInputRef,
	"TABLE_INPUT_REF":            RootNodeTypeTableInputRef,
	"PATTERN_INPUT_REF":          RootNodeTypePatternInputRef,
	"LOCAL_REF":                  RootNodeTypeLocalRef,
	"CORREL_VARIABLE":            RootNodeTypeCorrelVariable,
	"PATTERN_QUANTIFIER":         RootNodeTypePatternQuantifier,
	"ROW":                        RootNodeTypeRow,
	"COLUMN_LIST":                RootNodeTypeColumnList,
	"CAST":                       RootNodeTypeCast,
	"NEXT_VALUE":                 RootNodeTypeNextValue,
	"CURRENT_VALUE":              RootNodeTypeCurrentValue,
	"FLOOR":                      RootNodeTypeFloor,
	"CEIL":                       RootNodeTypeCeil,
	"TRIM":                       RootNodeTypeTrim,
	"LTRIM":                      RootNodeTypeLtrim,
	"RTRIM":                      RootNodeTypeRtrim,
	"EXTRACT":                    RootNodeTypeExtract,
	"JDBC_FN":                    RootNodeTypeJdbcFn,
	"MULTISET_VALUE_CONSTRUCTOR": RootNodeTypeMultisetValueConstructor,
	"MULTISET_QUERY_CONSTRUCTOR": RootNodeTypeMultisetQueryConstructor,
	"UNNEST":                     RootNodeTypeUnnest,
	"LATERAL":                    RootNodeTypeLateral,
	"COLLECTION_TABLE":           RootNodeTypeCollectionTable,
	"ARRAY_VALUE_CONSTRUCTOR":    RootNodeTypeArrayValueConstructor,
	"ARRAY_QUERY_CONSTRUCTOR":    RootNodeTypeArrayQueryConstructor,
	"MAP_VALUE_CONSTRUCTOR":      RootNodeTypeMapValueConstructor,
	"MAP_QUERY_CONSTRUCTOR":      RootNodeTypeMapQueryConstructor,
	"CURSOR":                     RootNodeTypeCursor,
	"LITERAL_CHAIN":              RootNodeTypeLiteralChain,
	"ESCAPE":                     RootNodeTypeEscape,
	"REINTERPRET":                RootNodeTypeReinterpret,
	"EXTEND":                     RootNodeTypeExtend,
	"CUBE":                       RootNodeTypeCube,
	"ROLLUP":                     RootNodeTypeRollup,
	"GROUPING_SETS":              RootNodeTypeGroupingSets,
	"GROUPING":                   RootNodeTypeGrouping,
	"GROUPING_ID":                RootNodeTypeGroupingId,
	"GROUP_ID":                   RootNodeTypeGroupId,
	"PATTERN_PERMUTE":            RootNodeTypePatternPermute,
	"PATTERN_EXCLUDED":           RootNodeTypePatternExcluded,
	"COUNT":                      RootNodeTypeCount,
	"SUM":                        RootNodeTypeSum,
	"SUM0":                       RootNodeTypeSum0,
	"MIN":                        RootNodeTypeMin,
	"MAX":                        RootNodeTypeMax,
	"LEAD":                       RootNodeTypeLead,
	"LAG":                        RootNodeTypeLag,
	"FIRST_VALUE":                RootNodeTypeFirstValue,
	"LAST_VALUE":                 RootNodeTypeLastValue,
	"COVAR_POP":                  RootNodeTypeCovarPop,
	"COVAR_SAMP":                 RootNodeTypeCovarSamp,
	"REGR_SXX":                   RootNodeTypeRegrSxx,
	"REGR_SYY":                   RootNodeTypeRegrSyy,
	"AVG":                        RootNodeTypeAvg,
	"STDDEV_POP":                 RootNodeTypeStddevPop,
	"STDDEV_SAMP":                RootNodeTypeStddevSamp,
	"VAR_POP":                    RootNodeTypeVarPop,
	"VAR_SAMP":                   RootNodeTypeVarSamp,
	"NTILE":                      RootNodeTypeNtile,
	"COLLECT":                    RootNodeTypeCollect,
	"FUSION":                     RootNodeTypeFusion,
	"SINGLE_VALUE":               RootNodeTypeSingleValue,
	"ROW_NUMBER":                 RootNodeTypeRowNumber,
	"RANK":                       RootNodeTypeRank,
	"PERCENT_RANK":               RootNodeTypePercentRank,
	"DENSE_RANK":                 RootNodeTypeDenseRank,
	"CUME_DIST":                  RootNodeTypeCumeDist,
	"TUMBLE":                     RootNodeTypeTumble,
	"TUMBLE_START":               RootNodeTypeTumbleStart,
	"TUMBLE_END":                 RootNodeTypeTumbleEnd,
	"HOP":                        RootNodeTypeHop,
	"HOP_START":                  RootNodeTypeHopStart,
	"HOP_END":                    RootNodeTypeHopEnd,
	"SESSION":                    RootNodeTypeSession,
	"SESSION_START":              RootNodeTypeSessionStart,
	"SESSION_END":                RootNodeTypeSessionEnd,
	"COLUMN_DECL":                RootNodeTypeColumnDecl,
	"CHECK":                      RootNodeTypeCheck,
	"UNIQUE":                     RootNodeTypeUnique,
	"PRIMARY_KEY":                RootNodeTypePrimaryKey,
	"FOREIGN_KEY":                RootNodeTypeForeignKey,
	"COMMIT":                     RootNodeTypeCommit,
	"ROLLBACK":                   RootNodeTypeRollback,
	"ALTER_SESSION":              RootNodeTypeAlterSession,
	"CREATE_SCHEMA":              RootNodeTypeCreateSchema,
	"CREATE_FOREIGN_SCHEMA":      RootNodeTypeCreateForeignSchema,
	"DROP_SCHEMA":                RootNodeTypeDropSchema,
	"CREATE_TABLE":               RootNodeTypeCreateTable,
	"ALTER_TABLE":                RootNodeTypeAlterTable,
	"DROP_TABLE":                 RootNodeTypeDropTable,
	"CREATE_VIEW":                RootNodeTypeCreateView,
	"ALTER_VIEW":                 RootNodeTypeAlterView,
	"DROP_VIEW":                  RootNodeTypeDropView,
	"CREATE_MATERIALIZED_VIEW":   RootNodeTypeCreateMaterializedView,
	"ALTER_MATERIALIZED_VIEW":    RootNodeTypeAlterMaterializedView,
	"DROP_MATERIALIZED_VIEW":     RootNodeTypeDropMaterializedView,
	"CREATE_SEQUENCE":            RootNodeTypeCreateSequence,
	"ALTER_SEQUENCE":             RootNodeTypeAlterSequence,
	"DROP_SEQUENCE":              RootNodeTypeDropSequence,
	"CREATE_INDEX":               RootNodeTypeCreateIndex,
	"ALTER_INDEX":                RootNodeTypeAlterIndex,
	"DROP_INDEX":                 RootNodeTypeDropIndex,
	"OTHER_DDL":                  RootNodeTypeOtherDdl,
}

// GetRootNodeTypeEnumValues Enumerates the set of values for RootNodeTypeEnum
func GetRootNodeTypeEnumValues() []RootNodeTypeEnum {
	values := make([]RootNodeTypeEnum, 0)
	for _, v := range mappingRootNodeType {
		values = append(values, v)
	}
	return values
}
