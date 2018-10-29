// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListProductFiltersRequest wrapper for the ListProductFilters operation
type ListProductFiltersRequest struct {

	// Unique Identifier for which productcode is required.
	ProductCode *string `mandatory:"true" contributesTo:"path" name:"productCode"`

	// Limit tells how many records to be returned.
	// Limit should be greater than zero and less than or equal to hundred (default=30).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-prev-page` or `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Time created is default ordered as descending. Display name is default ordered as ascending.
	SortBy ListProductFiltersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListProductFiltersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProductFiltersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProductFiltersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProductFiltersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListProductFiltersResponse wrapper for the ListProductFilters operation
type ListProductFiltersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ProductFilterSummary instances
	Items []ProductFilterSummary `presentIn:"body"`

	// Unique identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination token for the next page of items
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination token for the previous page of items
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListProductFiltersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProductFiltersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProductFiltersSortByEnum Enum with underlying type: string
type ListProductFiltersSortByEnum string

// Set of constants representing the allowable values for ListProductFiltersSortByEnum
const (
	ListProductFiltersSortByTimeCreated ListProductFiltersSortByEnum = "TIME_CREATED"
	ListProductFiltersSortByDisplayname ListProductFiltersSortByEnum = "DISPLAYNAME"
)

var mappingListProductFiltersSortBy = map[string]ListProductFiltersSortByEnum{
	"TIME_CREATED": ListProductFiltersSortByTimeCreated,
	"DISPLAYNAME":  ListProductFiltersSortByDisplayname,
}

// GetListProductFiltersSortByEnumValues Enumerates the set of values for ListProductFiltersSortByEnum
func GetListProductFiltersSortByEnumValues() []ListProductFiltersSortByEnum {
	values := make([]ListProductFiltersSortByEnum, 0)
	for _, v := range mappingListProductFiltersSortBy {
		values = append(values, v)
	}
	return values
}

// ListProductFiltersSortOrderEnum Enum with underlying type: string
type ListProductFiltersSortOrderEnum string

// Set of constants representing the allowable values for ListProductFiltersSortOrderEnum
const (
	ListProductFiltersSortOrderAsc  ListProductFiltersSortOrderEnum = "ASC"
	ListProductFiltersSortOrderDesc ListProductFiltersSortOrderEnum = "DESC"
)

var mappingListProductFiltersSortOrder = map[string]ListProductFiltersSortOrderEnum{
	"ASC":  ListProductFiltersSortOrderAsc,
	"DESC": ListProductFiltersSortOrderDesc,
}

// GetListProductFiltersSortOrderEnumValues Enumerates the set of values for ListProductFiltersSortOrderEnum
func GetListProductFiltersSortOrderEnumValues() []ListProductFiltersSortOrderEnum {
	values := make([]ListProductFiltersSortOrderEnum, 0)
	for _, v := range mappingListProductFiltersSortOrder {
		values = append(values, v)
	}
	return values
}
