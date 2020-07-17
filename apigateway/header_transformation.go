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

// HeaderTransformation Modifies an HTTP header as it passes through the gateway.
type HeaderTransformation struct {

	// The name of the header to set or modify.
	Name *string `mandatory:"true" json:"name"`

	// The new value of the header. Optional, may be omitted when action is set to REMOVE.
	Value *string `mandatory:"false" json:"value"`

	// SET will replace all the header if any exists else add it, APPEND will append the new value to the existing headers,
	// and REMOVE will remove it.
	Action HeaderTransformationActionEnum `mandatory:"false" json:"action,omitempty"`

	// Specifies the condition which should evaluate to true for the action to be performed. Optional, if omitted the
	// action is always performed. A context variable can be referenced in the condition. If the context variable resolves
	// to a non-empty value then the condition evaluates to true else false.
	Condition *string `mandatory:"false" json:"condition"`
}

func (m HeaderTransformation) String() string {
	return common.PointerString(m)
}

// HeaderTransformationActionEnum Enum with underlying type: string
type HeaderTransformationActionEnum string

// Set of constants representing the allowable values for HeaderTransformationActionEnum
const (
	HeaderTransformationActionSet    HeaderTransformationActionEnum = "SET"
	HeaderTransformationActionAppend HeaderTransformationActionEnum = "APPEND"
	HeaderTransformationActionRemove HeaderTransformationActionEnum = "REMOVE"
)

var mappingHeaderTransformationAction = map[string]HeaderTransformationActionEnum{
	"SET":    HeaderTransformationActionSet,
	"APPEND": HeaderTransformationActionAppend,
	"REMOVE": HeaderTransformationActionRemove,
}

// GetHeaderTransformationActionEnumValues Enumerates the set of values for HeaderTransformationActionEnum
func GetHeaderTransformationActionEnumValues() []HeaderTransformationActionEnum {
	values := make([]HeaderTransformationActionEnum, 0)
	for _, v := range mappingHeaderTransformationAction {
		values = append(values, v)
	}
	return values
}
