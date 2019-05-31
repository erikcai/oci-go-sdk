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

// ApiSpecificationResponsePolicies Global behaviour applied to all responses sent by the API.
type ApiSpecificationResponsePolicies struct {

	// Transform the headers on the outgoing API response.
	TransformHeaders *TransformHeadersPolicy `mandatory:"false" json:"transformHeaders"`

	// Transform the query parameter on the outgoing API response.
	TransformQueryParams *TransformQueryParamsPolicy `mandatory:"false" json:"transformQueryParams"`
}

func (m ApiSpecificationResponsePolicies) String() string {
	return common.PointerString(m)
}
