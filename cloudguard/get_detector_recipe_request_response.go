// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// GetDetectorRecipeRequest wrapper for the GetDetectorRecipe operation
type GetDetectorRecipeRequest struct {

	// DetectorRecipe OCID
	DetectorRecipeId *string `mandatory:"true" contributesTo:"path" name:"detectorRecipeId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDetectorRecipeRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDetectorRecipeRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDetectorRecipeRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDetectorRecipeResponse wrapper for the GetDetectorRecipe operation
type GetDetectorRecipeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DetectorRecipe instance
	DetectorRecipe `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDetectorRecipeResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDetectorRecipeResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
