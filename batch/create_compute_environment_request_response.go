// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package batch

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"net/http"
)

// CreateComputeEnvironmentRequest wrapper for the CreateComputeEnvironment operation
type CreateComputeEnvironmentRequest struct {

	// The compute environment is being created.
	CreateComputeEnvironmentDetails `contributesTo:"body"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of
	// a timeout or server error without risk of executing that same action again.
	// Retry tokens expire after 24 hours,
	// but can be invalidated before then due to conflicting operations
	// (for example, if a resource has been deleted and purged from the system,
	// then a retry of the original creation request may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateComputeEnvironmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateComputeEnvironmentRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateComputeEnvironmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateComputeEnvironmentResponse wrapper for the CreateComputeEnvironment operation
type CreateComputeEnvironmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ComputeEnvironment instance
	ComputeEnvironment `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response CreateComputeEnvironmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateComputeEnvironmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
