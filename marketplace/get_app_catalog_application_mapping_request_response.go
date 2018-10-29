// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetAppCatalogApplicationMappingRequest wrapper for the GetAppCatalogApplicationMapping operation
type GetAppCatalogApplicationMappingRequest struct {

	// Unique ID of the application.
	MarketplaceApplicationId *string `mandatory:"true" contributesTo:"query" name:"marketplaceApplicationId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAppCatalogApplicationMappingRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAppCatalogApplicationMappingRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAppCatalogApplicationMappingRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAppCatalogApplicationMappingResponse wrapper for the GetAppCatalogApplicationMapping operation
type GetAppCatalogApplicationMappingResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AppCatalogApplicationMap instance
	AppCatalogApplicationMap `presentIn:"body"`

	// Unique identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAppCatalogApplicationMappingResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAppCatalogApplicationMappingResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
