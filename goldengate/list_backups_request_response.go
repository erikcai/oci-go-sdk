// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"net/http"
)

// ListBackupsRequest wrapper for the ListBackups operation
type ListBackupsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the deployment in which to list resources.
	DeploymentId *string `mandatory:"false" contributesTo:"query" name:"deploymentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListBackupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBackupsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBackupsResponse wrapper for the ListBackups operation
type ListBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BackupCollection instances
	BackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBackupsLifecycleStateEnum Enum with underlying type: string
type ListBackupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListBackupsLifecycleStateEnum
const (
	ListBackupsLifecycleStateCreating ListBackupsLifecycleStateEnum = "CREATING"
	ListBackupsLifecycleStateUpdating ListBackupsLifecycleStateEnum = "UPDATING"
	ListBackupsLifecycleStateActive   ListBackupsLifecycleStateEnum = "ACTIVE"
	ListBackupsLifecycleStateInactive ListBackupsLifecycleStateEnum = "INACTIVE"
	ListBackupsLifecycleStateDeleting ListBackupsLifecycleStateEnum = "DELETING"
	ListBackupsLifecycleStateDeleted  ListBackupsLifecycleStateEnum = "DELETED"
	ListBackupsLifecycleStateFailed   ListBackupsLifecycleStateEnum = "FAILED"
)

var mappingListBackupsLifecycleState = map[string]ListBackupsLifecycleStateEnum{
	"CREATING": ListBackupsLifecycleStateCreating,
	"UPDATING": ListBackupsLifecycleStateUpdating,
	"ACTIVE":   ListBackupsLifecycleStateActive,
	"INACTIVE": ListBackupsLifecycleStateInactive,
	"DELETING": ListBackupsLifecycleStateDeleting,
	"DELETED":  ListBackupsLifecycleStateDeleted,
	"FAILED":   ListBackupsLifecycleStateFailed,
}

// GetListBackupsLifecycleStateEnumValues Enumerates the set of values for ListBackupsLifecycleStateEnum
func GetListBackupsLifecycleStateEnumValues() []ListBackupsLifecycleStateEnum {
	values := make([]ListBackupsLifecycleStateEnum, 0)
	for _, v := range mappingListBackupsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListBackupsSortOrderEnum Enum with underlying type: string
type ListBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListBackupsSortOrderEnum
const (
	ListBackupsSortOrderAsc  ListBackupsSortOrderEnum = "ASC"
	ListBackupsSortOrderDesc ListBackupsSortOrderEnum = "DESC"
)

var mappingListBackupsSortOrder = map[string]ListBackupsSortOrderEnum{
	"ASC":  ListBackupsSortOrderAsc,
	"DESC": ListBackupsSortOrderDesc,
}

// GetListBackupsSortOrderEnumValues Enumerates the set of values for ListBackupsSortOrderEnum
func GetListBackupsSortOrderEnumValues() []ListBackupsSortOrderEnum {
	values := make([]ListBackupsSortOrderEnum, 0)
	for _, v := range mappingListBackupsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListBackupsSortByEnum Enum with underlying type: string
type ListBackupsSortByEnum string

// Set of constants representing the allowable values for ListBackupsSortByEnum
const (
	ListBackupsSortByTimecreated ListBackupsSortByEnum = "timeCreated"
	ListBackupsSortByDisplayname ListBackupsSortByEnum = "displayName"
)

var mappingListBackupsSortBy = map[string]ListBackupsSortByEnum{
	"timeCreated": ListBackupsSortByTimecreated,
	"displayName": ListBackupsSortByDisplayname,
}

// GetListBackupsSortByEnumValues Enumerates the set of values for ListBackupsSortByEnum
func GetListBackupsSortByEnumValues() []ListBackupsSortByEnum {
	values := make([]ListBackupsSortByEnum, 0)
	for _, v := range mappingListBackupsSortBy {
		values = append(values, v)
	}
	return values
}
