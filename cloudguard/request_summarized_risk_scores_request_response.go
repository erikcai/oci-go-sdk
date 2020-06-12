// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// RequestSummarizedRiskScoresRequest wrapper for the RequestSummarizedRiskScores operation
type RequestSummarizedRiskScoresRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel RequestSummarizedRiskScoresAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

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

func (request RequestSummarizedRiskScoresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RequestSummarizedRiskScoresRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RequestSummarizedRiskScoresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// RequestSummarizedRiskScoresResponse wrapper for the RequestSummarizedRiskScores operation
type RequestSummarizedRiskScoresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RiskScoreAggregationCollection instances
	RiskScoreAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response RequestSummarizedRiskScoresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RequestSummarizedRiskScoresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RequestSummarizedRiskScoresAccessLevelEnum Enum with underlying type: string
type RequestSummarizedRiskScoresAccessLevelEnum string

// Set of constants representing the allowable values for RequestSummarizedRiskScoresAccessLevelEnum
const (
	RequestSummarizedRiskScoresAccessLevelRestricted RequestSummarizedRiskScoresAccessLevelEnum = "RESTRICTED"
	RequestSummarizedRiskScoresAccessLevelAccessible RequestSummarizedRiskScoresAccessLevelEnum = "ACCESSIBLE"
)

var mappingRequestSummarizedRiskScoresAccessLevel = map[string]RequestSummarizedRiskScoresAccessLevelEnum{
	"RESTRICTED": RequestSummarizedRiskScoresAccessLevelRestricted,
	"ACCESSIBLE": RequestSummarizedRiskScoresAccessLevelAccessible,
}

// GetRequestSummarizedRiskScoresAccessLevelEnumValues Enumerates the set of values for RequestSummarizedRiskScoresAccessLevelEnum
func GetRequestSummarizedRiskScoresAccessLevelEnumValues() []RequestSummarizedRiskScoresAccessLevelEnum {
	values := make([]RequestSummarizedRiskScoresAccessLevelEnum, 0)
	for _, v := range mappingRequestSummarizedRiskScoresAccessLevel {
		values = append(values, v)
	}
	return values
}
