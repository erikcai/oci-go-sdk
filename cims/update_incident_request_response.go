// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateIncidentRequest wrapper for the UpdateIncident operation
type UpdateIncidentRequest struct {
	IncidentId *string `mandatory:"true" contributesTo:"path" name:"incidentId"`

	UpdateIncidentDetails `contributesTo:"body"`

	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateIncidentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateIncidentRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateIncidentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateIncidentResponse wrapper for the UpdateIncident operation
type UpdateIncidentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Incident instance
	Incident `presentIn:"body"`

	// OPC Request Id
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// e-Tag
	Etag *string `presentIn:"header" name:"etag"`
}

func (response UpdateIncidentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateIncidentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
