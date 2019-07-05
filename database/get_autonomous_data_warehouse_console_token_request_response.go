// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetAutonomousDataWarehouseConsoleTokenRequest wrapper for the GetAutonomousDataWarehouseConsoleToken operation
type GetAutonomousDataWarehouseConsoleTokenRequest struct {

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousDataWarehouseId *string `mandatory:"true" contributesTo:"path" name:"autonomousDataWarehouseId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAutonomousDataWarehouseConsoleTokenRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAutonomousDataWarehouseConsoleTokenRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAutonomousDataWarehouseConsoleTokenRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAutonomousDataWarehouseConsoleTokenResponse wrapper for the GetAutonomousDataWarehouseConsoleToken operation
type GetAutonomousDataWarehouseConsoleTokenResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AutonomousDataWarehouseConsoleTokenDetails instance
	AutonomousDataWarehouseConsoleTokenDetails `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAutonomousDataWarehouseConsoleTokenResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAutonomousDataWarehouseConsoleTokenResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
