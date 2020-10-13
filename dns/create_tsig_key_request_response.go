// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v27/common"
	"net/http"
)

// CreateTsigKeyRequest wrapper for the CreateTsigKey operation
type CreateTsigKeyRequest struct {

	// Details for creating a new TSIG key.
	CreateTsigKeyDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope CreateTsigKeyScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateTsigKeyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateTsigKeyRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateTsigKeyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateTsigKeyResponse wrapper for the CreateTsigKey operation
type CreateTsigKeyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TsigKey instance
	TsigKey `presentIn:"body"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`

	// The full URI of the resource related to the request.
	Location *string `presentIn:"header" name:"location"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response CreateTsigKeyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateTsigKeyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// CreateTsigKeyScopeEnum Enum with underlying type: string
type CreateTsigKeyScopeEnum string

// Set of constants representing the allowable values for CreateTsigKeyScopeEnum
const (
	CreateTsigKeyScopeGlobal  CreateTsigKeyScopeEnum = "GLOBAL"
	CreateTsigKeyScopePrivate CreateTsigKeyScopeEnum = "PRIVATE"
)

var mappingCreateTsigKeyScope = map[string]CreateTsigKeyScopeEnum{
	"GLOBAL":  CreateTsigKeyScopeGlobal,
	"PRIVATE": CreateTsigKeyScopePrivate,
}

// GetCreateTsigKeyScopeEnumValues Enumerates the set of values for CreateTsigKeyScopeEnum
func GetCreateTsigKeyScopeEnumValues() []CreateTsigKeyScopeEnum {
	values := make([]CreateTsigKeyScopeEnum, 0)
	for _, v := range mappingCreateTsigKeyScope {
		values = append(values, v)
	}
	return values
}
