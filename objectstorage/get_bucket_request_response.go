// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetBucketRequest wrapper for the GetBucket operation
type GetBucketRequest struct {

	// The top-level namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The entity tag to match. For creating and committing a multipart upload to an object, this is the entity tag of the target object.
	// For uploading a part, this is the entity tag of the target part.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The entity tag to avoid matching. The only valid value is '*', which indicates that the request should fail if the object already exists.
	// For creating and committing a multipart upload, this is the entity tag of the target object. For uploading a part, this is the entity tag of the target part.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Bucket summary includes the 'namespace', 'name', 'compartmentId', 'createdBy', 'timeCreated',
	// and 'etag' fields. This parameter can also include 'estimatedCount' (Estimated number of objects) and 'estimatedSize'
	// (total Estimated size in bytes of all objects). For example 'estimatedCount,estimatedSize'
	Fields []GetBucketFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"csv"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetBucketRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetBucketRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetBucketRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetBucketResponse wrapper for the GetBucket operation
type GetBucketResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Bucket instance
	Bucket `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The current entity tag for the bucket.
	ETag *string `presentIn:"header" name:"etag"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response GetBucketResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetBucketResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetBucketFieldsEnum Enum with underlying type: string
type GetBucketFieldsEnum string

// Set of constants representing the allowable values for GetBucketFields
const (
	GetBucketFieldsEstimatedcount GetBucketFieldsEnum = "estimatedCount"
	GetBucketFieldsEstimatedsize  GetBucketFieldsEnum = "estimatedSize"
)

var mappingGetBucketFields = map[string]GetBucketFieldsEnum{
	"estimatedCount": GetBucketFieldsEstimatedcount,
	"estimatedSize":  GetBucketFieldsEstimatedsize,
}

// GetGetBucketFieldsEnumValues Enumerates the set of values for GetBucketFields
func GetGetBucketFieldsEnumValues() []GetBucketFieldsEnum {
	values := make([]GetBucketFieldsEnum, 0)
	for _, v := range mappingGetBucketFields {
		values = append(values, v)
	}
	return values
}
