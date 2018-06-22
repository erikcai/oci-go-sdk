// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListFlowLogConfigAttachmentsRequest wrapper for the ListFlowLogConfigAttachments operation
type ListFlowLogConfigAttachmentsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of a resource that has flow logs enabled.
	TargetEntityId *string `mandatory:"false" contributesTo:"query" name:"targetEntityId"`

	// The type of resource that has flow logs enabled.
	TargetEntityType ListFlowLogConfigAttachmentsTargetEntityTypeEnum `mandatory:"false" contributesTo:"query" name:"targetEntityType" omitEmpty:"true"`

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

func (request ListFlowLogConfigAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFlowLogConfigAttachmentsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFlowLogConfigAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFlowLogConfigAttachmentsResponse wrapper for the ListFlowLogConfigAttachments operation
type ListFlowLogConfigAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FlowLogConfigAttachment instances
	Items []FlowLogConfigAttachment `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListFlowLogConfigAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFlowLogConfigAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFlowLogConfigAttachmentsTargetEntityTypeEnum Enum with underlying type: string
type ListFlowLogConfigAttachmentsTargetEntityTypeEnum string

// Set of constants representing the allowable values for ListFlowLogConfigAttachmentsTargetEntityType
const (
	ListFlowLogConfigAttachmentsTargetEntityTypeSubnet ListFlowLogConfigAttachmentsTargetEntityTypeEnum = "SUBNET"
)

var mappingListFlowLogConfigAttachmentsTargetEntityType = map[string]ListFlowLogConfigAttachmentsTargetEntityTypeEnum{
	"SUBNET": ListFlowLogConfigAttachmentsTargetEntityTypeSubnet,
}

// GetListFlowLogConfigAttachmentsTargetEntityTypeEnumValues Enumerates the set of values for ListFlowLogConfigAttachmentsTargetEntityType
func GetListFlowLogConfigAttachmentsTargetEntityTypeEnumValues() []ListFlowLogConfigAttachmentsTargetEntityTypeEnum {
	values := make([]ListFlowLogConfigAttachmentsTargetEntityTypeEnum, 0)
	for _, v := range mappingListFlowLogConfigAttachmentsTargetEntityType {
		values = append(values, v)
	}
	return values
}
