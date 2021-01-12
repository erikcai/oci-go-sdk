// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v32/common"
	"io"
	"net/http"
)

// GetClientVpnEndpointUserProfileRequest wrapper for the GetClientVpnEndpointUserProfile operation
type GetClientVpnEndpointUserProfileRequest struct {

	// The OCID of the ClientVpnEndpoint.
	ClientVpnEndpointId *string `mandatory:"true" contributesTo:"path" name:"clientVpnEndpointId"`

	// The username of the ClientVpnEndpointUser.
	UserName *string `mandatory:"true" contributesTo:"path" name:"userName"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetClientVpnEndpointUserProfileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetClientVpnEndpointUserProfileRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetClientVpnEndpointUserProfileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetClientVpnEndpointUserProfileResponse wrapper for the GetClientVpnEndpointUserProfile operation
type GetClientVpnEndpointUserProfileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetClientVpnEndpointUserProfileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetClientVpnEndpointUserProfileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
