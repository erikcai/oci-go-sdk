// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"github.com/oracle/oci-go-sdk/v33/common"
	"net/http"
)

// SummarizeApplicationsRequest wrapper for the SummarizeApplications operation
type SummarizeApplicationsRequest struct {

	// unique Fleet identifier
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// fleet-unique identifier of the application entity
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// A filter to return only resources that entirely match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// type of application
	ApplicationType *string `mandatory:"false" contributesTo:"query" name:"applicationType"`

	// The vendor of the related jre resource
	JreVendor *string `mandatory:"false" contributesTo:"query" name:"jreVendor"`

	// The distribution of the related jre resource
	JreDistribution *string `mandatory:"false" contributesTo:"query" name:"jreDistribution"`

	// The version of the related jre resource
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The absolute path to the root of the installation
	InstallationPath *string `mandatory:"false" contributesTo:"query" name:"installationPath"`

	// the OCID of the related managed instance entity
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// the start of the time period in which resources are searched in RFC3339 format
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// the end of the time period in which resources are searched in RFC3339 format
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SummarizeApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort application views. Only one sort order may be provided. Default order for firstSeen and lastSeen is descending. Default order for displayName is ascending. If no value is specified timeLastSeen is default.
	SortBy SummarizeApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeApplicationsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeApplicationsResponse wrapper for the SummarizeApplications operation
type SummarizeApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApplicationViewCollection instances
	ApplicationViewCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeApplicationsSortOrderEnum Enum with underlying type: string
type SummarizeApplicationsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeApplicationsSortOrderEnum
const (
	SummarizeApplicationsSortOrderAsc  SummarizeApplicationsSortOrderEnum = "ASC"
	SummarizeApplicationsSortOrderDesc SummarizeApplicationsSortOrderEnum = "DESC"
)

var mappingSummarizeApplicationsSortOrder = map[string]SummarizeApplicationsSortOrderEnum{
	"ASC":  SummarizeApplicationsSortOrderAsc,
	"DESC": SummarizeApplicationsSortOrderDesc,
}

// GetSummarizeApplicationsSortOrderEnumValues Enumerates the set of values for SummarizeApplicationsSortOrderEnum
func GetSummarizeApplicationsSortOrderEnumValues() []SummarizeApplicationsSortOrderEnum {
	values := make([]SummarizeApplicationsSortOrderEnum, 0)
	for _, v := range mappingSummarizeApplicationsSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeApplicationsSortByEnum Enum with underlying type: string
type SummarizeApplicationsSortByEnum string

// Set of constants representing the allowable values for SummarizeApplicationsSortByEnum
const (
	SummarizeApplicationsSortByTimefirstseen SummarizeApplicationsSortByEnum = "timeFirstSeen"
	SummarizeApplicationsSortByTimelastseen  SummarizeApplicationsSortByEnum = "timeLastSeen"
	SummarizeApplicationsSortByDisplayname   SummarizeApplicationsSortByEnum = "displayName"
)

var mappingSummarizeApplicationsSortBy = map[string]SummarizeApplicationsSortByEnum{
	"timeFirstSeen": SummarizeApplicationsSortByTimefirstseen,
	"timeLastSeen":  SummarizeApplicationsSortByTimelastseen,
	"displayName":   SummarizeApplicationsSortByDisplayname,
}

// GetSummarizeApplicationsSortByEnumValues Enumerates the set of values for SummarizeApplicationsSortByEnum
func GetSummarizeApplicationsSortByEnumValues() []SummarizeApplicationsSortByEnum {
	values := make([]SummarizeApplicationsSortByEnum, 0)
	for _, v := range mappingSummarizeApplicationsSortBy {
		values = append(values, v)
	}
	return values
}
