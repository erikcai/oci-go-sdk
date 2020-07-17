// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// QueryParameterTransformation Modifies a query parameter as it passes through the gateway.
type QueryParameterTransformation struct {

	// The name of the query parameter to set or modify.
	Name *string `mandatory:"true" json:"name"`

	// The new value of the query parameter. Optional, may be omitted when action is set to REMOVE.
	Value *string `mandatory:"false" json:"value"`

	// SET will replace all the query parameters if any exists else add it, APPEND will append the new value to the
	// existing query parameters, and REMOVE will remove it.
	Action QueryParameterTransformationActionEnum `mandatory:"false" json:"action,omitempty"`

	// Specifies the condition which should evaluate to true for the action to be performed. Optional, if omitted the
	// action is always performed. A context variable can be referenced in the condition. If the context variable resolves
	// to a non-empty value then the condition evaluates to true else false.
	Condition *string `mandatory:"false" json:"condition"`
}

func (m QueryParameterTransformation) String() string {
	return common.PointerString(m)
}

// QueryParameterTransformationActionEnum Enum with underlying type: string
type QueryParameterTransformationActionEnum string

// Set of constants representing the allowable values for QueryParameterTransformationActionEnum
const (
	QueryParameterTransformationActionSet    QueryParameterTransformationActionEnum = "SET"
	QueryParameterTransformationActionAppend QueryParameterTransformationActionEnum = "APPEND"
	QueryParameterTransformationActionRemove QueryParameterTransformationActionEnum = "REMOVE"
)

var mappingQueryParameterTransformationAction = map[string]QueryParameterTransformationActionEnum{
	"SET":    QueryParameterTransformationActionSet,
	"APPEND": QueryParameterTransformationActionAppend,
	"REMOVE": QueryParameterTransformationActionRemove,
}

// GetQueryParameterTransformationActionEnumValues Enumerates the set of values for QueryParameterTransformationActionEnum
func GetQueryParameterTransformationActionEnumValues() []QueryParameterTransformationActionEnum {
	values := make([]QueryParameterTransformationActionEnum, 0)
	for _, v := range mappingQueryParameterTransformationAction {
		values = append(values, v)
	}
	return values
}
