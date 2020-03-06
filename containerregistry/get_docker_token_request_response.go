// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package containerregistry

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDockerTokenRequest wrapper for the GetDockerToken operation
type GetDockerTokenRequest struct {

	// Authorization header
	Authorization *string `mandatory:"true" contributesTo:"header" name:"Authorization"`

	// Domain name of target service server
	Service *string `mandatory:"false" contributesTo:"query" name:"service"`

	// Scope to restrict access to resources. Consists of Docker token access as described at https://docs.docker.com/registry/spec/auth/scope/
	Scope []string `contributesTo:"query" name:"scope" collectionFormat:"csv"`

	// Unique Oracle-assigned identifier for the request
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDockerTokenRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDockerTokenRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDockerTokenRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDockerTokenResponse wrapper for the GetDockerToken operation
type GetDockerTokenResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AccessTokenResponse instance
	AccessTokenResponse `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDockerTokenResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDockerTokenResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
