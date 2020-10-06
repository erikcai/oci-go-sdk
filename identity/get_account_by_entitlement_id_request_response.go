// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"net/http"
)

// GetAccountByEntitlementIdRequest wrapper for the GetAccountByEntitlementId operation
type GetAccountByEntitlementIdRequest struct {

	// The name of the whitelisted service caller. Such as 'PSM'
	ServiceName *string `mandatory:"true" contributesTo:"query" name:"serviceName"`

	// The service entitlement Id for OCI resources.
	EntitlementId *string `mandatory:"true" contributesTo:"path" name:"entitlementId"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAccountByEntitlementIdRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAccountByEntitlementIdRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAccountByEntitlementIdRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAccountByEntitlementIdResponse wrapper for the GetAccountByEntitlementId operation
type GetAccountByEntitlementIdResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Account instance
	Account `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAccountByEntitlementIdResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAccountByEntitlementIdResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
