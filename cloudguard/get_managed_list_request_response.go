// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetManagedListRequest wrapper for the GetManagedList operation
type GetManagedListRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The cloudguard list OCID to be passed in the request.
	ManagedListId *string `mandatory:"true" contributesTo:"path" name:"managedListId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder GetManagedListSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy GetManagedListSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetManagedListRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetManagedListRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetManagedListRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetManagedListResponse wrapper for the GetManagedList operation
type GetManagedListResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedList instances
	ManagedList `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetManagedListResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetManagedListResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetManagedListSortOrderEnum Enum with underlying type: string
type GetManagedListSortOrderEnum string

// Set of constants representing the allowable values for GetManagedListSortOrderEnum
const (
	GetManagedListSortOrderAsc  GetManagedListSortOrderEnum = "ASC"
	GetManagedListSortOrderDesc GetManagedListSortOrderEnum = "DESC"
)

var mappingGetManagedListSortOrder = map[string]GetManagedListSortOrderEnum{
	"ASC":  GetManagedListSortOrderAsc,
	"DESC": GetManagedListSortOrderDesc,
}

// GetGetManagedListSortOrderEnumValues Enumerates the set of values for GetManagedListSortOrderEnum
func GetGetManagedListSortOrderEnumValues() []GetManagedListSortOrderEnum {
	values := make([]GetManagedListSortOrderEnum, 0)
	for _, v := range mappingGetManagedListSortOrder {
		values = append(values, v)
	}
	return values
}

// GetManagedListSortByEnum Enum with underlying type: string
type GetManagedListSortByEnum string

// Set of constants representing the allowable values for GetManagedListSortByEnum
const (
	GetManagedListSortByTimecreated GetManagedListSortByEnum = "timeCreated"
	GetManagedListSortByDisplayname GetManagedListSortByEnum = "displayName"
)

var mappingGetManagedListSortBy = map[string]GetManagedListSortByEnum{
	"timeCreated": GetManagedListSortByTimecreated,
	"displayName": GetManagedListSortByDisplayname,
}

// GetGetManagedListSortByEnumValues Enumerates the set of values for GetManagedListSortByEnum
func GetGetManagedListSortByEnumValues() []GetManagedListSortByEnum {
	values := make([]GetManagedListSortByEnum, 0)
	for _, v := range mappingGetManagedListSortBy {
		values = append(values, v)
	}
	return values
}
