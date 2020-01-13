// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
type ListWorkRequestsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the resource associated with a work request
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Type of the resource associated with a work request
	ResourceType ListWorkRequestsResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"resourceType" omitEmpty:"true"`

	// MySQLaaS Work Request Status
	WorkRequestStatus ListWorkRequestsWorkRequestStatusEnum `mandatory:"false" contributesTo:"query" name:"workRequestStatus" omitEmpty:"true"`

	// The optional field to sort the results by.
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous "List" call. For information about pagination, see List
	// Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

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

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a specific request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Opaque token representing the next page of results.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
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
	ListWorkRequestsResourceTypeInstance      ListWorkRequestsResourceTypeEnum = "MYSQL_INSTANCE"
	ListWorkRequestsResourceTypeConfiguration ListWorkRequestsResourceTypeEnum = "MYSQL_CONFIGURATION"
)

var mappingListWorkRequestsResourceType = map[string]ListWorkRequestsResourceTypeEnum{
	"MYSQL_INSTANCE":      ListWorkRequestsResourceTypeInstance,
	"MYSQL_CONFIGURATION": ListWorkRequestsResourceTypeConfiguration,
}

// GetListWorkRequestsResourceTypeEnumValues Enumerates the set of values for ListWorkRequestsResourceTypeEnum
func GetListWorkRequestsResourceTypeEnumValues() []ListWorkRequestsResourceTypeEnum {
	values := make([]ListWorkRequestsResourceTypeEnum, 0)
	for _, v := range mappingListWorkRequestsResourceType {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestsWorkRequestStatusEnum Enum with underlying type: string
type ListWorkRequestsWorkRequestStatusEnum string

// Set of constants representing the allowable values for ListWorkRequestsWorkRequestStatusEnum
const (
	ListWorkRequestsWorkRequestStatusAccepted   ListWorkRequestsWorkRequestStatusEnum = "ACCEPTED"
	ListWorkRequestsWorkRequestStatusInProgress ListWorkRequestsWorkRequestStatusEnum = "IN_PROGRESS"
	ListWorkRequestsWorkRequestStatusFailed     ListWorkRequestsWorkRequestStatusEnum = "FAILED"
	ListWorkRequestsWorkRequestStatusSucceeded  ListWorkRequestsWorkRequestStatusEnum = "SUCCEEDED"
	ListWorkRequestsWorkRequestStatusCanceling  ListWorkRequestsWorkRequestStatusEnum = "CANCELING"
	ListWorkRequestsWorkRequestStatusCanceled   ListWorkRequestsWorkRequestStatusEnum = "CANCELED"
)

var mappingListWorkRequestsWorkRequestStatus = map[string]ListWorkRequestsWorkRequestStatusEnum{
	"ACCEPTED":    ListWorkRequestsWorkRequestStatusAccepted,
	"IN_PROGRESS": ListWorkRequestsWorkRequestStatusInProgress,
	"FAILED":      ListWorkRequestsWorkRequestStatusFailed,
	"SUCCEEDED":   ListWorkRequestsWorkRequestStatusSucceeded,
	"CANCELING":   ListWorkRequestsWorkRequestStatusCanceling,
	"CANCELED":    ListWorkRequestsWorkRequestStatusCanceled,
}

// GetListWorkRequestsWorkRequestStatusEnumValues Enumerates the set of values for ListWorkRequestsWorkRequestStatusEnum
func GetListWorkRequestsWorkRequestStatusEnumValues() []ListWorkRequestsWorkRequestStatusEnum {
	values := make([]ListWorkRequestsWorkRequestStatusEnum, 0)
	for _, v := range mappingListWorkRequestsWorkRequestStatus {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestsSortByEnum Enum with underlying type: string
type ListWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortByEnum
const (
	ListWorkRequestsSortById            ListWorkRequestsSortByEnum = "ID"
	ListWorkRequestsSortByOperationType ListWorkRequestsSortByEnum = "OPERATION_TYPE"
	ListWorkRequestsSortByStatus        ListWorkRequestsSortByEnum = "STATUS"
	ListWorkRequestsSortByTimeAccepted  ListWorkRequestsSortByEnum = "TIME_ACCEPTED"
	ListWorkRequestsSortByTimeStarted   ListWorkRequestsSortByEnum = "TIME_STARTED"
	ListWorkRequestsSortByTimeFinished  ListWorkRequestsSortByEnum = "TIME_FINISHED"
)

var mappingListWorkRequestsSortBy = map[string]ListWorkRequestsSortByEnum{
	"ID":             ListWorkRequestsSortById,
	"OPERATION_TYPE": ListWorkRequestsSortByOperationType,
	"STATUS":         ListWorkRequestsSortByStatus,
	"TIME_ACCEPTED":  ListWorkRequestsSortByTimeAccepted,
	"TIME_STARTED":   ListWorkRequestsSortByTimeStarted,
	"TIME_FINISHED":  ListWorkRequestsSortByTimeFinished,
}

// GetListWorkRequestsSortByEnumValues Enumerates the set of values for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumValues() []ListWorkRequestsSortByEnum {
	values := make([]ListWorkRequestsSortByEnum, 0)
	for _, v := range mappingListWorkRequestsSortBy {
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
