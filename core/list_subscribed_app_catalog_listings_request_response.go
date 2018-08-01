// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSubscribedAppCatalogListingsRequest wrapper for the ListSubscribedAppCatalogListings operation
type ListSubscribedAppCatalogListingsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

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
	SortBy ListSubscribedAppCatalogListingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListSubscribedAppCatalogListingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only the listings that matches the given listing id.
	ListingId *string `mandatory:"false" contributesTo:"query" name:"listingId"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubscribedAppCatalogListingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubscribedAppCatalogListingsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubscribedAppCatalogListingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSubscribedAppCatalogListingsResponse wrapper for the ListSubscribedAppCatalogListings operation
type ListSubscribedAppCatalogListingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SubscribedAppCatalogListingSummary instances
	Items []SubscribedAppCatalogListingSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSubscribedAppCatalogListingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubscribedAppCatalogListingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubscribedAppCatalogListingsSortByEnum Enum with underlying type: string
type ListSubscribedAppCatalogListingsSortByEnum string

// Set of constants representing the allowable values for ListSubscribedAppCatalogListingsSortBy
const (
	ListSubscribedAppCatalogListingsSortByTimecreated ListSubscribedAppCatalogListingsSortByEnum = "TIMECREATED"
	ListSubscribedAppCatalogListingsSortByDisplayname ListSubscribedAppCatalogListingsSortByEnum = "DISPLAYNAME"
)

var mappingListSubscribedAppCatalogListingsSortBy = map[string]ListSubscribedAppCatalogListingsSortByEnum{
	"TIMECREATED": ListSubscribedAppCatalogListingsSortByTimecreated,
	"DISPLAYNAME": ListSubscribedAppCatalogListingsSortByDisplayname,
}

// GetListSubscribedAppCatalogListingsSortByEnumValues Enumerates the set of values for ListSubscribedAppCatalogListingsSortBy
func GetListSubscribedAppCatalogListingsSortByEnumValues() []ListSubscribedAppCatalogListingsSortByEnum {
	values := make([]ListSubscribedAppCatalogListingsSortByEnum, 0)
	for _, v := range mappingListSubscribedAppCatalogListingsSortBy {
		values = append(values, v)
	}
	return values
}

// ListSubscribedAppCatalogListingsSortOrderEnum Enum with underlying type: string
type ListSubscribedAppCatalogListingsSortOrderEnum string

// Set of constants representing the allowable values for ListSubscribedAppCatalogListingsSortOrder
const (
	ListSubscribedAppCatalogListingsSortOrderAsc  ListSubscribedAppCatalogListingsSortOrderEnum = "ASC"
	ListSubscribedAppCatalogListingsSortOrderDesc ListSubscribedAppCatalogListingsSortOrderEnum = "DESC"
)

var mappingListSubscribedAppCatalogListingsSortOrder = map[string]ListSubscribedAppCatalogListingsSortOrderEnum{
	"ASC":  ListSubscribedAppCatalogListingsSortOrderAsc,
	"DESC": ListSubscribedAppCatalogListingsSortOrderDesc,
}

// GetListSubscribedAppCatalogListingsSortOrderEnumValues Enumerates the set of values for ListSubscribedAppCatalogListingsSortOrder
func GetListSubscribedAppCatalogListingsSortOrderEnumValues() []ListSubscribedAppCatalogListingsSortOrderEnum {
	values := make([]ListSubscribedAppCatalogListingsSortOrderEnum, 0)
	for _, v := range mappingListSubscribedAppCatalogListingsSortOrder {
		values = append(values, v)
	}
	return values
}
