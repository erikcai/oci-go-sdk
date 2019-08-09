// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListJobExecutionsRequest wrapper for the ListJobExecutions operation
type ListJobExecutionsRequest struct {

	// unique Catalog identifier
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique Job key.
	JobKey *string `mandatory:"true" contributesTo:"path" name:"jobKey"`

	// Job Execution Lifecycle State.
	LifecycleState JobExecutionStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the Resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the Resource was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// Id (OCID) of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// Id of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Job type.
	JobType JobTypeEnum `mandatory:"false" contributesTo:"query" name:"jobType" omitEmpty:"true"`

	// Sub-Type of this job execution.
	JobSubType *string `mandatory:"false" contributesTo:"query" name:"jobSubType"`

	// The unique key of the parent execution or null if this Job Execution has no parent.
	ParentKey *string `mandatory:"false" contributesTo:"query" name:"parentKey"`

	// Time that the Job Execution was started or in the case of a future time, the time in which the job will start.
	// An RFC3339 formatted datetime string.
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// Time that the Job Execution ended or null if the job is still running or hasn't run yet.
	// An RFC3339 formatted datetime string.
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// Error code returned from the Job Execution ot null if job is still running or didn't return an error.
	ErrorCode *string `mandatory:"false" contributesTo:"query" name:"errorCode"`

	// Error message returned from the Job Execution ot null if job is still running or didn't return an error.
	ErrorMessage *string `mandatory:"false" contributesTo:"query" name:"errorMessage"`

	// Process identifier related to the Job Execution.
	ProcessKey *string `mandatory:"false" contributesTo:"query" name:"processKey"`

	// The a URL of the job for accessing this resource and its status.
	ExternalUrl *string `mandatory:"false" contributesTo:"query" name:"externalUrl"`

	// Event that triggered the execution of this Job or null.
	EventKey *string `mandatory:"false" contributesTo:"query" name:"eventKey"`

	// Unique entity key.
	DataEntityKey *string `mandatory:"false" contributesTo:"query" name:"dataEntityKey"`

	// Used to control which fields are returned in a Job Execution summary response.
	Fields []ListJobExecutionsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListJobExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListJobExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListJobExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobExecutionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListJobExecutionsResponse wrapper for the ListJobExecutions operation
type ListJobExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []JobExecutionSummary instances
	Items []JobExecutionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobExecutionsFieldsEnum Enum with underlying type: string
type ListJobExecutionsFieldsEnum string

// Set of constants representing the allowable values for ListJobExecutionsFieldsEnum
const (
	ListJobExecutionsFieldsKey            ListJobExecutionsFieldsEnum = "key"
	ListJobExecutionsFieldsJobkey         ListJobExecutionsFieldsEnum = "jobKey"
	ListJobExecutionsFieldsJobtype        ListJobExecutionsFieldsEnum = "jobType"
	ListJobExecutionsFieldsParentkey      ListJobExecutionsFieldsEnum = "parentKey"
	ListJobExecutionsFieldsLifecyclestate ListJobExecutionsFieldsEnum = "lifecycleState"
	ListJobExecutionsFieldsTimecreated    ListJobExecutionsFieldsEnum = "timeCreated"
	ListJobExecutionsFieldsTimestarted    ListJobExecutionsFieldsEnum = "timeStarted"
	ListJobExecutionsFieldsTimeended      ListJobExecutionsFieldsEnum = "timeEnded"
	ListJobExecutionsFieldsUri            ListJobExecutionsFieldsEnum = "uri"
)

var mappingListJobExecutionsFields = map[string]ListJobExecutionsFieldsEnum{
	"key":            ListJobExecutionsFieldsKey,
	"jobKey":         ListJobExecutionsFieldsJobkey,
	"jobType":        ListJobExecutionsFieldsJobtype,
	"parentKey":      ListJobExecutionsFieldsParentkey,
	"lifecycleState": ListJobExecutionsFieldsLifecyclestate,
	"timeCreated":    ListJobExecutionsFieldsTimecreated,
	"timeStarted":    ListJobExecutionsFieldsTimestarted,
	"timeEnded":      ListJobExecutionsFieldsTimeended,
	"uri":            ListJobExecutionsFieldsUri,
}

// GetListJobExecutionsFieldsEnumValues Enumerates the set of values for ListJobExecutionsFieldsEnum
func GetListJobExecutionsFieldsEnumValues() []ListJobExecutionsFieldsEnum {
	values := make([]ListJobExecutionsFieldsEnum, 0)
	for _, v := range mappingListJobExecutionsFields {
		values = append(values, v)
	}
	return values
}

// ListJobExecutionsSortByEnum Enum with underlying type: string
type ListJobExecutionsSortByEnum string

// Set of constants representing the allowable values for ListJobExecutionsSortByEnum
const (
	ListJobExecutionsSortByTimecreated ListJobExecutionsSortByEnum = "TIMECREATED"
	ListJobExecutionsSortByDisplayname ListJobExecutionsSortByEnum = "DISPLAYNAME"
)

var mappingListJobExecutionsSortBy = map[string]ListJobExecutionsSortByEnum{
	"TIMECREATED": ListJobExecutionsSortByTimecreated,
	"DISPLAYNAME": ListJobExecutionsSortByDisplayname,
}

// GetListJobExecutionsSortByEnumValues Enumerates the set of values for ListJobExecutionsSortByEnum
func GetListJobExecutionsSortByEnumValues() []ListJobExecutionsSortByEnum {
	values := make([]ListJobExecutionsSortByEnum, 0)
	for _, v := range mappingListJobExecutionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListJobExecutionsSortOrderEnum Enum with underlying type: string
type ListJobExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListJobExecutionsSortOrderEnum
const (
	ListJobExecutionsSortOrderAsc  ListJobExecutionsSortOrderEnum = "ASC"
	ListJobExecutionsSortOrderDesc ListJobExecutionsSortOrderEnum = "DESC"
)

var mappingListJobExecutionsSortOrder = map[string]ListJobExecutionsSortOrderEnum{
	"ASC":  ListJobExecutionsSortOrderAsc,
	"DESC": ListJobExecutionsSortOrderDesc,
}

// GetListJobExecutionsSortOrderEnumValues Enumerates the set of values for ListJobExecutionsSortOrderEnum
func GetListJobExecutionsSortOrderEnumValues() []ListJobExecutionsSortOrderEnum {
	values := make([]ListJobExecutionsSortOrderEnum, 0)
	for _, v := range mappingListJobExecutionsSortOrder {
		values = append(values, v)
	}
	return values
}
