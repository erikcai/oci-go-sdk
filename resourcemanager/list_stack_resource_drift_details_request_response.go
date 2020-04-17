// Copyright (c) 2016, 2018, , Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListStackResourceDriftDetailsRequest wrapper for the ListStackResourceDriftDetails operation
type ListStackResourceDriftDetailsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack.
	StackId *string `mandatory:"true" contributesTo:"path" name:"stackId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only resources that match the given drift status. The value is case-insensitive.
	// Allowable values -
	//   - MODIFIED
	//   - IN_SYNC
	//   - DELETED
	ResourceDriftStatus StackResourceDriftSummaryResourceDriftStatusEnum `mandatory:"false" contributesTo:"query" name:"resourceDriftStatus" omitEmpty:"true"`

	// Display name on which to query.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Specifies the field on which to sort.
	// By default, `TIMECREATED` is ordered descending.
	// By default, `DISPLAYNAME` is ordered ascending. Note that you can sort only on one field.
	SortBy ListStackResourceDriftDetailsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListStackResourceDriftDetailsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStackResourceDriftDetailsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStackResourceDriftDetailsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStackResourceDriftDetailsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListStackResourceDriftDetailsResponse wrapper for the ListStackResourceDriftDetails operation
type ListStackResourceDriftDetailsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []StackResourceDriftCollection instances
	Items []StackResourceDriftCollection `presentIn:"body"`

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of paginated list items. If the `opc-next-page`
	// header appears in the response, additional pages of results remain.
	// To receive the next page, include the header value in the `page` param.
	// If the `opc-next-page` header does not appear in the response, there
	// are no more list items to get. For more information about list pagination,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStackResourceDriftDetailsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStackResourceDriftDetailsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStackResourceDriftDetailsSortByEnum Enum with underlying type: string
type ListStackResourceDriftDetailsSortByEnum string

// Set of constants representing the allowable values for ListStackResourceDriftDetailsSortByEnum
const (
	ListStackResourceDriftDetailsSortByTimecreated ListStackResourceDriftDetailsSortByEnum = "TIMECREATED"
	ListStackResourceDriftDetailsSortByDisplayname ListStackResourceDriftDetailsSortByEnum = "DISPLAYNAME"
)

var mappingListStackResourceDriftDetailsSortBy = map[string]ListStackResourceDriftDetailsSortByEnum{
	"TIMECREATED": ListStackResourceDriftDetailsSortByTimecreated,
	"DISPLAYNAME": ListStackResourceDriftDetailsSortByDisplayname,
}

// GetListStackResourceDriftDetailsSortByEnumValues Enumerates the set of values for ListStackResourceDriftDetailsSortByEnum
func GetListStackResourceDriftDetailsSortByEnumValues() []ListStackResourceDriftDetailsSortByEnum {
	values := make([]ListStackResourceDriftDetailsSortByEnum, 0)
	for _, v := range mappingListStackResourceDriftDetailsSortBy {
		values = append(values, v)
	}
	return values
}

// ListStackResourceDriftDetailsSortOrderEnum Enum with underlying type: string
type ListStackResourceDriftDetailsSortOrderEnum string

// Set of constants representing the allowable values for ListStackResourceDriftDetailsSortOrderEnum
const (
	ListStackResourceDriftDetailsSortOrderAsc  ListStackResourceDriftDetailsSortOrderEnum = "ASC"
	ListStackResourceDriftDetailsSortOrderDesc ListStackResourceDriftDetailsSortOrderEnum = "DESC"
)

var mappingListStackResourceDriftDetailsSortOrder = map[string]ListStackResourceDriftDetailsSortOrderEnum{
	"ASC":  ListStackResourceDriftDetailsSortOrderAsc,
	"DESC": ListStackResourceDriftDetailsSortOrderDesc,
}

// GetListStackResourceDriftDetailsSortOrderEnumValues Enumerates the set of values for ListStackResourceDriftDetailsSortOrderEnum
func GetListStackResourceDriftDetailsSortOrderEnumValues() []ListStackResourceDriftDetailsSortOrderEnum {
	values := make([]ListStackResourceDriftDetailsSortOrderEnum, 0)
	for _, v := range mappingListStackResourceDriftDetailsSortOrder {
		values = append(values, v)
	}
	return values
}
