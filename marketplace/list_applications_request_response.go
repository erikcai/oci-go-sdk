// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListApplicationsRequest wrapper for the ListApplications operation
type ListApplicationsRequest struct {

	// The keyword for which application search has to be done.
	Keyword *string `mandatory:"false" contributesTo:"query" name:"keyword"`

	// Gets the featured applications only if isFeatured is set to true.
	IsFeatured *bool `mandatory:"false" contributesTo:"query" name:"isFeatured"`

	// Given product name will be exactly matched with internal name of the oracle cloud service.
	Product *string `mandatory:"false" contributesTo:"query" name:"product"`

	// This is one of the filter options. It specifies the application categories for which search has to be done. Categories are dependent on products. Categories can be used for filtering only when product value is present.  Multiple values can be passed in comma separated.
	Categories *string `mandatory:"false" contributesTo:"query" name:"categories"`

	// This is one of the filter options. The product filter codes for which search has to be done. Filters are dependent on products. filters can be used for filtering only when product value is Not-Null. Multi-valued parameter. It is a semi colon separated list in the following formatEg: filters= filtercode1=filtervaluecode1,filtervaluecode2; filtercode2= filtervaluecode3
	Filters *string `mandatory:"false" contributesTo:"query" name:"filters"`

	// Application resource types.
	ResourceTypes ListApplicationsResourceTypesEnum `mandatory:"false" contributesTo:"query" name:"resourceTypes" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Limit tells how many records to be returned.
	// Limit should be greater than zero and less than or equal to hundred (default=30).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-prev-page` or `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Applications display order by field.
	SortBy ListApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

	// Unique identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination token for the next page of items
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination token for the previous page of items
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplicationsResourceTypesEnum Enum with underlying type: string
type ListApplicationsResourceTypesEnum string

// Set of constants representing the allowable values for ListApplicationsResourceTypesEnum
const (
	ListApplicationsResourceTypesTerraform       ListApplicationsResourceTypesEnum = "terraform"
	ListApplicationsResourceTypesOcimachineimage ListApplicationsResourceTypesEnum = "ocimachineimage"
)

var mappingListApplicationsResourceTypes = map[string]ListApplicationsResourceTypesEnum{
	"terraform":       ListApplicationsResourceTypesTerraform,
	"ocimachineimage": ListApplicationsResourceTypesOcimachineimage,
}

// GetListApplicationsResourceTypesEnumValues Enumerates the set of values for ListApplicationsResourceTypesEnum
func GetListApplicationsResourceTypesEnumValues() []ListApplicationsResourceTypesEnum {
	values := make([]ListApplicationsResourceTypesEnum, 0)
	for _, v := range mappingListApplicationsResourceTypes {
		values = append(values, v)
	}
	return values
}

// ListApplicationsSortByEnum Enum with underlying type: string
type ListApplicationsSortByEnum string

// Set of constants representing the allowable values for ListApplicationsSortByEnum
const (
	ListApplicationsSortByName        ListApplicationsSortByEnum = "NAME"
	ListApplicationsSortByReleasedate ListApplicationsSortByEnum = "RELEASEDATE"
	ListApplicationsSortByRating      ListApplicationsSortByEnum = "RATING"
)

var mappingListApplicationsSortBy = map[string]ListApplicationsSortByEnum{
	"NAME":        ListApplicationsSortByName,
	"RELEASEDATE": ListApplicationsSortByReleasedate,
	"RATING":      ListApplicationsSortByRating,
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
