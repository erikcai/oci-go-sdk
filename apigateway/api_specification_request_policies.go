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

	// If the request has been authenticated, validate that the given scopes are active for the request.
	AuthorizeScopes *AuthorizeScopesPolicy `mandatory:"false" json:"authorizeScopes"`

	// Validate the headers on the incoming API requests.
	ValidateHeaders *ValidateHeadersPolicy `mandatory:"false" json:"validateHeaders"`

	// Validate the query parameters on the incoming API requests.
	ValidateQueryParams *ValidateQueryParamsPolicy `mandatory:"false" json:"validateQueryParams"`

	// Transform the headers on the incoming API requests.
	TransformHeaders *TransformHeadersPolicy `mandatory:"false" json:"transformHeaders"`

	// Transform the query parameters on the incoming API requests.
	TransformQueryParams *TransformQueryParamsPolicy `mandatory:"false" json:"transformQueryParams"`
}

func (m ApiSpecificationRequestPolicies) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ApiSpecificationRequestPolicies) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Authentication       authenticationpolicy        `json:"authentication"`
		RateLimiting         *RateLimitingPolicy         `json:"rateLimiting"`
		AuthorizeScopes      *AuthorizeScopesPolicy      `json:"authorizeScopes"`
		ValidateHeaders      *ValidateHeadersPolicy      `json:"validateHeaders"`
		ValidateQueryParams  *ValidateQueryParamsPolicy  `json:"validateQueryParams"`
		TransformHeaders     *TransformHeadersPolicy     `json:"transformHeaders"`
		TransformQueryParams *TransformQueryParamsPolicy `json:"transformQueryParams"`
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
	m.AuthorizeScopes = model.AuthorizeScopes
	m.ValidateHeaders = model.ValidateHeaders
	m.ValidateQueryParams = model.ValidateQueryParams
	m.TransformHeaders = model.TransformHeaders
	m.TransformQueryParams = model.TransformQueryParams
	return
}
