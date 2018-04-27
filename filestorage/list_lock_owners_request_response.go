// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLockOwnersRequest wrapper for the ListLockOwners operation
type ListLockOwnersRequest struct {

	// The OCID of the file system.
	FileSystemId *string `mandatory:"true" contributesTo:"path" name:"fileSystemId"`

	// The OCID of a mount target. This feature is currently in preview and may change before public release. Do not use it for production workloads.
	MountTarget *string `mandatory:"false" contributesTo:"query" name:"mountTarget"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLockOwnersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLockOwnersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLockOwnersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLockOwnersResponse wrapper for the ListLockOwners operation
type ListLockOwnersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LockOwner instances
	Items []LockOwner `presentIn:"body"`

	// For pagination of a list of items. When paging through
	// a list, if this header appears in the response, then a
	// partial list might have been returned. Include this
	// value as the `page` parameter for the subsequent GET
	// request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLockOwnersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLockOwnersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
