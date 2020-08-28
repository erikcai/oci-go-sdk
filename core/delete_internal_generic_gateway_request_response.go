// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeleteInternalGenericGatewayRequest wrapper for the DeleteInternalGenericGateway operation
type DeleteInternalGenericGatewayRequest struct {

	// The OCID of the generic gateway
	InternalGenericGatewayId *string `mandatory:"true" contributesTo:"path" name:"internalGenericGatewayId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

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

func (request DeleteInternalGenericGatewayRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteInternalGenericGatewayRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteInternalGenericGatewayRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteInternalGenericGatewayResponse wrapper for the DeleteInternalGenericGateway operation
type DeleteInternalGenericGatewayResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteInternalGenericGatewayResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteInternalGenericGatewayResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
