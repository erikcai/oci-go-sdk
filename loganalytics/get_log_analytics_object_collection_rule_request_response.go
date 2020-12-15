// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// GetLogAnalyticsObjectCollectionRuleRequest wrapper for the GetLogAnalyticsObjectCollectionRule operation
type GetLogAnalyticsObjectCollectionRuleRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The Logging Analytics Object Collection Rule OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	LogAnalyticsObjectCollectionRuleId *string `mandatory:"true" contributesTo:"path" name:"logAnalyticsObjectCollectionRuleId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetLogAnalyticsObjectCollectionRuleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetLogAnalyticsObjectCollectionRuleRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetLogAnalyticsObjectCollectionRuleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetLogAnalyticsObjectCollectionRuleResponse wrapper for the GetLogAnalyticsObjectCollectionRule operation
type GetLogAnalyticsObjectCollectionRuleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LogAnalyticsObjectCollectionRule instance
	LogAnalyticsObjectCollectionRule `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetLogAnalyticsObjectCollectionRuleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetLogAnalyticsObjectCollectionRuleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
