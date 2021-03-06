// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package clouddeploy

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"net/http"
)

// ListDeploymentsRequest wrapper for the ListDeployments operation
type ListDeploymentsRequest struct {

	// unique Deployment identifier
	DeploymentId *string `mandatory:"false" contributesTo:"query" name:"deploymentId"`

	// The ID of the parent Pipeline.
	PipelineId *string `mandatory:"false" contributesTo:"query" name:"pipelineId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of the parent Application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// A filter to return only Deployments that matches the given lifecycleState.
	LifecycleState DeploymentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDeploymentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDeploymentsResponse wrapper for the ListDeployments operation
type ListDeploymentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentCollection instances
	DeploymentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentsSortOrderEnum Enum with underlying type: string
type ListDeploymentsSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentsSortOrderEnum
const (
	ListDeploymentsSortOrderAsc  ListDeploymentsSortOrderEnum = "ASC"
	ListDeploymentsSortOrderDesc ListDeploymentsSortOrderEnum = "DESC"
)

var mappingListDeploymentsSortOrder = map[string]ListDeploymentsSortOrderEnum{
	"ASC":  ListDeploymentsSortOrderAsc,
	"DESC": ListDeploymentsSortOrderDesc,
}

// GetListDeploymentsSortOrderEnumValues Enumerates the set of values for ListDeploymentsSortOrderEnum
func GetListDeploymentsSortOrderEnumValues() []ListDeploymentsSortOrderEnum {
	values := make([]ListDeploymentsSortOrderEnum, 0)
	for _, v := range mappingListDeploymentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDeploymentsSortByEnum Enum with underlying type: string
type ListDeploymentsSortByEnum string

// Set of constants representing the allowable values for ListDeploymentsSortByEnum
const (
	ListDeploymentsSortByTimecreated ListDeploymentsSortByEnum = "timeCreated"
	ListDeploymentsSortByDisplayname ListDeploymentsSortByEnum = "displayName"
)

var mappingListDeploymentsSortBy = map[string]ListDeploymentsSortByEnum{
	"timeCreated": ListDeploymentsSortByTimecreated,
	"displayName": ListDeploymentsSortByDisplayname,
}

// GetListDeploymentsSortByEnumValues Enumerates the set of values for ListDeploymentsSortByEnum
func GetListDeploymentsSortByEnumValues() []ListDeploymentsSortByEnum {
	values := make([]ListDeploymentsSortByEnum, 0)
	for _, v := range mappingListDeploymentsSortBy {
		values = append(values, v)
	}
	return values
}
