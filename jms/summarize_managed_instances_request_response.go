// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// SummarizeManagedInstancesRequest wrapper for the SummarizeManagedInstances operation
type SummarizeManagedInstancesRequest struct {

	// unique Fleet identifier
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// the OCID of the managed instance entity
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// type of managed instance
	ManagedInstanceType SummarizeManagedInstancesManagedInstanceTypeEnum `mandatory:"false" contributesTo:"query" name:"managedInstanceType" omitEmpty:"true"`

	// The vendor of the related jre resource
	JreVendor *string `mandatory:"false" contributesTo:"query" name:"jreVendor"`

	// The distribution of the related jre resource
	JreDistribution *string `mandatory:"false" contributesTo:"query" name:"jreDistribution"`

	// The version of the related jre resource
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The absolute path to the root of the installation
	InstallationPath *string `mandatory:"false" contributesTo:"query" name:"installationPath"`

	// fleet-unique identifier of the related application entity
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// the start of the time period in which resources are searched in RFC3339 format
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// the end of the time period in which resources are searched in RFC3339 format
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SummarizeManagedInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by the managed instance views. Only one sort order may be provided. Default order for firstSeen and lastSeen is descending. If no value is specified timeLastSeen is default.
	SortBy SummarizeManagedInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeManagedInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeManagedInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeManagedInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeManagedInstancesResponse wrapper for the SummarizeManagedInstances operation
type SummarizeManagedInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceViewCollection instances
	ManagedInstanceViewCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeManagedInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeManagedInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeManagedInstancesManagedInstanceTypeEnum Enum with underlying type: string
type SummarizeManagedInstancesManagedInstanceTypeEnum string

// Set of constants representing the allowable values for SummarizeManagedInstancesManagedInstanceTypeEnum
const (
	SummarizeManagedInstancesManagedInstanceTypePolarisAgent SummarizeManagedInstancesManagedInstanceTypeEnum = "POLARIS_AGENT"
)

var mappingSummarizeManagedInstancesManagedInstanceType = map[string]SummarizeManagedInstancesManagedInstanceTypeEnum{
	"POLARIS_AGENT": SummarizeManagedInstancesManagedInstanceTypePolarisAgent,
}

// GetSummarizeManagedInstancesManagedInstanceTypeEnumValues Enumerates the set of values for SummarizeManagedInstancesManagedInstanceTypeEnum
func GetSummarizeManagedInstancesManagedInstanceTypeEnumValues() []SummarizeManagedInstancesManagedInstanceTypeEnum {
	values := make([]SummarizeManagedInstancesManagedInstanceTypeEnum, 0)
	for _, v := range mappingSummarizeManagedInstancesManagedInstanceType {
		values = append(values, v)
	}
	return values
}

// SummarizeManagedInstancesSortOrderEnum Enum with underlying type: string
type SummarizeManagedInstancesSortOrderEnum string

// Set of constants representing the allowable values for SummarizeManagedInstancesSortOrderEnum
const (
	SummarizeManagedInstancesSortOrderAsc  SummarizeManagedInstancesSortOrderEnum = "ASC"
	SummarizeManagedInstancesSortOrderDesc SummarizeManagedInstancesSortOrderEnum = "DESC"
)

var mappingSummarizeManagedInstancesSortOrder = map[string]SummarizeManagedInstancesSortOrderEnum{
	"ASC":  SummarizeManagedInstancesSortOrderAsc,
	"DESC": SummarizeManagedInstancesSortOrderDesc,
}

// GetSummarizeManagedInstancesSortOrderEnumValues Enumerates the set of values for SummarizeManagedInstancesSortOrderEnum
func GetSummarizeManagedInstancesSortOrderEnumValues() []SummarizeManagedInstancesSortOrderEnum {
	values := make([]SummarizeManagedInstancesSortOrderEnum, 0)
	for _, v := range mappingSummarizeManagedInstancesSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeManagedInstancesSortByEnum Enum with underlying type: string
type SummarizeManagedInstancesSortByEnum string

// Set of constants representing the allowable values for SummarizeManagedInstancesSortByEnum
const (
	SummarizeManagedInstancesSortByTimefirstseen SummarizeManagedInstancesSortByEnum = "timeFirstSeen"
	SummarizeManagedInstancesSortByTimelastseen  SummarizeManagedInstancesSortByEnum = "timeLastSeen"
)

var mappingSummarizeManagedInstancesSortBy = map[string]SummarizeManagedInstancesSortByEnum{
	"timeFirstSeen": SummarizeManagedInstancesSortByTimefirstseen,
	"timeLastSeen":  SummarizeManagedInstancesSortByTimelastseen,
}

// GetSummarizeManagedInstancesSortByEnumValues Enumerates the set of values for SummarizeManagedInstancesSortByEnum
func GetSummarizeManagedInstancesSortByEnumValues() []SummarizeManagedInstancesSortByEnum {
	values := make([]SummarizeManagedInstancesSortByEnum, 0)
	for _, v := range mappingSummarizeManagedInstancesSortBy {
		values = append(values, v)
	}
	return values
}
