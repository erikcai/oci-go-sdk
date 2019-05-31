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

// ApiSpecificationRouteBackend The backend to forward requests to.
type ApiSpecificationRouteBackend interface {
}

type apispecificationroutebackend struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *apispecificationroutebackend) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerapispecificationroutebackend apispecificationroutebackend
	s := struct {
		Model Unmarshalerapispecificationroutebackend
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *apispecificationroutebackend) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "HTTP_BACKEND":
		mm := HttpBackend{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_FUNCTIONS_BACKEND":
		mm := OracleFunctionBackend{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m apispecificationroutebackend) String() string {
	return common.PointerString(m)
}
