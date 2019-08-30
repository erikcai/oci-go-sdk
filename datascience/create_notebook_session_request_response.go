// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// CreateNotebookSessionRequest wrapper for the CreateNotebookSession operation
type CreateNotebookSessionRequest struct {

	// Details for creating a new notebook session.
	CreateNotebookSessionDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or server error without risk of executing that same action again. Retry tokens expire after 24 hours, but can be invalidated before then due to conflicting operations. For example, if a resource has been deleted and purged from the system, then a retry of the original creation request might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateNotebookSessionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateNotebookSessionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateNotebookSessionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateNotebookSessionResponse wrapper for the CreateNotebookSession operation
type CreateNotebookSessionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The NotebookSession instance
	NotebookSession `presentIn:"body"`

	// The URI for the entity being described in the response body.
	Location *string `presentIn:"header" name:"location"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// OCID of the WorkRequest associated with this operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response CreateNotebookSessionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateNotebookSessionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
