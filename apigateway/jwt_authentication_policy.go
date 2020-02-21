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

// JwtAuthenticationPolicy An authentication policy that allows native JWT authentication
type JwtAuthenticationPolicy struct {

	// party that issued the certificate
	Issuers []string `mandatory:"true" json:"issuers"`

	// recipients that the authentication is intended for
	Audiences []string `mandatory:"true" json:"audiences"`

	PublicKeys PublicKeySet `mandatory:"true" json:"publicKeys"`

	// Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS
	// route authorization.
	IsAnonymousAccessAllowed *bool `mandatory:"false" json:"isAnonymousAccessAllowed"`

	// the header that contains the authentication
	TokenHeader *string `mandatory:"false" json:"tokenHeader"`

	// the query parameter that contains the authentication
	TokenQueryParam *string `mandatory:"false" json:"tokenQueryParam"`

	// scheme that is used to authenticate the token
	TokenAuthScheme *string `mandatory:"false" json:"tokenAuthScheme"`

	// A list of claims which should be validated to consider the JWT valid
	VerifyClaims []JsonWebTokenClaim `mandatory:"false" json:"verifyClaims"`

	// Maximum expected time difference between the system clocks
	// of the token issuer and the API Gateway
	MaxClockSkewInSeconds *float32 `mandatory:"false" json:"maxClockSkewInSeconds"`
}

//GetIsAnonymousAccessAllowed returns IsAnonymousAccessAllowed
func (m JwtAuthenticationPolicy) GetIsAnonymousAccessAllowed() *bool {
	return m.IsAnonymousAccessAllowed
}

func (m JwtAuthenticationPolicy) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m JwtAuthenticationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJwtAuthenticationPolicy JwtAuthenticationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeJwtAuthenticationPolicy
	}{
		"JWT_AUTHENTICATION",
		(MarshalTypeJwtAuthenticationPolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *JwtAuthenticationPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsAnonymousAccessAllowed *bool               `json:"isAnonymousAccessAllowed"`
		TokenHeader              *string             `json:"tokenHeader"`
		TokenQueryParam          *string             `json:"tokenQueryParam"`
		TokenAuthScheme          *string             `json:"tokenAuthScheme"`
		VerifyClaims             []JsonWebTokenClaim `json:"verifyClaims"`
		MaxClockSkewInSeconds    *float32            `json:"maxClockSkewInSeconds"`
		Issuers                  []string            `json:"issuers"`
		Audiences                []string            `json:"audiences"`
		PublicKeys               publickeyset        `json:"publicKeys"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsAnonymousAccessAllowed = model.IsAnonymousAccessAllowed

	m.TokenHeader = model.TokenHeader

	m.TokenQueryParam = model.TokenQueryParam

	m.TokenAuthScheme = model.TokenAuthScheme

	m.VerifyClaims = make([]JsonWebTokenClaim, len(model.VerifyClaims))
	for i, n := range model.VerifyClaims {
		m.VerifyClaims[i] = n
	}

	m.MaxClockSkewInSeconds = model.MaxClockSkewInSeconds

	m.Issuers = make([]string, len(model.Issuers))
	for i, n := range model.Issuers {
		m.Issuers[i] = n
	}

	m.Audiences = make([]string, len(model.Audiences))
	for i, n := range model.Audiences {
		m.Audiences[i] = n
	}

	nn, e = model.PublicKeys.UnmarshalPolymorphicJSON(model.PublicKeys.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PublicKeys = nn.(PublicKeySet)
	} else {
		m.PublicKeys = nil
	}
	return
}
