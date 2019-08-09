// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListJobsRequest wrapper for the ListJobs operation
type ListJobsRequest struct {

	// unique Catalog identifier
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Job Lifecycle State.
	LifecycleState JobLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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

	// Unique Job Definition key.
	JobDefinitionKey *string `mandatory:"false" contributesTo:"query" name:"jobDefinitionKey"`

	// Schedule specified in the cron expression format that has seven fields for second , minute , hour , day-of-month , month , day-of-week , year .
	// It can also include special characters like * for all and ? for any . There are also pre-defined schedules that can be specified using
	// special strings. For example, @hourly will run the job every hour.
	ScheduleCronExpression *int `mandatory:"false" contributesTo:"query" name:"scheduleCronExpression"`

	// Time of the day the execution is scheduled. An RFC3339 formatted datetime string.
	TimeScheduled *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduled"`

	// Date that the schedule should be operational. An RFC3339 formatted datetime string.
	TimeScheduleBegin *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduleBegin"`

	// Date that the schedule should end from being operational. An RFC3339 formatted datetime string.
	TimeScheduleEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduleEnd"`

	// Type of the Job Schedule.
	ScheduleType JobScheduleTypeEnum `mandatory:"false" contributesTo:"query" name:"scheduleType" omitEmpty:"true"`

	// Unique connection key.
	ConnectionKey *string `mandatory:"false" contributesTo:"query" name:"connectionKey"`

	// Used to control which fields are returned in a Job summary response.
	Fields []ListJobsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListJobsResponse wrapper for the ListJobs operation
type ListJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []JobSummary instances
	Items []JobSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobsFieldsEnum Enum with underlying type: string
type ListJobsFieldsEnum string

// Set of constants representing the allowable values for ListJobsFieldsEnum
const (
	ListJobsFieldsKey              ListJobsFieldsEnum = "key"
	ListJobsFieldsDisplayname      ListJobsFieldsEnum = "displayName"
	ListJobsFieldsDescription      ListJobsFieldsEnum = "description"
	ListJobsFieldsCatalogid        ListJobsFieldsEnum = "catalogId"
	ListJobsFieldsJobdefinitionkey ListJobsFieldsEnum = "jobDefinitionKey"
	ListJobsFieldsLifecyclestate   ListJobsFieldsEnum = "lifecycleState"
	ListJobsFieldsTimecreated      ListJobsFieldsEnum = "timeCreated"
	ListJobsFieldsJobtype          ListJobsFieldsEnum = "jobType"
	ListJobsFieldsScheduletype     ListJobsFieldsEnum = "scheduleType"
	ListJobsFieldsUri              ListJobsFieldsEnum = "uri"
)

var mappingListJobsFields = map[string]ListJobsFieldsEnum{
	"key":              ListJobsFieldsKey,
	"displayName":      ListJobsFieldsDisplayname,
	"description":      ListJobsFieldsDescription,
	"catalogId":        ListJobsFieldsCatalogid,
	"jobDefinitionKey": ListJobsFieldsJobdefinitionkey,
	"lifecycleState":   ListJobsFieldsLifecyclestate,
	"timeCreated":      ListJobsFieldsTimecreated,
	"jobType":          ListJobsFieldsJobtype,
	"scheduleType":     ListJobsFieldsScheduletype,
	"uri":              ListJobsFieldsUri,
}

// GetListJobsFieldsEnumValues Enumerates the set of values for ListJobsFieldsEnum
func GetListJobsFieldsEnumValues() []ListJobsFieldsEnum {
	values := make([]ListJobsFieldsEnum, 0)
	for _, v := range mappingListJobsFields {
		values = append(values, v)
	}
	return values
}

// ListJobsSortByEnum Enum with underlying type: string
type ListJobsSortByEnum string

// Set of constants representing the allowable values for ListJobsSortByEnum
const (
	ListJobsSortByTimecreated ListJobsSortByEnum = "TIMECREATED"
	ListJobsSortByDisplayname ListJobsSortByEnum = "DISPLAYNAME"
)

var mappingListJobsSortBy = map[string]ListJobsSortByEnum{
	"TIMECREATED": ListJobsSortByTimecreated,
	"DISPLAYNAME": ListJobsSortByDisplayname,
}

// GetListJobsSortByEnumValues Enumerates the set of values for ListJobsSortByEnum
func GetListJobsSortByEnumValues() []ListJobsSortByEnum {
	values := make([]ListJobsSortByEnum, 0)
	for _, v := range mappingListJobsSortBy {
		values = append(values, v)
	}
	return values
}

// ListJobsSortOrderEnum Enum with underlying type: string
type ListJobsSortOrderEnum string

// Set of constants representing the allowable values for ListJobsSortOrderEnum
const (
	ListJobsSortOrderAsc  ListJobsSortOrderEnum = "ASC"
	ListJobsSortOrderDesc ListJobsSortOrderEnum = "DESC"
)

var mappingListJobsSortOrder = map[string]ListJobsSortOrderEnum{
	"ASC":  ListJobsSortOrderAsc,
	"DESC": ListJobsSortOrderDesc,
}

// GetListJobsSortOrderEnumValues Enumerates the set of values for ListJobsSortOrderEnum
func GetListJobsSortOrderEnumValues() []ListJobsSortOrderEnum {
	values := make([]ListJobsSortOrderEnum, 0)
	for _, v := range mappingListJobsSortOrder {
		values = append(values, v)
	}
	return values
}
