// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetAnalyticsClusterMemoryEstimateRequest wrapper for the GetAnalyticsClusterMemoryEstimate operation
type GetAnalyticsClusterMemoryEstimateRequest struct {

	// The MySQLaaS Instance OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	InstanceId *string `mandatory:"true" contributesTo:"path" name:"instanceId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAnalyticsClusterMemoryEstimateRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAnalyticsClusterMemoryEstimateRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAnalyticsClusterMemoryEstimateRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAnalyticsClusterMemoryEstimateResponse wrapper for the GetAnalyticsClusterMemoryEstimate operation
type GetAnalyticsClusterMemoryEstimateResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AnalyticsClusterMemoryEstimates instance
	AnalyticsClusterMemoryEstimates `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a specific request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAnalyticsClusterMemoryEstimateResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAnalyticsClusterMemoryEstimateResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
