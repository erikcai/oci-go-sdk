// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operationsinsights

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// SummarizeSqlInsightsRequest wrapper for the SummarizeSqlInsights operation
type SummarizeSqlInsightsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D
	DatabaseType []SummarizeSqlInsightsDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Filter sqls by percentage of db time.
	DatabaseTimePctGreaterThan *float64 `mandatory:"false" contributesTo:"query" name:"databaseTimePctGreaterThan"`

	// Specify time period in ISO 8601 format with respect to current time.
	// Default is last 30 days represented by P30D.
	// If timeInterval is specified, then timeIntervalStart and timeIntervalEnd will be ignored.
	// Examples  P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months), . Maximum value allowed is 25 months prior to current time (P25M).
	AnalysisTimeInterval *string `mandatory:"false" contributesTo:"query" name:"analysisTimeInterval"`

	// Analysis start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// The minimum allowed value is 2 years prior to the current day.
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	// If analysisTimeInterval is specified, this parameter is ignored.
	TimeIntervalStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalStart"`

	// Analysis end time in UTC in ISO 8601 format(exclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd are used together.
	// If timeIntervalEnd is not specified, current time is used as timeIntervalEnd.
	TimeIntervalEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalEnd"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeSqlInsightsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeSqlInsightsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeSqlInsightsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeSqlInsightsResponse wrapper for the SummarizeSqlInsights operation
type SummarizeSqlInsightsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlInsightAggregationCollection instances
	SqlInsightAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeSqlInsightsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeSqlInsightsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeSqlInsightsDatabaseTypeEnum Enum with underlying type: string
type SummarizeSqlInsightsDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeSqlInsightsDatabaseTypeEnum
const (
	SummarizeSqlInsightsDatabaseTypeAdwS SummarizeSqlInsightsDatabaseTypeEnum = "ADW-S"
	SummarizeSqlInsightsDatabaseTypeAtpS SummarizeSqlInsightsDatabaseTypeEnum = "ATP-S"
	SummarizeSqlInsightsDatabaseTypeAdwD SummarizeSqlInsightsDatabaseTypeEnum = "ADW-D"
	SummarizeSqlInsightsDatabaseTypeAtpD SummarizeSqlInsightsDatabaseTypeEnum = "ATP-D"
)

var mappingSummarizeSqlInsightsDatabaseType = map[string]SummarizeSqlInsightsDatabaseTypeEnum{
	"ADW-S": SummarizeSqlInsightsDatabaseTypeAdwS,
	"ATP-S": SummarizeSqlInsightsDatabaseTypeAtpS,
	"ADW-D": SummarizeSqlInsightsDatabaseTypeAdwD,
	"ATP-D": SummarizeSqlInsightsDatabaseTypeAtpD,
}

// GetSummarizeSqlInsightsDatabaseTypeEnumValues Enumerates the set of values for SummarizeSqlInsightsDatabaseTypeEnum
func GetSummarizeSqlInsightsDatabaseTypeEnumValues() []SummarizeSqlInsightsDatabaseTypeEnum {
	values := make([]SummarizeSqlInsightsDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeSqlInsightsDatabaseType {
		values = append(values, v)
	}
	return values
}
