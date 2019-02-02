// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAutonomousPodsRequest wrapper for the ListAutonomousPods operation
type ListAutonomousPodsRequest struct {

	// The compartment OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The Autonomous Container Database OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	AutonomousDbSystemId *string `mandatory:"false" contributesTo:"query" name:"autonomousDbSystemId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	// **Note:** If you do not include the availability domain filter, the resources are grouped by availability domain, then sorted.
	SortBy ListAutonomousPodsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousPodsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState AutonomousPodSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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

func (request ListAutonomousPodsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousPodsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousPodsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAutonomousPodsResponse wrapper for the ListAutonomousPods operation
type ListAutonomousPodsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousPodSummary instances
	Items []AutonomousPodSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousPodsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousPodsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousPodsSortByEnum Enum with underlying type: string
type ListAutonomousPodsSortByEnum string

// Set of constants representing the allowable values for ListAutonomousPodsSortByEnum
const (
	ListAutonomousPodsSortByTimecreated ListAutonomousPodsSortByEnum = "TIMECREATED"
	ListAutonomousPodsSortByDisplayname ListAutonomousPodsSortByEnum = "DISPLAYNAME"
)

var mappingListAutonomousPodsSortBy = map[string]ListAutonomousPodsSortByEnum{
	"TIMECREATED": ListAutonomousPodsSortByTimecreated,
	"DISPLAYNAME": ListAutonomousPodsSortByDisplayname,
}

// GetListAutonomousPodsSortByEnumValues Enumerates the set of values for ListAutonomousPodsSortByEnum
func GetListAutonomousPodsSortByEnumValues() []ListAutonomousPodsSortByEnum {
	values := make([]ListAutonomousPodsSortByEnum, 0)
	for _, v := range mappingListAutonomousPodsSortBy {
		values = append(values, v)
	}
	return values
}

// ListAutonomousPodsSortOrderEnum Enum with underlying type: string
type ListAutonomousPodsSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousPodsSortOrderEnum
const (
	ListAutonomousPodsSortOrderAsc  ListAutonomousPodsSortOrderEnum = "ASC"
	ListAutonomousPodsSortOrderDesc ListAutonomousPodsSortOrderEnum = "DESC"
)

var mappingListAutonomousPodsSortOrder = map[string]ListAutonomousPodsSortOrderEnum{
	"ASC":  ListAutonomousPodsSortOrderAsc,
	"DESC": ListAutonomousPodsSortOrderDesc,
}

// GetListAutonomousPodsSortOrderEnumValues Enumerates the set of values for ListAutonomousPodsSortOrderEnum
func GetListAutonomousPodsSortOrderEnumValues() []ListAutonomousPodsSortOrderEnum {
	values := make([]ListAutonomousPodsSortOrderEnum, 0)
	for _, v := range mappingListAutonomousPodsSortOrder {
		values = append(values, v)
	}
	return values
}
