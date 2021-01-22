// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v33/common"
	"net/http"
)

// GetMonitorResultRequest wrapper for the GetMonitorResult operation
type GetMonitorResultRequest struct {

	// The APM Domain Id the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// The OCID of the Monitor.
	MonitorId *string `mandatory:"true" contributesTo:"path" name:"monitorId"`

	// The vantagePoint name.
	VantagePoint *string `mandatory:"true" contributesTo:"query" name:"vantagePoint"`

	// The result type har or screenshot or log.
	ResultType *string `mandatory:"true" contributesTo:"query" name:"resultType"`

	// The result content type zip or raw.
	ResultContentType *string `mandatory:"true" contributesTo:"query" name:"resultContentType"`

	// The object posted time.
	ExecutionTime *string `mandatory:"true" contributesTo:"path" name:"executionTime"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetMonitorResultRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetMonitorResultRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetMonitorResultRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetMonitorResultResponse wrapper for the GetMonitorResult operation
type GetMonitorResultResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The MonitorResult instance
	MonitorResult `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetMonitorResultResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetMonitorResultResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
