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

// FunctionNode auto generated description
type FunctionNode struct {

	// auto generated description
	ModelType FunctionNodeModelTypeEnum `mandatory:"true" json:"modelType"`

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// type
	Type FunctionNodeTypeEnum `mandatory:"false" json:"type,omitempty"`

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

	FuncLibrary *FunctionLibrary `mandatory:"false" json:"funcLibrary"`

	FunctionSignature *FunctionSignature `mandatory:"false" json:"functionSignature"`
}

func (m FunctionNode) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *FunctionNode) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key               *string                   `json:"key"`
		ModelVersion      *string                   `json:"modelVersion"`
		ParentRef         *ParentReference          `json:"parentRef"`
		Type              FunctionNodeTypeEnum      `json:"type"`
		ParsedString      *string                   `json:"parsedString"`
		StartPosition     *int                      `json:"startPosition"`
		EndPosition       *int                      `json:"endPosition"`
		OperatorName      *string                   `json:"operatorName"`
		IsBinaryOperator  *bool                     `json:"isBinaryOperator"`
		ObjectStatus      *int                      `json:"objectStatus"`
		Operands          []expressionnode          `json:"operands"`
		FuncLibrary       *FunctionLibrary          `json:"funcLibrary"`
		FunctionSignature *FunctionSignature        `json:"functionSignature"`
		ModelType         FunctionNodeModelTypeEnum `json:"modelType"`
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

	m.FuncLibrary = model.FuncLibrary

	m.FunctionSignature = model.FunctionSignature

	m.ModelType = model.ModelType
	return
}

// FunctionNodeModelTypeEnum Enum with underlying type: string
type FunctionNodeModelTypeEnum string

// Set of constants representing the allowable values for FunctionNodeModelTypeEnum
const (
	FunctionNodeModelTypeReferenceNode  FunctionNodeModelTypeEnum = "REFERENCE_NODE"
	FunctionNodeModelTypeOperatorNode   FunctionNodeModelTypeEnum = "OPERATOR_NODE"
	FunctionNodeModelTypeRootNode       FunctionNodeModelTypeEnum = "ROOT_NODE"
	FunctionNodeModelTypeValueNode      FunctionNodeModelTypeEnum = "VALUE_NODE"
	FunctionNodeModelTypeIdentifierNode FunctionNodeModelTypeEnum = "IDENTIFIER_NODE"
	FunctionNodeModelTypeParameterNode  FunctionNodeModelTypeEnum = "PARAMETER_NODE"
	FunctionNodeModelTypeFunctionNode   FunctionNodeModelTypeEnum = "FUNCTION_NODE"
)

var mappingFunctionNodeModelType = map[string]FunctionNodeModelTypeEnum{
	"REFERENCE_NODE":  FunctionNodeModelTypeReferenceNode,
	"OPERATOR_NODE":   FunctionNodeModelTypeOperatorNode,
	"ROOT_NODE":       FunctionNodeModelTypeRootNode,
	"VALUE_NODE":      FunctionNodeModelTypeValueNode,
	"IDENTIFIER_NODE": FunctionNodeModelTypeIdentifierNode,
	"PARAMETER_NODE":  FunctionNodeModelTypeParameterNode,
	"FUNCTION_NODE":   FunctionNodeModelTypeFunctionNode,
}

// GetFunctionNodeModelTypeEnumValues Enumerates the set of values for FunctionNodeModelTypeEnum
func GetFunctionNodeModelTypeEnumValues() []FunctionNodeModelTypeEnum {
	values := make([]FunctionNodeModelTypeEnum, 0)
	for _, v := range mappingFunctionNodeModelType {
		values = append(values, v)
	}
	return values
}

// FunctionNodeTypeEnum Enum with underlying type: string
type FunctionNodeTypeEnum string

// Set of constants representing the allowable values for FunctionNodeTypeEnum
const (
	FunctionNodeTypeOther                    FunctionNodeTypeEnum = "OTHER"
	FunctionNodeTypeSelect                   FunctionNodeTypeEnum = "SELECT"
	FunctionNodeTypeJoin                     FunctionNodeTypeEnum = "JOIN"
	FunctionNodeTypeIdentifier               FunctionNodeTypeEnum = "IDENTIFIER"
	FunctionNodeTypeLiteral                  FunctionNodeTypeEnum = "LITERAL"
	FunctionNodeTypeOtherFunction            FunctionNodeTypeEnum = "OTHER_FUNCTION"
	FunctionNodeTypeExplain                  FunctionNodeTypeEnum = "EXPLAIN"
	FunctionNodeTypeDescribeSchema           FunctionNodeTypeEnum = "DESCRIBE_SCHEMA"
	FunctionNodeTypeDescribeTable            FunctionNodeTypeEnum = "DESCRIBE_TABLE"
	FunctionNodeTypeInsert                   FunctionNodeTypeEnum = "INSERT"
	FunctionNodeTypeDelete                   FunctionNodeTypeEnum = "DELETE"
	FunctionNodeTypeUpdate                   FunctionNodeTypeEnum = "UPDATE"
	FunctionNodeTypeSetOption                FunctionNodeTypeEnum = "SET_OPTION"
	FunctionNodeTypeDynamicParam             FunctionNodeTypeEnum = "DYNAMIC_PARAM"
	FunctionNodeTypeOrderBy                  FunctionNodeTypeEnum = "ORDER_BY"
	FunctionNodeTypeWith                     FunctionNodeTypeEnum = "WITH"
	FunctionNodeTypeWithItem                 FunctionNodeTypeEnum = "WITH_ITEM"
	FunctionNodeTypeUnion                    FunctionNodeTypeEnum = "UNION"
	FunctionNodeTypeExcept                   FunctionNodeTypeEnum = "EXCEPT"
	FunctionNodeTypeIntersect                FunctionNodeTypeEnum = "INTERSECT"
	FunctionNodeTypeAs                       FunctionNodeTypeEnum = "AS"
	FunctionNodeTypeArgumentAssignment       FunctionNodeTypeEnum = "ARGUMENT_ASSIGNMENT"
	FunctionNodeTypeDefault                  FunctionNodeTypeEnum = "DEFAULT"
	FunctionNodeTypeOver                     FunctionNodeTypeEnum = "OVER"
	FunctionNodeTypeFilter                   FunctionNodeTypeEnum = "FILTER"
	FunctionNodeTypeWindow                   FunctionNodeTypeEnum = "WINDOW"
	FunctionNodeTypeMerge                    FunctionNodeTypeEnum = "MERGE"
	FunctionNodeTypeTablesample              FunctionNodeTypeEnum = "TABLESAMPLE"
	FunctionNodeTypeMatchRecognize           FunctionNodeTypeEnum = "MATCH_RECOGNIZE"
	FunctionNodeTypeTimes                    FunctionNodeTypeEnum = "TIMES"
	FunctionNodeTypeDivide                   FunctionNodeTypeEnum = "DIVIDE"
	FunctionNodeTypeMod                      FunctionNodeTypeEnum = "MOD"
	FunctionNodeTypePlus                     FunctionNodeTypeEnum = "PLUS"
	FunctionNodeTypeMinus                    FunctionNodeTypeEnum = "MINUS"
	FunctionNodeTypePatternAlter             FunctionNodeTypeEnum = "PATTERN_ALTER"
	FunctionNodeTypePatternConcat            FunctionNodeTypeEnum = "PATTERN_CONCAT"
	FunctionNodeTypeIn                       FunctionNodeTypeEnum = "IN"
	FunctionNodeTypeNotIn                    FunctionNodeTypeEnum = "NOT_IN"
	FunctionNodeTypeLessThan                 FunctionNodeTypeEnum = "LESS_THAN"
	FunctionNodeTypeGreaterThan              FunctionNodeTypeEnum = "GREATER_THAN"
	FunctionNodeTypeLessThanOrEqual          FunctionNodeTypeEnum = "LESS_THAN_OR_EQUAL"
	FunctionNodeTypeGreaterThanOrEqual       FunctionNodeTypeEnum = "GREATER_THAN_OR_EQUAL"
	FunctionNodeTypeEquals                   FunctionNodeTypeEnum = "EQUALS"
	FunctionNodeTypeNotEquals                FunctionNodeTypeEnum = "NOT_EQUALS"
	FunctionNodeTypeIsDistinctFrom           FunctionNodeTypeEnum = "IS_DISTINCT_FROM"
	FunctionNodeTypeIsNotDistinctFrom        FunctionNodeTypeEnum = "IS_NOT_DISTINCT_FROM"
	FunctionNodeTypeOr                       FunctionNodeTypeEnum = "OR"
	FunctionNodeTypeAnd                      FunctionNodeTypeEnum = "AND"
	FunctionNodeTypeDot                      FunctionNodeTypeEnum = "DOT"
	FunctionNodeTypeOverlaps                 FunctionNodeTypeEnum = "OVERLAPS"
	FunctionNodeTypeContains                 FunctionNodeTypeEnum = "CONTAINS"
	FunctionNodeTypePrecedes                 FunctionNodeTypeEnum = "PRECEDES"
	FunctionNodeTypeImmediatelyPrecedes      FunctionNodeTypeEnum = "IMMEDIATELY_PRECEDES"
	FunctionNodeTypeSucceeds                 FunctionNodeTypeEnum = "SUCCEEDS"
	FunctionNodeTypeImmediatelySucceeds      FunctionNodeTypeEnum = "IMMEDIATELY_SUCCEEDS"
	FunctionNodeTypePeriodEquals             FunctionNodeTypeEnum = "PERIOD_EQUALS"
	FunctionNodeTypeLike                     FunctionNodeTypeEnum = "LIKE"
	FunctionNodeTypeSimilar                  FunctionNodeTypeEnum = "SIMILAR"
	FunctionNodeTypeBetween                  FunctionNodeTypeEnum = "BETWEEN"
	FunctionNodeTypeCase                     FunctionNodeTypeEnum = "CASE"
	FunctionNodeTypeNullif                   FunctionNodeTypeEnum = "NULLIF"
	FunctionNodeTypeCoalesce                 FunctionNodeTypeEnum = "COALESCE"
	FunctionNodeTypeDecode                   FunctionNodeTypeEnum = "DECODE"
	FunctionNodeTypeNvl                      FunctionNodeTypeEnum = "NVL"
	FunctionNodeTypeGreatest                 FunctionNodeTypeEnum = "GREATEST"
	FunctionNodeTypeLeast                    FunctionNodeTypeEnum = "LEAST"
	FunctionNodeTypeTimestampAdd             FunctionNodeTypeEnum = "TIMESTAMP_ADD"
	FunctionNodeTypeTimestampDiff            FunctionNodeTypeEnum = "TIMESTAMP_DIFF"
	FunctionNodeTypeNot                      FunctionNodeTypeEnum = "NOT"
	FunctionNodeTypePlusPrefix               FunctionNodeTypeEnum = "PLUS_PREFIX"
	FunctionNodeTypeMinusPrefix              FunctionNodeTypeEnum = "MINUS_PREFIX"
	FunctionNodeTypeExists                   FunctionNodeTypeEnum = "EXISTS"
	FunctionNodeTypeSome                     FunctionNodeTypeEnum = "SOME"
	FunctionNodeTypeAll                      FunctionNodeTypeEnum = "ALL"
	FunctionNodeTypeValues                   FunctionNodeTypeEnum = "VALUES"
	FunctionNodeTypeExplicitTable            FunctionNodeTypeEnum = "EXPLICIT_TABLE"
	FunctionNodeTypeScalarQuery              FunctionNodeTypeEnum = "SCALAR_QUERY"
	FunctionNodeTypeProcedureCall            FunctionNodeTypeEnum = "PROCEDURE_CALL"
	FunctionNodeTypeNewSpecification         FunctionNodeTypeEnum = "NEW_SPECIFICATION"
	FunctionNodeTypeFinal                    FunctionNodeTypeEnum = "FINAL"
	FunctionNodeTypeRunning                  FunctionNodeTypeEnum = "RUNNING"
	FunctionNodeTypePrev                     FunctionNodeTypeEnum = "PREV"
	FunctionNodeTypeNext                     FunctionNodeTypeEnum = "NEXT"
	FunctionNodeTypeFirst                    FunctionNodeTypeEnum = "FIRST"
	FunctionNodeTypeLast                     FunctionNodeTypeEnum = "LAST"
	FunctionNodeTypeClassifier               FunctionNodeTypeEnum = "CLASSIFIER"
	FunctionNodeTypeMatchNumber              FunctionNodeTypeEnum = "MATCH_NUMBER"
	FunctionNodeTypeSkipToFirst              FunctionNodeTypeEnum = "SKIP_TO_FIRST"
	FunctionNodeTypeSkipToLast               FunctionNodeTypeEnum = "SKIP_TO_LAST"
	FunctionNodeTypeDescending               FunctionNodeTypeEnum = "DESCENDING"
	FunctionNodeTypeNullsFirst               FunctionNodeTypeEnum = "NULLS_FIRST"
	FunctionNodeTypeNullsLast                FunctionNodeTypeEnum = "NULLS_LAST"
	FunctionNodeTypeIsTrue                   FunctionNodeTypeEnum = "IS_TRUE"
	FunctionNodeTypeIsFalse                  FunctionNodeTypeEnum = "IS_FALSE"
	FunctionNodeTypeIsNotTrue                FunctionNodeTypeEnum = "IS_NOT_TRUE"
	FunctionNodeTypeIsNotFalse               FunctionNodeTypeEnum = "IS_NOT_FALSE"
	FunctionNodeTypeIsUnknown                FunctionNodeTypeEnum = "IS_UNKNOWN"
	FunctionNodeTypeIsNull                   FunctionNodeTypeEnum = "IS_NULL"
	FunctionNodeTypeIsNotNull                FunctionNodeTypeEnum = "IS_NOT_NULL"
	FunctionNodeTypePreceding                FunctionNodeTypeEnum = "PRECEDING"
	FunctionNodeTypeFollowing                FunctionNodeTypeEnum = "FOLLOWING"
	FunctionNodeTypeFieldAccess              FunctionNodeTypeEnum = "FIELD_ACCESS"
	FunctionNodeTypeInputRef                 FunctionNodeTypeEnum = "INPUT_REF"
	FunctionNodeTypeTableInputRef            FunctionNodeTypeEnum = "TABLE_INPUT_REF"
	FunctionNodeTypePatternInputRef          FunctionNodeTypeEnum = "PATTERN_INPUT_REF"
	FunctionNodeTypeLocalRef                 FunctionNodeTypeEnum = "LOCAL_REF"
	FunctionNodeTypeCorrelVariable           FunctionNodeTypeEnum = "CORREL_VARIABLE"
	FunctionNodeTypePatternQuantifier        FunctionNodeTypeEnum = "PATTERN_QUANTIFIER"
	FunctionNodeTypeRow                      FunctionNodeTypeEnum = "ROW"
	FunctionNodeTypeColumnList               FunctionNodeTypeEnum = "COLUMN_LIST"
	FunctionNodeTypeCast                     FunctionNodeTypeEnum = "CAST"
	FunctionNodeTypeNextValue                FunctionNodeTypeEnum = "NEXT_VALUE"
	FunctionNodeTypeCurrentValue             FunctionNodeTypeEnum = "CURRENT_VALUE"
	FunctionNodeTypeFloor                    FunctionNodeTypeEnum = "FLOOR"
	FunctionNodeTypeCeil                     FunctionNodeTypeEnum = "CEIL"
	FunctionNodeTypeTrim                     FunctionNodeTypeEnum = "TRIM"
	FunctionNodeTypeLtrim                    FunctionNodeTypeEnum = "LTRIM"
	FunctionNodeTypeRtrim                    FunctionNodeTypeEnum = "RTRIM"
	FunctionNodeTypeExtract                  FunctionNodeTypeEnum = "EXTRACT"
	FunctionNodeTypeJdbcFn                   FunctionNodeTypeEnum = "JDBC_FN"
	FunctionNodeTypeMultisetValueConstructor FunctionNodeTypeEnum = "MULTISET_VALUE_CONSTRUCTOR"
	FunctionNodeTypeMultisetQueryConstructor FunctionNodeTypeEnum = "MULTISET_QUERY_CONSTRUCTOR"
	FunctionNodeTypeUnnest                   FunctionNodeTypeEnum = "UNNEST"
	FunctionNodeTypeLateral                  FunctionNodeTypeEnum = "LATERAL"
	FunctionNodeTypeCollectionTable          FunctionNodeTypeEnum = "COLLECTION_TABLE"
	FunctionNodeTypeArrayValueConstructor    FunctionNodeTypeEnum = "ARRAY_VALUE_CONSTRUCTOR"
	FunctionNodeTypeArrayQueryConstructor    FunctionNodeTypeEnum = "ARRAY_QUERY_CONSTRUCTOR"
	FunctionNodeTypeMapValueConstructor      FunctionNodeTypeEnum = "MAP_VALUE_CONSTRUCTOR"
	FunctionNodeTypeMapQueryConstructor      FunctionNodeTypeEnum = "MAP_QUERY_CONSTRUCTOR"
	FunctionNodeTypeCursor                   FunctionNodeTypeEnum = "CURSOR"
	FunctionNodeTypeLiteralChain             FunctionNodeTypeEnum = "LITERAL_CHAIN"
	FunctionNodeTypeEscape                   FunctionNodeTypeEnum = "ESCAPE"
	FunctionNodeTypeReinterpret              FunctionNodeTypeEnum = "REINTERPRET"
	FunctionNodeTypeExtend                   FunctionNodeTypeEnum = "EXTEND"
	FunctionNodeTypeCube                     FunctionNodeTypeEnum = "CUBE"
	FunctionNodeTypeRollup                   FunctionNodeTypeEnum = "ROLLUP"
	FunctionNodeTypeGroupingSets             FunctionNodeTypeEnum = "GROUPING_SETS"
	FunctionNodeTypeGrouping                 FunctionNodeTypeEnum = "GROUPING"
	FunctionNodeTypeGroupingId               FunctionNodeTypeEnum = "GROUPING_ID"
	FunctionNodeTypeGroupId                  FunctionNodeTypeEnum = "GROUP_ID"
	FunctionNodeTypePatternPermute           FunctionNodeTypeEnum = "PATTERN_PERMUTE"
	FunctionNodeTypePatternExcluded          FunctionNodeTypeEnum = "PATTERN_EXCLUDED"
	FunctionNodeTypeCount                    FunctionNodeTypeEnum = "COUNT"
	FunctionNodeTypeSum                      FunctionNodeTypeEnum = "SUM"
	FunctionNodeTypeSum0                     FunctionNodeTypeEnum = "SUM0"
	FunctionNodeTypeMin                      FunctionNodeTypeEnum = "MIN"
	FunctionNodeTypeMax                      FunctionNodeTypeEnum = "MAX"
	FunctionNodeTypeLead                     FunctionNodeTypeEnum = "LEAD"
	FunctionNodeTypeLag                      FunctionNodeTypeEnum = "LAG"
	FunctionNodeTypeFirstValue               FunctionNodeTypeEnum = "FIRST_VALUE"
	FunctionNodeTypeLastValue                FunctionNodeTypeEnum = "LAST_VALUE"
	FunctionNodeTypeCovarPop                 FunctionNodeTypeEnum = "COVAR_POP"
	FunctionNodeTypeCovarSamp                FunctionNodeTypeEnum = "COVAR_SAMP"
	FunctionNodeTypeRegrSxx                  FunctionNodeTypeEnum = "REGR_SXX"
	FunctionNodeTypeRegrSyy                  FunctionNodeTypeEnum = "REGR_SYY"
	FunctionNodeTypeAvg                      FunctionNodeTypeEnum = "AVG"
	FunctionNodeTypeStddevPop                FunctionNodeTypeEnum = "STDDEV_POP"
	FunctionNodeTypeStddevSamp               FunctionNodeTypeEnum = "STDDEV_SAMP"
	FunctionNodeTypeVarPop                   FunctionNodeTypeEnum = "VAR_POP"
	FunctionNodeTypeVarSamp                  FunctionNodeTypeEnum = "VAR_SAMP"
	FunctionNodeTypeNtile                    FunctionNodeTypeEnum = "NTILE"
	FunctionNodeTypeCollect                  FunctionNodeTypeEnum = "COLLECT"
	FunctionNodeTypeFusion                   FunctionNodeTypeEnum = "FUSION"
	FunctionNodeTypeSingleValue              FunctionNodeTypeEnum = "SINGLE_VALUE"
	FunctionNodeTypeRowNumber                FunctionNodeTypeEnum = "ROW_NUMBER"
	FunctionNodeTypeRank                     FunctionNodeTypeEnum = "RANK"
	FunctionNodeTypePercentRank              FunctionNodeTypeEnum = "PERCENT_RANK"
	FunctionNodeTypeDenseRank                FunctionNodeTypeEnum = "DENSE_RANK"
	FunctionNodeTypeCumeDist                 FunctionNodeTypeEnum = "CUME_DIST"
	FunctionNodeTypeTumble                   FunctionNodeTypeEnum = "TUMBLE"
	FunctionNodeTypeTumbleStart              FunctionNodeTypeEnum = "TUMBLE_START"
	FunctionNodeTypeTumbleEnd                FunctionNodeTypeEnum = "TUMBLE_END"
	FunctionNodeTypeHop                      FunctionNodeTypeEnum = "HOP"
	FunctionNodeTypeHopStart                 FunctionNodeTypeEnum = "HOP_START"
	FunctionNodeTypeHopEnd                   FunctionNodeTypeEnum = "HOP_END"
	FunctionNodeTypeSession                  FunctionNodeTypeEnum = "SESSION"
	FunctionNodeTypeSessionStart             FunctionNodeTypeEnum = "SESSION_START"
	FunctionNodeTypeSessionEnd               FunctionNodeTypeEnum = "SESSION_END"
	FunctionNodeTypeColumnDecl               FunctionNodeTypeEnum = "COLUMN_DECL"
	FunctionNodeTypeCheck                    FunctionNodeTypeEnum = "CHECK"
	FunctionNodeTypeUnique                   FunctionNodeTypeEnum = "UNIQUE"
	FunctionNodeTypePrimaryKey               FunctionNodeTypeEnum = "PRIMARY_KEY"
	FunctionNodeTypeForeignKey               FunctionNodeTypeEnum = "FOREIGN_KEY"
	FunctionNodeTypeCommit                   FunctionNodeTypeEnum = "COMMIT"
	FunctionNodeTypeRollback                 FunctionNodeTypeEnum = "ROLLBACK"
	FunctionNodeTypeAlterSession             FunctionNodeTypeEnum = "ALTER_SESSION"
	FunctionNodeTypeCreateSchema             FunctionNodeTypeEnum = "CREATE_SCHEMA"
	FunctionNodeTypeCreateForeignSchema      FunctionNodeTypeEnum = "CREATE_FOREIGN_SCHEMA"
	FunctionNodeTypeDropSchema               FunctionNodeTypeEnum = "DROP_SCHEMA"
	FunctionNodeTypeCreateTable              FunctionNodeTypeEnum = "CREATE_TABLE"
	FunctionNodeTypeAlterTable               FunctionNodeTypeEnum = "ALTER_TABLE"
	FunctionNodeTypeDropTable                FunctionNodeTypeEnum = "DROP_TABLE"
	FunctionNodeTypeCreateView               FunctionNodeTypeEnum = "CREATE_VIEW"
	FunctionNodeTypeAlterView                FunctionNodeTypeEnum = "ALTER_VIEW"
	FunctionNodeTypeDropView                 FunctionNodeTypeEnum = "DROP_VIEW"
	FunctionNodeTypeCreateMaterializedView   FunctionNodeTypeEnum = "CREATE_MATERIALIZED_VIEW"
	FunctionNodeTypeAlterMaterializedView    FunctionNodeTypeEnum = "ALTER_MATERIALIZED_VIEW"
	FunctionNodeTypeDropMaterializedView     FunctionNodeTypeEnum = "DROP_MATERIALIZED_VIEW"
	FunctionNodeTypeCreateSequence           FunctionNodeTypeEnum = "CREATE_SEQUENCE"
	FunctionNodeTypeAlterSequence            FunctionNodeTypeEnum = "ALTER_SEQUENCE"
	FunctionNodeTypeDropSequence             FunctionNodeTypeEnum = "DROP_SEQUENCE"
	FunctionNodeTypeCreateIndex              FunctionNodeTypeEnum = "CREATE_INDEX"
	FunctionNodeTypeAlterIndex               FunctionNodeTypeEnum = "ALTER_INDEX"
	FunctionNodeTypeDropIndex                FunctionNodeTypeEnum = "DROP_INDEX"
	FunctionNodeTypeOtherDdl                 FunctionNodeTypeEnum = "OTHER_DDL"
)

var mappingFunctionNodeType = map[string]FunctionNodeTypeEnum{
	"OTHER":           FunctionNodeTypeOther,
	"SELECT":          FunctionNodeTypeSelect,
	"JOIN":            FunctionNodeTypeJoin,
	"IDENTIFIER":      FunctionNodeTypeIdentifier,
	"LITERAL":         FunctionNodeTypeLiteral,
	"OTHER_FUNCTION":  FunctionNodeTypeOtherFunction,
	"EXPLAIN":         FunctionNodeTypeExplain,
	"DESCRIBE_SCHEMA": FunctionNodeTypeDescribeSchema,
	"DESCRIBE_TABLE":  FunctionNodeTypeDescribeTable,
	"INSERT":          FunctionNodeTypeInsert,
	"DELETE":          FunctionNodeTypeDelete,
	"UPDATE":          FunctionNodeTypeUpdate,
	"SET_OPTION":      FunctionNodeTypeSetOption,
	"DYNAMIC_PARAM":   FunctionNodeTypeDynamicParam,
	"ORDER_BY":        FunctionNodeTypeOrderBy,
	"WITH":            FunctionNodeTypeWith,
	"WITH_ITEM":       FunctionNodeTypeWithItem,
	"UNION":           FunctionNodeTypeUnion,
	"EXCEPT":          FunctionNodeTypeExcept,
	"INTERSECT":       FunctionNodeTypeIntersect,
	"AS":              FunctionNodeTypeAs,
	"ARGUMENT_ASSIGNMENT":   FunctionNodeTypeArgumentAssignment,
	"DEFAULT":               FunctionNodeTypeDefault,
	"OVER":                  FunctionNodeTypeOver,
	"FILTER":                FunctionNodeTypeFilter,
	"WINDOW":                FunctionNodeTypeWindow,
	"MERGE":                 FunctionNodeTypeMerge,
	"TABLESAMPLE":           FunctionNodeTypeTablesample,
	"MATCH_RECOGNIZE":       FunctionNodeTypeMatchRecognize,
	"TIMES":                 FunctionNodeTypeTimes,
	"DIVIDE":                FunctionNodeTypeDivide,
	"MOD":                   FunctionNodeTypeMod,
	"PLUS":                  FunctionNodeTypePlus,
	"MINUS":                 FunctionNodeTypeMinus,
	"PATTERN_ALTER":         FunctionNodeTypePatternAlter,
	"PATTERN_CONCAT":        FunctionNodeTypePatternConcat,
	"IN":                    FunctionNodeTypeIn,
	"NOT_IN":                FunctionNodeTypeNotIn,
	"LESS_THAN":             FunctionNodeTypeLessThan,
	"GREATER_THAN":          FunctionNodeTypeGreaterThan,
	"LESS_THAN_OR_EQUAL":    FunctionNodeTypeLessThanOrEqual,
	"GREATER_THAN_OR_EQUAL": FunctionNodeTypeGreaterThanOrEqual,
	"EQUALS":                FunctionNodeTypeEquals,
	"NOT_EQUALS":            FunctionNodeTypeNotEquals,
	"IS_DISTINCT_FROM":      FunctionNodeTypeIsDistinctFrom,
	"IS_NOT_DISTINCT_FROM":  FunctionNodeTypeIsNotDistinctFrom,
	"OR":                         FunctionNodeTypeOr,
	"AND":                        FunctionNodeTypeAnd,
	"DOT":                        FunctionNodeTypeDot,
	"OVERLAPS":                   FunctionNodeTypeOverlaps,
	"CONTAINS":                   FunctionNodeTypeContains,
	"PRECEDES":                   FunctionNodeTypePrecedes,
	"IMMEDIATELY_PRECEDES":       FunctionNodeTypeImmediatelyPrecedes,
	"SUCCEEDS":                   FunctionNodeTypeSucceeds,
	"IMMEDIATELY_SUCCEEDS":       FunctionNodeTypeImmediatelySucceeds,
	"PERIOD_EQUALS":              FunctionNodeTypePeriodEquals,
	"LIKE":                       FunctionNodeTypeLike,
	"SIMILAR":                    FunctionNodeTypeSimilar,
	"BETWEEN":                    FunctionNodeTypeBetween,
	"CASE":                       FunctionNodeTypeCase,
	"NULLIF":                     FunctionNodeTypeNullif,
	"COALESCE":                   FunctionNodeTypeCoalesce,
	"DECODE":                     FunctionNodeTypeDecode,
	"NVL":                        FunctionNodeTypeNvl,
	"GREATEST":                   FunctionNodeTypeGreatest,
	"LEAST":                      FunctionNodeTypeLeast,
	"TIMESTAMP_ADD":              FunctionNodeTypeTimestampAdd,
	"TIMESTAMP_DIFF":             FunctionNodeTypeTimestampDiff,
	"NOT":                        FunctionNodeTypeNot,
	"PLUS_PREFIX":                FunctionNodeTypePlusPrefix,
	"MINUS_PREFIX":               FunctionNodeTypeMinusPrefix,
	"EXISTS":                     FunctionNodeTypeExists,
	"SOME":                       FunctionNodeTypeSome,
	"ALL":                        FunctionNodeTypeAll,
	"VALUES":                     FunctionNodeTypeValues,
	"EXPLICIT_TABLE":             FunctionNodeTypeExplicitTable,
	"SCALAR_QUERY":               FunctionNodeTypeScalarQuery,
	"PROCEDURE_CALL":             FunctionNodeTypeProcedureCall,
	"NEW_SPECIFICATION":          FunctionNodeTypeNewSpecification,
	"FINAL":                      FunctionNodeTypeFinal,
	"RUNNING":                    FunctionNodeTypeRunning,
	"PREV":                       FunctionNodeTypePrev,
	"NEXT":                       FunctionNodeTypeNext,
	"FIRST":                      FunctionNodeTypeFirst,
	"LAST":                       FunctionNodeTypeLast,
	"CLASSIFIER":                 FunctionNodeTypeClassifier,
	"MATCH_NUMBER":               FunctionNodeTypeMatchNumber,
	"SKIP_TO_FIRST":              FunctionNodeTypeSkipToFirst,
	"SKIP_TO_LAST":               FunctionNodeTypeSkipToLast,
	"DESCENDING":                 FunctionNodeTypeDescending,
	"NULLS_FIRST":                FunctionNodeTypeNullsFirst,
	"NULLS_LAST":                 FunctionNodeTypeNullsLast,
	"IS_TRUE":                    FunctionNodeTypeIsTrue,
	"IS_FALSE":                   FunctionNodeTypeIsFalse,
	"IS_NOT_TRUE":                FunctionNodeTypeIsNotTrue,
	"IS_NOT_FALSE":               FunctionNodeTypeIsNotFalse,
	"IS_UNKNOWN":                 FunctionNodeTypeIsUnknown,
	"IS_NULL":                    FunctionNodeTypeIsNull,
	"IS_NOT_NULL":                FunctionNodeTypeIsNotNull,
	"PRECEDING":                  FunctionNodeTypePreceding,
	"FOLLOWING":                  FunctionNodeTypeFollowing,
	"FIELD_ACCESS":               FunctionNodeTypeFieldAccess,
	"INPUT_REF":                  FunctionNodeTypeInputRef,
	"TABLE_INPUT_REF":            FunctionNodeTypeTableInputRef,
	"PATTERN_INPUT_REF":          FunctionNodeTypePatternInputRef,
	"LOCAL_REF":                  FunctionNodeTypeLocalRef,
	"CORREL_VARIABLE":            FunctionNodeTypeCorrelVariable,
	"PATTERN_QUANTIFIER":         FunctionNodeTypePatternQuantifier,
	"ROW":                        FunctionNodeTypeRow,
	"COLUMN_LIST":                FunctionNodeTypeColumnList,
	"CAST":                       FunctionNodeTypeCast,
	"NEXT_VALUE":                 FunctionNodeTypeNextValue,
	"CURRENT_VALUE":              FunctionNodeTypeCurrentValue,
	"FLOOR":                      FunctionNodeTypeFloor,
	"CEIL":                       FunctionNodeTypeCeil,
	"TRIM":                       FunctionNodeTypeTrim,
	"LTRIM":                      FunctionNodeTypeLtrim,
	"RTRIM":                      FunctionNodeTypeRtrim,
	"EXTRACT":                    FunctionNodeTypeExtract,
	"JDBC_FN":                    FunctionNodeTypeJdbcFn,
	"MULTISET_VALUE_CONSTRUCTOR": FunctionNodeTypeMultisetValueConstructor,
	"MULTISET_QUERY_CONSTRUCTOR": FunctionNodeTypeMultisetQueryConstructor,
	"UNNEST":                     FunctionNodeTypeUnnest,
	"LATERAL":                    FunctionNodeTypeLateral,
	"COLLECTION_TABLE":           FunctionNodeTypeCollectionTable,
	"ARRAY_VALUE_CONSTRUCTOR":    FunctionNodeTypeArrayValueConstructor,
	"ARRAY_QUERY_CONSTRUCTOR":    FunctionNodeTypeArrayQueryConstructor,
	"MAP_VALUE_CONSTRUCTOR":      FunctionNodeTypeMapValueConstructor,
	"MAP_QUERY_CONSTRUCTOR":      FunctionNodeTypeMapQueryConstructor,
	"CURSOR":                     FunctionNodeTypeCursor,
	"LITERAL_CHAIN":              FunctionNodeTypeLiteralChain,
	"ESCAPE":                     FunctionNodeTypeEscape,
	"REINTERPRET":                FunctionNodeTypeReinterpret,
	"EXTEND":                     FunctionNodeTypeExtend,
	"CUBE":                       FunctionNodeTypeCube,
	"ROLLUP":                     FunctionNodeTypeRollup,
	"GROUPING_SETS":              FunctionNodeTypeGroupingSets,
	"GROUPING":                   FunctionNodeTypeGrouping,
	"GROUPING_ID":                FunctionNodeTypeGroupingId,
	"GROUP_ID":                   FunctionNodeTypeGroupId,
	"PATTERN_PERMUTE":            FunctionNodeTypePatternPermute,
	"PATTERN_EXCLUDED":           FunctionNodeTypePatternExcluded,
	"COUNT":                      FunctionNodeTypeCount,
	"SUM":                        FunctionNodeTypeSum,
	"SUM0":                       FunctionNodeTypeSum0,
	"MIN":                        FunctionNodeTypeMin,
	"MAX":                        FunctionNodeTypeMax,
	"LEAD":                       FunctionNodeTypeLead,
	"LAG":                        FunctionNodeTypeLag,
	"FIRST_VALUE":                FunctionNodeTypeFirstValue,
	"LAST_VALUE":                 FunctionNodeTypeLastValue,
	"COVAR_POP":                  FunctionNodeTypeCovarPop,
	"COVAR_SAMP":                 FunctionNodeTypeCovarSamp,
	"REGR_SXX":                   FunctionNodeTypeRegrSxx,
	"REGR_SYY":                   FunctionNodeTypeRegrSyy,
	"AVG":                        FunctionNodeTypeAvg,
	"STDDEV_POP":                 FunctionNodeTypeStddevPop,
	"STDDEV_SAMP":                FunctionNodeTypeStddevSamp,
	"VAR_POP":                    FunctionNodeTypeVarPop,
	"VAR_SAMP":                   FunctionNodeTypeVarSamp,
	"NTILE":                      FunctionNodeTypeNtile,
	"COLLECT":                    FunctionNodeTypeCollect,
	"FUSION":                     FunctionNodeTypeFusion,
	"SINGLE_VALUE":               FunctionNodeTypeSingleValue,
	"ROW_NUMBER":                 FunctionNodeTypeRowNumber,
	"RANK":                       FunctionNodeTypeRank,
	"PERCENT_RANK":               FunctionNodeTypePercentRank,
	"DENSE_RANK":                 FunctionNodeTypeDenseRank,
	"CUME_DIST":                  FunctionNodeTypeCumeDist,
	"TUMBLE":                     FunctionNodeTypeTumble,
	"TUMBLE_START":               FunctionNodeTypeTumbleStart,
	"TUMBLE_END":                 FunctionNodeTypeTumbleEnd,
	"HOP":                        FunctionNodeTypeHop,
	"HOP_START":                  FunctionNodeTypeHopStart,
	"HOP_END":                    FunctionNodeTypeHopEnd,
	"SESSION":                    FunctionNodeTypeSession,
	"SESSION_START":              FunctionNodeTypeSessionStart,
	"SESSION_END":                FunctionNodeTypeSessionEnd,
	"COLUMN_DECL":                FunctionNodeTypeColumnDecl,
	"CHECK":                      FunctionNodeTypeCheck,
	"UNIQUE":                     FunctionNodeTypeUnique,
	"PRIMARY_KEY":                FunctionNodeTypePrimaryKey,
	"FOREIGN_KEY":                FunctionNodeTypeForeignKey,
	"COMMIT":                     FunctionNodeTypeCommit,
	"ROLLBACK":                   FunctionNodeTypeRollback,
	"ALTER_SESSION":              FunctionNodeTypeAlterSession,
	"CREATE_SCHEMA":              FunctionNodeTypeCreateSchema,
	"CREATE_FOREIGN_SCHEMA":      FunctionNodeTypeCreateForeignSchema,
	"DROP_SCHEMA":                FunctionNodeTypeDropSchema,
	"CREATE_TABLE":               FunctionNodeTypeCreateTable,
	"ALTER_TABLE":                FunctionNodeTypeAlterTable,
	"DROP_TABLE":                 FunctionNodeTypeDropTable,
	"CREATE_VIEW":                FunctionNodeTypeCreateView,
	"ALTER_VIEW":                 FunctionNodeTypeAlterView,
	"DROP_VIEW":                  FunctionNodeTypeDropView,
	"CREATE_MATERIALIZED_VIEW":   FunctionNodeTypeCreateMaterializedView,
	"ALTER_MATERIALIZED_VIEW":    FunctionNodeTypeAlterMaterializedView,
	"DROP_MATERIALIZED_VIEW":     FunctionNodeTypeDropMaterializedView,
	"CREATE_SEQUENCE":            FunctionNodeTypeCreateSequence,
	"ALTER_SEQUENCE":             FunctionNodeTypeAlterSequence,
	"DROP_SEQUENCE":              FunctionNodeTypeDropSequence,
	"CREATE_INDEX":               FunctionNodeTypeCreateIndex,
	"ALTER_INDEX":                FunctionNodeTypeAlterIndex,
	"DROP_INDEX":                 FunctionNodeTypeDropIndex,
	"OTHER_DDL":                  FunctionNodeTypeOtherDdl,
}

// GetFunctionNodeTypeEnumValues Enumerates the set of values for FunctionNodeTypeEnum
func GetFunctionNodeTypeEnumValues() []FunctionNodeTypeEnum {
	values := make([]FunctionNodeTypeEnum, 0)
	for _, v := range mappingFunctionNodeType {
		values = append(values, v)
	}
	return values
}
