// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datasecurity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDataSecurityInstanceRequest wrapper for the GetDataSecurityInstance operation
type GetDataSecurityInstanceRequest struct {

	// unique data security instance identifier
	DataSecurityInstanceId *string `mandatory:"true" contributesTo:"path" name:"dataSecurityInstanceId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDataSecurityInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDataSecurityInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDataSecurityInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDataSecurityInstanceResponse wrapper for the GetDataSecurityInstance operation
type GetDataSecurityInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DataSecurityInstance instance
	DataSecurityInstance `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDataSecurityInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDataSecurityInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
