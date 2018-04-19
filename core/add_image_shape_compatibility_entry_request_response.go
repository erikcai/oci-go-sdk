// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// AddImageShapeCompatibilityEntryRequest wrapper for the AddImageShapeCompatibilityEntry operation
type AddImageShapeCompatibilityEntryRequest struct {

	// The OCID of the image.
	ImageId *string `mandatory:"true" contributesTo:"path" name:"imageId"`

	// Shape name.
	ShapeName *string `mandatory:"true" contributesTo:"path" name:"shapeName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request AddImageShapeCompatibilityEntryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request AddImageShapeCompatibilityEntryRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request AddImageShapeCompatibilityEntryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// AddImageShapeCompatibilityEntryResponse wrapper for the AddImageShapeCompatibilityEntry operation
type AddImageShapeCompatibilityEntryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ImageShapeCompatibilityEntry instance
	ImageShapeCompatibilityEntry `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response AddImageShapeCompatibilityEntryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response AddImageShapeCompatibilityEntryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}