// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package sch

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListServiceConnectorsRequest wrapper for the ListServiceConnectors operation
type ListServiceConnectorsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListServiceConnectorsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListServiceConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListServiceConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListServiceConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServiceConnectorsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServiceConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListServiceConnectorsResponse wrapper for the ListServiceConnectors operation
type ListServiceConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServiceConnectorCollection instances
	ServiceConnectorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListServiceConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServiceConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServiceConnectorsLifecycleStateEnum Enum with underlying type: string
type ListServiceConnectorsLifecycleStateEnum string

// Set of constants representing the allowable values for ListServiceConnectorsLifecycleStateEnum
const (
	ListServiceConnectorsLifecycleStateCreating ListServiceConnectorsLifecycleStateEnum = "CREATING"
	ListServiceConnectorsLifecycleStateUpdating ListServiceConnectorsLifecycleStateEnum = "UPDATING"
	ListServiceConnectorsLifecycleStateActive   ListServiceConnectorsLifecycleStateEnum = "ACTIVE"
	ListServiceConnectorsLifecycleStateInactive ListServiceConnectorsLifecycleStateEnum = "INACTIVE"
	ListServiceConnectorsLifecycleStateDeleting ListServiceConnectorsLifecycleStateEnum = "DELETING"
	ListServiceConnectorsLifecycleStateDeleted  ListServiceConnectorsLifecycleStateEnum = "DELETED"
	ListServiceConnectorsLifecycleStateFailed   ListServiceConnectorsLifecycleStateEnum = "FAILED"
)

var mappingListServiceConnectorsLifecycleState = map[string]ListServiceConnectorsLifecycleStateEnum{
	"CREATING": ListServiceConnectorsLifecycleStateCreating,
	"UPDATING": ListServiceConnectorsLifecycleStateUpdating,
	"ACTIVE":   ListServiceConnectorsLifecycleStateActive,
	"INACTIVE": ListServiceConnectorsLifecycleStateInactive,
	"DELETING": ListServiceConnectorsLifecycleStateDeleting,
	"DELETED":  ListServiceConnectorsLifecycleStateDeleted,
	"FAILED":   ListServiceConnectorsLifecycleStateFailed,
}

// GetListServiceConnectorsLifecycleStateEnumValues Enumerates the set of values for ListServiceConnectorsLifecycleStateEnum
func GetListServiceConnectorsLifecycleStateEnumValues() []ListServiceConnectorsLifecycleStateEnum {
	values := make([]ListServiceConnectorsLifecycleStateEnum, 0)
	for _, v := range mappingListServiceConnectorsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListServiceConnectorsSortOrderEnum Enum with underlying type: string
type ListServiceConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListServiceConnectorsSortOrderEnum
const (
	ListServiceConnectorsSortOrderAsc  ListServiceConnectorsSortOrderEnum = "ASC"
	ListServiceConnectorsSortOrderDesc ListServiceConnectorsSortOrderEnum = "DESC"
)

var mappingListServiceConnectorsSortOrder = map[string]ListServiceConnectorsSortOrderEnum{
	"ASC":  ListServiceConnectorsSortOrderAsc,
	"DESC": ListServiceConnectorsSortOrderDesc,
}

// GetListServiceConnectorsSortOrderEnumValues Enumerates the set of values for ListServiceConnectorsSortOrderEnum
func GetListServiceConnectorsSortOrderEnumValues() []ListServiceConnectorsSortOrderEnum {
	values := make([]ListServiceConnectorsSortOrderEnum, 0)
	for _, v := range mappingListServiceConnectorsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListServiceConnectorsSortByEnum Enum with underlying type: string
type ListServiceConnectorsSortByEnum string

// Set of constants representing the allowable values for ListServiceConnectorsSortByEnum
const (
	ListServiceConnectorsSortByTimecreated ListServiceConnectorsSortByEnum = "timeCreated"
	ListServiceConnectorsSortByDisplayname ListServiceConnectorsSortByEnum = "displayName"
)

var mappingListServiceConnectorsSortBy = map[string]ListServiceConnectorsSortByEnum{
	"timeCreated": ListServiceConnectorsSortByTimecreated,
	"displayName": ListServiceConnectorsSortByDisplayname,
}

// GetListServiceConnectorsSortByEnumValues Enumerates the set of values for ListServiceConnectorsSortByEnum
func GetListServiceConnectorsSortByEnumValues() []ListServiceConnectorsSortByEnum {
	values := make([]ListServiceConnectorsSortByEnum, 0)
	for _, v := range mappingListServiceConnectorsSortBy {
		values = append(values, v)
	}
	return values
}
