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

// TransformQueryParamsPolicyItem Sets or deletes a query parameter.
type TransformQueryParamsPolicyItem struct {

	// Query parameter to perform the operation on.
	Key *string `mandatory:"true" json:"key"`

	// Operation to be performed.
	Operation TransformQueryParamsPolicyItemOperationEnum `mandatory:"true" json:"operation"`

	// Transform the query parameter to this value based on the operation type.
	Value *string `mandatory:"false" json:"value"`
}

func (m TransformQueryParamsPolicyItem) String() string {
	return common.PointerString(m)
}

// TransformQueryParamsPolicyItemOperationEnum Enum with underlying type: string
type TransformQueryParamsPolicyItemOperationEnum string

// Set of constants representing the allowable values for TransformQueryParamsPolicyItemOperationEnum
const (
	TransformQueryParamsPolicyItemOperationSet    TransformQueryParamsPolicyItemOperationEnum = "SET"
	TransformQueryParamsPolicyItemOperationDelete TransformQueryParamsPolicyItemOperationEnum = "DELETE"
)

var mappingTransformQueryParamsPolicyItemOperation = map[string]TransformQueryParamsPolicyItemOperationEnum{
	"SET":    TransformQueryParamsPolicyItemOperationSet,
	"DELETE": TransformQueryParamsPolicyItemOperationDelete,
}

// GetTransformQueryParamsPolicyItemOperationEnumValues Enumerates the set of values for TransformQueryParamsPolicyItemOperationEnum
func GetTransformQueryParamsPolicyItemOperationEnumValues() []TransformQueryParamsPolicyItemOperationEnum {
	values := make([]TransformQueryParamsPolicyItemOperationEnum, 0)
	for _, v := range mappingTransformQueryParamsPolicyItemOperation {
		values = append(values, v)
	}
	return values
}
