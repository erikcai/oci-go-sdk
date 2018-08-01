// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// SubscribeAppCatalogListingRequest wrapper for the SubscribeAppCatalogListing operation
type SubscribeAppCatalogListingRequest struct {

	// The OCID of the listing.
	ListingId *string `mandatory:"true" contributesTo:"path" name:"listingId"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Listing Resource Version.
	ResourceVersion *string `mandatory:"true" contributesTo:"path" name:"resourceVersion"`

	// Request for subscribe to listing resource version.
	SubscribeListingResourceVersionDetails AppCatalogListingResourceVersionAgreements `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SubscribeAppCatalogListingRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SubscribeAppCatalogListingRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SubscribeAppCatalogListingRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SubscribeAppCatalogListingResponse wrapper for the SubscribeAppCatalogListing operation
type SubscribeAppCatalogListingResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SubscribeAppCatalogListingResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SubscribeAppCatalogListingResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
