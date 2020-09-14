// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operationsinsights

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// SummarizeSqlStatisticsRequest wrapper for the SummarizeSqlStatistics operation
type SummarizeSqlStatisticsRequest struct {

	// The OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D
	DatabaseType []SummarizeSqlStatisticsDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Filter sqls by percentage of db time.
	DatabaseTimePctGreaterThan *float64 `mandatory:"false" contributesTo:"query" name:"databaseTimePctGreaterThan"`

	// One or more unique SQL_IDs for a SQL Statement.
	SqlIdentifier []string `contributesTo:"query" name:"sqlIdentifier" collectionFormat:"multi"`

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

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeSqlStatisticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to use when sorting SQL statistics.
	// Example: databaseTimeInSec
	SortBy SummarizeSqlStatisticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter sqls by one or more performance categories.
	Category []SummarizeSqlStatisticsCategoryEnum `contributesTo:"query" name:"category" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeSqlStatisticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeSqlStatisticsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeSqlStatisticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeSqlStatisticsResponse wrapper for the SummarizeSqlStatistics operation
type SummarizeSqlStatisticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlStatisticAggregationCollection instances
	SqlStatisticAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeSqlStatisticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeSqlStatisticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeSqlStatisticsDatabaseTypeEnum Enum with underlying type: string
type SummarizeSqlStatisticsDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeSqlStatisticsDatabaseTypeEnum
const (
	SummarizeSqlStatisticsDatabaseTypeAdwS SummarizeSqlStatisticsDatabaseTypeEnum = "ADW-S"
	SummarizeSqlStatisticsDatabaseTypeAtpS SummarizeSqlStatisticsDatabaseTypeEnum = "ATP-S"
	SummarizeSqlStatisticsDatabaseTypeAdwD SummarizeSqlStatisticsDatabaseTypeEnum = "ADW-D"
	SummarizeSqlStatisticsDatabaseTypeAtpD SummarizeSqlStatisticsDatabaseTypeEnum = "ATP-D"
)

var mappingSummarizeSqlStatisticsDatabaseType = map[string]SummarizeSqlStatisticsDatabaseTypeEnum{
	"ADW-S": SummarizeSqlStatisticsDatabaseTypeAdwS,
	"ATP-S": SummarizeSqlStatisticsDatabaseTypeAtpS,
	"ADW-D": SummarizeSqlStatisticsDatabaseTypeAdwD,
	"ATP-D": SummarizeSqlStatisticsDatabaseTypeAtpD,
}

// GetSummarizeSqlStatisticsDatabaseTypeEnumValues Enumerates the set of values for SummarizeSqlStatisticsDatabaseTypeEnum
func GetSummarizeSqlStatisticsDatabaseTypeEnumValues() []SummarizeSqlStatisticsDatabaseTypeEnum {
	values := make([]SummarizeSqlStatisticsDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeSqlStatisticsDatabaseType {
		values = append(values, v)
	}
	return values
}

// SummarizeSqlStatisticsSortOrderEnum Enum with underlying type: string
type SummarizeSqlStatisticsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeSqlStatisticsSortOrderEnum
const (
	SummarizeSqlStatisticsSortOrderAsc  SummarizeSqlStatisticsSortOrderEnum = "ASC"
	SummarizeSqlStatisticsSortOrderDesc SummarizeSqlStatisticsSortOrderEnum = "DESC"
)

var mappingSummarizeSqlStatisticsSortOrder = map[string]SummarizeSqlStatisticsSortOrderEnum{
	"ASC":  SummarizeSqlStatisticsSortOrderAsc,
	"DESC": SummarizeSqlStatisticsSortOrderDesc,
}

// GetSummarizeSqlStatisticsSortOrderEnumValues Enumerates the set of values for SummarizeSqlStatisticsSortOrderEnum
func GetSummarizeSqlStatisticsSortOrderEnumValues() []SummarizeSqlStatisticsSortOrderEnum {
	values := make([]SummarizeSqlStatisticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeSqlStatisticsSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeSqlStatisticsSortByEnum Enum with underlying type: string
type SummarizeSqlStatisticsSortByEnum string

// Set of constants representing the allowable values for SummarizeSqlStatisticsSortByEnum
const (
	SummarizeSqlStatisticsSortByDatabasetimeinsec                  SummarizeSqlStatisticsSortByEnum = "databaseTimeInSec"
	SummarizeSqlStatisticsSortByExecutionsperhour                  SummarizeSqlStatisticsSortByEnum = "executionsPerHour"
	SummarizeSqlStatisticsSortByExecutionscount                    SummarizeSqlStatisticsSortByEnum = "executionsCount"
	SummarizeSqlStatisticsSortByCputimeinsec                       SummarizeSqlStatisticsSortByEnum = "cpuTimeInSec"
	SummarizeSqlStatisticsSortByIotimeinsec                        SummarizeSqlStatisticsSortByEnum = "ioTimeInSec"
	SummarizeSqlStatisticsSortByInefficientwaittimeinsec           SummarizeSqlStatisticsSortByEnum = "inefficientWaitTimeInSec"
	SummarizeSqlStatisticsSortByResponsetimeinsec                  SummarizeSqlStatisticsSortByEnum = "responseTimeInSec"
	SummarizeSqlStatisticsSortByPlancount                          SummarizeSqlStatisticsSortByEnum = "planCount"
	SummarizeSqlStatisticsSortByVariability                        SummarizeSqlStatisticsSortByEnum = "variability"
	SummarizeSqlStatisticsSortByAverageactivesessions              SummarizeSqlStatisticsSortByEnum = "averageActiveSessions"
	SummarizeSqlStatisticsSortByDatabasetimepct                    SummarizeSqlStatisticsSortByEnum = "databaseTimePct"
	SummarizeSqlStatisticsSortByInefficiencyinpct                  SummarizeSqlStatisticsSortByEnum = "inefficiencyInPct"
	SummarizeSqlStatisticsSortByChangeincputimeinpct               SummarizeSqlStatisticsSortByEnum = "changeInCpuTimeInPct"
	SummarizeSqlStatisticsSortByChangeiniotimeinpct                SummarizeSqlStatisticsSortByEnum = "changeInIoTimeInPct"
	SummarizeSqlStatisticsSortByChangeininefficientwaittimeinpct   SummarizeSqlStatisticsSortByEnum = "changeInInefficientWaitTimeInPct"
	SummarizeSqlStatisticsSortByChangeinresponsetimeinpct          SummarizeSqlStatisticsSortByEnum = "changeInResponseTimeInPct"
	SummarizeSqlStatisticsSortByChangeinaverageactivesessionsinpct SummarizeSqlStatisticsSortByEnum = "changeInAverageActiveSessionsInPct"
	SummarizeSqlStatisticsSortByChangeinexecutionsperhourinpct     SummarizeSqlStatisticsSortByEnum = "changeInExecutionsPerHourInPct"
	SummarizeSqlStatisticsSortByChangeininefficiencyinpct          SummarizeSqlStatisticsSortByEnum = "changeInInefficiencyInPct"
)

var mappingSummarizeSqlStatisticsSortBy = map[string]SummarizeSqlStatisticsSortByEnum{
	"databaseTimeInSec":                  SummarizeSqlStatisticsSortByDatabasetimeinsec,
	"executionsPerHour":                  SummarizeSqlStatisticsSortByExecutionsperhour,
	"executionsCount":                    SummarizeSqlStatisticsSortByExecutionscount,
	"cpuTimeInSec":                       SummarizeSqlStatisticsSortByCputimeinsec,
	"ioTimeInSec":                        SummarizeSqlStatisticsSortByIotimeinsec,
	"inefficientWaitTimeInSec":           SummarizeSqlStatisticsSortByInefficientwaittimeinsec,
	"responseTimeInSec":                  SummarizeSqlStatisticsSortByResponsetimeinsec,
	"planCount":                          SummarizeSqlStatisticsSortByPlancount,
	"variability":                        SummarizeSqlStatisticsSortByVariability,
	"averageActiveSessions":              SummarizeSqlStatisticsSortByAverageactivesessions,
	"databaseTimePct":                    SummarizeSqlStatisticsSortByDatabasetimepct,
	"inefficiencyInPct":                  SummarizeSqlStatisticsSortByInefficiencyinpct,
	"changeInCpuTimeInPct":               SummarizeSqlStatisticsSortByChangeincputimeinpct,
	"changeInIoTimeInPct":                SummarizeSqlStatisticsSortByChangeiniotimeinpct,
	"changeInInefficientWaitTimeInPct":   SummarizeSqlStatisticsSortByChangeininefficientwaittimeinpct,
	"changeInResponseTimeInPct":          SummarizeSqlStatisticsSortByChangeinresponsetimeinpct,
	"changeInAverageActiveSessionsInPct": SummarizeSqlStatisticsSortByChangeinaverageactivesessionsinpct,
	"changeInExecutionsPerHourInPct":     SummarizeSqlStatisticsSortByChangeinexecutionsperhourinpct,
	"changeInInefficiencyInPct":          SummarizeSqlStatisticsSortByChangeininefficiencyinpct,
}

// GetSummarizeSqlStatisticsSortByEnumValues Enumerates the set of values for SummarizeSqlStatisticsSortByEnum
func GetSummarizeSqlStatisticsSortByEnumValues() []SummarizeSqlStatisticsSortByEnum {
	values := make([]SummarizeSqlStatisticsSortByEnum, 0)
	for _, v := range mappingSummarizeSqlStatisticsSortBy {
		values = append(values, v)
	}
	return values
}

// SummarizeSqlStatisticsCategoryEnum Enum with underlying type: string
type SummarizeSqlStatisticsCategoryEnum string

// Set of constants representing the allowable values for SummarizeSqlStatisticsCategoryEnum
const (
	SummarizeSqlStatisticsCategoryDegrading                                            SummarizeSqlStatisticsCategoryEnum = "DEGRADING"
	SummarizeSqlStatisticsCategoryVariant                                              SummarizeSqlStatisticsCategoryEnum = "VARIANT"
	SummarizeSqlStatisticsCategoryInefficient                                          SummarizeSqlStatisticsCategoryEnum = "INEFFICIENT"
	SummarizeSqlStatisticsCategoryChangingPlans                                        SummarizeSqlStatisticsCategoryEnum = "CHANGING_PLANS"
	SummarizeSqlStatisticsCategoryDegradingVariant                                     SummarizeSqlStatisticsCategoryEnum = "DEGRADING_VARIANT"
	SummarizeSqlStatisticsCategoryDegradingInefficient                                 SummarizeSqlStatisticsCategoryEnum = "DEGRADING_INEFFICIENT"
	SummarizeSqlStatisticsCategoryDegradingChangingPlans                               SummarizeSqlStatisticsCategoryEnum = "DEGRADING_CHANGING_PLANS"
	SummarizeSqlStatisticsCategoryDegradingIncreasingIo                                SummarizeSqlStatisticsCategoryEnum = "DEGRADING_INCREASING_IO"
	SummarizeSqlStatisticsCategoryDegradingIncreasingCpu                               SummarizeSqlStatisticsCategoryEnum = "DEGRADING_INCREASING_CPU"
	SummarizeSqlStatisticsCategoryDegradingIncreasingInefficientWait                   SummarizeSqlStatisticsCategoryEnum = "DEGRADING_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingIo                SummarizeSqlStatisticsCategoryEnum = "DEGRADING_CHANGING_PLANS_AND_INCREASING_IO"
	SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingCpu               SummarizeSqlStatisticsCategoryEnum = "DEGRADING_CHANGING_PLANS_AND_INCREASING_CPU"
	SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingInefficientWait   SummarizeSqlStatisticsCategoryEnum = "DEGRADING_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryVariantInefficient                                   SummarizeSqlStatisticsCategoryEnum = "VARIANT_INEFFICIENT"
	SummarizeSqlStatisticsCategoryVariantChangingPlans                                 SummarizeSqlStatisticsCategoryEnum = "VARIANT_CHANGING_PLANS"
	SummarizeSqlStatisticsCategoryVariantIncreasingIo                                  SummarizeSqlStatisticsCategoryEnum = "VARIANT_INCREASING_IO"
	SummarizeSqlStatisticsCategoryVariantIncreasingCpu                                 SummarizeSqlStatisticsCategoryEnum = "VARIANT_INCREASING_CPU"
	SummarizeSqlStatisticsCategoryVariantIncreasingInefficientWait                     SummarizeSqlStatisticsCategoryEnum = "VARIANT_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingIo                  SummarizeSqlStatisticsCategoryEnum = "VARIANT_CHANGING_PLANS_AND_INCREASING_IO"
	SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingCpu                 SummarizeSqlStatisticsCategoryEnum = "VARIANT_CHANGING_PLANS_AND_INCREASING_CPU"
	SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingInefficientWait     SummarizeSqlStatisticsCategoryEnum = "VARIANT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryInefficientChangingPlans                             SummarizeSqlStatisticsCategoryEnum = "INEFFICIENT_CHANGING_PLANS"
	SummarizeSqlStatisticsCategoryInefficientIncreasingInefficientWait                 SummarizeSqlStatisticsCategoryEnum = "INEFFICIENT_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryInefficientChangingPlansAndIncreasingInefficientWait SummarizeSqlStatisticsCategoryEnum = "INEFFICIENT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT"
)

var mappingSummarizeSqlStatisticsCategory = map[string]SummarizeSqlStatisticsCategoryEnum{
	"DEGRADING":                                   SummarizeSqlStatisticsCategoryDegrading,
	"VARIANT":                                     SummarizeSqlStatisticsCategoryVariant,
	"INEFFICIENT":                                 SummarizeSqlStatisticsCategoryInefficient,
	"CHANGING_PLANS":                              SummarizeSqlStatisticsCategoryChangingPlans,
	"DEGRADING_VARIANT":                           SummarizeSqlStatisticsCategoryDegradingVariant,
	"DEGRADING_INEFFICIENT":                       SummarizeSqlStatisticsCategoryDegradingInefficient,
	"DEGRADING_CHANGING_PLANS":                    SummarizeSqlStatisticsCategoryDegradingChangingPlans,
	"DEGRADING_INCREASING_IO":                     SummarizeSqlStatisticsCategoryDegradingIncreasingIo,
	"DEGRADING_INCREASING_CPU":                    SummarizeSqlStatisticsCategoryDegradingIncreasingCpu,
	"DEGRADING_INCREASING_INEFFICIENT_WAIT":       SummarizeSqlStatisticsCategoryDegradingIncreasingInefficientWait,
	"DEGRADING_CHANGING_PLANS_AND_INCREASING_IO":  SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingIo,
	"DEGRADING_CHANGING_PLANS_AND_INCREASING_CPU": SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingCpu,
	"DEGRADING_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT": SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingInefficientWait,
	"VARIANT_INEFFICIENT":                                        SummarizeSqlStatisticsCategoryVariantInefficient,
	"VARIANT_CHANGING_PLANS":                                     SummarizeSqlStatisticsCategoryVariantChangingPlans,
	"VARIANT_INCREASING_IO":                                      SummarizeSqlStatisticsCategoryVariantIncreasingIo,
	"VARIANT_INCREASING_CPU":                                     SummarizeSqlStatisticsCategoryVariantIncreasingCpu,
	"VARIANT_INCREASING_INEFFICIENT_WAIT":                        SummarizeSqlStatisticsCategoryVariantIncreasingInefficientWait,
	"VARIANT_CHANGING_PLANS_AND_INCREASING_IO":                   SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingIo,
	"VARIANT_CHANGING_PLANS_AND_INCREASING_CPU":                  SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingCpu,
	"VARIANT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT":     SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingInefficientWait,
	"INEFFICIENT_CHANGING_PLANS":                                 SummarizeSqlStatisticsCategoryInefficientChangingPlans,
	"INEFFICIENT_INCREASING_INEFFICIENT_WAIT":                    SummarizeSqlStatisticsCategoryInefficientIncreasingInefficientWait,
	"INEFFICIENT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT": SummarizeSqlStatisticsCategoryInefficientChangingPlansAndIncreasingInefficientWait,
}

// GetSummarizeSqlStatisticsCategoryEnumValues Enumerates the set of values for SummarizeSqlStatisticsCategoryEnum
func GetSummarizeSqlStatisticsCategoryEnumValues() []SummarizeSqlStatisticsCategoryEnum {
	values := make([]SummarizeSqlStatisticsCategoryEnum, 0)
	for _, v := range mappingSummarizeSqlStatisticsCategory {
		values = append(values, v)
	}
	return values
}
