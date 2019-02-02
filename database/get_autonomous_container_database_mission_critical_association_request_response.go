// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetAutonomousContainerDatabaseMissionCriticalAssociationRequest wrapper for the GetAutonomousContainerDatabaseMissionCriticalAssociation operation
type GetAutonomousContainerDatabaseMissionCriticalAssociationRequest struct {

	// The Autonomous Container Database OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"true" contributesTo:"path" name:"autonomousContainerDatabaseId"`

	// The Autonomous Container Database Mission Critical Association OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseMissionCriticalAssociationId *string `mandatory:"true" contributesTo:"path" name:"autonomousContainerDatabaseMissionCriticalAssociationId"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAutonomousContainerDatabaseMissionCriticalAssociationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAutonomousContainerDatabaseMissionCriticalAssociationRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAutonomousContainerDatabaseMissionCriticalAssociationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAutonomousContainerDatabaseMissionCriticalAssociationResponse wrapper for the GetAutonomousContainerDatabaseMissionCriticalAssociation operation
type GetAutonomousContainerDatabaseMissionCriticalAssociationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AutonomousContainerDatabaseMissionCriticalAssociation instance
	AutonomousContainerDatabaseMissionCriticalAssociation `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAutonomousContainerDatabaseMissionCriticalAssociationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAutonomousContainerDatabaseMissionCriticalAssociationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
