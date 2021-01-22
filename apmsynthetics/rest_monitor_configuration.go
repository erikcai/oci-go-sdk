// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Synthetic
//
// API for APM Synthetic service. Use this API to query Synthethetic scripts and monitors.
//

package apmsynthetics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v33/common"
)

// RestMonitorConfiguration Request configuration for REST monitor type.
type RestMonitorConfiguration struct {

	// If isFailureRetried enabled, then if call is failed then it will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	// If redirection enabled, then redirects will be allowed while accessing target URL.
	IsRedirectionEnabled *bool `mandatory:"false" json:"isRedirectionEnabled"`

	// If certificate validation enabled, then call will fail for certificate errors.
	IsCertificateValidationEnabled *bool `mandatory:"false" json:"isCertificateValidationEnabled"`

	ReqAuthenticationDetails *RequestAuthenticationDetails `mandatory:"false" json:"reqAuthenticationDetails"`

	// List of request headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]`
	RequestHeaders []Header `mandatory:"false" json:"requestHeaders"`

	// List of request query params. Example: `[{"paramName": "sortOrder", "paramValue": "asc"}]`
	RequestQueryParams []RequestQueryParam `mandatory:"false" json:"requestQueryParams"`

	// Request post body content.
	RequestPostBody *string `mandatory:"false" json:"requestPostBody"`

	// Verify response content against regular expression based string.
	// If response content does not match with the verifyResponseContent, then it will be considered as failure.
	VerifyResponseContent *string `mandatory:"false" json:"verifyResponseContent"`

	// Expected HTTP response codes. For status code range, please set values like 2xx, 3xx.
	VerifyResponseCodes []string `mandatory:"false" json:"verifyResponseCodes"`

	// Request HTTP method.
	RequestMethod RequestMethodsEnum `mandatory:"false" json:"requestMethod,omitempty"`

	// Request http authentication scheme.
	ReqAuthenticationScheme RequestAuthenticationSchemesEnum `mandatory:"false" json:"reqAuthenticationScheme,omitempty"`
}

//GetIsFailureRetried returns IsFailureRetried
func (m RestMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

func (m RestMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RestMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRestMonitorConfiguration RestMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeRestMonitorConfiguration
	}{
		"REST_CONFIG",
		(MarshalTypeRestMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
