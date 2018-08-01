// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAppCatalogListingsRequest wrapper for the ListAppCatalogListings operation
type ListAppCatalogListingsRequest struct {

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListAppCatalogListingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListAppCatalogListingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only the publisher that matches the given publisher name exactly.
	PublisherName *string `mandatory:"false" contributesTo:"query" name:"publisherName"`

	// A filter to return only publishers that match the given publisher type exactly. Valid types are OCI, ORACLE, TRUSTED, STANDARD.
	PublisherType *string `mandatory:"false" contributesTo:"query" name:"publisherType"`

	// A filter to return only the listing that matches the given listing name exactly.
	ListingName *string `mandatory:"false" contributesTo:"query" name:"listingName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAppCatalogListingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAppCatalogListingsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAppCatalogListingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAppCatalogListingsResponse wrapper for the ListAppCatalogListings operation
type ListAppCatalogListingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AppCatalogListingSummary instances
	Items []AppCatalogListingSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAppCatalogListingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAppCatalogListingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAppCatalogListingsSortByEnum Enum with underlying type: string
type ListAppCatalogListingsSortByEnum string

// Set of constants representing the allowable values for ListAppCatalogListingsSortBy
const (
	ListAppCatalogListingsSortByTimecreated ListAppCatalogListingsSortByEnum = "TIMECREATED"
	ListAppCatalogListingsSortByDisplayname ListAppCatalogListingsSortByEnum = "DISPLAYNAME"
)

var mappingListAppCatalogListingsSortBy = map[string]ListAppCatalogListingsSortByEnum{
	"TIMECREATED": ListAppCatalogListingsSortByTimecreated,
	"DISPLAYNAME": ListAppCatalogListingsSortByDisplayname,
}

// GetListAppCatalogListingsSortByEnumValues Enumerates the set of values for ListAppCatalogListingsSortBy
func GetListAppCatalogListingsSortByEnumValues() []ListAppCatalogListingsSortByEnum {
	values := make([]ListAppCatalogListingsSortByEnum, 0)
	for _, v := range mappingListAppCatalogListingsSortBy {
		values = append(values, v)
	}
	return values
}

// ListAppCatalogListingsSortOrderEnum Enum with underlying type: string
type ListAppCatalogListingsSortOrderEnum string

// Set of constants representing the allowable values for ListAppCatalogListingsSortOrder
const (
	ListAppCatalogListingsSortOrderAsc  ListAppCatalogListingsSortOrderEnum = "ASC"
	ListAppCatalogListingsSortOrderDesc ListAppCatalogListingsSortOrderEnum = "DESC"
)

var mappingListAppCatalogListingsSortOrder = map[string]ListAppCatalogListingsSortOrderEnum{
	"ASC":  ListAppCatalogListingsSortOrderAsc,
	"DESC": ListAppCatalogListingsSortOrderDesc,
}

// GetListAppCatalogListingsSortOrderEnumValues Enumerates the set of values for ListAppCatalogListingsSortOrder
func GetListAppCatalogListingsSortOrderEnumValues() []ListAppCatalogListingsSortOrderEnum {
	values := make([]ListAppCatalogListingsSortOrderEnum, 0)
	for _, v := range mappingListAppCatalogListingsSortOrder {
		values = append(values, v)
	}
	return values
}
