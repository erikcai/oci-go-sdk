// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListConsoleConnectionsRequest wrapper for the ListConsoleConnections operation
type ListConsoleConnectionsRequest struct {

	// The database node OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbNodeId *string `mandatory:"true" contributesTo:"path" name:"dbNodeId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListConsoleConnectionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListConsoleConnectionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ConsoleConnectionSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConsoleConnectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConsoleConnectionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConsoleConnectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListConsoleConnectionsResponse wrapper for the ListConsoleConnections operation
type ListConsoleConnectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ConsoleConnectionSummary instances
	Items []ConsoleConnectionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConsoleConnectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConsoleConnectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConsoleConnectionsSortByEnum Enum with underlying type: string
type ListConsoleConnectionsSortByEnum string

// Set of constants representing the allowable values for ListConsoleConnectionsSortByEnum
const (
	ListConsoleConnectionsSortByTimecreated ListConsoleConnectionsSortByEnum = "TIMECREATED"
)

var mappingListConsoleConnectionsSortBy = map[string]ListConsoleConnectionsSortByEnum{
	"TIMECREATED": ListConsoleConnectionsSortByTimecreated,
}

// GetListConsoleConnectionsSortByEnumValues Enumerates the set of values for ListConsoleConnectionsSortByEnum
func GetListConsoleConnectionsSortByEnumValues() []ListConsoleConnectionsSortByEnum {
	values := make([]ListConsoleConnectionsSortByEnum, 0)
	for _, v := range mappingListConsoleConnectionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListConsoleConnectionsSortOrderEnum Enum with underlying type: string
type ListConsoleConnectionsSortOrderEnum string

// Set of constants representing the allowable values for ListConsoleConnectionsSortOrderEnum
const (
	ListConsoleConnectionsSortOrderAsc  ListConsoleConnectionsSortOrderEnum = "ASC"
	ListConsoleConnectionsSortOrderDesc ListConsoleConnectionsSortOrderEnum = "DESC"
)

var mappingListConsoleConnectionsSortOrder = map[string]ListConsoleConnectionsSortOrderEnum{
	"ASC":  ListConsoleConnectionsSortOrderAsc,
	"DESC": ListConsoleConnectionsSortOrderDesc,
}

// GetListConsoleConnectionsSortOrderEnumValues Enumerates the set of values for ListConsoleConnectionsSortOrderEnum
func GetListConsoleConnectionsSortOrderEnumValues() []ListConsoleConnectionsSortOrderEnum {
	values := make([]ListConsoleConnectionsSortOrderEnum, 0)
	for _, v := range mappingListConsoleConnectionsSortOrder {
		values = append(values, v)
	}
	return values
}
