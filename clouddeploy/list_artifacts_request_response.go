// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package clouddeploy

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"net/http"
)

// ListArtifactsRequest wrapper for the ListArtifacts operation
type ListArtifactsRequest struct {

	// The ID of the Artifact for filtering.
	ArtifactId *string `mandatory:"false" contributesTo:"query" name:"artifactId"`

	// The ID of the parent Application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only Artifacts that matches the given lifecycleState.
	LifecycleState ArtifactLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListArtifactsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListArtifactsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListArtifactsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListArtifactsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListArtifactsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListArtifactsResponse wrapper for the ListArtifacts operation
type ListArtifactsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ArtifactCollection instances
	ArtifactCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListArtifactsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListArtifactsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListArtifactsSortOrderEnum Enum with underlying type: string
type ListArtifactsSortOrderEnum string

// Set of constants representing the allowable values for ListArtifactsSortOrderEnum
const (
	ListArtifactsSortOrderAsc  ListArtifactsSortOrderEnum = "ASC"
	ListArtifactsSortOrderDesc ListArtifactsSortOrderEnum = "DESC"
)

var mappingListArtifactsSortOrder = map[string]ListArtifactsSortOrderEnum{
	"ASC":  ListArtifactsSortOrderAsc,
	"DESC": ListArtifactsSortOrderDesc,
}

// GetListArtifactsSortOrderEnumValues Enumerates the set of values for ListArtifactsSortOrderEnum
func GetListArtifactsSortOrderEnumValues() []ListArtifactsSortOrderEnum {
	values := make([]ListArtifactsSortOrderEnum, 0)
	for _, v := range mappingListArtifactsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListArtifactsSortByEnum Enum with underlying type: string
type ListArtifactsSortByEnum string

// Set of constants representing the allowable values for ListArtifactsSortByEnum
const (
	ListArtifactsSortByTimecreated ListArtifactsSortByEnum = "timeCreated"
	ListArtifactsSortByDisplayname ListArtifactsSortByEnum = "displayName"
)

var mappingListArtifactsSortBy = map[string]ListArtifactsSortByEnum{
	"timeCreated": ListArtifactsSortByTimecreated,
	"displayName": ListArtifactsSortByDisplayname,
}

// GetListArtifactsSortByEnumValues Enumerates the set of values for ListArtifactsSortByEnum
func GetListArtifactsSortByEnumValues() []ListArtifactsSortByEnum {
	values := make([]ListArtifactsSortByEnum, 0)
	for _, v := range mappingListArtifactsSortBy {
		values = append(values, v)
	}
	return values
}
