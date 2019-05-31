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

// CustomAuthenticationPolicy Use a function to validate custom headers or query parameters sent with the request authentication.
type CustomAuthenticationPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle function resource.
	FunctionId *string `mandatory:"true" json:"functionId"`

	Headers []string `mandatory:"false" json:"headers"`

	QueryParams []string `mandatory:"false" json:"queryParams"`
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
