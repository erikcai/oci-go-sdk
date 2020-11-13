// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocas

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"net/http"
)

// DetectLanguageEntitiesRequest wrapper for the DetectLanguageEntities operation
type DetectLanguageEntitiesRequest struct {

	// The details to make a Entity detect call
	// Example: `{"text": "If an emerging growth company, indicate by check mark if the registrant has elected not
	//            to use the extended transition period for complying"}`
	DetectLanguageEntitiesDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If this parameter is set to true, user will get only PII(Personally identifiable information) entities
	// like PhoneNumber, Email, Person etc. Default value is false.
	IsPii *bool `mandatory:"false" contributesTo:"query" name:"isPii"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DetectLanguageEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DetectLanguageEntitiesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DetectLanguageEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DetectLanguageEntitiesResponse wrapper for the DetectLanguageEntities operation
type DetectLanguageEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DetectLanguageEntitiesResult instance
	DetectLanguageEntitiesResult `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DetectLanguageEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DetectLanguageEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}