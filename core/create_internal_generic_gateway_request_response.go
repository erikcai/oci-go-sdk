// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"net/http"
)

// CreateInternalGenericGatewayRequest wrapper for the CreateInternalGenericGateway operation
type CreateInternalGenericGatewayRequest struct {

	// Requesting to create an internal generic gateway.
	CreateInternalGenericGatewayDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This is the operation name used for authorization. This is only used when the API is called by a service as part of another API.
	InternalAuthzOperationName *string `mandatory:"false" contributesTo:"header" name:"internal-authz-operation-name"`

	// This is the resource kind used for authorization. This is only used when the API is called by a service as part of another API.
	InternalAuthzResourceKind *string `mandatory:"false" contributesTo:"header" name:"internal-authz-resource-kind"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateInternalGenericGatewayRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateInternalGenericGatewayRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateInternalGenericGatewayRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateInternalGenericGatewayResponse wrapper for the CreateInternalGenericGateway operation
type CreateInternalGenericGatewayResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InternalGenericGateway instance
	InternalGenericGateway `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateInternalGenericGatewayResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateInternalGenericGatewayResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
