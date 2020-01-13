// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListWorkRequestResourcesRequest wrapper for the ListWorkRequestResources operation
type ListWorkRequestResourcesRequest struct {

	// the ID of the WorkRequest
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// The name of the Availability Domain.
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// MySQLaaS Work Request Resource Action Type
	ResourceActionType ListWorkRequestResourcesResourceActionTypeEnum `mandatory:"false" contributesTo:"query" name:"resourceActionType" omitEmpty:"true"`

	// The optional field to sort the results by.
	SortBy ListWorkRequestResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListWorkRequestResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListWorkRequestResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestResourcesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkRequestResourcesResponse wrapper for the ListWorkRequestResources operation
type ListWorkRequestResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkRequestResource instances
	Items []WorkRequestResource `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a specific request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Opaque token representing the next page of results.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestResourcesResourceActionTypeEnum Enum with underlying type: string
type ListWorkRequestResourcesResourceActionTypeEnum string

// Set of constants representing the allowable values for ListWorkRequestResourcesResourceActionTypeEnum
const (
	ListWorkRequestResourcesResourceActionTypeCreated    ListWorkRequestResourcesResourceActionTypeEnum = "CREATED"
	ListWorkRequestResourcesResourceActionTypeUpdated    ListWorkRequestResourcesResourceActionTypeEnum = "UPDATED"
	ListWorkRequestResourcesResourceActionTypeDeleted    ListWorkRequestResourcesResourceActionTypeEnum = "DELETED"
	ListWorkRequestResourcesResourceActionTypeInProgress ListWorkRequestResourcesResourceActionTypeEnum = "IN_PROGRESS"
	ListWorkRequestResourcesResourceActionTypeFailed     ListWorkRequestResourcesResourceActionTypeEnum = "FAILED"
)

var mappingListWorkRequestResourcesResourceActionType = map[string]ListWorkRequestResourcesResourceActionTypeEnum{
	"CREATED":     ListWorkRequestResourcesResourceActionTypeCreated,
	"UPDATED":     ListWorkRequestResourcesResourceActionTypeUpdated,
	"DELETED":     ListWorkRequestResourcesResourceActionTypeDeleted,
	"IN_PROGRESS": ListWorkRequestResourcesResourceActionTypeInProgress,
	"FAILED":      ListWorkRequestResourcesResourceActionTypeFailed,
}

// GetListWorkRequestResourcesResourceActionTypeEnumValues Enumerates the set of values for ListWorkRequestResourcesResourceActionTypeEnum
func GetListWorkRequestResourcesResourceActionTypeEnumValues() []ListWorkRequestResourcesResourceActionTypeEnum {
	values := make([]ListWorkRequestResourcesResourceActionTypeEnum, 0)
	for _, v := range mappingListWorkRequestResourcesResourceActionType {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestResourcesSortByEnum Enum with underlying type: string
type ListWorkRequestResourcesSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestResourcesSortByEnum
const (
	ListWorkRequestResourcesSortByActionType ListWorkRequestResourcesSortByEnum = "ACTION_TYPE"
	ListWorkRequestResourcesSortByEntityType ListWorkRequestResourcesSortByEnum = "ENTITY_TYPE"
)

var mappingListWorkRequestResourcesSortBy = map[string]ListWorkRequestResourcesSortByEnum{
	"ACTION_TYPE": ListWorkRequestResourcesSortByActionType,
	"ENTITY_TYPE": ListWorkRequestResourcesSortByEntityType,
}

// GetListWorkRequestResourcesSortByEnumValues Enumerates the set of values for ListWorkRequestResourcesSortByEnum
func GetListWorkRequestResourcesSortByEnumValues() []ListWorkRequestResourcesSortByEnum {
	values := make([]ListWorkRequestResourcesSortByEnum, 0)
	for _, v := range mappingListWorkRequestResourcesSortBy {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestResourcesSortOrderEnum Enum with underlying type: string
type ListWorkRequestResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestResourcesSortOrderEnum
const (
	ListWorkRequestResourcesSortOrderAsc  ListWorkRequestResourcesSortOrderEnum = "ASC"
	ListWorkRequestResourcesSortOrderDesc ListWorkRequestResourcesSortOrderEnum = "DESC"
)

var mappingListWorkRequestResourcesSortOrder = map[string]ListWorkRequestResourcesSortOrderEnum{
	"ASC":  ListWorkRequestResourcesSortOrderAsc,
	"DESC": ListWorkRequestResourcesSortOrderDesc,
}

// GetListWorkRequestResourcesSortOrderEnumValues Enumerates the set of values for ListWorkRequestResourcesSortOrderEnum
func GetListWorkRequestResourcesSortOrderEnumValues() []ListWorkRequestResourcesSortOrderEnum {
	values := make([]ListWorkRequestResourcesSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestResourcesSortOrder {
		values = append(values, v)
	}
	return values
}
