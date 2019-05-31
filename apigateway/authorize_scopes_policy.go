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

// AuthorizeScopesPolicy If authentication has been performed, validate that the given scopes are active for the request.
type AuthorizeScopesPolicy struct {

	// The scopes that should be active for an authorized request.
	RequiredScopes []string `mandatory:"false" json:"requiredScopes"`
}

func (m AuthorizeScopesPolicy) String() string {
	return common.PointerString(m)
}
