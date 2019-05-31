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

// ApiSpecificationRouteRequestPolicies Behaviour applied to any requests received by the API on this route.
type ApiSpecificationRouteRequestPolicies struct {

	// If the request has been authenticated, validate that the given scopes are active for the request on a
	// specific route.
	AuthorizeScopes *AuthorizeScopesPolicy `mandatory:"false" json:"authorizeScopes"`

	// Validate the headers on the incoming API requests on a specific route.
	ValidateHeaders *ValidateHeadersPolicy `mandatory:"false" json:"validateHeaders"`

	// Validate the query parameters on the incoming API requests on a specific route.
	ValidateQueryParams *ValidateQueryParamsPolicy `mandatory:"false" json:"validateQueryParams"`

	// Transform the headers on the incoming API requests on a specific route.
	TransformHeaders *TransformHeadersPolicy `mandatory:"false" json:"transformHeaders"`

	// Transform the query parameters on the incoming API requests on a specific route.
	TransformQueryParams *TransformQueryParamsPolicy `mandatory:"false" json:"transformQueryParams"`
}

func (m ApiSpecificationRouteRequestPolicies) String() string {
	return common.PointerString(m)
}
