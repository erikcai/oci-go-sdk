// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetMysqlaasInstanceLimitsCheckRequest wrapper for the GetMysqlaasInstanceLimitsCheck operation
type GetMysqlaasInstanceLimitsCheckRequest struct {

	// OCID of the MySQL DbSystem that we are trying to create. The DbSystem does not exist yet, so this
	// is just for logging purposes. The id will only be valid if the check passes.
	DbSystemId *string `mandatory:"true" contributesTo:"query" name:"dbSystemId"`

	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The Availability Domain where the MySQLaaS instance is going to be located.
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	ShapeName *string `mandatory:"true" contributesTo:"query" name:"shapeName"`

	DataStorageSizeInGBs *int `mandatory:"true" contributesTo:"query" name:"dataStorageSizeInGBs"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The shape determines resources to allocate to the MySQLaaS Analyics
	// Cluster nodes - CPU cores, memory.
	AnalyticsShapeName *string `mandatory:"false" contributesTo:"query" name:"analyticsShapeName"`

	// The number of analytics-processing nodes provisioned for the
	// MySQLaaS Analytics Cluster.
	AnalyticsClusterSize *int `mandatory:"false" contributesTo:"query" name:"analyticsClusterSize"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetMysqlaasInstanceLimitsCheckRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetMysqlaasInstanceLimitsCheckRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetMysqlaasInstanceLimitsCheckRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetMysqlaasInstanceLimitsCheckResponse wrapper for the GetMysqlaasInstanceLimitsCheck operation
type GetMysqlaasInstanceLimitsCheckResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The MysqlaasInstanceLimitsCheck instance
	MysqlaasInstanceLimitsCheck `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetMysqlaasInstanceLimitsCheckResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetMysqlaasInstanceLimitsCheckResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
