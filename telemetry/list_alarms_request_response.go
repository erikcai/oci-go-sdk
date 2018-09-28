// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package telemetry

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAlarmsRequest wrapper for the ListAlarms operation
type ListAlarmsRequest struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the
	// resources monitored by the metric that you are searching for. Use tenancyId to search in
	// the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer part of the request identifier token. If you need to contact Oracle about a particular
	// request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works, see List Pagination (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Default: 1000
	// Example: 500
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to return only resources that match the given display name exactly.
	// Use this filter to list an alarm by name. Alternatively, when you know the alarm OCID, use the GetAlarm operation.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAlarmsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAlarmsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAlarmsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAlarmsResponse wrapper for the ListAlarms operation
type ListAlarmsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AlarmSummary instances
	Items []AlarmSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAlarmsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAlarmsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
