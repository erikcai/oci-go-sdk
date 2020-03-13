// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDataSafePrivateEndpointsRequest wrapper for the ListDataSafePrivateEndpoints operation
type ListDataSafePrivateEndpointsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of the VCN in which Data Safe Private Endpoints were provisioned.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ListDataSafePrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataSafePrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListDataSafePrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSafePrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSafePrivateEndpointsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSafePrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDataSafePrivateEndpointsResponse wrapper for the ListDataSafePrivateEndpoints operation
type ListDataSafePrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DataSafePrivateEndpointSummary instances
	Items []DataSafePrivateEndpointSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `DataSafePrivateEndpoints`s. If this header appears in the response, then this
	// is a partial list of data safe PrivateEndpoints. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of data safe PrivateEndpoints.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataSafePrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSafePrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSafePrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsLifecycleStateEnum
const (
	ListDataSafePrivateEndpointsLifecycleStateCreating ListDataSafePrivateEndpointsLifecycleStateEnum = "CREATING"
	ListDataSafePrivateEndpointsLifecycleStateUpdating ListDataSafePrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListDataSafePrivateEndpointsLifecycleStateActive   ListDataSafePrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListDataSafePrivateEndpointsLifecycleStateDeleting ListDataSafePrivateEndpointsLifecycleStateEnum = "DELETING"
	ListDataSafePrivateEndpointsLifecycleStateDeleted  ListDataSafePrivateEndpointsLifecycleStateEnum = "DELETED"
	ListDataSafePrivateEndpointsLifecycleStateFailed   ListDataSafePrivateEndpointsLifecycleStateEnum = "FAILED"
	ListDataSafePrivateEndpointsLifecycleStateNa       ListDataSafePrivateEndpointsLifecycleStateEnum = "NA"
)

var mappingListDataSafePrivateEndpointsLifecycleState = map[string]ListDataSafePrivateEndpointsLifecycleStateEnum{
	"CREATING": ListDataSafePrivateEndpointsLifecycleStateCreating,
	"UPDATING": ListDataSafePrivateEndpointsLifecycleStateUpdating,
	"ACTIVE":   ListDataSafePrivateEndpointsLifecycleStateActive,
	"DELETING": ListDataSafePrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListDataSafePrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListDataSafePrivateEndpointsLifecycleStateFailed,
	"NA":       ListDataSafePrivateEndpointsLifecycleStateNa,
}

// GetListDataSafePrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsLifecycleStateEnum
func GetListDataSafePrivateEndpointsLifecycleStateEnumValues() []ListDataSafePrivateEndpointsLifecycleStateEnum {
	values := make([]ListDataSafePrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDataSafePrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsSortOrderEnum
const (
	ListDataSafePrivateEndpointsSortOrderAsc  ListDataSafePrivateEndpointsSortOrderEnum = "ASC"
	ListDataSafePrivateEndpointsSortOrderDesc ListDataSafePrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListDataSafePrivateEndpointsSortOrder = map[string]ListDataSafePrivateEndpointsSortOrderEnum{
	"ASC":  ListDataSafePrivateEndpointsSortOrderAsc,
	"DESC": ListDataSafePrivateEndpointsSortOrderDesc,
}

// GetListDataSafePrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsSortOrderEnum
func GetListDataSafePrivateEndpointsSortOrderEnumValues() []ListDataSafePrivateEndpointsSortOrderEnum {
	values := make([]ListDataSafePrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDataSafePrivateEndpointsSortByEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsSortByEnum
const (
	ListDataSafePrivateEndpointsSortByTimecreated ListDataSafePrivateEndpointsSortByEnum = "TIMECREATED"
	ListDataSafePrivateEndpointsSortByDisplayname ListDataSafePrivateEndpointsSortByEnum = "DISPLAYNAME"
)

var mappingListDataSafePrivateEndpointsSortBy = map[string]ListDataSafePrivateEndpointsSortByEnum{
	"TIMECREATED": ListDataSafePrivateEndpointsSortByTimecreated,
	"DISPLAYNAME": ListDataSafePrivateEndpointsSortByDisplayname,
}

// GetListDataSafePrivateEndpointsSortByEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsSortByEnum
func GetListDataSafePrivateEndpointsSortByEnumValues() []ListDataSafePrivateEndpointsSortByEnum {
	values := make([]ListDataSafePrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsSortBy {
		values = append(values, v)
	}
	return values
}
