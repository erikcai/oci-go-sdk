// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package compdocsapi

import (
	"github.com/oracle/oci-go-sdk/v33/common"
	"net/http"
)

// ListComplianceDocumentsRequest wrapper for the ListComplianceDocuments operation
type ListComplianceDocumentsRequest struct {

	// A filter that returns only compliance documents with a document type that exactly matches the specified document type.
	Type ComplianceDocumentTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// A filter that returns only compliance documents with a platform type that exactly match the specified platform.
	Platform ComplianceDocumentPlatformEnum `mandatory:"false" contributesTo:"query" name:"platform" omitEmpty:"true"`

	// A filter that returns only compliance documents with a name that exactly match the specified name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter that returns only compliance documents with a lifecycle state that currently exactly matches
	// the specified lifecycle state.
	LifecycleState ComplianceDocumentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListComplianceDocumentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order. The default order for `TIMECREATED` is descending. The default order for `DISPLAYNAME`
	// is ascending. If no value is specified, results are sorted by the time created by default.
	SortBy ListComplianceDocumentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListComplianceDocumentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListComplianceDocumentsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListComplianceDocumentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListComplianceDocumentsResponse wrapper for the ListComplianceDocuments operation
type ListComplianceDocumentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ComplianceDocumentSummary instances
	Items []ComplianceDocumentSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListComplianceDocumentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListComplianceDocumentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListComplianceDocumentsSortOrderEnum Enum with underlying type: string
type ListComplianceDocumentsSortOrderEnum string

// Set of constants representing the allowable values for ListComplianceDocumentsSortOrderEnum
const (
	ListComplianceDocumentsSortOrderAsc  ListComplianceDocumentsSortOrderEnum = "ASC"
	ListComplianceDocumentsSortOrderDesc ListComplianceDocumentsSortOrderEnum = "DESC"
)

var mappingListComplianceDocumentsSortOrder = map[string]ListComplianceDocumentsSortOrderEnum{
	"ASC":  ListComplianceDocumentsSortOrderAsc,
	"DESC": ListComplianceDocumentsSortOrderDesc,
}

// GetListComplianceDocumentsSortOrderEnumValues Enumerates the set of values for ListComplianceDocumentsSortOrderEnum
func GetListComplianceDocumentsSortOrderEnumValues() []ListComplianceDocumentsSortOrderEnum {
	values := make([]ListComplianceDocumentsSortOrderEnum, 0)
	for _, v := range mappingListComplianceDocumentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListComplianceDocumentsSortByEnum Enum with underlying type: string
type ListComplianceDocumentsSortByEnum string

// Set of constants representing the allowable values for ListComplianceDocumentsSortByEnum
const (
	ListComplianceDocumentsSortByTimecreated ListComplianceDocumentsSortByEnum = "timeCreated"
	ListComplianceDocumentsSortByName        ListComplianceDocumentsSortByEnum = "name"
)

var mappingListComplianceDocumentsSortBy = map[string]ListComplianceDocumentsSortByEnum{
	"timeCreated": ListComplianceDocumentsSortByTimecreated,
	"name":        ListComplianceDocumentsSortByName,
}

// GetListComplianceDocumentsSortByEnumValues Enumerates the set of values for ListComplianceDocumentsSortByEnum
func GetListComplianceDocumentsSortByEnumValues() []ListComplianceDocumentsSortByEnum {
	values := make([]ListComplianceDocumentsSortByEnum, 0)
	for _, v := range mappingListComplianceDocumentsSortBy {
		values = append(values, v)
	}
	return values
}
