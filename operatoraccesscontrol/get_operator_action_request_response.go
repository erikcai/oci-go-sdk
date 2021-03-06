// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operatoraccesscontrol

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"net/http"
)

// GetOperatorActionRequest wrapper for the GetOperatorAction operation
type GetOperatorActionRequest struct {

	// Unique Oracle supplied identifier associated with the operator action.
	OperatorActionId *string `mandatory:"true" contributesTo:"path" name:"operatorActionId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetOperatorActionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetOperatorActionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetOperatorActionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetOperatorActionResponse wrapper for the GetOperatorAction operation
type GetOperatorActionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OperatorAction instance
	OperatorAction `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetOperatorActionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetOperatorActionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
