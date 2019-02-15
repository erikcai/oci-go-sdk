// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package cache

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListCachingConsumptionRequest wrapper for the ListCachingConsumption operation
type ListCachingConsumptionRequest struct {

	// The scope of the limit consumption.
	Scope ListCachingConsumptionScopeEnum `mandatory:"true" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCachingConsumptionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCachingConsumptionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCachingConsumptionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCachingConsumptionResponse wrapper for the ListCachingConsumption operation
type ListCachingConsumptionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []LimitConsumption instance
	Items []LimitConsumption `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCachingConsumptionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCachingConsumptionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCachingConsumptionScopeEnum Enum with underlying type: string
type ListCachingConsumptionScopeEnum string

// Set of constants representing the allowable values for ListCachingConsumptionScopeEnum
const (
	ListCachingConsumptionScopeRegion ListCachingConsumptionScopeEnum = "REGION"
)

var mappingListCachingConsumptionScope = map[string]ListCachingConsumptionScopeEnum{
	"REGION": ListCachingConsumptionScopeRegion,
}

// GetListCachingConsumptionScopeEnumValues Enumerates the set of values for ListCachingConsumptionScopeEnum
func GetListCachingConsumptionScopeEnumValues() []ListCachingConsumptionScopeEnum {
	values := make([]ListCachingConsumptionScopeEnum, 0)
	for _, v := range mappingListCachingConsumptionScope {
		values = append(values, v)
	}
	return values
}
