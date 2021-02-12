// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"net/http"
)

// CreateInternalPublicIpRequest wrapper for the CreateInternalPublicIp operation
type CreateInternalPublicIpRequest struct {

	// Create internal public IP details.
	CreateInternalPublicIpDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// This is the operation name used for authorization. This is only used when the API is called by a service as part of another API.
	InternalAuthzOperationName *string `mandatory:"false" contributesTo:"header" name:"internal-authz-operation-name"`

	// This is the resource kind used for authorization. This is only used when the API is called by a service as part of another API.
	InternalAuthzResourceKind *string `mandatory:"false" contributesTo:"header" name:"internal-authz-resource-kind"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateInternalPublicIpRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateInternalPublicIpRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateInternalPublicIpRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateInternalPublicIpResponse wrapper for the CreateInternalPublicIp operation
type CreateInternalPublicIpResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InternalPublicIp instance
	InternalPublicIp `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateInternalPublicIpResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateInternalPublicIpResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
