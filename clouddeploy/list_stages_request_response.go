// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v33/common"
	"net/http"
)

// ListStagesRequest wrapper for the ListStages operation
type ListStagesRequest struct {

	// The ID of the Stage to be queried
	StageId *string `mandatory:"false" contributesTo:"query" name:"stageId"`

	// The ID of the parent Pipeline.
	PipelineId *string `mandatory:"false" contributesTo:"query" name:"pipelineId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only Stages that matches the given lifecycleState.
	LifecycleState StageLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListStagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListStagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStagesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListStagesResponse wrapper for the ListStages operation
type ListStagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StageCollection instances
	StageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStagesSortOrderEnum Enum with underlying type: string
type ListStagesSortOrderEnum string

// Set of constants representing the allowable values for ListStagesSortOrderEnum
const (
	ListStagesSortOrderAsc  ListStagesSortOrderEnum = "ASC"
	ListStagesSortOrderDesc ListStagesSortOrderEnum = "DESC"
)

var mappingListStagesSortOrder = map[string]ListStagesSortOrderEnum{
	"ASC":  ListStagesSortOrderAsc,
	"DESC": ListStagesSortOrderDesc,
}

// GetListStagesSortOrderEnumValues Enumerates the set of values for ListStagesSortOrderEnum
func GetListStagesSortOrderEnumValues() []ListStagesSortOrderEnum {
	values := make([]ListStagesSortOrderEnum, 0)
	for _, v := range mappingListStagesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListStagesSortByEnum Enum with underlying type: string
type ListStagesSortByEnum string

// Set of constants representing the allowable values for ListStagesSortByEnum
const (
	ListStagesSortByTimecreated ListStagesSortByEnum = "timeCreated"
	ListStagesSortByDisplayname ListStagesSortByEnum = "displayName"
)

var mappingListStagesSortBy = map[string]ListStagesSortByEnum{
	"timeCreated": ListStagesSortByTimecreated,
	"displayName": ListStagesSortByDisplayname,
}

// GetListStagesSortByEnumValues Enumerates the set of values for ListStagesSortByEnum
func GetListStagesSortByEnumValues() []ListStagesSortByEnum {
	values := make([]ListStagesSortByEnum, 0)
	for _, v := range mappingListStagesSortBy {
		values = append(values, v)
	}
	return values
}
