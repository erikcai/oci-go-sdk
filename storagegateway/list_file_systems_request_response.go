// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListFileSystemsRequest wrapper for the ListFileSystems operation
type ListFileSystemsRequest struct {

	// The storage gateway OCID.
	StorageGatewayId *string `mandatory:"true" contributesTo:"path" name:"storageGatewayId"`

	// The value of the `opc-next-page` response header from the previous "List" call. For information about
	// pagination, see List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The optional field to sort the results by.
	SortBy ListFileSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The optional order in which to sort the results.
	SortOrder ListFileSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The optional filter to apply to the results.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The optional filter to apply to the results.
	LifecycleState LifecycleState `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFileSystemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFileSystemsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFileSystemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFileSystemsResponse wrapper for the ListFileSystems operation
type ListFileSystemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FileSystemSummary instances
	Items []FileSystemSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle
	// about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of list of items. When paging through a list, provide this value as
	// the `page` parameter for the subsequent request to page backwards.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header
	// appears in the response, then there are additional items still to get. Include this
	// value as the `page` parameter for the subsequent GET request. For information about
	// pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFileSystemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFileSystemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFileSystemsSortByEnum Enum with underlying type: string
type ListFileSystemsSortByEnum string

// Set of constants representing the allowable values for ListFileSystemsSortByEnum
const (
	ListFileSystemsSortByTimecreated ListFileSystemsSortByEnum = "TIMECREATED"
	ListFileSystemsSortByDisplayname ListFileSystemsSortByEnum = "DISPLAYNAME"
)

var mappingListFileSystemsSortBy = map[string]ListFileSystemsSortByEnum{
	"TIMECREATED": ListFileSystemsSortByTimecreated,
	"DISPLAYNAME": ListFileSystemsSortByDisplayname,
}

// GetListFileSystemsSortByEnumValues Enumerates the set of values for ListFileSystemsSortByEnum
func GetListFileSystemsSortByEnumValues() []ListFileSystemsSortByEnum {
	values := make([]ListFileSystemsSortByEnum, 0)
	for _, v := range mappingListFileSystemsSortBy {
		values = append(values, v)
	}
	return values
}

// ListFileSystemsSortOrderEnum Enum with underlying type: string
type ListFileSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListFileSystemsSortOrderEnum
const (
	ListFileSystemsSortOrderAsc  ListFileSystemsSortOrderEnum = "ASC"
	ListFileSystemsSortOrderDesc ListFileSystemsSortOrderEnum = "DESC"
)

var mappingListFileSystemsSortOrder = map[string]ListFileSystemsSortOrderEnum{
	"ASC":  ListFileSystemsSortOrderAsc,
	"DESC": ListFileSystemsSortOrderDesc,
}

// GetListFileSystemsSortOrderEnumValues Enumerates the set of values for ListFileSystemsSortOrderEnum
func GetListFileSystemsSortOrderEnumValues() []ListFileSystemsSortOrderEnum {
	values := make([]ListFileSystemsSortOrderEnum, 0)
	for _, v := range mappingListFileSystemsSortOrder {
		values = append(values, v)
	}
	return values
}
