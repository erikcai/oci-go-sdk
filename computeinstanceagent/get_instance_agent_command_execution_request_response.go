// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computeinstanceagent

import (
	"github.com/oracle/oci-go-sdk/v29/common"
	"net/http"
)

// GetInstanceAgentCommandExecutionRequest wrapper for the GetInstanceAgentCommandExecution operation
type GetInstanceAgentCommandExecutionRequest struct {

	// The OCID of the command.
	InstanceAgentCommandId *string `mandatory:"true" contributesTo:"path" name:"instanceAgentCommandId"`

	// The OCID of the instance.
	InstanceId *string `mandatory:"true" contributesTo:"query" name:"instanceId"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetInstanceAgentCommandExecutionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetInstanceAgentCommandExecutionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetInstanceAgentCommandExecutionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetInstanceAgentCommandExecutionResponse wrapper for the GetInstanceAgentCommandExecution operation
type GetInstanceAgentCommandExecutionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InstanceAgentCommandExecution instance
	InstanceAgentCommandExecution `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetInstanceAgentCommandExecutionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetInstanceAgentCommandExecutionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
