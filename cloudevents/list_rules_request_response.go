// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package cloudevents

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListRulesRequest wrapper for the ListRules operation
type ListRulesRequest struct {

	// Compartment OCID
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return. The value must be between 1 and 50. The default is 10.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState RuleLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Only rules that have display names matching this parameter will be returned
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// which rule attribute to sort by (defaults to timeCreated)
	SortBy ListRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder *string `mandatory:"false" contributesTo:"query" name:"sortOrder"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRulesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRulesResponse wrapper for the ListRules operation
type ListRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []RuleSummary instances
	Items []RuleSummary `presentIn:"body"`

	// Pagination token
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRulesSortByEnum Enum with underlying type: string
type ListRulesSortByEnum string

// Set of constants representing the allowable values for ListRulesSortBy
const (
	ListRulesSortByTimeCreated ListRulesSortByEnum = "TIME_CREATED"
	ListRulesSortById          ListRulesSortByEnum = "ID"
	ListRulesSortByDisplayName ListRulesSortByEnum = "DISPLAY_NAME"
)

var mappingListRulesSortBy = map[string]ListRulesSortByEnum{
	"TIME_CREATED": ListRulesSortByTimeCreated,
	"ID":           ListRulesSortById,
	"DISPLAY_NAME": ListRulesSortByDisplayName,
}

// GetListRulesSortByEnumValues Enumerates the set of values for ListRulesSortBy
func GetListRulesSortByEnumValues() []ListRulesSortByEnum {
	values := make([]ListRulesSortByEnum, 0)
	for _, v := range mappingListRulesSortBy {
		values = append(values, v)
	}
	return values
}
