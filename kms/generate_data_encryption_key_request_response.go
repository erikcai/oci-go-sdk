// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package kms

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GenerateDataEncryptionKeyRequest wrapper for the GenerateDataEncryptionKey operation
type GenerateDataEncryptionKeyRequest struct {

	// GenerateKeyDetails
	GenerateKeyDetails `contributesTo:"body"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GenerateDataEncryptionKeyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GenerateDataEncryptionKeyRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GenerateDataEncryptionKeyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GenerateDataEncryptionKeyResponse wrapper for the GenerateDataEncryptionKey operation
type GenerateDataEncryptionKeyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The GeneratedKey instance
	GeneratedKey `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GenerateDataEncryptionKeyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GenerateDataEncryptionKeyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
