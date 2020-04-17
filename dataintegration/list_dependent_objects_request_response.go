// Copyright (c) 2016, 2018, , Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDependentObjectsRequest wrapper for the ListDependentObjects operation
type ListDependentObjectsRequest struct {

	// DIS workspace id
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// DIS application key
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// This parameter allows users to specify which fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// This filter parameter can be used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This filter parameter can be used to filter by the identifier of the published object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// This filter parameter can be used to filter by the object type of the object.
	// This parameter can be suffixed with an optional filter operator InSubtree.
	// For DIS APIs we will filter based on type Task.
	Type []string `contributesTo:"query" name:"type" collectionFormat:"multi"`

	// This is used in association with type parameter. If value is true,
	// then type all sub types of the given type parameter is considered.
	// If value is false, then sub types are not considered. Default is false.
	TypeInSubtree *string `mandatory:"false" contributesTo:"query" name:"typeInSubtree"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This parameter is used to control the sort order.  Supported values are `ASC` (ascending) and `DESC` (descending).
	SortOrder ListDependentObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This parameter allows users to specify a sort field.  Supported sort fields are `name`, `identifier`, `timeCreated`, and `timeUpdated`.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListDependentObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDependentObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDependentObjectsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDependentObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDependentObjectsResponse wrapper for the ListDependentObjects operation
type ListDependentObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DependentObjectSummaryCollection instances
	DependentObjectSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDependentObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDependentObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDependentObjectsSortOrderEnum Enum with underlying type: string
type ListDependentObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListDependentObjectsSortOrderEnum
const (
	ListDependentObjectsSortOrderAsc  ListDependentObjectsSortOrderEnum = "ASC"
	ListDependentObjectsSortOrderDesc ListDependentObjectsSortOrderEnum = "DESC"
)

var mappingListDependentObjectsSortOrder = map[string]ListDependentObjectsSortOrderEnum{
	"ASC":  ListDependentObjectsSortOrderAsc,
	"DESC": ListDependentObjectsSortOrderDesc,
}

// GetListDependentObjectsSortOrderEnumValues Enumerates the set of values for ListDependentObjectsSortOrderEnum
func GetListDependentObjectsSortOrderEnumValues() []ListDependentObjectsSortOrderEnum {
	values := make([]ListDependentObjectsSortOrderEnum, 0)
	for _, v := range mappingListDependentObjectsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDependentObjectsSortByEnum Enum with underlying type: string
type ListDependentObjectsSortByEnum string

// Set of constants representing the allowable values for ListDependentObjectsSortByEnum
const (
	ListDependentObjectsSortByTimeCreated ListDependentObjectsSortByEnum = "TIME_CREATED"
	ListDependentObjectsSortByDisplayName ListDependentObjectsSortByEnum = "DISPLAY_NAME"
)

var mappingListDependentObjectsSortBy = map[string]ListDependentObjectsSortByEnum{
	"TIME_CREATED": ListDependentObjectsSortByTimeCreated,
	"DISPLAY_NAME": ListDependentObjectsSortByDisplayName,
}

// GetListDependentObjectsSortByEnumValues Enumerates the set of values for ListDependentObjectsSortByEnum
func GetListDependentObjectsSortByEnumValues() []ListDependentObjectsSortByEnum {
	values := make([]ListDependentObjectsSortByEnum, 0)
	for _, v := range mappingListDependentObjectsSortBy {
		values = append(values, v)
	}
	return values
}
