// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListCategoriesRequest wrapper for the ListCategories operation
type ListCategoriesRequest struct {

	// Unique Identifier for which productcode is required.
	ProductCode *string `mandatory:"true" contributesTo:"path" name:"productCode"`

	// This is one of the filter option. It specifies the application categories for which search has to be done. Multiple values can be passed in comma separated.
	ResourceTypes *string `mandatory:"false" contributesTo:"query" name:"resourceTypes"`

	// Limit tells how many records to be returned.
	// Limit should be greater than zero and less than or equal to hundred (default=30).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-prev-page` or `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Time created is default ordered as descending. Display name is default ordered as ascending.
	SortBy ListCategoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListCategoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCategoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCategoriesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCategoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCategoriesResponse wrapper for the ListCategories operation
type ListCategoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CategorySummary instances
	Items []CategorySummary `presentIn:"body"`

	// Unique identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination token for the next page of items
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination token for the previous page of items
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCategoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCategoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCategoriesSortByEnum Enum with underlying type: string
type ListCategoriesSortByEnum string

// Set of constants representing the allowable values for ListCategoriesSortByEnum
const (
	ListCategoriesSortByTimeCreated ListCategoriesSortByEnum = "TIME_CREATED"
	ListCategoriesSortByDisplayname ListCategoriesSortByEnum = "DISPLAYNAME"
)

var mappingListCategoriesSortBy = map[string]ListCategoriesSortByEnum{
	"TIME_CREATED": ListCategoriesSortByTimeCreated,
	"DISPLAYNAME":  ListCategoriesSortByDisplayname,
}

// GetListCategoriesSortByEnumValues Enumerates the set of values for ListCategoriesSortByEnum
func GetListCategoriesSortByEnumValues() []ListCategoriesSortByEnum {
	values := make([]ListCategoriesSortByEnum, 0)
	for _, v := range mappingListCategoriesSortBy {
		values = append(values, v)
	}
	return values
}

// ListCategoriesSortOrderEnum Enum with underlying type: string
type ListCategoriesSortOrderEnum string

// Set of constants representing the allowable values for ListCategoriesSortOrderEnum
const (
	ListCategoriesSortOrderAsc  ListCategoriesSortOrderEnum = "ASC"
	ListCategoriesSortOrderDesc ListCategoriesSortOrderEnum = "DESC"
)

var mappingListCategoriesSortOrder = map[string]ListCategoriesSortOrderEnum{
	"ASC":  ListCategoriesSortOrderAsc,
	"DESC": ListCategoriesSortOrderDesc,
}

// GetListCategoriesSortOrderEnumValues Enumerates the set of values for ListCategoriesSortOrderEnum
func GetListCategoriesSortOrderEnumValues() []ListCategoriesSortOrderEnum {
	values := make([]ListCategoriesSortOrderEnum, 0)
	for _, v := range mappingListCategoriesSortOrder {
		values = append(values, v)
	}
	return values
}
