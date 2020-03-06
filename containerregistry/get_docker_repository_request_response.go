// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package containerregistry

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDockerRepositoryRequest wrapper for the GetDockerRepository operation
type GetDockerRepositoryRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository
	// Example: `ocid1.dockerrepo.oc1..exampleuniqueID`
	DockerRepositoryId *string `mandatory:"true" contributesTo:"path" name:"dockerRepositoryId"`

	// Unique Oracle-assigned identifier for the request
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDockerRepositoryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDockerRepositoryRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDockerRepositoryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDockerRepositoryResponse wrapper for the GetDockerRepository operation
type GetDockerRepositoryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DockerRepository instance
	DockerRepository `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDockerRepositoryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDockerRepositoryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
