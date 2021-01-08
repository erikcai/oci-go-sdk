// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// ListPipelinesRequest wrapper for the ListPipelines operation
type ListPipelinesRequest struct {

	// The ID of the parent Pipeline.
	PipelineId *string `mandatory:"false" contributesTo:"query" name:"pipelineId"`

	// The ID of the parent Application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only Pipelines that matches the given lifecycleState.
	LifecycleState PipelineLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPipelinesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPipelinesResponse wrapper for the ListPipelines operation
type ListPipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PipelineCollection instances
	PipelineCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPipelinesSortOrderEnum Enum with underlying type: string
type ListPipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListPipelinesSortOrderEnum
const (
	ListPipelinesSortOrderAsc  ListPipelinesSortOrderEnum = "ASC"
	ListPipelinesSortOrderDesc ListPipelinesSortOrderEnum = "DESC"
)

var mappingListPipelinesSortOrder = map[string]ListPipelinesSortOrderEnum{
	"ASC":  ListPipelinesSortOrderAsc,
	"DESC": ListPipelinesSortOrderDesc,
}

// GetListPipelinesSortOrderEnumValues Enumerates the set of values for ListPipelinesSortOrderEnum
func GetListPipelinesSortOrderEnumValues() []ListPipelinesSortOrderEnum {
	values := make([]ListPipelinesSortOrderEnum, 0)
	for _, v := range mappingListPipelinesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListPipelinesSortByEnum Enum with underlying type: string
type ListPipelinesSortByEnum string

// Set of constants representing the allowable values for ListPipelinesSortByEnum
const (
	ListPipelinesSortByTimecreated ListPipelinesSortByEnum = "timeCreated"
	ListPipelinesSortByDisplayname ListPipelinesSortByEnum = "displayName"
)

var mappingListPipelinesSortBy = map[string]ListPipelinesSortByEnum{
	"timeCreated": ListPipelinesSortByTimecreated,
	"displayName": ListPipelinesSortByDisplayname,
}

// GetListPipelinesSortByEnumValues Enumerates the set of values for ListPipelinesSortByEnum
func GetListPipelinesSortByEnumValues() []ListPipelinesSortByEnum {
	values := make([]ListPipelinesSortByEnum, 0)
	for _, v := range mappingListPipelinesSortBy {
		values = append(values, v)
	}
	return values
}
