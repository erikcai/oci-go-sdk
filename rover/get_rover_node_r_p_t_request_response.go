// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"github.com/oracle/oci-go-sdk/v29/common"
	"net/http"
)

// GetRoverNodeRPTRequest wrapper for the GetRoverNodeRPT operation
type GetRoverNodeRPTRequest struct {

	// Unique RoverNode identifier
	RoverNodeId *string `mandatory:"true" contributesTo:"path" name:"roverNodeId"`

	// The Java Web Token which is a signature of the request that is signed with the resource's private key
	// This is meant solely in the context of getRPT
	Jwt *string `mandatory:"true" contributesTo:"header" name:"jwt"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetRoverNodeRPTRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetRoverNodeRPTRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetRoverNodeRPTRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetRoverNodeRPTResponse wrapper for the GetRoverNodeRPT operation
type GetRoverNodeRPTResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The GetRoverNodeRpt instance
	GetRoverNodeRpt `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetRoverNodeRPTResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetRoverNodeRPTResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
