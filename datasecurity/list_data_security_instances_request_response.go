// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datasecurity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDataSecurityInstancesRequest wrapper for the ListDataSecurityInstances operation
type ListDataSecurityInstancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState DataSecurityInstanceSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataSecurityInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListDataSecurityInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSecurityInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSecurityInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSecurityInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDataSecurityInstancesResponse wrapper for the ListDataSecurityInstances operation
type ListDataSecurityInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DataSecurityInstanceSummary instances
	Items []DataSecurityInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `DataSecurityInstance`s. If this header appears in the response, then this
	// is a partial list of data security instances. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of data security nstances.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataSecurityInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSecurityInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSecurityInstancesSortOrderEnum Enum with underlying type: string
type ListDataSecurityInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListDataSecurityInstancesSortOrderEnum
const (
	ListDataSecurityInstancesSortOrderAsc  ListDataSecurityInstancesSortOrderEnum = "ASC"
	ListDataSecurityInstancesSortOrderDesc ListDataSecurityInstancesSortOrderEnum = "DESC"
)

var mappingListDataSecurityInstancesSortOrder = map[string]ListDataSecurityInstancesSortOrderEnum{
	"ASC":  ListDataSecurityInstancesSortOrderAsc,
	"DESC": ListDataSecurityInstancesSortOrderDesc,
}

// GetListDataSecurityInstancesSortOrderEnumValues Enumerates the set of values for ListDataSecurityInstancesSortOrderEnum
func GetListDataSecurityInstancesSortOrderEnumValues() []ListDataSecurityInstancesSortOrderEnum {
	values := make([]ListDataSecurityInstancesSortOrderEnum, 0)
	for _, v := range mappingListDataSecurityInstancesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDataSecurityInstancesSortByEnum Enum with underlying type: string
type ListDataSecurityInstancesSortByEnum string

// Set of constants representing the allowable values for ListDataSecurityInstancesSortByEnum
const (
	ListDataSecurityInstancesSortByTimecreated ListDataSecurityInstancesSortByEnum = "TIMECREATED"
	ListDataSecurityInstancesSortByDisplayname ListDataSecurityInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListDataSecurityInstancesSortBy = map[string]ListDataSecurityInstancesSortByEnum{
	"TIMECREATED": ListDataSecurityInstancesSortByTimecreated,
	"DISPLAYNAME": ListDataSecurityInstancesSortByDisplayname,
}

// GetListDataSecurityInstancesSortByEnumValues Enumerates the set of values for ListDataSecurityInstancesSortByEnum
func GetListDataSecurityInstancesSortByEnumValues() []ListDataSecurityInstancesSortByEnum {
	values := make([]ListDataSecurityInstancesSortByEnum, 0)
	for _, v := range mappingListDataSecurityInstancesSortBy {
		values = append(values, v)
	}
	return values
}
