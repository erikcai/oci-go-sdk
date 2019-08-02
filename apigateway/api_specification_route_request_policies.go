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

	// If the request has been authenticated, validate that the given scope
	// is active for the request on a specific route.
	AuthorizeScope *AuthorizeScopePolicy `mandatory:"false" json:"authorizeScope"`

	// Enable CORS (Cross-Origin-Resource-Sharing) request handling.
	Cors *CorsPolicy `mandatory:"false" json:"cors"`
}

func (m ApiSpecificationRouteRequestPolicies) String() string {
	return common.PointerString(m)
}
