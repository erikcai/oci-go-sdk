// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"net/http"
)

// ListOutboundConnectorsRequest wrapper for the ListOutboundConnectors operation
type ListOutboundConnectorsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListOutboundConnectorsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Example: `My resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter results by OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can choose either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by display name, results are
	// shown in ascending order.
	SortBy ListOutboundConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListOutboundConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOutboundConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOutboundConnectorsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOutboundConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOutboundConnectorsResponse wrapper for the ListOutboundConnectors operation
type ListOutboundConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []OutboundConnectorSummary instances
	Items []OutboundConnectorSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListOutboundConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOutboundConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOutboundConnectorsLifecycleStateEnum Enum with underlying type: string
type ListOutboundConnectorsLifecycleStateEnum string

// Set of constants representing the allowable values for ListOutboundConnectorsLifecycleStateEnum
const (
	ListOutboundConnectorsLifecycleStateCreating ListOutboundConnectorsLifecycleStateEnum = "CREATING"
	ListOutboundConnectorsLifecycleStateActive   ListOutboundConnectorsLifecycleStateEnum = "ACTIVE"
	ListOutboundConnectorsLifecycleStateDeleting ListOutboundConnectorsLifecycleStateEnum = "DELETING"
	ListOutboundConnectorsLifecycleStateDeleted  ListOutboundConnectorsLifecycleStateEnum = "DELETED"
	ListOutboundConnectorsLifecycleStateFailed   ListOutboundConnectorsLifecycleStateEnum = "FAILED"
)

var mappingListOutboundConnectorsLifecycleState = map[string]ListOutboundConnectorsLifecycleStateEnum{
	"CREATING": ListOutboundConnectorsLifecycleStateCreating,
	"ACTIVE":   ListOutboundConnectorsLifecycleStateActive,
	"DELETING": ListOutboundConnectorsLifecycleStateDeleting,
	"DELETED":  ListOutboundConnectorsLifecycleStateDeleted,
	"FAILED":   ListOutboundConnectorsLifecycleStateFailed,
}

// GetListOutboundConnectorsLifecycleStateEnumValues Enumerates the set of values for ListOutboundConnectorsLifecycleStateEnum
func GetListOutboundConnectorsLifecycleStateEnumValues() []ListOutboundConnectorsLifecycleStateEnum {
	values := make([]ListOutboundConnectorsLifecycleStateEnum, 0)
	for _, v := range mappingListOutboundConnectorsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListOutboundConnectorsSortByEnum Enum with underlying type: string
type ListOutboundConnectorsSortByEnum string

// Set of constants representing the allowable values for ListOutboundConnectorsSortByEnum
const (
	ListOutboundConnectorsSortByTimecreated ListOutboundConnectorsSortByEnum = "TIMECREATED"
	ListOutboundConnectorsSortByDisplayname ListOutboundConnectorsSortByEnum = "DISPLAYNAME"
)

var mappingListOutboundConnectorsSortBy = map[string]ListOutboundConnectorsSortByEnum{
	"TIMECREATED": ListOutboundConnectorsSortByTimecreated,
	"DISPLAYNAME": ListOutboundConnectorsSortByDisplayname,
}

// GetListOutboundConnectorsSortByEnumValues Enumerates the set of values for ListOutboundConnectorsSortByEnum
func GetListOutboundConnectorsSortByEnumValues() []ListOutboundConnectorsSortByEnum {
	values := make([]ListOutboundConnectorsSortByEnum, 0)
	for _, v := range mappingListOutboundConnectorsSortBy {
		values = append(values, v)
	}
	return values
}

// ListOutboundConnectorsSortOrderEnum Enum with underlying type: string
type ListOutboundConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListOutboundConnectorsSortOrderEnum
const (
	ListOutboundConnectorsSortOrderAsc  ListOutboundConnectorsSortOrderEnum = "ASC"
	ListOutboundConnectorsSortOrderDesc ListOutboundConnectorsSortOrderEnum = "DESC"
)

var mappingListOutboundConnectorsSortOrder = map[string]ListOutboundConnectorsSortOrderEnum{
	"ASC":  ListOutboundConnectorsSortOrderAsc,
	"DESC": ListOutboundConnectorsSortOrderDesc,
}

// GetListOutboundConnectorsSortOrderEnumValues Enumerates the set of values for ListOutboundConnectorsSortOrderEnum
func GetListOutboundConnectorsSortOrderEnumValues() []ListOutboundConnectorsSortOrderEnum {
	values := make([]ListOutboundConnectorsSortOrderEnum, 0)
	for _, v := range mappingListOutboundConnectorsSortOrder {
		values = append(values, v)
	}
	return values
}
