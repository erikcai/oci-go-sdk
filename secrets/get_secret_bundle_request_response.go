// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package secrets

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetSecretBundleRequest wrapper for the GetSecretBundle operation
type GetSecretBundleRequest struct {

	// Secret OCID
	SecretId *string `mandatory:"true" contributesTo:"path" name:"secretId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The Version Number of the secret.
	VersionNumber *int64 `mandatory:"false" contributesTo:"query" name:"versionNumber"`

	// The name of the secret version.
	SecretVersionName *string `mandatory:"false" contributesTo:"query" name:"secretVersionName"`

	// Secret Stage like CURRENT, PENDING etc.
	Stage GetSecretBundleStageEnum `mandatory:"false" contributesTo:"query" name:"stage" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetSecretBundleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetSecretBundleRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetSecretBundleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetSecretBundleResponse wrapper for the GetSecretBundle operation
type GetSecretBundleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SecretBundle instance
	SecretBundle `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetSecretBundleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetSecretBundleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetSecretBundleStageEnum Enum with underlying type: string
type GetSecretBundleStageEnum string

// Set of constants representing the allowable values for GetSecretBundleStageEnum
const (
	GetSecretBundleStageCurrent    GetSecretBundleStageEnum = "CURRENT"
	GetSecretBundleStagePending    GetSecretBundleStageEnum = "PENDING"
	GetSecretBundleStageLatest     GetSecretBundleStageEnum = "LATEST"
	GetSecretBundleStagePrevious   GetSecretBundleStageEnum = "PREVIOUS"
	GetSecretBundleStageDeprecated GetSecretBundleStageEnum = "DEPRECATED"
)

var mappingGetSecretBundleStage = map[string]GetSecretBundleStageEnum{
	"CURRENT":    GetSecretBundleStageCurrent,
	"PENDING":    GetSecretBundleStagePending,
	"LATEST":     GetSecretBundleStageLatest,
	"PREVIOUS":   GetSecretBundleStagePrevious,
	"DEPRECATED": GetSecretBundleStageDeprecated,
}

// GetGetSecretBundleStageEnumValues Enumerates the set of values for GetSecretBundleStageEnum
func GetGetSecretBundleStageEnumValues() []GetSecretBundleStageEnum {
	values := make([]GetSecretBundleStageEnum, 0)
	for _, v := range mappingGetSecretBundleStage {
		values = append(values, v)
	}
	return values
}