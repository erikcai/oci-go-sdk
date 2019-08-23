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

// ApiSpecificationRouteRequestPolicies Behaviour applied to any requests received by the API on this route.
type ApiSpecificationRouteRequestPolicies struct {

	// If the request has been authenticated, confirm that a valid scope
	// is active for the request on a specific route.
	Authorization RouteAuthorizationPolicy `mandatory:"false" json:"authorization"`

	// Enable CORS (Cross-Origin-Resource-Sharing) request handling.
	Cors *CorsPolicy `mandatory:"false" json:"cors"`
}

func (m ApiSpecificationRouteRequestPolicies) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ApiSpecificationRouteRequestPolicies) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Authorization routeauthorizationpolicy `json:"authorization"`
		Cors          *CorsPolicy              `json:"cors"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	nn, e := model.Authorization.UnmarshalPolymorphicJSON(model.Authorization.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Authorization = nn.(RouteAuthorizationPolicy)
	} else {
		m.Authorization = nil
	}
	m.Cors = model.Cors
	return
}
