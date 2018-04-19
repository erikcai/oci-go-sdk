// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package orchestration

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
	Type LogEntryTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// A filter to return logs only of the given level or of greater severity.
	Level LogEntryLevelEnum `mandatory:"false" contributesTo:"query" name:"level" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder GetJobLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call. For information about pagination, see
	// List Pagination (https://docs.us-phoenix-1.oraclecloud.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The timestamp before which logs will not be returned. If logs are ordered descending, then start is chronologically after the logs returned.
	StartTime *common.SDKTime `mandatory:"false" contributesTo:"query" name:"startTime"`

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

	// The []LogEntry instance
	Items []LogEntry `presentIn:"body"`

	// Unique identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
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
