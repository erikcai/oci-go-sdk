// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetJobLogsRequest wrapper for the GetJobLogs operation
type GetJobLogsRequest struct {

	// Job OCID
	JobId *string `mandatory:"true" contributesTo:"path" name:"jobId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return logs only of the given type.
	Type []string `contributesTo:"query" name:"type" collectionFormat:"multi"`

	// A filter to return logs only of the given level or of greater severity.
	LevelGreaterThanOrEqualTo LogEntryLevelEnum `mandatory:"false" contributesTo:"query" name:"levelGreaterThanOrEqualTo" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder GetJobLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call. For information about pagination, see
	// List Pagination (https://docs.us-phoenix-1.oraclecloud.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call. For information about
	// pagination, see List Pagination (https://docs.us-phoenix-1.oraclecloud.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The timestamp before which logs will not be returned.
	TimestampGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timestampGreaterThanOrEqualTo"`

	// The timestamp after which logs will not be returned.
	TimestampLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timestampLessThanOrEqualTo"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetJobLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJobLogsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetJobLogsResponse wrapper for the GetJobLogs operation
type GetJobLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LogEntry instances
	Items []LogEntry `presentIn:"body"`

	// Unique identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination token for the next page of items
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response GetJobLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJobLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetJobLogsSortOrderEnum Enum with underlying type: string
type GetJobLogsSortOrderEnum string

// Set of constants representing the allowable values for GetJobLogsSortOrder
const (
	GetJobLogsSortOrderAsc  GetJobLogsSortOrderEnum = "ASC"
	GetJobLogsSortOrderDesc GetJobLogsSortOrderEnum = "DESC"
)

var mappingGetJobLogsSortOrder = map[string]GetJobLogsSortOrderEnum{
	"ASC":  GetJobLogsSortOrderAsc,
	"DESC": GetJobLogsSortOrderDesc,
}

// GetGetJobLogsSortOrderEnumValues Enumerates the set of values for GetJobLogsSortOrder
func GetGetJobLogsSortOrderEnumValues() []GetJobLogsSortOrderEnum {
	values := make([]GetJobLogsSortOrderEnum, 0)
	for _, v := range mappingGetJobLogsSortOrder {
		values = append(values, v)
	}
	return values
}
