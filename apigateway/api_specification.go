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

// ApiSpecification The logical configuration of the API exposed by a deployment.
type ApiSpecification struct {

	// Policies that should be applied to the incoming API requests.
	RequestPolicies *ApiSpecificationRequestPolicies `mandatory:"false" json:"requestPolicies"`

	// Policies that should be applied to the outgoing API response.
	ResponsePolicies *ApiSpecificationResponsePolicies `mandatory:"false" json:"responsePolicies"`

	// A list of routes that this API exposes.
	Routes []ApiSpecificationRoute `mandatory:"false" json:"routes"`
}

func (m ApiSpecification) String() string {
	return common.PointerString(m)
}
