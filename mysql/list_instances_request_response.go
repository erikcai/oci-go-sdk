// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListInstancesRequest wrapper for the ListInstances operation
type ListInstancesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If true, return only MySQLaaS Instances with Analytics enabled, if false
	// return only MySQLaaS Instances without Analytics enabled. If not
	// present, return all MySQLaaS Instances.
	IsAnalyticsEnabled *bool `mandatory:"false" contributesTo:"query" name:"isAnalyticsEnabled"`

	// The name of the Availability Domain.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The requested MySQLaaS Configuration instance.
	ConfigurationId *string `mandatory:"false" contributesTo:"query" name:"configurationId"`

	// Filter instances if they are using the latest revision of the
	// Configuration they are associated with.
	IsUpToDate *bool `mandatory:"false" contributesTo:"query" name:"isUpToDate"`

	// The DbSystem OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle State
	LifecycleState ListInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending. Display name is default ordered as ascending.
	SortBy ListInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListInstancesResponse wrapper for the ListInstances operation
type ListInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InstanceSummary instances
	Items []InstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a specific request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Opaque token representing the next page of results.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstancesLifecycleStateEnum Enum with underlying type: string
type ListInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListInstancesLifecycleStateEnum
const (
	ListInstancesLifecycleStateCreating ListInstancesLifecycleStateEnum = "CREATING"
	ListInstancesLifecycleStateActive   ListInstancesLifecycleStateEnum = "ACTIVE"
	ListInstancesLifecycleStateInactive ListInstancesLifecycleStateEnum = "INACTIVE"
	ListInstancesLifecycleStateUpdating ListInstancesLifecycleStateEnum = "UPDATING"
	ListInstancesLifecycleStateDeleting ListInstancesLifecycleStateEnum = "DELETING"
	ListInstancesLifecycleStateDeleted  ListInstancesLifecycleStateEnum = "DELETED"
	ListInstancesLifecycleStateFailed   ListInstancesLifecycleStateEnum = "FAILED"
)

var mappingListInstancesLifecycleState = map[string]ListInstancesLifecycleStateEnum{
	"CREATING": ListInstancesLifecycleStateCreating,
	"ACTIVE":   ListInstancesLifecycleStateActive,
	"INACTIVE": ListInstancesLifecycleStateInactive,
	"UPDATING": ListInstancesLifecycleStateUpdating,
	"DELETING": ListInstancesLifecycleStateDeleting,
	"DELETED":  ListInstancesLifecycleStateDeleted,
	"FAILED":   ListInstancesLifecycleStateFailed,
}

// GetListInstancesLifecycleStateEnumValues Enumerates the set of values for ListInstancesLifecycleStateEnum
func GetListInstancesLifecycleStateEnumValues() []ListInstancesLifecycleStateEnum {
	values := make([]ListInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListInstancesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListInstancesSortByEnum Enum with underlying type: string
type ListInstancesSortByEnum string

// Set of constants representing the allowable values for ListInstancesSortByEnum
const (
	ListInstancesSortByDisplayname ListInstancesSortByEnum = "displayName"
	ListInstancesSortByTimecreated ListInstancesSortByEnum = "timeCreated"
)

var mappingListInstancesSortBy = map[string]ListInstancesSortByEnum{
	"displayName": ListInstancesSortByDisplayname,
	"timeCreated": ListInstancesSortByTimecreated,
}

// GetListInstancesSortByEnumValues Enumerates the set of values for ListInstancesSortByEnum
func GetListInstancesSortByEnumValues() []ListInstancesSortByEnum {
	values := make([]ListInstancesSortByEnum, 0)
	for _, v := range mappingListInstancesSortBy {
		values = append(values, v)
	}
	return values
}

// ListInstancesSortOrderEnum Enum with underlying type: string
type ListInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListInstancesSortOrderEnum
const (
	ListInstancesSortOrderAsc  ListInstancesSortOrderEnum = "ASC"
	ListInstancesSortOrderDesc ListInstancesSortOrderEnum = "DESC"
)

var mappingListInstancesSortOrder = map[string]ListInstancesSortOrderEnum{
	"ASC":  ListInstancesSortOrderAsc,
	"DESC": ListInstancesSortOrderDesc,
}

// GetListInstancesSortOrderEnumValues Enumerates the set of values for ListInstancesSortOrderEnum
func GetListInstancesSortOrderEnumValues() []ListInstancesSortOrderEnum {
	values := make([]ListInstancesSortOrderEnum, 0)
	for _, v := range mappingListInstancesSortOrder {
		values = append(values, v)
	}
	return values
}
