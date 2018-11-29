// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAutonomousDbSystemsRequest wrapper for the ListAutonomousDbSystems operation
type ListAutonomousDbSystemsRequest struct {

	// The compartment OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	// **Note:** If you do not include the availability domain filter, the resources are grouped by availability domain, then sorted.
	SortBy ListAutonomousDbSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousDbSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState AutonomousDbSystemSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given availability domain exactly.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousDbSystemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousDbSystemsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDbSystemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAutonomousDbSystemsResponse wrapper for the ListAutonomousDbSystems operation
type ListAutonomousDbSystemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousDbSystemSummary instances
	Items []AutonomousDbSystemSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousDbSystemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousDbSystemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousDbSystemsSortByEnum Enum with underlying type: string
type ListAutonomousDbSystemsSortByEnum string

// Set of constants representing the allowable values for ListAutonomousDbSystemsSortByEnum
const (
	ListAutonomousDbSystemsSortByTimecreated ListAutonomousDbSystemsSortByEnum = "TIMECREATED"
	ListAutonomousDbSystemsSortByDisplayname ListAutonomousDbSystemsSortByEnum = "DISPLAYNAME"
)

var mappingListAutonomousDbSystemsSortBy = map[string]ListAutonomousDbSystemsSortByEnum{
	"TIMECREATED": ListAutonomousDbSystemsSortByTimecreated,
	"DISPLAYNAME": ListAutonomousDbSystemsSortByDisplayname,
}

// GetListAutonomousDbSystemsSortByEnumValues Enumerates the set of values for ListAutonomousDbSystemsSortByEnum
func GetListAutonomousDbSystemsSortByEnumValues() []ListAutonomousDbSystemsSortByEnum {
	values := make([]ListAutonomousDbSystemsSortByEnum, 0)
	for _, v := range mappingListAutonomousDbSystemsSortBy {
		values = append(values, v)
	}
	return values
}

// ListAutonomousDbSystemsSortOrderEnum Enum with underlying type: string
type ListAutonomousDbSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousDbSystemsSortOrderEnum
const (
	ListAutonomousDbSystemsSortOrderAsc  ListAutonomousDbSystemsSortOrderEnum = "ASC"
	ListAutonomousDbSystemsSortOrderDesc ListAutonomousDbSystemsSortOrderEnum = "DESC"
)

var mappingListAutonomousDbSystemsSortOrder = map[string]ListAutonomousDbSystemsSortOrderEnum{
	"ASC":  ListAutonomousDbSystemsSortOrderAsc,
	"DESC": ListAutonomousDbSystemsSortOrderDesc,
}

// GetListAutonomousDbSystemsSortOrderEnumValues Enumerates the set of values for ListAutonomousDbSystemsSortOrderEnum
func GetListAutonomousDbSystemsSortOrderEnumValues() []ListAutonomousDbSystemsSortOrderEnum {
	values := make([]ListAutonomousDbSystemsSortOrderEnum, 0)
	for _, v := range mappingListAutonomousDbSystemsSortOrder {
		values = append(values, v)
	}
	return values
}