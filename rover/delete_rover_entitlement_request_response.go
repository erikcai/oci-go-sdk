// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"github.com/oracle/oci-go-sdk/v27/common"
	"net/http"
)

// DeleteRoverEntitlementRequest wrapper for the DeleteRoverEntitlement operation
type DeleteRoverEntitlementRequest struct {

	// Id of the Rover Device Entitlement
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy DeleteRoverEntitlementSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteRoverEntitlementRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteRoverEntitlementRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteRoverEntitlementRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteRoverEntitlementResponse wrapper for the DeleteRoverEntitlement operation
type DeleteRoverEntitlementResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteRoverEntitlementResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteRoverEntitlementResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// DeleteRoverEntitlementSortByEnum Enum with underlying type: string
type DeleteRoverEntitlementSortByEnum string

// Set of constants representing the allowable values for DeleteRoverEntitlementSortByEnum
const (
	DeleteRoverEntitlementSortByTimecreated DeleteRoverEntitlementSortByEnum = "timeCreated"
	DeleteRoverEntitlementSortByDisplayname DeleteRoverEntitlementSortByEnum = "displayName"
)

var mappingDeleteRoverEntitlementSortBy = map[string]DeleteRoverEntitlementSortByEnum{
	"timeCreated": DeleteRoverEntitlementSortByTimecreated,
	"displayName": DeleteRoverEntitlementSortByDisplayname,
}

// GetDeleteRoverEntitlementSortByEnumValues Enumerates the set of values for DeleteRoverEntitlementSortByEnum
func GetDeleteRoverEntitlementSortByEnumValues() []DeleteRoverEntitlementSortByEnum {
	values := make([]DeleteRoverEntitlementSortByEnum, 0)
	for _, v := range mappingDeleteRoverEntitlementSortBy {
		values = append(values, v)
	}
	return values
}
