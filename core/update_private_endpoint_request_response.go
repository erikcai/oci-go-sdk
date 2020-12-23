// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// UpdatePrivateEndpointRequest wrapper for the UpdatePrivateEndpoint operation
type UpdatePrivateEndpointRequest struct {

	// The private endpoint's OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	PrivateEndpointId *string `mandatory:"true" contributesTo:"path" name:"privateEndpointId"`

	// Details for updating a private endpoint.
	UpdatePrivateEndpointDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource. The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicates that this request is a dry-run.
	// If set to true, nothing will be created, but only the validation will be performed.
	IsDryRun *bool `mandatory:"false" contributesTo:"query" name:"isDryRun"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdatePrivateEndpointRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdatePrivateEndpointRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdatePrivateEndpointRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdatePrivateEndpointResponse wrapper for the UpdatePrivateEndpoint operation
type UpdatePrivateEndpointResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PrivateEndpoint instance
	PrivateEndpoint `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The OCID of the work request. Use GetWorkRequest (https://docs.cloud.oracle.com/api/#/en/workrequests/20160918/WorkRequest/GetWorkRequest)
	// with this ID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response UpdatePrivateEndpointResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdatePrivateEndpointResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
