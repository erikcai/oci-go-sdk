// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateBlockchainPlatformRequest wrapper for the UpdateBlockchainPlatform operation
type UpdateBlockchainPlatformRequest struct {

	// The Blockchain Platform fields to update. Fields that are not provided will not be updated.
	UpdateBlockchainPlatformDetails `contributesTo:"body"`

	// unique service identifier
	BlockchainPlatformId *string `mandatory:"true" contributesTo:"path" name:"blockchainPlatformId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateBlockchainPlatformRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateBlockchainPlatformRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateBlockchainPlatformRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateBlockchainPlatformResponse wrapper for the UpdateBlockchainPlatform operation
type UpdateBlockchainPlatformResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateBlockchainPlatformResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateBlockchainPlatformResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}