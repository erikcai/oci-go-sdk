// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListConfigurationRevisionsRequest wrapper for the ListConfigurationRevisions operation
type ListConfigurationRevisionsRequest struct {

	// The OCID of the MySQLaaS Configuration.
	ConfigurationId *string `mandatory:"true" contributesTo:"path" name:"configurationId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending.
	SortBy ListConfigurationRevisionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConfigurationRevisionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous "List" call. For information about pagination, see List
	// Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConfigurationRevisionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConfigurationRevisionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConfigurationRevisionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListConfigurationRevisionsResponse wrapper for the ListConfigurationRevisions operation
type ListConfigurationRevisionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ConfigurationRevisionSummary instances
	Items []ConfigurationRevisionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a specific request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Opaque token representing the next page of results.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConfigurationRevisionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConfigurationRevisionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConfigurationRevisionsSortByEnum Enum with underlying type: string
type ListConfigurationRevisionsSortByEnum string

// Set of constants representing the allowable values for ListConfigurationRevisionsSortByEnum
const (
	ListConfigurationRevisionsSortByTimecreated ListConfigurationRevisionsSortByEnum = "timeCreated"
)

var mappingListConfigurationRevisionsSortBy = map[string]ListConfigurationRevisionsSortByEnum{
	"timeCreated": ListConfigurationRevisionsSortByTimecreated,
}

// GetListConfigurationRevisionsSortByEnumValues Enumerates the set of values for ListConfigurationRevisionsSortByEnum
func GetListConfigurationRevisionsSortByEnumValues() []ListConfigurationRevisionsSortByEnum {
	values := make([]ListConfigurationRevisionsSortByEnum, 0)
	for _, v := range mappingListConfigurationRevisionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListConfigurationRevisionsSortOrderEnum Enum with underlying type: string
type ListConfigurationRevisionsSortOrderEnum string

// Set of constants representing the allowable values for ListConfigurationRevisionsSortOrderEnum
const (
	ListConfigurationRevisionsSortOrderAsc  ListConfigurationRevisionsSortOrderEnum = "ASC"
	ListConfigurationRevisionsSortOrderDesc ListConfigurationRevisionsSortOrderEnum = "DESC"
)

var mappingListConfigurationRevisionsSortOrder = map[string]ListConfigurationRevisionsSortOrderEnum{
	"ASC":  ListConfigurationRevisionsSortOrderAsc,
	"DESC": ListConfigurationRevisionsSortOrderDesc,
}

// GetListConfigurationRevisionsSortOrderEnumValues Enumerates the set of values for ListConfigurationRevisionsSortOrderEnum
func GetListConfigurationRevisionsSortOrderEnumValues() []ListConfigurationRevisionsSortOrderEnum {
	values := make([]ListConfigurationRevisionsSortOrderEnum, 0)
	for _, v := range mappingListConfigurationRevisionsSortOrder {
		values = append(values, v)
	}
	return values
}
