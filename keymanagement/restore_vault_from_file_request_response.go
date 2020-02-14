// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// RestoreVaultFromFileRequest wrapper for the RestoreVaultFromFile operation
type RestoreVaultFromFileRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The content length of the body.
	ContentLength *int64 `mandatory:"true" contributesTo:"header" name:"content-length"`

	// The encrypted payload to upload to restore the vault.
	RestoreVaultFromFileDetails io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// For optimistic concurrency control. In the PUT or DELETE call for a
	// resource, set the `if-match` parameter to the value of the etag from a
	// previous GET or POST response for that resource. The resource will be
	// updated or deleted only if the etag you provide matches the resource's
	// current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The base-64 encoded MD5 hash of the body, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.15.
	// If the Content-MD5 header is present, KMS server will performs an integrity check on the body of the HTTP request by computing the MD5
	// hash for the body and comparing it to the MD5 hash supplied in the header. If the two hashes do not match, the object is rejected and
	// an HTTP-400 Unmatched Content MD5 error is returned with the message: "The computed MD5 of the request body (ACTUAL_MD5) does not match the Content-MD5 header (HEADER_MD5)"
	ContentMd5 *string `mandatory:"false" contributesTo:"header" name:"content-md5"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case
	// of a timeout or server error without risk of executing that same action
	// again. Retry tokens expire after 24 hours, but can be invalidated
	// before then due to conflicting operations (e.g., if a resource has been
	// deleted and purged from the system, then a retry of the original
	// creation request may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RestoreVaultFromFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RestoreVaultFromFileRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RestoreVaultFromFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// RestoreVaultFromFileResponse wrapper for the RestoreVaultFromFile operation
type RestoreVaultFromFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Vault instance
	Vault `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// The base-64 encoded MD5 hash of the request body as computed by the server.
	OpcContentMd5 *string `presentIn:"header" name:"opc-content-md5"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Work request id to track progress of the restore operation
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response RestoreVaultFromFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RestoreVaultFromFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
