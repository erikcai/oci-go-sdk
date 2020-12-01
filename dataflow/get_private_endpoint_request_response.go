// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataflow

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"net/http"
)

// GetPrivateEndpointRequest wrapper for the GetPrivateEndpoint operation
type GetPrivateEndpointRequest struct {

	// The unique ID for a private endpoint.
	PrivateEndpointId *string `mandatory:"true" contributesTo:"path" name:"privateEndpointId"`

	// Unique identifier for the request. If provided, the returned request ID will include this value.
	// Otherwise, a random request ID will be generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPrivateEndpointRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPrivateEndpointRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPrivateEndpointRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetPrivateEndpointResponse wrapper for the GetPrivateEndpoint operation
type GetPrivateEndpointResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PrivateEndpoint instance
	PrivateEndpoint `presentIn:"body"`

	// For optimistic concurrency control.
	// See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPrivateEndpointResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPrivateEndpointResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
