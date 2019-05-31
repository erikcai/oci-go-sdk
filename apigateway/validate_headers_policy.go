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

// ValidateHeadersPolicy Performs a list of header validations on the request or response.
type ValidateHeadersPolicy struct {

	// Validate headers on the request or response.
	Validations []ValidateHeadersPolicyItem `mandatory:"false" json:"validations"`
}

func (m ValidateHeadersPolicy) String() string {
	return common.PointerString(m)
}
