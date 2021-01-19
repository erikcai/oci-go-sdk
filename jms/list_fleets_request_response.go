// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"github.com/oracle/oci-go-sdk/v33/common"
	"net/http"
)

// ListFleetsRequest wrapper for the ListFleets operation
type ListFleetsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only fleets that have the specified identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources their lifecycleState matches the specified lifecycleState.
	LifecycleState ListFleetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that entirely match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListFleetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort fleets. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListFleetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFleetsResponse wrapper for the ListFleets operation
type ListFleetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FleetCollection instances
	FleetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetsLifecycleStateEnum Enum with underlying type: string
type ListFleetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFleetsLifecycleStateEnum
const (
	ListFleetsLifecycleStateCreating ListFleetsLifecycleStateEnum = "CREATING"
	ListFleetsLifecycleStateUpdating ListFleetsLifecycleStateEnum = "UPDATING"
	ListFleetsLifecycleStateActive   ListFleetsLifecycleStateEnum = "ACTIVE"
	ListFleetsLifecycleStateDeleting ListFleetsLifecycleStateEnum = "DELETING"
	ListFleetsLifecycleStateDeleted  ListFleetsLifecycleStateEnum = "DELETED"
	ListFleetsLifecycleStateFailed   ListFleetsLifecycleStateEnum = "FAILED"
)

var mappingListFleetsLifecycleState = map[string]ListFleetsLifecycleStateEnum{
	"CREATING": ListFleetsLifecycleStateCreating,
	"UPDATING": ListFleetsLifecycleStateUpdating,
	"ACTIVE":   ListFleetsLifecycleStateActive,
	"DELETING": ListFleetsLifecycleStateDeleting,
	"DELETED":  ListFleetsLifecycleStateDeleted,
	"FAILED":   ListFleetsLifecycleStateFailed,
}

// GetListFleetsLifecycleStateEnumValues Enumerates the set of values for ListFleetsLifecycleStateEnum
func GetListFleetsLifecycleStateEnumValues() []ListFleetsLifecycleStateEnum {
	values := make([]ListFleetsLifecycleStateEnum, 0)
	for _, v := range mappingListFleetsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListFleetsSortOrderEnum Enum with underlying type: string
type ListFleetsSortOrderEnum string

// Set of constants representing the allowable values for ListFleetsSortOrderEnum
const (
	ListFleetsSortOrderAsc  ListFleetsSortOrderEnum = "ASC"
	ListFleetsSortOrderDesc ListFleetsSortOrderEnum = "DESC"
)

var mappingListFleetsSortOrder = map[string]ListFleetsSortOrderEnum{
	"ASC":  ListFleetsSortOrderAsc,
	"DESC": ListFleetsSortOrderDesc,
}

// GetListFleetsSortOrderEnumValues Enumerates the set of values for ListFleetsSortOrderEnum
func GetListFleetsSortOrderEnumValues() []ListFleetsSortOrderEnum {
	values := make([]ListFleetsSortOrderEnum, 0)
	for _, v := range mappingListFleetsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListFleetsSortByEnum Enum with underlying type: string
type ListFleetsSortByEnum string

// Set of constants representing the allowable values for ListFleetsSortByEnum
const (
	ListFleetsSortByTimecreated ListFleetsSortByEnum = "timeCreated"
	ListFleetsSortByDisplayname ListFleetsSortByEnum = "displayName"
)

var mappingListFleetsSortBy = map[string]ListFleetsSortByEnum{
	"timeCreated": ListFleetsSortByTimecreated,
	"displayName": ListFleetsSortByDisplayname,
}

// GetListFleetsSortByEnumValues Enumerates the set of values for ListFleetsSortByEnum
func GetListFleetsSortByEnumValues() []ListFleetsSortByEnum {
	values := make([]ListFleetsSortByEnum, 0)
	for _, v := range mappingListFleetsSortBy {
		values = append(values, v)
	}
	return values
}
