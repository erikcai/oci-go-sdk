// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"net/http"
)

// ListJobExecutionsRequest wrapper for the ListJobExecutions operation
type ListJobExecutionsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Job id
	JobId *string `mandatory:"false" contributesTo:"query" name:"jobId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseGroupId"`

	// Status of the job execution
	Status *string `mandatory:"false" contributesTo:"query" name:"status"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for `TIMECREATED` is descending. Default order for `NAME`
	// is ascending. The `NAME` sort order is case sensitive.
	SortBy ListJobExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListJobExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJobExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobExecutionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListJobExecutionsResponse wrapper for the ListJobExecutions operation
type ListJobExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JobExecutionCollection instances
	JobExecutionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobExecutionsSortByEnum Enum with underlying type: string
type ListJobExecutionsSortByEnum string

// Set of constants representing the allowable values for ListJobExecutionsSortByEnum
const (
	ListJobExecutionsSortByTimecreated ListJobExecutionsSortByEnum = "TIMECREATED"
	ListJobExecutionsSortByName        ListJobExecutionsSortByEnum = "NAME"
)

var mappingListJobExecutionsSortBy = map[string]ListJobExecutionsSortByEnum{
	"TIMECREATED": ListJobExecutionsSortByTimecreated,
	"NAME":        ListJobExecutionsSortByName,
}

// GetListJobExecutionsSortByEnumValues Enumerates the set of values for ListJobExecutionsSortByEnum
func GetListJobExecutionsSortByEnumValues() []ListJobExecutionsSortByEnum {
	values := make([]ListJobExecutionsSortByEnum, 0)
	for _, v := range mappingListJobExecutionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListJobExecutionsSortOrderEnum Enum with underlying type: string
type ListJobExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListJobExecutionsSortOrderEnum
const (
	ListJobExecutionsSortOrderAsc  ListJobExecutionsSortOrderEnum = "ASC"
	ListJobExecutionsSortOrderDesc ListJobExecutionsSortOrderEnum = "DESC"
)

var mappingListJobExecutionsSortOrder = map[string]ListJobExecutionsSortOrderEnum{
	"ASC":  ListJobExecutionsSortOrderAsc,
	"DESC": ListJobExecutionsSortOrderDesc,
}

// GetListJobExecutionsSortOrderEnumValues Enumerates the set of values for ListJobExecutionsSortOrderEnum
func GetListJobExecutionsSortOrderEnumValues() []ListJobExecutionsSortOrderEnum {
	values := make([]ListJobExecutionsSortOrderEnum, 0)
	for _, v := range mappingListJobExecutionsSortOrder {
		values = append(values, v)
	}
	return values
}