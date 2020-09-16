// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"net/http"
)

// CreateConfigurationSourceProviderRequest wrapper for the CreateConfigurationSourceProvider operation
type CreateConfigurationSourceProviderRequest struct {

	// The properties for creating a ConfigurationSourceProvider.
	CreateConfigurationSourceProviderDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of retrying the same action. Retry tokens expire after
	// 24 hours, but can be invalidated before then due to conflicting operations. For example,
	// if a resource has been deleted and purged from the system, then a retry of the original
	// creation request may be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateConfigurationSourceProviderRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateConfigurationSourceProviderRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateConfigurationSourceProviderRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateConfigurationSourceProviderResponse wrapper for the CreateConfigurationSourceProvider operation
type CreateConfigurationSourceProviderResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ConfigurationSourceProvider instance
	ConfigurationSourceProvider `presentIn:"body"`

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response CreateConfigurationSourceProviderResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateConfigurationSourceProviderResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
