// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package orchestration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListStacksRequest wrapper for the ListStacks operation
type ListStacksRequest struct {

	// May be either a comparment OCID to list stacks within a compartment or stack OCID to list a single stack
	Id *string `mandatory:"true" contributesTo:"query" name:"id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState StackLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Display name
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided. Time created is default ordered as descending. Display name is default ordered as ascending.
	SortBy ListStacksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListStacksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call. For information about pagination, see
	// List Pagination (https://docs.us-phoenix-1.oraclecloud.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call. For information about
	// pagination, see List Pagination (https://docs.us-phoenix-1.oraclecloud.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStacksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStacksRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStacksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListStacksResponse wrapper for the ListStacks operation
type ListStacksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []StackSummary instances
	Items []StackSummary `presentIn:"body"`

	// Unique identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination token for the next page of items
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStacksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStacksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStacksSortByEnum Enum with underlying type: string
type ListStacksSortByEnum string

// Set of constants representing the allowable values for ListStacksSortBy
const (
	ListStacksSortByTimecreated ListStacksSortByEnum = "TIMECREATED"
	ListStacksSortByDisplayname ListStacksSortByEnum = "DISPLAYNAME"
)

var mappingListStacksSortBy = map[string]ListStacksSortByEnum{
	"TIMECREATED": ListStacksSortByTimecreated,
	"DISPLAYNAME": ListStacksSortByDisplayname,
}

// GetListStacksSortByEnumValues Enumerates the set of values for ListStacksSortBy
func GetListStacksSortByEnumValues() []ListStacksSortByEnum {
	values := make([]ListStacksSortByEnum, 0)
	for _, v := range mappingListStacksSortBy {
		values = append(values, v)
	}
	return values
}

// ListStacksSortOrderEnum Enum with underlying type: string
type ListStacksSortOrderEnum string

// Set of constants representing the allowable values for ListStacksSortOrder
const (
	ListStacksSortOrderAsc  ListStacksSortOrderEnum = "ASC"
	ListStacksSortOrderDesc ListStacksSortOrderEnum = "DESC"
)

var mappingListStacksSortOrder = map[string]ListStacksSortOrderEnum{
	"ASC":  ListStacksSortOrderAsc,
	"DESC": ListStacksSortOrderDesc,
}

// GetListStacksSortOrderEnumValues Enumerates the set of values for ListStacksSortOrder
func GetListStacksSortOrderEnumValues() []ListStacksSortOrderEnum {
	values := make([]ListStacksSortOrderEnum, 0)
	for _, v := range mappingListStacksSortOrder {
		values = append(values, v)
	}
	return values
}
