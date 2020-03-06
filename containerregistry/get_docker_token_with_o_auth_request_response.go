// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package containerregistry

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDockerTokenWithOAuthRequest wrapper for the GetDockerTokenWithOAuth operation
type GetDockerTokenWithOAuthRequest struct {

	// Token request
	TokenRequest `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDockerTokenWithOAuthRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDockerTokenWithOAuthRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDockerTokenWithOAuthRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDockerTokenWithOAuthResponse wrapper for the GetDockerTokenWithOAuth operation
type GetDockerTokenWithOAuthResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AccessTokenResponse instance
	AccessTokenResponse `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDockerTokenWithOAuthResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDockerTokenWithOAuthResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
