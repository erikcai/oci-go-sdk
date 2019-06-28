// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListIncidentsRequest wrapper for the ListIncidents operation
type ListIncidentsRequest struct {
	Csi *string `mandatory:"true" contributesTo:"query" name:"csi"`

	Email *string `mandatory:"false" contributesTo:"query" name:"email"`

	Limit *string `mandatory:"false" contributesTo:"query" name:"limit"`

	SortBy *string `mandatory:"false" contributesTo:"query" name:"sortBy"`

	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIncidentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIncidentsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIncidentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListIncidentsResponse wrapper for the ListIncidents operation
type ListIncidentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []Incident instances
	Items []Incident `presentIn:"body"`

	// OPC Request Id
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// OPC next page
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// e-Tag
	Etag *string `presentIn:"header" name:"etag"`
}

func (response ListIncidentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIncidentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
