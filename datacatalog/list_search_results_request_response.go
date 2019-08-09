// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSearchResultsRequest wrapper for the ListSearchResults operation
type ListSearchResultsRequest struct {

	// unique Catalog identifier
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// The information used to create an extended search results.
	QuerySearchDetails SearchQuery `contributesTo:"body"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Immutable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState LifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A search timeout string (for example, timeout=4000ms), bounding the search request to be executed within the
	// specified time value and bail with the hits accumulated up to that point when expired.
	// Defaults to no timeout.
	Timeout *string `mandatory:"false" contributesTo:"query" name:"timeout"`

	// The query string that allows the user to specify a keyword or keyword qualified by fieldname for search criteria.
	// For example , queryString=name:first_name .
	QueryString *string `mandatory:"false" contributesTo:"query" name:"queryString"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListSearchResultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSearchResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSearchResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSearchResultsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSearchResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSearchResultsResponse wrapper for the ListSearchResults operation
type ListSearchResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SearchResultCollection instances
	SearchResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSearchResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSearchResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSearchResultsSortByEnum Enum with underlying type: string
type ListSearchResultsSortByEnum string

// Set of constants representing the allowable values for ListSearchResultsSortByEnum
const (
	ListSearchResultsSortByTimecreated ListSearchResultsSortByEnum = "TIMECREATED"
	ListSearchResultsSortByDisplayname ListSearchResultsSortByEnum = "DISPLAYNAME"
)

var mappingListSearchResultsSortBy = map[string]ListSearchResultsSortByEnum{
	"TIMECREATED": ListSearchResultsSortByTimecreated,
	"DISPLAYNAME": ListSearchResultsSortByDisplayname,
}

// GetListSearchResultsSortByEnumValues Enumerates the set of values for ListSearchResultsSortByEnum
func GetListSearchResultsSortByEnumValues() []ListSearchResultsSortByEnum {
	values := make([]ListSearchResultsSortByEnum, 0)
	for _, v := range mappingListSearchResultsSortBy {
		values = append(values, v)
	}
	return values
}

// ListSearchResultsSortOrderEnum Enum with underlying type: string
type ListSearchResultsSortOrderEnum string

// Set of constants representing the allowable values for ListSearchResultsSortOrderEnum
const (
	ListSearchResultsSortOrderAsc  ListSearchResultsSortOrderEnum = "ASC"
	ListSearchResultsSortOrderDesc ListSearchResultsSortOrderEnum = "DESC"
)

var mappingListSearchResultsSortOrder = map[string]ListSearchResultsSortOrderEnum{
	"ASC":  ListSearchResultsSortOrderAsc,
	"DESC": ListSearchResultsSortOrderDesc,
}

// GetListSearchResultsSortOrderEnumValues Enumerates the set of values for ListSearchResultsSortOrderEnum
func GetListSearchResultsSortOrderEnumValues() []ListSearchResultsSortOrderEnum {
	values := make([]ListSearchResultsSortOrderEnum, 0)
	for _, v := range mappingListSearchResultsSortOrder {
		values = append(values, v)
	}
	return values
}
