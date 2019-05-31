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

// HttpBackend Send the request to an HTTP backend.
type HttpBackend struct {
	Url *string `mandatory:"true" json:"url"`
}

func (m HttpBackend) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HttpBackend) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpBackend HttpBackend
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeHttpBackend
	}{
		"HTTP_BACKEND",
		(MarshalTypeHttpBackend)(m),
	}

	return json.Marshal(&s)
}
