// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v32/common"
	"net/http"
)

// ListEnvironmentsRequest wrapper for the ListEnvironments operation
type ListEnvironmentsRequest struct {

	// The ID of the parent Application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// ID of the environment for filtering.
	EnvironmentId *string `mandatory:"false" contributesTo:"query" name:"environmentId"`

	// A filter to return only Environments that matches the given lifecycleState.
	LifecycleState EnvironmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListEnvironmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListEnvironmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEnvironmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEnvironmentsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEnvironmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListEnvironmentsResponse wrapper for the ListEnvironments operation
type ListEnvironmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EnvironmentCollection instances
	EnvironmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEnvironmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEnvironmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEnvironmentsSortOrderEnum Enum with underlying type: string
type ListEnvironmentsSortOrderEnum string

// Set of constants representing the allowable values for ListEnvironmentsSortOrderEnum
const (
	ListEnvironmentsSortOrderAsc  ListEnvironmentsSortOrderEnum = "ASC"
	ListEnvironmentsSortOrderDesc ListEnvironmentsSortOrderEnum = "DESC"
)

var mappingListEnvironmentsSortOrder = map[string]ListEnvironmentsSortOrderEnum{
	"ASC":  ListEnvironmentsSortOrderAsc,
	"DESC": ListEnvironmentsSortOrderDesc,
}

// GetListEnvironmentsSortOrderEnumValues Enumerates the set of values for ListEnvironmentsSortOrderEnum
func GetListEnvironmentsSortOrderEnumValues() []ListEnvironmentsSortOrderEnum {
	values := make([]ListEnvironmentsSortOrderEnum, 0)
	for _, v := range mappingListEnvironmentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListEnvironmentsSortByEnum Enum with underlying type: string
type ListEnvironmentsSortByEnum string

// Set of constants representing the allowable values for ListEnvironmentsSortByEnum
const (
	ListEnvironmentsSortByTimecreated ListEnvironmentsSortByEnum = "timeCreated"
	ListEnvironmentsSortByDisplayname ListEnvironmentsSortByEnum = "displayName"
)

var mappingListEnvironmentsSortBy = map[string]ListEnvironmentsSortByEnum{
	"timeCreated": ListEnvironmentsSortByTimecreated,
	"displayName": ListEnvironmentsSortByDisplayname,
}

// GetListEnvironmentsSortByEnumValues Enumerates the set of values for ListEnvironmentsSortByEnum
func GetListEnvironmentsSortByEnumValues() []ListEnvironmentsSortByEnum {
	values := make([]ListEnvironmentsSortByEnum, 0)
	for _, v := range mappingListEnvironmentsSortBy {
		values = append(values, v)
	}
	return values
}
