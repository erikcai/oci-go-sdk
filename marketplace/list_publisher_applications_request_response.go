// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListPublisherApplicationsRequest wrapper for the ListPublisherApplications operation
type ListPublisherApplicationsRequest struct {

	// Unique identifier of the publisher for which applications are required.
	PublisherId *string `mandatory:"true" contributesTo:"path" name:"publisherId"`

	// Given product name will be exactly matched with internal name of the oracle cloud service.
	Product *string `mandatory:"false" contributesTo:"query" name:"product"`

	// This is one of the filter options. It specifies the application categories for which search has to be done. Categories are dependent on products. Categories can be used for filtering only when product value is present.  Multiple values can be passed in comma separated.
	Categories *string `mandatory:"false" contributesTo:"query" name:"categories"`

	// This is one of the filter options. The product filter codes for which search has to be done. Filters are dependent on products. filters can be used for filtering only when product value is Not-Null. Multi-valued parameter. It is a semi colon separated list in the following formatEg: filters= filtercode1=filtervaluecode1,filtervaluecode2; filtercode2= filtervaluecode3
	Filters *string `mandatory:"false" contributesTo:"query" name:"filters"`

	// Limit tells how many records to be returned.
	// Limit should be greater than zero and less than or equal to hundred (default=30).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-prev-page` or `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Applications display order by field.
	SortBy ListPublisherApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListPublisherApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Application resource types.
	ResourceTypes ListPublisherApplicationsResourceTypesEnum `mandatory:"false" contributesTo:"query" name:"resourceTypes" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPublisherApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPublisherApplicationsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPublisherApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPublisherApplicationsResponse wrapper for the ListPublisherApplications operation
type ListPublisherApplicationsResponse struct {

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

func (response ListPublisherApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPublisherApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPublisherApplicationsSortByEnum Enum with underlying type: string
type ListPublisherApplicationsSortByEnum string

// Set of constants representing the allowable values for ListPublisherApplicationsSortByEnum
const (
	ListPublisherApplicationsSortByName        ListPublisherApplicationsSortByEnum = "NAME"
	ListPublisherApplicationsSortByReleasedate ListPublisherApplicationsSortByEnum = "RELEASEDATE"
	ListPublisherApplicationsSortByRating      ListPublisherApplicationsSortByEnum = "RATING"
)

var mappingListPublisherApplicationsSortBy = map[string]ListPublisherApplicationsSortByEnum{
	"NAME":        ListPublisherApplicationsSortByName,
	"RELEASEDATE": ListPublisherApplicationsSortByReleasedate,
	"RATING":      ListPublisherApplicationsSortByRating,
}

// GetListPublisherApplicationsSortByEnumValues Enumerates the set of values for ListPublisherApplicationsSortByEnum
func GetListPublisherApplicationsSortByEnumValues() []ListPublisherApplicationsSortByEnum {
	values := make([]ListPublisherApplicationsSortByEnum, 0)
	for _, v := range mappingListPublisherApplicationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListPublisherApplicationsSortOrderEnum Enum with underlying type: string
type ListPublisherApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListPublisherApplicationsSortOrderEnum
const (
	ListPublisherApplicationsSortOrderAsc  ListPublisherApplicationsSortOrderEnum = "ASC"
	ListPublisherApplicationsSortOrderDesc ListPublisherApplicationsSortOrderEnum = "DESC"
)

var mappingListPublisherApplicationsSortOrder = map[string]ListPublisherApplicationsSortOrderEnum{
	"ASC":  ListPublisherApplicationsSortOrderAsc,
	"DESC": ListPublisherApplicationsSortOrderDesc,
}

// GetListPublisherApplicationsSortOrderEnumValues Enumerates the set of values for ListPublisherApplicationsSortOrderEnum
func GetListPublisherApplicationsSortOrderEnumValues() []ListPublisherApplicationsSortOrderEnum {
	values := make([]ListPublisherApplicationsSortOrderEnum, 0)
	for _, v := range mappingListPublisherApplicationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListPublisherApplicationsResourceTypesEnum Enum with underlying type: string
type ListPublisherApplicationsResourceTypesEnum string

// Set of constants representing the allowable values for ListPublisherApplicationsResourceTypesEnum
const (
	ListPublisherApplicationsResourceTypesTerraform       ListPublisherApplicationsResourceTypesEnum = "terraform"
	ListPublisherApplicationsResourceTypesOcimachineimage ListPublisherApplicationsResourceTypesEnum = "ocimachineimage"
)

var mappingListPublisherApplicationsResourceTypes = map[string]ListPublisherApplicationsResourceTypesEnum{
	"terraform":       ListPublisherApplicationsResourceTypesTerraform,
	"ocimachineimage": ListPublisherApplicationsResourceTypesOcimachineimage,
}

// GetListPublisherApplicationsResourceTypesEnumValues Enumerates the set of values for ListPublisherApplicationsResourceTypesEnum
func GetListPublisherApplicationsResourceTypesEnumValues() []ListPublisherApplicationsResourceTypesEnum {
	values := make([]ListPublisherApplicationsResourceTypesEnum, 0)
	for _, v := range mappingListPublisherApplicationsResourceTypes {
		values = append(values, v)
	}
	return values
}
