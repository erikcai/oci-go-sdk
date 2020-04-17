// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLogRulesRequest wrapper for the ListLogRules operation
type ListLogRulesRequest struct {

	// Compartment OCID to list resources in. Please see compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// OCID of logRule
	LogRuleId *string `mandatory:"false" contributesTo:"query" name:"logRuleId"`

	// Resource name
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Lifecycle state of the log rule
	LifecycleState ListLogRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListLogRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogRulesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogRulesResponse wrapper for the ListLogRules operation
type ListLogRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LogRuleSummary instances
	Items []LogRuleSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLogRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogRulesLifecycleStateEnum Enum with underlying type: string
type ListLogRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogRulesLifecycleStateEnum
const (
	ListLogRulesLifecycleStateCreating ListLogRulesLifecycleStateEnum = "CREATING"
	ListLogRulesLifecycleStateActive   ListLogRulesLifecycleStateEnum = "ACTIVE"
	ListLogRulesLifecycleStateUpdating ListLogRulesLifecycleStateEnum = "UPDATING"
	ListLogRulesLifecycleStateInactive ListLogRulesLifecycleStateEnum = "INACTIVE"
	ListLogRulesLifecycleStateDeleting ListLogRulesLifecycleStateEnum = "DELETING"
)

var mappingListLogRulesLifecycleState = map[string]ListLogRulesLifecycleStateEnum{
	"CREATING": ListLogRulesLifecycleStateCreating,
	"ACTIVE":   ListLogRulesLifecycleStateActive,
	"UPDATING": ListLogRulesLifecycleStateUpdating,
	"INACTIVE": ListLogRulesLifecycleStateInactive,
	"DELETING": ListLogRulesLifecycleStateDeleting,
}

// GetListLogRulesLifecycleStateEnumValues Enumerates the set of values for ListLogRulesLifecycleStateEnum
func GetListLogRulesLifecycleStateEnumValues() []ListLogRulesLifecycleStateEnum {
	values := make([]ListLogRulesLifecycleStateEnum, 0)
	for _, v := range mappingListLogRulesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListLogRulesSortByEnum Enum with underlying type: string
type ListLogRulesSortByEnum string

// Set of constants representing the allowable values for ListLogRulesSortByEnum
const (
	ListLogRulesSortByTimecreated ListLogRulesSortByEnum = "timeCreated"
	ListLogRulesSortByDisplayname ListLogRulesSortByEnum = "displayName"
)

var mappingListLogRulesSortBy = map[string]ListLogRulesSortByEnum{
	"timeCreated": ListLogRulesSortByTimecreated,
	"displayName": ListLogRulesSortByDisplayname,
}

// GetListLogRulesSortByEnumValues Enumerates the set of values for ListLogRulesSortByEnum
func GetListLogRulesSortByEnumValues() []ListLogRulesSortByEnum {
	values := make([]ListLogRulesSortByEnum, 0)
	for _, v := range mappingListLogRulesSortBy {
		values = append(values, v)
	}
	return values
}

// ListLogRulesSortOrderEnum Enum with underlying type: string
type ListLogRulesSortOrderEnum string

// Set of constants representing the allowable values for ListLogRulesSortOrderEnum
const (
	ListLogRulesSortOrderAsc  ListLogRulesSortOrderEnum = "ASC"
	ListLogRulesSortOrderDesc ListLogRulesSortOrderEnum = "DESC"
)

var mappingListLogRulesSortOrder = map[string]ListLogRulesSortOrderEnum{
	"ASC":  ListLogRulesSortOrderAsc,
	"DESC": ListLogRulesSortOrderDesc,
}

// GetListLogRulesSortOrderEnumValues Enumerates the set of values for ListLogRulesSortOrderEnum
func GetListLogRulesSortOrderEnumValues() []ListLogRulesSortOrderEnum {
	values := make([]ListLogRulesSortOrderEnum, 0)
	for _, v := range mappingListLogRulesSortOrder {
		values = append(values, v)
	}
	return values
}
