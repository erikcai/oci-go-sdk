// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetLocalPeeringConnectionRequest wrapper for the GetLocalPeeringConnection operation
type GetLocalPeeringConnectionRequest struct {

	// The OCID of the local peering connection. This feature is currently in preview and may change before public release. Do not use it for production workloads.
	LocalPeeringConnectionId *string `mandatory:"true" contributesTo:"path" name:"localPeeringConnectionId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetLocalPeeringConnectionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetLocalPeeringConnectionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetLocalPeeringConnectionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetLocalPeeringConnectionResponse wrapper for the GetLocalPeeringConnection operation
type GetLocalPeeringConnectionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LocalPeeringConnection instance
	LocalPeeringConnection `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetLocalPeeringConnectionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetLocalPeeringConnectionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
