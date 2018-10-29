// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetInstallConfigurationRequest wrapper for the GetInstallConfiguration operation
type GetInstallConfigurationRequest struct {

	// Unique ID of the Application.
	ApplicationId *string `mandatory:"true" contributesTo:"path" name:"applicationId"`

	// It specifies the Package Version Id in case of Applications with multiple Package versions.
	PackageVersionId *string `mandatory:"false" contributesTo:"query" name:"packageVersionId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetInstallConfigurationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetInstallConfigurationRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetInstallConfigurationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetInstallConfigurationResponse wrapper for the GetInstallConfiguration operation
type GetInstallConfigurationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InstallConfiguration instance
	InstallConfiguration `presentIn:"body"`

	// Unique identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetInstallConfigurationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetInstallConfigurationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
