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

// TransformHeadersPolicy Performs a list of header operations on the request or response.
type TransformHeadersPolicy struct {

	// Operation to be performed on the request or response headers.
	Items []TransformHeadersPolicyItem `mandatory:"false" json:"items"`
}

func (m TransformHeadersPolicy) String() string {
	return common.PointerString(m)
}
