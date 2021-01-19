// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"github.com/oracle/oci-go-sdk/v33/common"
	"net/http"
)

// SummarizeInstallationsRequest wrapper for the SummarizeInstallations operation
type SummarizeInstallationsRequest struct {

	// unique Fleet identifier
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

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
	SortOrder SummarizeInstallationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort installation views. Only one sort order may be provided. Default order for firstSeen, lastSeen and jreVersion is descending. Default order for jreDistribution and jreVendor is ascending. If no value is specified timeLastSeen is default.
	SortBy SummarizeInstallationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeInstallationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeInstallationsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeInstallationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeInstallationsResponse wrapper for the SummarizeInstallations operation
type SummarizeInstallationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstallationViewCollection instances
	InstallationViewCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeInstallationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeInstallationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeInstallationsSortOrderEnum Enum with underlying type: string
type SummarizeInstallationsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeInstallationsSortOrderEnum
const (
	SummarizeInstallationsSortOrderAsc  SummarizeInstallationsSortOrderEnum = "ASC"
	SummarizeInstallationsSortOrderDesc SummarizeInstallationsSortOrderEnum = "DESC"
)

var mappingSummarizeInstallationsSortOrder = map[string]SummarizeInstallationsSortOrderEnum{
	"ASC":  SummarizeInstallationsSortOrderAsc,
	"DESC": SummarizeInstallationsSortOrderDesc,
}

// GetSummarizeInstallationsSortOrderEnumValues Enumerates the set of values for SummarizeInstallationsSortOrderEnum
func GetSummarizeInstallationsSortOrderEnumValues() []SummarizeInstallationsSortOrderEnum {
	values := make([]SummarizeInstallationsSortOrderEnum, 0)
	for _, v := range mappingSummarizeInstallationsSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeInstallationsSortByEnum Enum with underlying type: string
type SummarizeInstallationsSortByEnum string

// Set of constants representing the allowable values for SummarizeInstallationsSortByEnum
const (
	SummarizeInstallationsSortByTimefirstseen   SummarizeInstallationsSortByEnum = "timeFirstSeen"
	SummarizeInstallationsSortByTimelastseen    SummarizeInstallationsSortByEnum = "timeLastSeen"
	SummarizeInstallationsSortByJreversion      SummarizeInstallationsSortByEnum = "jreVersion"
	SummarizeInstallationsSortByJredistribution SummarizeInstallationsSortByEnum = "jreDistribution"
	SummarizeInstallationsSortByJrevendor       SummarizeInstallationsSortByEnum = "jreVendor"
	SummarizeInstallationsSortByPath            SummarizeInstallationsSortByEnum = "path"
)

var mappingSummarizeInstallationsSortBy = map[string]SummarizeInstallationsSortByEnum{
	"timeFirstSeen":   SummarizeInstallationsSortByTimefirstseen,
	"timeLastSeen":    SummarizeInstallationsSortByTimelastseen,
	"jreVersion":      SummarizeInstallationsSortByJreversion,
	"jreDistribution": SummarizeInstallationsSortByJredistribution,
	"jreVendor":       SummarizeInstallationsSortByJrevendor,
	"path":            SummarizeInstallationsSortByPath,
}

// GetSummarizeInstallationsSortByEnumValues Enumerates the set of values for SummarizeInstallationsSortByEnum
func GetSummarizeInstallationsSortByEnumValues() []SummarizeInstallationsSortByEnum {
	values := make([]SummarizeInstallationsSortByEnum, 0)
	for _, v := range mappingSummarizeInstallationsSortBy {
		values = append(values, v)
	}
	return values
}
