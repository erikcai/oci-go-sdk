// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
type ListWorkRequestsRequest struct {

	// <b>Filter</b> results by the OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID. Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Filter results by the OCID of the resource associated with the work request.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Filter results by the type of the resource associated with the work request.
	ResourceType ListWorkRequestsResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"resourceType" omitEmpty:"true"`

	// Filter results by work request status.
	Status ListWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

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
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one value only.
	// By default, when you sort by time fields, results are shown
	// in descending order. All other fields default to ascending order.
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkRequestsResponse wrapper for the ListWorkRequests operation
type ListWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkRequestSummary instances
	Items []WorkRequestSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
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

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsResourceTypeEnum Enum with underlying type: string
type ListWorkRequestsResourceTypeEnum string

// Set of constants representing the allowable values for ListWorkRequestsResourceTypeEnum
const (
	ListWorkRequestsResourceTypeNotebookSession ListWorkRequestsResourceTypeEnum = "NOTEBOOK_SESSION"
	ListWorkRequestsResourceTypeProject         ListWorkRequestsResourceTypeEnum = "PROJECT"
)

var mappingListWorkRequestsResourceType = map[string]ListWorkRequestsResourceTypeEnum{
	"NOTEBOOK_SESSION": ListWorkRequestsResourceTypeNotebookSession,
	"PROJECT":          ListWorkRequestsResourceTypeProject,
}

// GetListWorkRequestsResourceTypeEnumValues Enumerates the set of values for ListWorkRequestsResourceTypeEnum
func GetListWorkRequestsResourceTypeEnumValues() []ListWorkRequestsResourceTypeEnum {
	values := make([]ListWorkRequestsResourceTypeEnum, 0)
	for _, v := range mappingListWorkRequestsResourceType {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestsStatusEnum Enum with underlying type: string
type ListWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListWorkRequestsStatusEnum
const (
	ListWorkRequestsStatusAccepted   ListWorkRequestsStatusEnum = "ACCEPTED"
	ListWorkRequestsStatusInProgress ListWorkRequestsStatusEnum = "IN_PROGRESS"
	ListWorkRequestsStatusFailed     ListWorkRequestsStatusEnum = "FAILED"
	ListWorkRequestsStatusSucceeded  ListWorkRequestsStatusEnum = "SUCCEEDED"
	ListWorkRequestsStatusCanceling  ListWorkRequestsStatusEnum = "CANCELING"
	ListWorkRequestsStatusCanceled   ListWorkRequestsStatusEnum = "CANCELED"
)

var mappingListWorkRequestsStatus = map[string]ListWorkRequestsStatusEnum{
	"ACCEPTED":    ListWorkRequestsStatusAccepted,
	"IN_PROGRESS": ListWorkRequestsStatusInProgress,
	"FAILED":      ListWorkRequestsStatusFailed,
	"SUCCEEDED":   ListWorkRequestsStatusSucceeded,
	"CANCELING":   ListWorkRequestsStatusCanceling,
	"CANCELED":    ListWorkRequestsStatusCanceled,
}

// GetListWorkRequestsStatusEnumValues Enumerates the set of values for ListWorkRequestsStatusEnum
func GetListWorkRequestsStatusEnumValues() []ListWorkRequestsStatusEnum {
	values := make([]ListWorkRequestsStatusEnum, 0)
	for _, v := range mappingListWorkRequestsStatus {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestsSortOrderEnum Enum with underlying type: string
type ListWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortOrderEnum
const (
	ListWorkRequestsSortOrderAsc  ListWorkRequestsSortOrderEnum = "ASC"
	ListWorkRequestsSortOrderDesc ListWorkRequestsSortOrderEnum = "DESC"
)

var mappingListWorkRequestsSortOrder = map[string]ListWorkRequestsSortOrderEnum{
	"ASC":  ListWorkRequestsSortOrderAsc,
	"DESC": ListWorkRequestsSortOrderDesc,
}

// GetListWorkRequestsSortOrderEnumValues Enumerates the set of values for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumValues() []ListWorkRequestsSortOrderEnum {
	values := make([]ListWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestsSortByEnum Enum with underlying type: string
type ListWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortByEnum
const (
	ListWorkRequestsSortById            ListWorkRequestsSortByEnum = "id"
	ListWorkRequestsSortByOperationtype ListWorkRequestsSortByEnum = "operationType"
	ListWorkRequestsSortByStatus        ListWorkRequestsSortByEnum = "status"
	ListWorkRequestsSortByTimeaccepted  ListWorkRequestsSortByEnum = "timeAccepted"
	ListWorkRequestsSortByTimestarted   ListWorkRequestsSortByEnum = "timeStarted"
	ListWorkRequestsSortByTimefinished  ListWorkRequestsSortByEnum = "timeFinished"
)

var mappingListWorkRequestsSortBy = map[string]ListWorkRequestsSortByEnum{
	"id":            ListWorkRequestsSortById,
	"operationType": ListWorkRequestsSortByOperationtype,
	"status":        ListWorkRequestsSortByStatus,
	"timeAccepted":  ListWorkRequestsSortByTimeaccepted,
	"timeStarted":   ListWorkRequestsSortByTimestarted,
	"timeFinished":  ListWorkRequestsSortByTimefinished,
}

// GetListWorkRequestsSortByEnumValues Enumerates the set of values for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumValues() []ListWorkRequestsSortByEnum {
	values := make([]ListWorkRequestsSortByEnum, 0)
	for _, v := range mappingListWorkRequestsSortBy {
		values = append(values, v)
	}
	return values
}
