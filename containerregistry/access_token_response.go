// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Registry Extension API
//
// API for managing images and repositories in Oracle Cloud Infrastructure Registry.
//

package containerregistry

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AccessTokenResponse Access token response. See https://docs.docker.com/registry/spec/auth/token/
type AccessTokenResponse struct {

	// The duration in seconds since the token was issued that it will remain valid. When omitted, this defaults to 60 seconds. For compatibility with older clients, a token should never be returned with less than 60 seconds to live.
	ExpiresIn *int `mandatory:"true" json:"expires_in"`

	// The resource in question.
	Scope *string `mandatory:"true" json:"scope"`

	// For compatibility with OAuth 2.0, we will also accept token under the name access_token.
	AccessToken *string `mandatory:"false" json:"access_token"`

	// An opaque Bearer token that clients should supply to subsequent requests in the Authorization header.
	Token *string `mandatory:"false" json:"token"`
}

func (m AccessTokenResponse) String() string {
	return common.PointerString(m)
}
