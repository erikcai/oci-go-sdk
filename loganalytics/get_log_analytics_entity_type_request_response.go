// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"net/http"
)

// GetLogAnalyticsEntityTypeRequest wrapper for the GetLogAnalyticsEntityType operation
type GetLogAnalyticsEntityTypeRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Log analytics entity type name.
	EntityTypeName *string `mandatory:"true" contributesTo:"path" name:"entityTypeName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetLogAnalyticsEntityTypeRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetLogAnalyticsEntityTypeRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetLogAnalyticsEntityTypeRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetLogAnalyticsEntityTypeResponse wrapper for the GetLogAnalyticsEntityType operation
type GetLogAnalyticsEntityTypeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LogAnalyticsEntityType instance
	LogAnalyticsEntityType `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetLogAnalyticsEntityTypeResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetLogAnalyticsEntityTypeResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
