// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListGlossaryTermRelationshipsRequest wrapper for the ListGlossaryTermRelationships operation
type ListGlossaryTermRelationshipsRequest struct {

	// unique Catalog identifier
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique Glossary key.
	GlossaryKey *string `mandatory:"true" contributesTo:"path" name:"glossaryKey"`

	// Unique Glossary Term key.
	TermKey *string `mandatory:"true" contributesTo:"path" name:"termKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState LifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Used to control which fields are returned in a Term Relationship summary response.
	Fields []ListGlossaryTermRelationshipsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListGlossaryTermRelationshipsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListGlossaryTermRelationshipsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGlossaryTermRelationshipsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGlossaryTermRelationshipsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGlossaryTermRelationshipsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListGlossaryTermRelationshipsResponse wrapper for the ListGlossaryTermRelationships operation
type ListGlossaryTermRelationshipsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []TermRelationshipSummary instances
	Items []TermRelationshipSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGlossaryTermRelationshipsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGlossaryTermRelationshipsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGlossaryTermRelationshipsFieldsEnum Enum with underlying type: string
type ListGlossaryTermRelationshipsFieldsEnum string

// Set of constants representing the allowable values for ListGlossaryTermRelationshipsFieldsEnum
const (
	ListGlossaryTermRelationshipsFieldsKey            ListGlossaryTermRelationshipsFieldsEnum = "key"
	ListGlossaryTermRelationshipsFieldsDisplayname    ListGlossaryTermRelationshipsFieldsEnum = "displayName"
	ListGlossaryTermRelationshipsFieldsDescription    ListGlossaryTermRelationshipsFieldsEnum = "description"
	ListGlossaryTermRelationshipsFieldsRelatedtermkey ListGlossaryTermRelationshipsFieldsEnum = "relatedTermKey"
	ListGlossaryTermRelationshipsFieldsParenttermkey  ListGlossaryTermRelationshipsFieldsEnum = "parentTermKey"
	ListGlossaryTermRelationshipsFieldsLifecyclestate ListGlossaryTermRelationshipsFieldsEnum = "lifecycleState"
	ListGlossaryTermRelationshipsFieldsTimecreated    ListGlossaryTermRelationshipsFieldsEnum = "timeCreated"
	ListGlossaryTermRelationshipsFieldsUri            ListGlossaryTermRelationshipsFieldsEnum = "uri"
)

var mappingListGlossaryTermRelationshipsFields = map[string]ListGlossaryTermRelationshipsFieldsEnum{
	"key":            ListGlossaryTermRelationshipsFieldsKey,
	"displayName":    ListGlossaryTermRelationshipsFieldsDisplayname,
	"description":    ListGlossaryTermRelationshipsFieldsDescription,
	"relatedTermKey": ListGlossaryTermRelationshipsFieldsRelatedtermkey,
	"parentTermKey":  ListGlossaryTermRelationshipsFieldsParenttermkey,
	"lifecycleState": ListGlossaryTermRelationshipsFieldsLifecyclestate,
	"timeCreated":    ListGlossaryTermRelationshipsFieldsTimecreated,
	"uri":            ListGlossaryTermRelationshipsFieldsUri,
}

// GetListGlossaryTermRelationshipsFieldsEnumValues Enumerates the set of values for ListGlossaryTermRelationshipsFieldsEnum
func GetListGlossaryTermRelationshipsFieldsEnumValues() []ListGlossaryTermRelationshipsFieldsEnum {
	values := make([]ListGlossaryTermRelationshipsFieldsEnum, 0)
	for _, v := range mappingListGlossaryTermRelationshipsFields {
		values = append(values, v)
	}
	return values
}

// ListGlossaryTermRelationshipsSortByEnum Enum with underlying type: string
type ListGlossaryTermRelationshipsSortByEnum string

// Set of constants representing the allowable values for ListGlossaryTermRelationshipsSortByEnum
const (
	ListGlossaryTermRelationshipsSortByTimecreated ListGlossaryTermRelationshipsSortByEnum = "TIMECREATED"
	ListGlossaryTermRelationshipsSortByDisplayname ListGlossaryTermRelationshipsSortByEnum = "DISPLAYNAME"
)

var mappingListGlossaryTermRelationshipsSortBy = map[string]ListGlossaryTermRelationshipsSortByEnum{
	"TIMECREATED": ListGlossaryTermRelationshipsSortByTimecreated,
	"DISPLAYNAME": ListGlossaryTermRelationshipsSortByDisplayname,
}

// GetListGlossaryTermRelationshipsSortByEnumValues Enumerates the set of values for ListGlossaryTermRelationshipsSortByEnum
func GetListGlossaryTermRelationshipsSortByEnumValues() []ListGlossaryTermRelationshipsSortByEnum {
	values := make([]ListGlossaryTermRelationshipsSortByEnum, 0)
	for _, v := range mappingListGlossaryTermRelationshipsSortBy {
		values = append(values, v)
	}
	return values
}

// ListGlossaryTermRelationshipsSortOrderEnum Enum with underlying type: string
type ListGlossaryTermRelationshipsSortOrderEnum string

// Set of constants representing the allowable values for ListGlossaryTermRelationshipsSortOrderEnum
const (
	ListGlossaryTermRelationshipsSortOrderAsc  ListGlossaryTermRelationshipsSortOrderEnum = "ASC"
	ListGlossaryTermRelationshipsSortOrderDesc ListGlossaryTermRelationshipsSortOrderEnum = "DESC"
)

var mappingListGlossaryTermRelationshipsSortOrder = map[string]ListGlossaryTermRelationshipsSortOrderEnum{
	"ASC":  ListGlossaryTermRelationshipsSortOrderAsc,
	"DESC": ListGlossaryTermRelationshipsSortOrderDesc,
}

// GetListGlossaryTermRelationshipsSortOrderEnumValues Enumerates the set of values for ListGlossaryTermRelationshipsSortOrderEnum
func GetListGlossaryTermRelationshipsSortOrderEnumValues() []ListGlossaryTermRelationshipsSortOrderEnum {
	values := make([]ListGlossaryTermRelationshipsSortOrderEnum, 0)
	for _, v := range mappingListGlossaryTermRelationshipsSortOrder {
		values = append(values, v)
	}
	return values
}
