// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetApiDeploymentSpecificationRequest wrapper for the GetApiDeploymentSpecification operation
type GetApiDeploymentSpecificationRequest struct {

	// The ocid of the API.
	ApiId *string `mandatory:"true" contributesTo:"path" name:"apiId"`

	// The client request id for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetApiDeploymentSpecificationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetApiDeploymentSpecificationRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetApiDeploymentSpecificationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetApiDeploymentSpecificationResponse wrapper for the GetApiDeploymentSpecification operation
type GetApiDeploymentSpecificationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ApiSpecification instance
	ApiSpecification `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// id.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetApiDeploymentSpecificationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetApiDeploymentSpecificationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
