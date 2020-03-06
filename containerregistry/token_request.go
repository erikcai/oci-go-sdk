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

// TokenRequest OAuth token request. See https://docs.docker.com/registry/spec/auth/oauth/
type TokenRequest struct {

	// String identifying the client.
	ClientId *string `mandatory:"true" json:"client_id"`

	// Type of grant used to get token.
	GrantType *string `mandatory:"true" json:"grant_type"`

	// The name of the service which hosts the resource to get access for.
	Service *string `mandatory:"true" json:"service"`

	// Access which is being requested.
	AccessType *string `mandatory:"false" json:"access_type"`

	// The password to use for authentication when grant type "password" is used.
	Password *string `mandatory:"false" json:"password"`

	// The refresh token to use for authentication when grant type "refresh_token" is used.
	RefreshToken *string `mandatory:"false" json:"refresh_token"`

	// The resource in question.
	Scope *string `mandatory:"false" json:"scope"`

	// The username to use for authentication when grant type "password" is used.
	Username *string `mandatory:"false" json:"username"`
}

func (m TokenRequest) String() string {
	return common.PointerString(m)
}
