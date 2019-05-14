// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetOdaInstanceRequest wrapper for the GetOdaInstance operation
type GetOdaInstanceRequest struct {

	// unique ODA instance identifier
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Internal use only.
	OpcOboToken *string `mandatory:"false" contributesTo:"header" name:"opc-obo-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetOdaInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetOdaInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetOdaInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetOdaInstanceResponse wrapper for the GetOdaInstance operation
type GetOdaInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OdaInstance instance
	OdaInstance `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Internal use only
	OpcHsmApp *string `presentIn:"header" name:"opc-hsm-app"`
}

func (response GetOdaInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetOdaInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
