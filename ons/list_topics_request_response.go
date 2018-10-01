// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTopicsRequest wrapper for the ListTopics operation
type ListTopicsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. You can provide no more than one sort order. If no params are provided, by default, there is no order.
	SortBy ListTopicsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTopicsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState NotificationTopicSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTopicsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTopicsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTopicsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTopicsResponse wrapper for the ListTopics operation
type ListTopicsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NotificationTopicSummary instances
	Items []NotificationTopicSummary `presentIn:"body"`

	// Pagination token
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTopicsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTopicsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTopicsSortByEnum Enum with underlying type: string
type ListTopicsSortByEnum string

// Set of constants representing the allowable values for ListTopicsSortByEnum
const (
	ListTopicsSortByTimecreated    ListTopicsSortByEnum = "TIMECREATED"
	ListTopicsSortByLifecyclestate ListTopicsSortByEnum = "LIFECYCLESTATE"
)

var mappingListTopicsSortBy = map[string]ListTopicsSortByEnum{
	"TIMECREATED":    ListTopicsSortByTimecreated,
	"LIFECYCLESTATE": ListTopicsSortByLifecyclestate,
}

// GetListTopicsSortByEnumValues Enumerates the set of values for ListTopicsSortByEnum
func GetListTopicsSortByEnumValues() []ListTopicsSortByEnum {
	values := make([]ListTopicsSortByEnum, 0)
	for _, v := range mappingListTopicsSortBy {
		values = append(values, v)
	}
	return values
}

// ListTopicsSortOrderEnum Enum with underlying type: string
type ListTopicsSortOrderEnum string

// Set of constants representing the allowable values for ListTopicsSortOrderEnum
const (
	ListTopicsSortOrderAsc  ListTopicsSortOrderEnum = "ASC"
	ListTopicsSortOrderDesc ListTopicsSortOrderEnum = "DESC"
)

var mappingListTopicsSortOrder = map[string]ListTopicsSortOrderEnum{
	"ASC":  ListTopicsSortOrderAsc,
	"DESC": ListTopicsSortOrderDesc,
}

// GetListTopicsSortOrderEnumValues Enumerates the set of values for ListTopicsSortOrderEnum
func GetListTopicsSortOrderEnumValues() []ListTopicsSortOrderEnum {
	values := make([]ListTopicsSortOrderEnum, 0)
	for _, v := range mappingListTopicsSortOrder {
		values = append(values, v)
	}
	return values
}
