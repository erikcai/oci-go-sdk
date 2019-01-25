// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTagDefaultsRequest wrapper for the ListTagDefaults operation
type ListTagDefaultsRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for TAGDEFINITIONNAME is ascending.
	SortBy ListTagDefaultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListTagDefaultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the specified OCID exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID of the Tag Definition.
	TagDefinitionId *string `mandatory:"false" contributesTo:"query" name:"tagDefinitionId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTagDefaultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTagDefaultsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTagDefaultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTagDefaultsResponse wrapper for the ListTagDefaults operation
type ListTagDefaultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []TagDefaultSummary instances
	Items []TagDefaultSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of Tag Default values. When paging through a list, if this header appears in
	// the response, then a partial list might have been returned. Include this value as the `page` parameter
	// for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTagDefaultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTagDefaultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTagDefaultsSortByEnum Enum with underlying type: string
type ListTagDefaultsSortByEnum string

// Set of constants representing the allowable values for ListTagDefaultsSortByEnum
const (
	ListTagDefaultsSortByTagdefinitionname ListTagDefaultsSortByEnum = "TAGDEFINITIONNAME"
	ListTagDefaultsSortByName              ListTagDefaultsSortByEnum = "NAME"
)

var mappingListTagDefaultsSortBy = map[string]ListTagDefaultsSortByEnum{
	"TAGDEFINITIONNAME": ListTagDefaultsSortByTagdefinitionname,
	"NAME":              ListTagDefaultsSortByName,
}

// GetListTagDefaultsSortByEnumValues Enumerates the set of values for ListTagDefaultsSortByEnum
func GetListTagDefaultsSortByEnumValues() []ListTagDefaultsSortByEnum {
	values := make([]ListTagDefaultsSortByEnum, 0)
	for _, v := range mappingListTagDefaultsSortBy {
		values = append(values, v)
	}
	return values
}

// ListTagDefaultsSortOrderEnum Enum with underlying type: string
type ListTagDefaultsSortOrderEnum string

// Set of constants representing the allowable values for ListTagDefaultsSortOrderEnum
const (
	ListTagDefaultsSortOrderAsc  ListTagDefaultsSortOrderEnum = "ASC"
	ListTagDefaultsSortOrderDesc ListTagDefaultsSortOrderEnum = "DESC"
)

var mappingListTagDefaultsSortOrder = map[string]ListTagDefaultsSortOrderEnum{
	"ASC":  ListTagDefaultsSortOrderAsc,
	"DESC": ListTagDefaultsSortOrderDesc,
}

// GetListTagDefaultsSortOrderEnumValues Enumerates the set of values for ListTagDefaultsSortOrderEnum
func GetListTagDefaultsSortOrderEnumValues() []ListTagDefaultsSortOrderEnum {
	values := make([]ListTagDefaultsSortOrderEnum, 0)
	for _, v := range mappingListTagDefaultsSortOrder {
		values = append(values, v)
	}
	return values
}
