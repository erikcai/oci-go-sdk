// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"net/http"
)

// ListRoverNodesRequest wrapper for the ListRoverNodes operation
type ListRoverNodesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListRoverNodesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRoverNodesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListRoverNodesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRoverNodesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRoverNodesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRoverNodesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRoverNodesResponse wrapper for the ListRoverNodes operation
type ListRoverNodesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RoverNodeCollection instances
	RoverNodeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRoverNodesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRoverNodesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRoverNodesLifecycleStateEnum Enum with underlying type: string
type ListRoverNodesLifecycleStateEnum string

// Set of constants representing the allowable values for ListRoverNodesLifecycleStateEnum
const (
	ListRoverNodesLifecycleStateCreating ListRoverNodesLifecycleStateEnum = "CREATING"
	ListRoverNodesLifecycleStateUpdating ListRoverNodesLifecycleStateEnum = "UPDATING"
	ListRoverNodesLifecycleStateActive   ListRoverNodesLifecycleStateEnum = "ACTIVE"
	ListRoverNodesLifecycleStateDeleting ListRoverNodesLifecycleStateEnum = "DELETING"
	ListRoverNodesLifecycleStateDeleted  ListRoverNodesLifecycleStateEnum = "DELETED"
	ListRoverNodesLifecycleStateFailed   ListRoverNodesLifecycleStateEnum = "FAILED"
)

var mappingListRoverNodesLifecycleState = map[string]ListRoverNodesLifecycleStateEnum{
	"CREATING": ListRoverNodesLifecycleStateCreating,
	"UPDATING": ListRoverNodesLifecycleStateUpdating,
	"ACTIVE":   ListRoverNodesLifecycleStateActive,
	"DELETING": ListRoverNodesLifecycleStateDeleting,
	"DELETED":  ListRoverNodesLifecycleStateDeleted,
	"FAILED":   ListRoverNodesLifecycleStateFailed,
}

// GetListRoverNodesLifecycleStateEnumValues Enumerates the set of values for ListRoverNodesLifecycleStateEnum
func GetListRoverNodesLifecycleStateEnumValues() []ListRoverNodesLifecycleStateEnum {
	values := make([]ListRoverNodesLifecycleStateEnum, 0)
	for _, v := range mappingListRoverNodesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListRoverNodesSortOrderEnum Enum with underlying type: string
type ListRoverNodesSortOrderEnum string

// Set of constants representing the allowable values for ListRoverNodesSortOrderEnum
const (
	ListRoverNodesSortOrderAsc  ListRoverNodesSortOrderEnum = "ASC"
	ListRoverNodesSortOrderDesc ListRoverNodesSortOrderEnum = "DESC"
)

var mappingListRoverNodesSortOrder = map[string]ListRoverNodesSortOrderEnum{
	"ASC":  ListRoverNodesSortOrderAsc,
	"DESC": ListRoverNodesSortOrderDesc,
}

// GetListRoverNodesSortOrderEnumValues Enumerates the set of values for ListRoverNodesSortOrderEnum
func GetListRoverNodesSortOrderEnumValues() []ListRoverNodesSortOrderEnum {
	values := make([]ListRoverNodesSortOrderEnum, 0)
	for _, v := range mappingListRoverNodesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListRoverNodesSortByEnum Enum with underlying type: string
type ListRoverNodesSortByEnum string

// Set of constants representing the allowable values for ListRoverNodesSortByEnum
const (
	ListRoverNodesSortByTimecreated ListRoverNodesSortByEnum = "timeCreated"
	ListRoverNodesSortByDisplayname ListRoverNodesSortByEnum = "displayName"
)

var mappingListRoverNodesSortBy = map[string]ListRoverNodesSortByEnum{
	"timeCreated": ListRoverNodesSortByTimecreated,
	"displayName": ListRoverNodesSortByDisplayname,
}

// GetListRoverNodesSortByEnumValues Enumerates the set of values for ListRoverNodesSortByEnum
func GetListRoverNodesSortByEnumValues() []ListRoverNodesSortByEnum {
	values := make([]ListRoverNodesSortByEnum, 0)
	for _, v := range mappingListRoverNodesSortBy {
		values = append(values, v)
	}
	return values
}
