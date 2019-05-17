// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListCloudSyncsRequest wrapper for the ListCloudSyncs operation
type ListCloudSyncsRequest struct {

	// The storage gateway OCID.
	StorageGatewayId *string `mandatory:"true" contributesTo:"path" name:"storageGatewayId"`

	// The value of the `opc-next-page` response header from the previous "List" call. For information about
	// pagination, see List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The optional field to sort the results by.
	SortBy ListCloudSyncsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The optional order in which to sort the results.
	SortOrder ListCloudSyncsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The optional filter to apply to the results.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The optional filter to apply to the results.
	LifecycleState LifecycleState `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudSyncsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudSyncsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudSyncsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCloudSyncsResponse wrapper for the ListCloudSyncs operation
type ListCloudSyncsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CloudSyncSummary instances
	Items []CloudSyncSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of list of items. When paging through a list, provide this value as the `page` parameter for
	// the subsequent request to page backwards.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudSyncsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudSyncsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudSyncsSortByEnum Enum with underlying type: string
type ListCloudSyncsSortByEnum string

// Set of constants representing the allowable values for ListCloudSyncsSortByEnum
const (
	ListCloudSyncsSortByTimecreated ListCloudSyncsSortByEnum = "TIMECREATED"
	ListCloudSyncsSortByDisplayname ListCloudSyncsSortByEnum = "DISPLAYNAME"
)

var mappingListCloudSyncsSortBy = map[string]ListCloudSyncsSortByEnum{
	"TIMECREATED": ListCloudSyncsSortByTimecreated,
	"DISPLAYNAME": ListCloudSyncsSortByDisplayname,
}

// GetListCloudSyncsSortByEnumValues Enumerates the set of values for ListCloudSyncsSortByEnum
func GetListCloudSyncsSortByEnumValues() []ListCloudSyncsSortByEnum {
	values := make([]ListCloudSyncsSortByEnum, 0)
	for _, v := range mappingListCloudSyncsSortBy {
		values = append(values, v)
	}
	return values
}

// ListCloudSyncsSortOrderEnum Enum with underlying type: string
type ListCloudSyncsSortOrderEnum string

// Set of constants representing the allowable values for ListCloudSyncsSortOrderEnum
const (
	ListCloudSyncsSortOrderAsc  ListCloudSyncsSortOrderEnum = "ASC"
	ListCloudSyncsSortOrderDesc ListCloudSyncsSortOrderEnum = "DESC"
)

var mappingListCloudSyncsSortOrder = map[string]ListCloudSyncsSortOrderEnum{
	"ASC":  ListCloudSyncsSortOrderAsc,
	"DESC": ListCloudSyncsSortOrderDesc,
}

// GetListCloudSyncsSortOrderEnumValues Enumerates the set of values for ListCloudSyncsSortOrderEnum
func GetListCloudSyncsSortOrderEnumValues() []ListCloudSyncsSortOrderEnum {
	values := make([]ListCloudSyncsSortOrderEnum, 0)
	for _, v := range mappingListCloudSyncsSortOrder {
		values = append(values, v)
	}
	return values
}
