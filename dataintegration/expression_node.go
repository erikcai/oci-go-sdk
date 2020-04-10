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

// ExpressionNode auto generated description
type ExpressionNode interface {
}

type expressionnode struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *expressionnode) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexpressionnode expressionnode
	s := struct {
		Model Unmarshalerexpressionnode
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *expressionnode) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "REFERENCE_NODE":
		mm := ReferenceNode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IDENTIFIER_NODE":
		mm := IdentifierNode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VALUE_NODE":
		mm := ValueNode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OPERATOR_NODE":
		mm := OperatorNode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m expressionnode) String() string {
	return common.PointerString(m)
}

// ExpressionNodeModelTypeEnum Enum with underlying type: string
type ExpressionNodeModelTypeEnum string

// Set of constants representing the allowable values for ExpressionNodeModelTypeEnum
const (
	ExpressionNodeModelTypeReferenceNode  ExpressionNodeModelTypeEnum = "REFERENCE_NODE"
	ExpressionNodeModelTypeOperatorNode   ExpressionNodeModelTypeEnum = "OPERATOR_NODE"
	ExpressionNodeModelTypeRootNode       ExpressionNodeModelTypeEnum = "ROOT_NODE"
	ExpressionNodeModelTypeValueNode      ExpressionNodeModelTypeEnum = "VALUE_NODE"
	ExpressionNodeModelTypeIdentifierNode ExpressionNodeModelTypeEnum = "IDENTIFIER_NODE"
	ExpressionNodeModelTypeParameterNode  ExpressionNodeModelTypeEnum = "PARAMETER_NODE"
	ExpressionNodeModelTypeFunctionNode   ExpressionNodeModelTypeEnum = "FUNCTION_NODE"
)

var mappingExpressionNodeModelType = map[string]ExpressionNodeModelTypeEnum{
	"REFERENCE_NODE":  ExpressionNodeModelTypeReferenceNode,
	"OPERATOR_NODE":   ExpressionNodeModelTypeOperatorNode,
	"ROOT_NODE":       ExpressionNodeModelTypeRootNode,
	"VALUE_NODE":      ExpressionNodeModelTypeValueNode,
	"IDENTIFIER_NODE": ExpressionNodeModelTypeIdentifierNode,
	"PARAMETER_NODE":  ExpressionNodeModelTypeParameterNode,
	"FUNCTION_NODE":   ExpressionNodeModelTypeFunctionNode,
}

// GetExpressionNodeModelTypeEnumValues Enumerates the set of values for ExpressionNodeModelTypeEnum
func GetExpressionNodeModelTypeEnumValues() []ExpressionNodeModelTypeEnum {
	values := make([]ExpressionNodeModelTypeEnum, 0)
	for _, v := range mappingExpressionNodeModelType {
		values = append(values, v)
	}
	return values
}
