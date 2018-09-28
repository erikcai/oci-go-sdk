// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListStreamingConsumptionRequest wrapper for the ListStreamingConsumption operation
type ListStreamingConsumptionRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The scope of the limit consumption. This should match ServiceLimits.name in the Limits API.
	Scope ListStreamingConsumptionScopeEnum `mandatory:"true" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The availability domain used to fiter.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStreamingConsumptionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStreamingConsumptionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStreamingConsumptionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListStreamingConsumptionResponse wrapper for the ListStreamingConsumption operation
type ListStreamingConsumptionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []LimitConsumption instance
	Items []LimitConsumption `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListStreamingConsumptionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStreamingConsumptionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStreamingConsumptionScopeEnum Enum with underlying type: string
type ListStreamingConsumptionScopeEnum string

// Set of constants representing the allowable values for ListStreamingConsumptionScope
const (
	ListStreamingConsumptionScopeRegion             ListStreamingConsumptionScopeEnum = "REGION"
	ListStreamingConsumptionScopeAvailabilityDomain ListStreamingConsumptionScopeEnum = "AVAILABILITY_DOMAIN"
)

var mappingListStreamingConsumptionScope = map[string]ListStreamingConsumptionScopeEnum{
	"REGION":              ListStreamingConsumptionScopeRegion,
	"AVAILABILITY_DOMAIN": ListStreamingConsumptionScopeAvailabilityDomain,
}

// GetListStreamingConsumptionScopeEnumValues Enumerates the set of values for ListStreamingConsumptionScope
func GetListStreamingConsumptionScopeEnumValues() []ListStreamingConsumptionScopeEnum {
	values := make([]ListStreamingConsumptionScopeEnum, 0)
	for _, v := range mappingListStreamingConsumptionScope {
		values = append(values, v)
	}
	return values
}
