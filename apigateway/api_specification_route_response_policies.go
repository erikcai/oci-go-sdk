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

// ApiSpecificationRouteResponsePolicies Behaviour applied to any responses sent by the API for requests on this route.
type ApiSpecificationRouteResponsePolicies struct {

	// Transform the headers on the outgoing API response on a specific route.
	TransformHeaders *TransformHeadersPolicy `mandatory:"false" json:"transformHeaders"`

	// Transform the query parameters on the outgoing API response on a specific route.
	TransformQueryParams *TransformQueryParamsPolicy `mandatory:"false" json:"transformQueryParams"`
}

func (m ApiSpecificationRouteResponsePolicies) String() string {
	return common.PointerString(m)
}
