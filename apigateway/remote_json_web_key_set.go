// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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

// RemoteJsonWebKeySet a public key that is retrieved at run-time from a remote location
type RemoteJsonWebKeySet struct {

	// The uri from which to retrieve the key.  Must be a uri accessible without authentication.
	Uri *string `mandatory:"true" json:"uri"`

	// The duration for which the JWKS should be cached before it is fetched again.
	MaxCacheDurationInHours *int `mandatory:"false" json:"maxCacheDurationInHours"`
}

func (m RemoteJsonWebKeySet) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RemoteJsonWebKeySet) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRemoteJsonWebKeySet RemoteJsonWebKeySet
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRemoteJsonWebKeySet
	}{
		"REMOTE_JWKS",
		(MarshalTypeRemoteJsonWebKeySet)(m),
	}

	return json.Marshal(&s)
}
