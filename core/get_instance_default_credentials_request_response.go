// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetInstanceDefaultCredentialsRequest wrapper for the GetInstanceDefaultCredentials operation
type GetInstanceDefaultCredentialsRequest struct {

	// The OCID of the instance.
	InstanceId *string `mandatory:"true" contributesTo:"path" name:"instanceId"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetInstanceDefaultCredentialsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetInstanceDefaultCredentialsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetInstanceDefaultCredentialsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetInstanceDefaultCredentialsResponse wrapper for the GetInstanceDefaultCredentials operation
type GetInstanceDefaultCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InstanceCredentials instance
	InstanceCredentials `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetInstanceDefaultCredentialsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetInstanceDefaultCredentialsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
