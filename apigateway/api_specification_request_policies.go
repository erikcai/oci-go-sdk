// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
//

package apigateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ApiSpecificationRequestPolicies Global behaviour applied to all requests received by the API.
type ApiSpecificationRequestPolicies struct {

	// Authentication policy for the incoming API requests.
	Authentication AuthenticationPolicy `mandatory:"false" json:"authentication"`

	// Rate limiting policy for the incoming API requests.
	RateLimiting *RateLimitingPolicy `mandatory:"false" json:"rateLimiting"`

	// Enable CORS (Cross-Origin-Resource-Sharing) request handling.
	Cors *CorsPolicy `mandatory:"false" json:"cors"`
}

func (m ApiSpecificationRequestPolicies) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ApiSpecificationRequestPolicies) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Authentication authenticationpolicy `json:"authentication"`
		RateLimiting   *RateLimitingPolicy  `json:"rateLimiting"`
		Cors           *CorsPolicy          `json:"cors"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	nn, e := model.Authentication.UnmarshalPolymorphicJSON(model.Authentication.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Authentication = nn.(AuthenticationPolicy)
	} else {
		m.Authentication = nil
	}
	m.RateLimiting = model.RateLimiting
	m.Cors = model.Cors
	return
}
