// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListModelsRequest wrapper for the ListModels operation
type ListModelsRequest struct {

	// <b>Filter</b> results by the OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID. Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by the OCID of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListModelsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filters results by the OCID of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending.
	SortOrder ListModelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by display name, results are
	// shown in ascending order.
	SortBy ListModelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListModelsResponse wrapper for the ListModels operation
type ListModelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModelSummary instances
	Items []ModelSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list in forward direction, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list in backword direction, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListModelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelsLifecycleStateEnum Enum with underlying type: string
type ListModelsLifecycleStateEnum string

// Set of constants representing the allowable values for ListModelsLifecycleStateEnum
const (
	ListModelsLifecycleStateActive   ListModelsLifecycleStateEnum = "ACTIVE"
	ListModelsLifecycleStateDeleted  ListModelsLifecycleStateEnum = "DELETED"
	ListModelsLifecycleStateFailed   ListModelsLifecycleStateEnum = "FAILED"
	ListModelsLifecycleStateInactive ListModelsLifecycleStateEnum = "INACTIVE"
)

var mappingListModelsLifecycleState = map[string]ListModelsLifecycleStateEnum{
	"ACTIVE":   ListModelsLifecycleStateActive,
	"DELETED":  ListModelsLifecycleStateDeleted,
	"FAILED":   ListModelsLifecycleStateFailed,
	"INACTIVE": ListModelsLifecycleStateInactive,
}

// GetListModelsLifecycleStateEnumValues Enumerates the set of values for ListModelsLifecycleStateEnum
func GetListModelsLifecycleStateEnumValues() []ListModelsLifecycleStateEnum {
	values := make([]ListModelsLifecycleStateEnum, 0)
	for _, v := range mappingListModelsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListModelsSortOrderEnum Enum with underlying type: string
type ListModelsSortOrderEnum string

// Set of constants representing the allowable values for ListModelsSortOrderEnum
const (
	ListModelsSortOrderAsc  ListModelsSortOrderEnum = "ASC"
	ListModelsSortOrderDesc ListModelsSortOrderEnum = "DESC"
)

var mappingListModelsSortOrder = map[string]ListModelsSortOrderEnum{
	"ASC":  ListModelsSortOrderAsc,
	"DESC": ListModelsSortOrderDesc,
}

// GetListModelsSortOrderEnumValues Enumerates the set of values for ListModelsSortOrderEnum
func GetListModelsSortOrderEnumValues() []ListModelsSortOrderEnum {
	values := make([]ListModelsSortOrderEnum, 0)
	for _, v := range mappingListModelsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListModelsSortByEnum Enum with underlying type: string
type ListModelsSortByEnum string

// Set of constants representing the allowable values for ListModelsSortByEnum
const (
	ListModelsSortByTimecreated    ListModelsSortByEnum = "timeCreated"
	ListModelsSortByDisplayname    ListModelsSortByEnum = "displayName"
	ListModelsSortByLifecyclestate ListModelsSortByEnum = "lifecycleState"
)

var mappingListModelsSortBy = map[string]ListModelsSortByEnum{
	"timeCreated":    ListModelsSortByTimecreated,
	"displayName":    ListModelsSortByDisplayname,
	"lifecycleState": ListModelsSortByLifecyclestate,
}

// GetListModelsSortByEnumValues Enumerates the set of values for ListModelsSortByEnum
func GetListModelsSortByEnumValues() []ListModelsSortByEnum {
	values := make([]ListModelsSortByEnum, 0)
	for _, v := range mappingListModelsSortBy {
		values = append(values, v)
	}
	return values
}
