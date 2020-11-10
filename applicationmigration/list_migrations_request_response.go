// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package applicationmigration

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"net/http"
)

// ListMigrationsRequest wrapper for the ListMigrations operation
type ListMigrationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a compartment. Retrieves details of objects in the specified compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on which to query for a migration.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListMigrationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field on which to sort.
	// By default, `TIMECREATED` is ordered descending.
	// By default, `DISPLAYNAME` is ordered ascending. Note that you can sort only on one field.
	SortBy ListMigrationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Display name on which to query.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// This field is not supported. Do not use.
	LifecycleState ListMigrationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMigrationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMigrationsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMigrationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListMigrationsResponse wrapper for the ListMigrations operation
type ListMigrationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MigrationSummary instances
	Items []MigrationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMigrationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMigrationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMigrationsSortOrderEnum Enum with underlying type: string
type ListMigrationsSortOrderEnum string

// Set of constants representing the allowable values for ListMigrationsSortOrderEnum
const (
	ListMigrationsSortOrderAsc  ListMigrationsSortOrderEnum = "ASC"
	ListMigrationsSortOrderDesc ListMigrationsSortOrderEnum = "DESC"
)

var mappingListMigrationsSortOrder = map[string]ListMigrationsSortOrderEnum{
	"ASC":  ListMigrationsSortOrderAsc,
	"DESC": ListMigrationsSortOrderDesc,
}

// GetListMigrationsSortOrderEnumValues Enumerates the set of values for ListMigrationsSortOrderEnum
func GetListMigrationsSortOrderEnumValues() []ListMigrationsSortOrderEnum {
	values := make([]ListMigrationsSortOrderEnum, 0)
	for _, v := range mappingListMigrationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListMigrationsSortByEnum Enum with underlying type: string
type ListMigrationsSortByEnum string

// Set of constants representing the allowable values for ListMigrationsSortByEnum
const (
	ListMigrationsSortByTimecreated ListMigrationsSortByEnum = "TIMECREATED"
	ListMigrationsSortByDisplayname ListMigrationsSortByEnum = "DISPLAYNAME"
)

var mappingListMigrationsSortBy = map[string]ListMigrationsSortByEnum{
	"TIMECREATED": ListMigrationsSortByTimecreated,
	"DISPLAYNAME": ListMigrationsSortByDisplayname,
}

// GetListMigrationsSortByEnumValues Enumerates the set of values for ListMigrationsSortByEnum
func GetListMigrationsSortByEnumValues() []ListMigrationsSortByEnum {
	values := make([]ListMigrationsSortByEnum, 0)
	for _, v := range mappingListMigrationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListMigrationsLifecycleStateEnum Enum with underlying type: string
type ListMigrationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListMigrationsLifecycleStateEnum
const (
	ListMigrationsLifecycleStateCreating  ListMigrationsLifecycleStateEnum = "CREATING"
	ListMigrationsLifecycleStateActive    ListMigrationsLifecycleStateEnum = "ACTIVE"
	ListMigrationsLifecycleStateInactive  ListMigrationsLifecycleStateEnum = "INACTIVE"
	ListMigrationsLifecycleStateUpdating  ListMigrationsLifecycleStateEnum = "UPDATING"
	ListMigrationsLifecycleStateSucceeded ListMigrationsLifecycleStateEnum = "SUCCEEDED"
	ListMigrationsLifecycleStateDeleting  ListMigrationsLifecycleStateEnum = "DELETING"
	ListMigrationsLifecycleStateDeleted   ListMigrationsLifecycleStateEnum = "DELETED"
)

var mappingListMigrationsLifecycleState = map[string]ListMigrationsLifecycleStateEnum{
	"CREATING":  ListMigrationsLifecycleStateCreating,
	"ACTIVE":    ListMigrationsLifecycleStateActive,
	"INACTIVE":  ListMigrationsLifecycleStateInactive,
	"UPDATING":  ListMigrationsLifecycleStateUpdating,
	"SUCCEEDED": ListMigrationsLifecycleStateSucceeded,
	"DELETING":  ListMigrationsLifecycleStateDeleting,
	"DELETED":   ListMigrationsLifecycleStateDeleted,
}

// GetListMigrationsLifecycleStateEnumValues Enumerates the set of values for ListMigrationsLifecycleStateEnum
func GetListMigrationsLifecycleStateEnumValues() []ListMigrationsLifecycleStateEnum {
	values := make([]ListMigrationsLifecycleStateEnum, 0)
	for _, v := range mappingListMigrationsLifecycleState {
		values = append(values, v)
	}
	return values
}
