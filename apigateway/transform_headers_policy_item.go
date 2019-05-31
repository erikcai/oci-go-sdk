// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TransformHeadersPolicyItem Sets or deletes a http header.
type TransformHeadersPolicyItem struct {

	// Header to perform the operation on.
	Header *string `mandatory:"true" json:"header"`

	// Operation to be performed.
	Operation TransformHeadersPolicyItemOperationEnum `mandatory:"true" json:"operation"`

	// Transform the header to this value based on the operation type.
	Value *string `mandatory:"false" json:"value"`
}

func (m TransformHeadersPolicyItem) String() string {
	return common.PointerString(m)
}

// TransformHeadersPolicyItemOperationEnum Enum with underlying type: string
type TransformHeadersPolicyItemOperationEnum string

// Set of constants representing the allowable values for TransformHeadersPolicyItemOperationEnum
const (
	TransformHeadersPolicyItemOperationSet       TransformHeadersPolicyItemOperationEnum = "SET"
	TransformHeadersPolicyItemOperationAppend    TransformHeadersPolicyItemOperationEnum = "APPEND"
	TransformHeadersPolicyItemOperationOverwrite TransformHeadersPolicyItemOperationEnum = "OVERWRITE"
	TransformHeadersPolicyItemOperationDelete    TransformHeadersPolicyItemOperationEnum = "DELETE"
)

var mappingTransformHeadersPolicyItemOperation = map[string]TransformHeadersPolicyItemOperationEnum{
	"SET":       TransformHeadersPolicyItemOperationSet,
	"APPEND":    TransformHeadersPolicyItemOperationAppend,
	"OVERWRITE": TransformHeadersPolicyItemOperationOverwrite,
	"DELETE":    TransformHeadersPolicyItemOperationDelete,
}

// GetTransformHeadersPolicyItemOperationEnumValues Enumerates the set of values for TransformHeadersPolicyItemOperationEnum
func GetTransformHeadersPolicyItemOperationEnumValues() []TransformHeadersPolicyItemOperationEnum {
	values := make([]TransformHeadersPolicyItemOperationEnum, 0)
	for _, v := range mappingTransformHeadersPolicyItemOperation {
		values = append(values, v)
	}
	return values
}
