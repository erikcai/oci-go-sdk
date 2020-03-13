// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDataSafePrivateEndpointRequest wrapper for the GetDataSafePrivateEndpoint operation
type GetDataSafePrivateEndpointRequest struct {

	// unique data safe private endpoint identifier
	DataSafePrivateEndpointId *string `mandatory:"true" contributesTo:"path" name:"dataSafePrivateEndpointId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDataSafePrivateEndpointRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDataSafePrivateEndpointRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDataSafePrivateEndpointRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDataSafePrivateEndpointResponse wrapper for the GetDataSafePrivateEndpoint operation
type GetDataSafePrivateEndpointResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DataSafePrivateEndpoint instance
	DataSafePrivateEndpoint `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDataSafePrivateEndpointResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDataSafePrivateEndpointResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
