// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetAutonomousPodMissionCriticalAssociationRequest wrapper for the GetAutonomousPodMissionCriticalAssociation operation
type GetAutonomousPodMissionCriticalAssociationRequest struct {

	// The Autonomous Pod OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	AutonomousPodId *string `mandatory:"true" contributesTo:"path" name:"autonomousPodId"`

	// The Autonomous Pod Mission Critical Association OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	AutonomousPodMissionCriticalAssociationId *string `mandatory:"true" contributesTo:"path" name:"autonomousPodMissionCriticalAssociationId"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAutonomousPodMissionCriticalAssociationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAutonomousPodMissionCriticalAssociationRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAutonomousPodMissionCriticalAssociationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAutonomousPodMissionCriticalAssociationResponse wrapper for the GetAutonomousPodMissionCriticalAssociation operation
type GetAutonomousPodMissionCriticalAssociationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AutonomousPodMissionCriticalAssociation instance
	AutonomousPodMissionCriticalAssociation `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAutonomousPodMissionCriticalAssociationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAutonomousPodMissionCriticalAssociationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
