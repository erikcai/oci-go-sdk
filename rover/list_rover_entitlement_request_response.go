// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"net/http"
)

// ListRoverEntitlementRequest wrapper for the ListRoverEntitlement operation
type ListRoverEntitlementRequest struct {

	// compartment id
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// filtering by Rover Device Entitlement id
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// filtering by displayName
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListRoverEntitlementSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRoverEntitlementRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRoverEntitlementRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRoverEntitlementRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRoverEntitlementResponse wrapper for the ListRoverEntitlement operation
type ListRoverEntitlementResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []RoverEntitlementSummary instances
	Items []RoverEntitlementSummary `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRoverEntitlementResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRoverEntitlementResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRoverEntitlementSortByEnum Enum with underlying type: string
type ListRoverEntitlementSortByEnum string

// Set of constants representing the allowable values for ListRoverEntitlementSortByEnum
const (
	ListRoverEntitlementSortByTimecreated ListRoverEntitlementSortByEnum = "timeCreated"
	ListRoverEntitlementSortByDisplayname ListRoverEntitlementSortByEnum = "displayName"
)

var mappingListRoverEntitlementSortBy = map[string]ListRoverEntitlementSortByEnum{
	"timeCreated": ListRoverEntitlementSortByTimecreated,
	"displayName": ListRoverEntitlementSortByDisplayname,
}

// GetListRoverEntitlementSortByEnumValues Enumerates the set of values for ListRoverEntitlementSortByEnum
func GetListRoverEntitlementSortByEnumValues() []ListRoverEntitlementSortByEnum {
	values := make([]ListRoverEntitlementSortByEnum, 0)
	for _, v := range mappingListRoverEntitlementSortBy {
		values = append(values, v)
	}
	return values
}
