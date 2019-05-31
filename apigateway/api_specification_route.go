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

// ApiSpecificationRoute A single route that forwards requests to a particular backend and may contain some additional policies.
type ApiSpecificationRoute struct {

	// A URL path pattern that must be matched on this route. The path pattern may contain a subset of RFC 6570 identifiers
	// to allow wildcard and parameterized matching.
	Path *string `mandatory:"true" json:"path"`

	// The backend of the route defined in the API.
	Backend ApiSpecificationRouteBackend `mandatory:"true" json:"backend"`

	// A list of allowed methods on this route.
	Methods []ApiSpecificationRouteMethodsEnum `mandatory:"false" json:"methods,omitempty"`

	// Policies to apply on the incoming API requests on a specific route.
	RequestPolicies *ApiSpecificationRouteRequestPolicies `mandatory:"false" json:"requestPolicies"`

	// Policies to apply on the outgoing API response on a specific route.
	ResponsePolicies *ApiSpecificationRouteResponsePolicies `mandatory:"false" json:"responsePolicies"`
}

func (m ApiSpecificationRoute) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ApiSpecificationRoute) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Methods          []ApiSpecificationRouteMethodsEnum     `json:"methods"`
		RequestPolicies  *ApiSpecificationRouteRequestPolicies  `json:"requestPolicies"`
		ResponsePolicies *ApiSpecificationRouteResponsePolicies `json:"responsePolicies"`
		Path             *string                                `json:"path"`
		Backend          apispecificationroutebackend           `json:"backend"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.Methods = make([]ApiSpecificationRouteMethodsEnum, len(model.Methods))
	for i, n := range model.Methods {
		m.Methods[i] = n
	}
	m.RequestPolicies = model.RequestPolicies
	m.ResponsePolicies = model.ResponsePolicies
	m.Path = model.Path
	nn, e := model.Backend.UnmarshalPolymorphicJSON(model.Backend.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Backend = nn.(ApiSpecificationRouteBackend)
	} else {
		m.Backend = nil
	}
	return
}

// ApiSpecificationRouteMethodsEnum Enum with underlying type: string
type ApiSpecificationRouteMethodsEnum string

// Set of constants representing the allowable values for ApiSpecificationRouteMethodsEnum
const (
	ApiSpecificationRouteMethodsAny     ApiSpecificationRouteMethodsEnum = "ANY"
	ApiSpecificationRouteMethodsHead    ApiSpecificationRouteMethodsEnum = "HEAD"
	ApiSpecificationRouteMethodsGet     ApiSpecificationRouteMethodsEnum = "GET"
	ApiSpecificationRouteMethodsPost    ApiSpecificationRouteMethodsEnum = "POST"
	ApiSpecificationRouteMethodsPut     ApiSpecificationRouteMethodsEnum = "PUT"
	ApiSpecificationRouteMethodsPatch   ApiSpecificationRouteMethodsEnum = "PATCH"
	ApiSpecificationRouteMethodsDelete  ApiSpecificationRouteMethodsEnum = "DELETE"
	ApiSpecificationRouteMethodsOptions ApiSpecificationRouteMethodsEnum = "OPTIONS"
)

var mappingApiSpecificationRouteMethods = map[string]ApiSpecificationRouteMethodsEnum{
	"ANY":     ApiSpecificationRouteMethodsAny,
	"HEAD":    ApiSpecificationRouteMethodsHead,
	"GET":     ApiSpecificationRouteMethodsGet,
	"POST":    ApiSpecificationRouteMethodsPost,
	"PUT":     ApiSpecificationRouteMethodsPut,
	"PATCH":   ApiSpecificationRouteMethodsPatch,
	"DELETE":  ApiSpecificationRouteMethodsDelete,
	"OPTIONS": ApiSpecificationRouteMethodsOptions,
}

// GetApiSpecificationRouteMethodsEnumValues Enumerates the set of values for ApiSpecificationRouteMethodsEnum
func GetApiSpecificationRouteMethodsEnumValues() []ApiSpecificationRouteMethodsEnum {
	values := make([]ApiSpecificationRouteMethodsEnum, 0)
	for _, v := range mappingApiSpecificationRouteMethods {
		values = append(values, v)
	}
	return values
}