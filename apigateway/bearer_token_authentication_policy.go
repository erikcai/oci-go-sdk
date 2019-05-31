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

// BearerTokenAuthenticationPolicy Use a function to validate the bearer token sent with the request authentication.
type BearerTokenAuthenticationPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle function resource.
	FunctionId *string `mandatory:"true" json:"functionId"`
}

func (m BearerTokenAuthenticationPolicy) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m BearerTokenAuthenticationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBearerTokenAuthenticationPolicy BearerTokenAuthenticationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeBearerTokenAuthenticationPolicy
	}{
		"BEARER_TOKEN_AUTHENTICATION",
		(MarshalTypeBearerTokenAuthenticationPolicy)(m),
	}

	return json.Marshal(&s)
}
