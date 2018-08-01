// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAppCatalogPublishersRequest wrapper for the ListAppCatalogPublishers operation
type ListAppCatalogPublishersRequest struct {

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the publisher that matches the given publisher name exactly.
	PublisherName *string `mandatory:"false" contributesTo:"query" name:"publisherName"`

	// A filter to return only publishers that match the given publisher type exactly. Valid types are OCI, ORACLE, TRUSTED, STANDARD.
	PublisherType *string `mandatory:"false" contributesTo:"query" name:"publisherType"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListAppCatalogPublishersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAppCatalogPublishersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAppCatalogPublishersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAppCatalogPublishersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAppCatalogPublishersResponse wrapper for the ListAppCatalogPublishers operation
type ListAppCatalogPublishersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AppCatalogPublisherSummary instances
	Items []AppCatalogPublisherSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAppCatalogPublishersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAppCatalogPublishersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAppCatalogPublishersSortOrderEnum Enum with underlying type: string
type ListAppCatalogPublishersSortOrderEnum string

// Set of constants representing the allowable values for ListAppCatalogPublishersSortOrder
const (
	ListAppCatalogPublishersSortOrderAsc  ListAppCatalogPublishersSortOrderEnum = "ASC"
	ListAppCatalogPublishersSortOrderDesc ListAppCatalogPublishersSortOrderEnum = "DESC"
)

var mappingListAppCatalogPublishersSortOrder = map[string]ListAppCatalogPublishersSortOrderEnum{
	"ASC":  ListAppCatalogPublishersSortOrderAsc,
	"DESC": ListAppCatalogPublishersSortOrderDesc,
}

// GetListAppCatalogPublishersSortOrderEnumValues Enumerates the set of values for ListAppCatalogPublishersSortOrder
func GetListAppCatalogPublishersSortOrderEnumValues() []ListAppCatalogPublishersSortOrderEnum {
	values := make([]ListAppCatalogPublishersSortOrderEnum, 0)
	for _, v := range mappingListAppCatalogPublishersSortOrder {
		values = append(values, v)
	}
	return values
}
