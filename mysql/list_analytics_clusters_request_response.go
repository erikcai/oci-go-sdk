// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAnalyticsClustersRequest wrapper for the ListAnalyticsClusters operation
type ListAnalyticsClustersRequest struct {

	// The name of the Availability Domain.
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle State
	LifecycleState ListAnalyticsClustersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending. Display name is default ordered as ascending.
	SortBy ListAnalyticsClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAnalyticsClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListAnalyticsClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAnalyticsClustersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAnalyticsClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAnalyticsClustersResponse wrapper for the ListAnalyticsClusters operation
type ListAnalyticsClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AnalyticsClusterSummary instances
	Items []AnalyticsClusterSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a specific request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Opaque token representing the next page of results.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAnalyticsClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAnalyticsClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAnalyticsClustersLifecycleStateEnum Enum with underlying type: string
type ListAnalyticsClustersLifecycleStateEnum string

// Set of constants representing the allowable values for ListAnalyticsClustersLifecycleStateEnum
const (
	ListAnalyticsClustersLifecycleStateCreating ListAnalyticsClustersLifecycleStateEnum = "CREATING"
	ListAnalyticsClustersLifecycleStateActive   ListAnalyticsClustersLifecycleStateEnum = "ACTIVE"
	ListAnalyticsClustersLifecycleStateInactive ListAnalyticsClustersLifecycleStateEnum = "INACTIVE"
	ListAnalyticsClustersLifecycleStateUpdating ListAnalyticsClustersLifecycleStateEnum = "UPDATING"
	ListAnalyticsClustersLifecycleStateDeleting ListAnalyticsClustersLifecycleStateEnum = "DELETING"
	ListAnalyticsClustersLifecycleStateDeleted  ListAnalyticsClustersLifecycleStateEnum = "DELETED"
	ListAnalyticsClustersLifecycleStateFailed   ListAnalyticsClustersLifecycleStateEnum = "FAILED"
)

var mappingListAnalyticsClustersLifecycleState = map[string]ListAnalyticsClustersLifecycleStateEnum{
	"CREATING": ListAnalyticsClustersLifecycleStateCreating,
	"ACTIVE":   ListAnalyticsClustersLifecycleStateActive,
	"INACTIVE": ListAnalyticsClustersLifecycleStateInactive,
	"UPDATING": ListAnalyticsClustersLifecycleStateUpdating,
	"DELETING": ListAnalyticsClustersLifecycleStateDeleting,
	"DELETED":  ListAnalyticsClustersLifecycleStateDeleted,
	"FAILED":   ListAnalyticsClustersLifecycleStateFailed,
}

// GetListAnalyticsClustersLifecycleStateEnumValues Enumerates the set of values for ListAnalyticsClustersLifecycleStateEnum
func GetListAnalyticsClustersLifecycleStateEnumValues() []ListAnalyticsClustersLifecycleStateEnum {
	values := make([]ListAnalyticsClustersLifecycleStateEnum, 0)
	for _, v := range mappingListAnalyticsClustersLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListAnalyticsClustersSortByEnum Enum with underlying type: string
type ListAnalyticsClustersSortByEnum string

// Set of constants representing the allowable values for ListAnalyticsClustersSortByEnum
const (
	ListAnalyticsClustersSortByDisplayname ListAnalyticsClustersSortByEnum = "displayName"
	ListAnalyticsClustersSortByTimecreated ListAnalyticsClustersSortByEnum = "timeCreated"
)

var mappingListAnalyticsClustersSortBy = map[string]ListAnalyticsClustersSortByEnum{
	"displayName": ListAnalyticsClustersSortByDisplayname,
	"timeCreated": ListAnalyticsClustersSortByTimecreated,
}

// GetListAnalyticsClustersSortByEnumValues Enumerates the set of values for ListAnalyticsClustersSortByEnum
func GetListAnalyticsClustersSortByEnumValues() []ListAnalyticsClustersSortByEnum {
	values := make([]ListAnalyticsClustersSortByEnum, 0)
	for _, v := range mappingListAnalyticsClustersSortBy {
		values = append(values, v)
	}
	return values
}

// ListAnalyticsClustersSortOrderEnum Enum with underlying type: string
type ListAnalyticsClustersSortOrderEnum string

// Set of constants representing the allowable values for ListAnalyticsClustersSortOrderEnum
const (
	ListAnalyticsClustersSortOrderAsc  ListAnalyticsClustersSortOrderEnum = "ASC"
	ListAnalyticsClustersSortOrderDesc ListAnalyticsClustersSortOrderEnum = "DESC"
)

var mappingListAnalyticsClustersSortOrder = map[string]ListAnalyticsClustersSortOrderEnum{
	"ASC":  ListAnalyticsClustersSortOrderAsc,
	"DESC": ListAnalyticsClustersSortOrderDesc,
}

// GetListAnalyticsClustersSortOrderEnumValues Enumerates the set of values for ListAnalyticsClustersSortOrderEnum
func GetListAnalyticsClustersSortOrderEnumValues() []ListAnalyticsClustersSortOrderEnum {
	values := make([]ListAnalyticsClustersSortOrderEnum, 0)
	for _, v := range mappingListAnalyticsClustersSortOrder {
		values = append(values, v)
	}
	return values
}
