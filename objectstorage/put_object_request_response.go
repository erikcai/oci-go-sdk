// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// PutObjectRequest wrapper for the PutObject operation
type PutObjectRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The name of the object. Avoid entering confidential information.
	// Example: `test/object1.log`
	ObjectName *string `mandatory:"true" contributesTo:"path" name:"objectName"`

	// The content length of the body.
	ContentLength *int64 `mandatory:"true" contributesTo:"header" name:"Content-Length"`

	// The object to upload to the object store.
	PutObjectBody io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// The entity tag (ETag) to match. For creating and committing a multipart upload to an object, this is the entity tag of the target object.
	// For uploading a part, this is the entity tag of the target part.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The entity tag (ETag) to avoid matching. The only valid value is '*', which indicates that the request should fail if the object
	// already exists. For creating and committing a multipart upload, this is the entity tag of the target object. For uploading a
	// part, this is the entity tag of the target part.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// 100-continue
	Expect *string `mandatory:"false" contributesTo:"header" name:"Expect"`

	// The optional base-64 header that defines the encoded MD5 hash of the body. If the optional Content-MD5 header is present, Object
	// Storage performs an integrity check on the body of the HTTP request by computing the MD5 hash for the body and comparing it to the
	// MD5 hash supplied in the header. If the two hashes do not match, the object is rejected and an HTTP-400 Unmatched Content MD5 error
	// is returned with the message:
	// "The computed MD5 of the request body (ACTUAL_MD5) does not match the Content-MD5 header (HEADER_MD5)"
	ContentMD5 *string `mandatory:"false" contributesTo:"header" name:"Content-MD5"`

	// The optional Content-Type header that defines the standard MIME type format of the object. Content type defaults to
	// 'application/octet-stream' if not specified in the PutObject call. Specifying values for this header has no effect
	// on Object Storage behavior. Programs that read the object determine what to do based on the value provided. For example,
	// you could use this header to identify and perform special operations on text only objects.
	ContentType *string `mandatory:"false" contributesTo:"header" name:"Content-Type"`

	// The optional Content-Language header that defines the content language of the object to upload. Specifying
	// values for this header has no effect on Object Storage behavior. Programs that read the object determine what
	// to do based on the value provided. For example, you could use this header to identify and differentiate objects
	// based on a particular language.
	ContentLanguage *string `mandatory:"false" contributesTo:"header" name:"Content-Language"`

	// The optional Content-Encoding header that defines the content encodings that were applied to the object to
	// upload. Specifying values for this header has no effect on Object Storage behavior. Programs that read the
	// object determine what to do based on the value provided. For example, you could use this header to determine
	// what decoding mechanisms need to be applied to obtain the media-type specified by the Content-Type header of
	// the object.
	ContentEncoding *string `mandatory:"false" contributesTo:"header" name:"Content-Encoding"`

	// The optional Content-Disposition header that defines presentational information for the object to be
	// returned in GetObject and HeadObject responses. Specifying values for this header has no effect on Object
	// Storage behavior. Programs that read the object determine what to do based on the value provided.
	// For example, you could use this header to let users download objects with custom filenames in a browser.
	ContentDisposition *string `mandatory:"false" contributesTo:"header" name:"Content-Disposition"`

	// The optional Cache-Control header that defines the caching behavior value to be returned in GetObject and
	// HeadObject responses. Specifying values for this header has no effect on Object Storage behavior. Programs
	// that read the object determine what to do based on the value provided.
	// For example, you could use this header to identify objects that require caching restrictions.
	CacheControl *string `mandatory:"false" contributesTo:"header" name:"Cache-Control"`

	// Optional user-defined metadata key and value.
	OpcMeta map[string]string `mandatory:"false" contributesTo:"header-collection" prefix:"opc-meta-"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PutObjectRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PutObjectRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PutObjectRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// PutObjectResponse wrapper for the PutObject operation
type PutObjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The base-64 encoded MD5 hash of the request body as computed by the server.
	OpcContentMd5 *string `presentIn:"header" name:"opc-content-md5"`

	// The entity tag (ETag) for the object.
	ETag *string `presentIn:"header" name:"etag"`

	// The time the object was modified, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29).
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`
}

func (response PutObjectResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PutObjectResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
