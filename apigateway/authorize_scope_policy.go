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

// AuthorizeScopePolicy If authentication has been performed, validate that the given scope is active for the request.
type AuthorizeScopePolicy struct {

	// A user whose scope includes any of these access ranges is allowed on
	// this route. Access ranges are case-sensitive.
	AllowedScope []string `mandatory:"false" json:"allowedScope"`
}

func (m AuthorizeScopePolicy) String() string {
	return common.PointerString(m)
}
