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

// TransformQueryParamsPolicy Performs a list of query parameter operations on the request or response.
type TransformQueryParamsPolicy struct {

	// Operation to be performed on the request or response query parameters.
	Items []TransformQueryParamsPolicyItem `mandatory:"false" json:"items"`
}

func (m TransformQueryParamsPolicy) String() string {
	return common.PointerString(m)
}
