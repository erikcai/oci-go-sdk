// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package cache

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// CancelWorkRequestRequest wrapper for the CancelWorkRequest operation
type CancelWorkRequestRequest struct {

	// The OCID of the work request.
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// Use the `if-match` parameter to use optimistic concurrency control. In the `PUT` or `DELETE` call
	// for a resource, set the `if-match` parameter to the value of the `etag`
	// from a previous `GET` or `POST` response for that resource. The resource
	// is updated or deleted only if the `etag` matches the resource's
	// current `etag` value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"If-Match"`

	// A unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CancelWorkRequestRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CancelWorkRequestRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CancelWorkRequestRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CancelWorkRequestResponse wrapper for the CancelWorkRequest operation
type CancelWorkRequestResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CancelWorkRequestResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CancelWorkRequestResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
