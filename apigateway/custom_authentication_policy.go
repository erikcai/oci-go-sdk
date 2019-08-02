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

// CustomAuthenticationPolicy Use a function to validate a custom header or query parameter sent with the request authentication.
// A valid policy must specify either tokenHeader or tokenQueryParam.
type CustomAuthenticationPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle function resource.
	FunctionId *string `mandatory:"true" json:"functionId"`

	// The name of the header containing the authentication token.
	TokenHeader *string `mandatory:"false" json:"tokenHeader"`

	// The name of the query parameter containing the authentication token.
	TokenQueryParam *string `mandatory:"false" json:"tokenQueryParam"`
}

func (m CustomAuthenticationPolicy) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CustomAuthenticationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCustomAuthenticationPolicy CustomAuthenticationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCustomAuthenticationPolicy
	}{
		"CUSTOM_AUTHENTICATION",
		(MarshalTypeCustomAuthenticationPolicy)(m),
	}

	return json.Marshal(&s)
}
