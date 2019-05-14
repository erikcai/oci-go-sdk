// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListOdaInstancesRequest wrapper for the ListOdaInstances operation
type ListOdaInstancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle state of the ODA instance
	LifecycleState ListOdaInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOdaInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListOdaInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Internal use only.
	OpcOboToken *string `mandatory:"false" contributesTo:"header" name:"opc-obo-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOdaInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOdaInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOdaInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOdaInstancesResponse wrapper for the ListOdaInstances operation
type ListOdaInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []OdaInstanceSummary instances
	Items []OdaInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of ODA instances. If this header appears in the response, then this
	// is a partial list of Digital Assistant instances. Include this value as the `page` parameter
	// in a subsequent GET request to get the next batch of ODA instances.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOdaInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOdaInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOdaInstancesLifecycleStateEnum Enum with underlying type: string
type ListOdaInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListOdaInstancesLifecycleStateEnum
const (
	ListOdaInstancesLifecycleStateCreating ListOdaInstancesLifecycleStateEnum = "CREATING"
	ListOdaInstancesLifecycleStateUpdating ListOdaInstancesLifecycleStateEnum = "UPDATING"
	ListOdaInstancesLifecycleStateActive   ListOdaInstancesLifecycleStateEnum = "ACTIVE"
	ListOdaInstancesLifecycleStateInactive ListOdaInstancesLifecycleStateEnum = "INACTIVE"
	ListOdaInstancesLifecycleStateDeleting ListOdaInstancesLifecycleStateEnum = "DELETING"
	ListOdaInstancesLifecycleStateDeleted  ListOdaInstancesLifecycleStateEnum = "DELETED"
	ListOdaInstancesLifecycleStateFailed   ListOdaInstancesLifecycleStateEnum = "FAILED"
)

var mappingListOdaInstancesLifecycleState = map[string]ListOdaInstancesLifecycleStateEnum{
	"CREATING": ListOdaInstancesLifecycleStateCreating,
	"UPDATING": ListOdaInstancesLifecycleStateUpdating,
	"ACTIVE":   ListOdaInstancesLifecycleStateActive,
	"INACTIVE": ListOdaInstancesLifecycleStateInactive,
	"DELETING": ListOdaInstancesLifecycleStateDeleting,
	"DELETED":  ListOdaInstancesLifecycleStateDeleted,
	"FAILED":   ListOdaInstancesLifecycleStateFailed,
}

// GetListOdaInstancesLifecycleStateEnumValues Enumerates the set of values for ListOdaInstancesLifecycleStateEnum
func GetListOdaInstancesLifecycleStateEnumValues() []ListOdaInstancesLifecycleStateEnum {
	values := make([]ListOdaInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListOdaInstancesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListOdaInstancesSortOrderEnum Enum with underlying type: string
type ListOdaInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListOdaInstancesSortOrderEnum
const (
	ListOdaInstancesSortOrderAsc  ListOdaInstancesSortOrderEnum = "ASC"
	ListOdaInstancesSortOrderDesc ListOdaInstancesSortOrderEnum = "DESC"
)

var mappingListOdaInstancesSortOrder = map[string]ListOdaInstancesSortOrderEnum{
	"ASC":  ListOdaInstancesSortOrderAsc,
	"DESC": ListOdaInstancesSortOrderDesc,
}

// GetListOdaInstancesSortOrderEnumValues Enumerates the set of values for ListOdaInstancesSortOrderEnum
func GetListOdaInstancesSortOrderEnumValues() []ListOdaInstancesSortOrderEnum {
	values := make([]ListOdaInstancesSortOrderEnum, 0)
	for _, v := range mappingListOdaInstancesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListOdaInstancesSortByEnum Enum with underlying type: string
type ListOdaInstancesSortByEnum string

// Set of constants representing the allowable values for ListOdaInstancesSortByEnum
const (
	ListOdaInstancesSortByTimecreated ListOdaInstancesSortByEnum = "TIMECREATED"
	ListOdaInstancesSortByDisplayname ListOdaInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListOdaInstancesSortBy = map[string]ListOdaInstancesSortByEnum{
	"TIMECREATED": ListOdaInstancesSortByTimecreated,
	"DISPLAYNAME": ListOdaInstancesSortByDisplayname,
}

// GetListOdaInstancesSortByEnumValues Enumerates the set of values for ListOdaInstancesSortByEnum
func GetListOdaInstancesSortByEnumValues() []ListOdaInstancesSortByEnum {
	values := make([]ListOdaInstancesSortByEnum, 0)
	for _, v := range mappingListOdaInstancesSortBy {
		values = append(values, v)
	}
	return values
}
