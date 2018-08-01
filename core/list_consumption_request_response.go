// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListConsumptionRequest wrapper for the ListConsumption operation
type ListConsumptionRequest struct {

	// The scope of the limit consumption.
	Scope ListConsumptionScopeEnum `mandatory:"true" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the Availability Domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConsumptionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConsumptionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConsumptionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListConsumptionResponse wrapper for the ListConsumption operation
type ListConsumptionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []LimitConsumption instance
	Items []LimitConsumption `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListConsumptionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConsumptionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConsumptionScopeEnum Enum with underlying type: string
type ListConsumptionScopeEnum string

// Set of constants representing the allowable values for ListConsumptionScope
const (
	ListConsumptionScopeRegion ListConsumptionScopeEnum = "REGION"
	ListConsumptionScopeAd     ListConsumptionScopeEnum = "AD"
)

var mappingListConsumptionScope = map[string]ListConsumptionScopeEnum{
	"REGION": ListConsumptionScopeRegion,
	"AD":     ListConsumptionScopeAd,
}

// GetListConsumptionScopeEnumValues Enumerates the set of values for ListConsumptionScope
func GetListConsumptionScopeEnumValues() []ListConsumptionScopeEnum {
	values := make([]ListConsumptionScopeEnum, 0)
	for _, v := range mappingListConsumptionScope {
		values = append(values, v)
	}
	return values
}
