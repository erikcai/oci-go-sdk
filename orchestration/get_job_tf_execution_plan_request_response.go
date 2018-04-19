// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package orchestration

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// GetJobTfExecutionPlanRequest wrapper for the GetJobTfExecutionPlan operation
type GetJobTfExecutionPlanRequest struct {

	// Job OCID
	JobId *string `mandatory:"true" contributesTo:"path" name:"jobId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetJobTfExecutionPlanRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJobTfExecutionPlanRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobTfExecutionPlanRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetJobTfExecutionPlanResponse wrapper for the GetJobTfExecutionPlan operation
type GetJobTfExecutionPlanResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetJobTfExecutionPlanResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJobTfExecutionPlanResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
