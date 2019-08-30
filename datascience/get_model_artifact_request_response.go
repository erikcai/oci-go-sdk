// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// GetModelArtifactRequest wrapper for the GetModelArtifact operation
type GetModelArtifactRequest struct {

	// The OCID of the model.
	ModelId *string `mandatory:"true" contributesTo:"path" name:"modelId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetModelArtifactRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetModelArtifactRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetModelArtifactRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetModelArtifactResponse wrapper for the GetModelArtifact operation
type GetModelArtifactResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The object size in bytes.
	ContentLength *int64 `presentIn:"header" name:"content-length"`

	// The content disposition of the body.
	ContentDisposition *string `presentIn:"header" name:"content-disposition"`

	// Content-MD5 header, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.15.
	// Unavailable for objects uploaded using multipart upload.
	ContentMd5 *string `presentIn:"header" name:"content-md5"`

	// The artifact modification time, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.29.
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`
}

func (response GetModelArtifactResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetModelArtifactResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
