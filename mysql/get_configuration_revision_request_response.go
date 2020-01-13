// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetConfigurationRevisionRequest wrapper for the GetConfigurationRevision operation
type GetConfigurationRevisionRequest struct {

	// The OCID of the MySQLaaS Configuration.
	ConfigurationId *string `mandatory:"true" contributesTo:"path" name:"configurationId"`

	// A client-opaque identifier for a specific revision of a
	// Configuration. The current implementation is a integer sequence number,
	// but that specific form should not be relied on.
	RevisionId *string `mandatory:"true" contributesTo:"path" name:"revisionId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetConfigurationRevisionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetConfigurationRevisionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetConfigurationRevisionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetConfigurationRevisionResponse wrapper for the GetConfigurationRevision operation
type GetConfigurationRevisionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ConfigurationRevision instance
	ConfigurationRevision `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a specific request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response GetConfigurationRevisionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetConfigurationRevisionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
