// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetLocalPeeringGatewayRequest wrapper for the GetLocalPeeringGateway operation
type GetLocalPeeringGatewayRequest struct {

	// The OCID of the local peering gateway.
	LocalPeeringGatewayId *string `mandatory:"true" contributesTo:"path" name:"localPeeringGatewayId"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetLocalPeeringGatewayRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetLocalPeeringGatewayRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetLocalPeeringGatewayRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetLocalPeeringGatewayResponse wrapper for the GetLocalPeeringGateway operation
type GetLocalPeeringGatewayResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LocalPeeringGateway instance
	LocalPeeringGateway `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetLocalPeeringGatewayResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetLocalPeeringGatewayResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
