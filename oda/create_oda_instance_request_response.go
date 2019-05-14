// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// CreateOdaInstanceRequest wrapper for the CreateOdaInstance operation
type CreateOdaInstanceRequest struct {

	// Details for the new ODA instance.
	CreateOdaInstanceDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Internal use only
	OpcHsmDomain *string `mandatory:"false" contributesTo:"header" name:"opc-hsm-domain"`

	// Internal use only.
	OpcOboToken *string `mandatory:"false" contributesTo:"header" name:"opc-obo-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateOdaInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateOdaInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateOdaInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateOdaInstanceResponse wrapper for the CreateOdaInstance operation
type CreateOdaInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OdaInstance instance
	OdaInstance `presentIn:"body"`

	// Fully qualified URL for the newly created instance
	Location *string `presentIn:"header" name:"location"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the asynchronous request. You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateOdaInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateOdaInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
