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

// AuthenticationPolicy Information on how to authenticate incoming requests.
type AuthenticationPolicy interface {
}

type authenticationpolicy struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *authenticationpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerauthenticationpolicy authenticationpolicy
	s := struct {
		Model Unmarshalerauthenticationpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *authenticationpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CUSTOM_AUTHENTICATION":
		mm := CustomAuthenticationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m authenticationpolicy) String() string {
	return common.PointerString(m)
}
