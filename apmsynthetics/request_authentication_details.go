// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Synthetic
//
// API for APM Synthetic service. Use this API to query Synthethetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// RequestAuthenticationDetails Request HTTP authentication details.
type RequestAuthenticationDetails struct {

	// Request http oauth scheme.
	OauthScheme OAuthSchemesEnum `mandatory:"false" json:"oauthScheme,omitempty"`

	// username for authentication.
	AuthUserName *string `mandatory:"false" json:"authUserName"`

	// user password for authetication.
	AuthUserPassword *string `mandatory:"false" json:"authUserPassword"`

	// Authentication token.
	AuthToken *string `mandatory:"false" json:"authToken"`

	// URL to get authetication token.
	AuthUrl *string `mandatory:"false" json:"authUrl"`

	// List of authentication headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]`
	AuthHeaders []Header `mandatory:"false" json:"authHeaders"`

	// Request method.
	AuthRequestMethod RequestMethodsEnum `mandatory:"false" json:"authRequestMethod,omitempty"`

	// Request post body.
	AuthRequestPostBody *string `mandatory:"false" json:"authRequestPostBody"`
}

func (m RequestAuthenticationDetails) String() string {
	return common.PointerString(m)
}
