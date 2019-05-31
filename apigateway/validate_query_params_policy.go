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

// ValidateQueryParamsPolicy Performs a list of query parameter validations on the request or response.
type ValidateQueryParamsPolicy struct {

	// Validate query parameters on the request or response.
	Validations []ValidateQueryParamsPolicyItem `mandatory:"false" json:"validations"`
}

func (m ValidateQueryParamsPolicy) String() string {
	return common.PointerString(m)
}
