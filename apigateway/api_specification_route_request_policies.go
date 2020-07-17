// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ApiSpecificationRouteRequestPolicies Behavior applied to any requests received by the API on this route.
type ApiSpecificationRouteRequestPolicies struct {
	Authorization RouteAuthorizationPolicy `mandatory:"false" json:"authorization"`

	Cors *CorsPolicy `mandatory:"false" json:"cors"`

	// Perform these transformations on the HTTP request headers in the order listed.
	HeaderTransformations []HeaderTransformation `mandatory:"false" json:"headerTransformations"`

	// Perform these transformations on the HTTP path query parameters in the order listed.
	QueryParameterTransformations []QueryParameterTransformation `mandatory:"false" json:"queryParameterTransformations"`
}

func (m ApiSpecificationRouteRequestPolicies) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ApiSpecificationRouteRequestPolicies) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Authorization                 routeauthorizationpolicy       `json:"authorization"`
		Cors                          *CorsPolicy                    `json:"cors"`
		HeaderTransformations         []HeaderTransformation         `json:"headerTransformations"`
		QueryParameterTransformations []QueryParameterTransformation `json:"queryParameterTransformations"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Authorization.UnmarshalPolymorphicJSON(model.Authorization.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Authorization = nn.(RouteAuthorizationPolicy)
	} else {
		m.Authorization = nil
	}

	m.Cors = model.Cors

	m.HeaderTransformations = make([]HeaderTransformation, len(model.HeaderTransformations))
	for i, n := range model.HeaderTransformations {
		m.HeaderTransformations[i] = n
	}

	m.QueryParameterTransformations = make([]QueryParameterTransformation, len(model.QueryParameterTransformations))
	for i, n := range model.QueryParameterTransformations {
		m.QueryParameterTransformations[i] = n
	}

	return
}
