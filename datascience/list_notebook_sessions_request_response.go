// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListNotebookSessionsRequest wrapper for the ListNotebookSessions operation
type ListNotebookSessionsRequest struct {

	// <b>Filter</b> results by the OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID. Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by the OCID of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListNotebookSessionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filters results by the OCID of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending.
	SortOrder ListNotebookSessionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by display name, results are
	// shown in ascending order.
	SortBy ListNotebookSessionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNotebookSessionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNotebookSessionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNotebookSessionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListNotebookSessionsResponse wrapper for the ListNotebookSessions operation
type ListNotebookSessionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NotebookSessionSummary instances
	Items []NotebookSessionSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list in backword direction, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListNotebookSessionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNotebookSessionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNotebookSessionsLifecycleStateEnum Enum with underlying type: string
type ListNotebookSessionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListNotebookSessionsLifecycleStateEnum
const (
	ListNotebookSessionsLifecycleStateCreating ListNotebookSessionsLifecycleStateEnum = "CREATING"
	ListNotebookSessionsLifecycleStateActive   ListNotebookSessionsLifecycleStateEnum = "ACTIVE"
	ListNotebookSessionsLifecycleStateDeleting ListNotebookSessionsLifecycleStateEnum = "DELETING"
	ListNotebookSessionsLifecycleStateDeleted  ListNotebookSessionsLifecycleStateEnum = "DELETED"
	ListNotebookSessionsLifecycleStateFailed   ListNotebookSessionsLifecycleStateEnum = "FAILED"
	ListNotebookSessionsLifecycleStateInactive ListNotebookSessionsLifecycleStateEnum = "INACTIVE"
	ListNotebookSessionsLifecycleStateUpdating ListNotebookSessionsLifecycleStateEnum = "UPDATING"
)

var mappingListNotebookSessionsLifecycleState = map[string]ListNotebookSessionsLifecycleStateEnum{
	"CREATING": ListNotebookSessionsLifecycleStateCreating,
	"ACTIVE":   ListNotebookSessionsLifecycleStateActive,
	"DELETING": ListNotebookSessionsLifecycleStateDeleting,
	"DELETED":  ListNotebookSessionsLifecycleStateDeleted,
	"FAILED":   ListNotebookSessionsLifecycleStateFailed,
	"INACTIVE": ListNotebookSessionsLifecycleStateInactive,
	"UPDATING": ListNotebookSessionsLifecycleStateUpdating,
}

// GetListNotebookSessionsLifecycleStateEnumValues Enumerates the set of values for ListNotebookSessionsLifecycleStateEnum
func GetListNotebookSessionsLifecycleStateEnumValues() []ListNotebookSessionsLifecycleStateEnum {
	values := make([]ListNotebookSessionsLifecycleStateEnum, 0)
	for _, v := range mappingListNotebookSessionsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListNotebookSessionsSortOrderEnum Enum with underlying type: string
type ListNotebookSessionsSortOrderEnum string

// Set of constants representing the allowable values for ListNotebookSessionsSortOrderEnum
const (
	ListNotebookSessionsSortOrderAsc  ListNotebookSessionsSortOrderEnum = "ASC"
	ListNotebookSessionsSortOrderDesc ListNotebookSessionsSortOrderEnum = "DESC"
)

var mappingListNotebookSessionsSortOrder = map[string]ListNotebookSessionsSortOrderEnum{
	"ASC":  ListNotebookSessionsSortOrderAsc,
	"DESC": ListNotebookSessionsSortOrderDesc,
}

// GetListNotebookSessionsSortOrderEnumValues Enumerates the set of values for ListNotebookSessionsSortOrderEnum
func GetListNotebookSessionsSortOrderEnumValues() []ListNotebookSessionsSortOrderEnum {
	values := make([]ListNotebookSessionsSortOrderEnum, 0)
	for _, v := range mappingListNotebookSessionsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListNotebookSessionsSortByEnum Enum with underlying type: string
type ListNotebookSessionsSortByEnum string

// Set of constants representing the allowable values for ListNotebookSessionsSortByEnum
const (
	ListNotebookSessionsSortByTimecreated ListNotebookSessionsSortByEnum = "timeCreated"
	ListNotebookSessionsSortByDisplayname ListNotebookSessionsSortByEnum = "displayName"
)

var mappingListNotebookSessionsSortBy = map[string]ListNotebookSessionsSortByEnum{
	"timeCreated": ListNotebookSessionsSortByTimecreated,
	"displayName": ListNotebookSessionsSortByDisplayname,
}

// GetListNotebookSessionsSortByEnumValues Enumerates the set of values for ListNotebookSessionsSortByEnum
func GetListNotebookSessionsSortByEnumValues() []ListNotebookSessionsSortByEnum {
	values := make([]ListNotebookSessionsSortByEnum, 0)
	for _, v := range mappingListNotebookSessionsSortBy {
		values = append(values, v)
	}
	return values
}
