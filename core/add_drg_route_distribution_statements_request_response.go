// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"net/http"
)

// AddDrgRouteDistributionStatementsRequest wrapper for the AddDrgRouteDistributionStatements operation
type AddDrgRouteDistributionStatementsRequest struct {

	// The OCID of the Route Distribution
	DrgRouteDistributionId *string `mandatory:"true" contributesTo:"path" name:"drgRouteDistributionId"`

	// Request with one or more route distribution statements to be inserted into the Route Distribution
	AddDrgRouteDistributionStatementsDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request AddDrgRouteDistributionStatementsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request AddDrgRouteDistributionStatementsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request AddDrgRouteDistributionStatementsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// AddDrgRouteDistributionStatementsResponse wrapper for the AddDrgRouteDistributionStatements operation
type AddDrgRouteDistributionStatementsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []DrgRouteDistributionStatement instance
	Items []DrgRouteDistributionStatement `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response AddDrgRouteDistributionStatementsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response AddDrgRouteDistributionStatementsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
