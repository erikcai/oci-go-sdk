// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package secrets

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSecretBundleVersionsRequest wrapper for the ListSecretBundleVersions operation
type ListSecretBundleVersionsRequest struct {

	// Secret OCID
	SecretId *string `mandatory:"true" contributesTo:"path" name:"secretId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated "List" call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call. For information about
	// pagination, see List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Time created is default ordered as descending. Display name is default ordered as ascending.
	SortBy ListSecretBundleVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListSecretBundleVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecretBundleVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecretBundleVersionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecretBundleVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSecretBundleVersionsResponse wrapper for the ListSecretBundleVersions operation
type ListSecretBundleVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SecretBundleVersionSummary instances
	Items []SecretBundleVersionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination token for the next page of items
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSecretBundleVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecretBundleVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecretBundleVersionsSortByEnum Enum with underlying type: string
type ListSecretBundleVersionsSortByEnum string

// Set of constants representing the allowable values for ListSecretBundleVersionsSortByEnum
const (
	ListSecretBundleVersionsSortByVersionNumber ListSecretBundleVersionsSortByEnum = "VERSION_NUMBER"
)

var mappingListSecretBundleVersionsSortBy = map[string]ListSecretBundleVersionsSortByEnum{
	"VERSION_NUMBER": ListSecretBundleVersionsSortByVersionNumber,
}

// GetListSecretBundleVersionsSortByEnumValues Enumerates the set of values for ListSecretBundleVersionsSortByEnum
func GetListSecretBundleVersionsSortByEnumValues() []ListSecretBundleVersionsSortByEnum {
	values := make([]ListSecretBundleVersionsSortByEnum, 0)
	for _, v := range mappingListSecretBundleVersionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListSecretBundleVersionsSortOrderEnum Enum with underlying type: string
type ListSecretBundleVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListSecretBundleVersionsSortOrderEnum
const (
	ListSecretBundleVersionsSortOrderAsc  ListSecretBundleVersionsSortOrderEnum = "ASC"
	ListSecretBundleVersionsSortOrderDesc ListSecretBundleVersionsSortOrderEnum = "DESC"
)

var mappingListSecretBundleVersionsSortOrder = map[string]ListSecretBundleVersionsSortOrderEnum{
	"ASC":  ListSecretBundleVersionsSortOrderAsc,
	"DESC": ListSecretBundleVersionsSortOrderDesc,
}

// GetListSecretBundleVersionsSortOrderEnumValues Enumerates the set of values for ListSecretBundleVersionsSortOrderEnum
func GetListSecretBundleVersionsSortOrderEnumValues() []ListSecretBundleVersionsSortOrderEnum {
	values := make([]ListSecretBundleVersionsSortOrderEnum, 0)
	for _, v := range mappingListSecretBundleVersionsSortOrder {
		values = append(values, v)
	}
	return values
}
