// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataflow

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// ListApplicationsRequest wrapper for the ListApplications operation
type ListApplicationsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request. If provided, the returned request ID will include this value.
	// Otherwise, a random request ID will be generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of results to return in a paginated `List` call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from the last `List` call
	// to sent back to server for getting the next page of results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field used to sort the results. Multiple fields are not supported.
	SortBy ListApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The ordering of results in ascending or descending order.
	SortOrder ListApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The query parameter for the Spark application name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `mandatory:"false" contributesTo:"query" name:"ownerPrincipalId"`

	// The displayName prefix.
	DisplayNameStartsWith *string `mandatory:"false" contributesTo:"query" name:"displayNameStartsWith"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplicationsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListApplicationsResponse wrapper for the ListApplications operation
type ListApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ApplicationSummary instances
	Items []ApplicationSummary `presentIn:"body"`

	// Retrieves the previous page of results.
	// When this header appears in the response, previous pages of results exist.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Retrieves the next page of results. When this header appears in the response,
	// additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplicationsSortByEnum Enum with underlying type: string
type ListApplicationsSortByEnum string

// Set of constants representing the allowable values for ListApplicationsSortByEnum
const (
	ListApplicationsSortByTimecreated ListApplicationsSortByEnum = "timeCreated"
	ListApplicationsSortByDisplayname ListApplicationsSortByEnum = "displayName"
	ListApplicationsSortByLanguage    ListApplicationsSortByEnum = "language"
)

var mappingListApplicationsSortBy = map[string]ListApplicationsSortByEnum{
	"timeCreated": ListApplicationsSortByTimecreated,
	"displayName": ListApplicationsSortByDisplayname,
	"language":    ListApplicationsSortByLanguage,
}

// GetListApplicationsSortByEnumValues Enumerates the set of values for ListApplicationsSortByEnum
func GetListApplicationsSortByEnumValues() []ListApplicationsSortByEnum {
	values := make([]ListApplicationsSortByEnum, 0)
	for _, v := range mappingListApplicationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListApplicationsSortOrderEnum Enum with underlying type: string
type ListApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListApplicationsSortOrderEnum
const (
	ListApplicationsSortOrderAsc  ListApplicationsSortOrderEnum = "ASC"
	ListApplicationsSortOrderDesc ListApplicationsSortOrderEnum = "DESC"
)

var mappingListApplicationsSortOrder = map[string]ListApplicationsSortOrderEnum{
	"ASC":  ListApplicationsSortOrderAsc,
	"DESC": ListApplicationsSortOrderDesc,
}

// GetListApplicationsSortOrderEnumValues Enumerates the set of values for ListApplicationsSortOrderEnum
func GetListApplicationsSortOrderEnumValues() []ListApplicationsSortOrderEnum {
	values := make([]ListApplicationsSortOrderEnum, 0)
	for _, v := range mappingListApplicationsSortOrder {
		values = append(values, v)
	}
	return values
}
