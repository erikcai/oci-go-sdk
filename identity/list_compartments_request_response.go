// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListCompartmentsRequest wrapper for the ListCompartments operation
type ListCompartmentsRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// if ACCESSIBLE is specified as accessLevel, those compartments under the compartment id
	// where the client has at least one resource INSPECT access (or greater) directly
	// or indirectly (the access permissions are in sub-compartments) are returned;
	// otherwise, all compartments directly under the tenancy id are returned.
	// And if ACCESSIBLE is specified as accessLevel, it will restrict the fields returned if the user doesn't have inspect access on that compartment"
	// The current default is 'ANY' (which will be deprecated at some point in the future).
	AccessLevel ListCompartmentsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The flag query parameter applies to fields that reference hierarchical entities,
	// Every compartment has a parent compartment, which may be self-referential for tencancy root compartments.
	// Specifies whether or not resources that are associated with descendants of the specified value are matched:
	// False means no traversal of the <property> hierarchy, and is the default value.
	// True traverses the <property> hierarchy as far down as possible.
	// Only Service Principal is allowed to use this query parameters.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListCompartmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListCompartmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState CompartmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCompartmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCompartmentsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCompartmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCompartmentsResponse wrapper for the ListCompartments operation
type ListCompartmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []Compartment instances
	Items []Compartment `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCompartmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCompartmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCompartmentsAccessLevelEnum Enum with underlying type: string
type ListCompartmentsAccessLevelEnum string

// Set of constants representing the allowable values for ListCompartmentsAccessLevelEnum
const (
	ListCompartmentsAccessLevelAny        ListCompartmentsAccessLevelEnum = "ANY"
	ListCompartmentsAccessLevelAccessible ListCompartmentsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCompartmentsAccessLevel = map[string]ListCompartmentsAccessLevelEnum{
	"ANY":        ListCompartmentsAccessLevelAny,
	"ACCESSIBLE": ListCompartmentsAccessLevelAccessible,
}

// GetListCompartmentsAccessLevelEnumValues Enumerates the set of values for ListCompartmentsAccessLevelEnum
func GetListCompartmentsAccessLevelEnumValues() []ListCompartmentsAccessLevelEnum {
	values := make([]ListCompartmentsAccessLevelEnum, 0)
	for _, v := range mappingListCompartmentsAccessLevel {
		values = append(values, v)
	}
	return values
}

// ListCompartmentsSortByEnum Enum with underlying type: string
type ListCompartmentsSortByEnum string

// Set of constants representing the allowable values for ListCompartmentsSortByEnum
const (
	ListCompartmentsSortByTimecreated ListCompartmentsSortByEnum = "TIMECREATED"
	ListCompartmentsSortByName        ListCompartmentsSortByEnum = "NAME"
)

var mappingListCompartmentsSortBy = map[string]ListCompartmentsSortByEnum{
	"TIMECREATED": ListCompartmentsSortByTimecreated,
	"NAME":        ListCompartmentsSortByName,
}

// GetListCompartmentsSortByEnumValues Enumerates the set of values for ListCompartmentsSortByEnum
func GetListCompartmentsSortByEnumValues() []ListCompartmentsSortByEnum {
	values := make([]ListCompartmentsSortByEnum, 0)
	for _, v := range mappingListCompartmentsSortBy {
		values = append(values, v)
	}
	return values
}

// ListCompartmentsSortOrderEnum Enum with underlying type: string
type ListCompartmentsSortOrderEnum string

// Set of constants representing the allowable values for ListCompartmentsSortOrderEnum
const (
	ListCompartmentsSortOrderAsc  ListCompartmentsSortOrderEnum = "ASC"
	ListCompartmentsSortOrderDesc ListCompartmentsSortOrderEnum = "DESC"
)

var mappingListCompartmentsSortOrder = map[string]ListCompartmentsSortOrderEnum{
	"ASC":  ListCompartmentsSortOrderAsc,
	"DESC": ListCompartmentsSortOrderDesc,
}

// GetListCompartmentsSortOrderEnumValues Enumerates the set of values for ListCompartmentsSortOrderEnum
func GetListCompartmentsSortOrderEnumValues() []ListCompartmentsSortOrderEnum {
	values := make([]ListCompartmentsSortOrderEnum, 0)
	for _, v := range mappingListCompartmentsSortOrder {
		values = append(values, v)
	}
	return values
}
