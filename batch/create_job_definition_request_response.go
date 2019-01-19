// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package batch

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// CreateJobDefinitionRequest wrapper for the CreateJobDefinition operation
type CreateJobDefinitionRequest struct {

	// Request to create a job definition.
	CreateJobDefinitionDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of
	// a timeout or server error without risk of executing that same action again.
	// Retry tokens expire after 24 hours,
	// but can be invalidated before then due to conflicting operations
	// (for example, if a resource has been deleted and purged from the system,
	// then a retry of the original creation request may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateJobDefinitionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateJobDefinitionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateJobDefinitionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateJobDefinitionResponse wrapper for the CreateJobDefinition operation
type CreateJobDefinitionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The JobDefinition instance
	JobDefinition `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response CreateJobDefinitionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateJobDefinitionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
