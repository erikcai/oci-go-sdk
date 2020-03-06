// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package containerregistry

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDockerTagsRequest wrapper for the ListDockerTags operation
type ListDockerTagsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the resource
	// belongs to. Use tenancyId to search in the root compartment.
	// Example: `ocid1.compartment.oc1..exampleuniqueID`
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A repository name
	// Filters results to the provided repository
	Repository *string `mandatory:"false" contributesTo:"query" name:"repository"`

	// A glob pattern as referenced by https://docs.oracle.com/javase/tutorial/essential/io/fileOps.html#glob
	// Filters results to docker tags which match this pattern
	// Example `2020.02.02*`
	TagMatchesGlob *string `mandatory:"false" contributesTo:"query" name:"tagMatchesGlob"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Starting position for search results (e.g. "docker:latest")
	TagGreaterThanOrEqualTo *string `mandatory:"false" contributesTo:"query" name:"tagGreaterThanOrEqualTo"`

	// The optional field to sort the results by.
	SortBy ListDockerTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The optional order in which to sort the results.
	SortOrder ListDockerTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A RFC 3339 timestamp.
	// Filter results to docker tags that have been tagged after the provided timestamp
	TimeTaggedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeTaggedGreaterThanOrEqualTo"`

	// A RFC 3339 timestamp.
	// Filter results to docker tags that have been tagged before the provided timestamp
	TimeTaggedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeTaggedLessThan"`

	// Unique Oracle-assigned identifier for the request
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDockerTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDockerTagsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDockerTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDockerTagsResponse wrapper for the ListDockerTags operation
type ListDockerTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DockerTagCollection instances
	DockerTagCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of
	// results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDockerTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDockerTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDockerTagsSortByEnum Enum with underlying type: string
type ListDockerTagsSortByEnum string

// Set of constants representing the allowable values for ListDockerTagsSortByEnum
const (
	ListDockerTagsSortByTag        ListDockerTagsSortByEnum = "tag"
	ListDockerTagsSortByTimetagged ListDockerTagsSortByEnum = "timeTagged"
)

var mappingListDockerTagsSortBy = map[string]ListDockerTagsSortByEnum{
	"tag":        ListDockerTagsSortByTag,
	"timeTagged": ListDockerTagsSortByTimetagged,
}

// GetListDockerTagsSortByEnumValues Enumerates the set of values for ListDockerTagsSortByEnum
func GetListDockerTagsSortByEnumValues() []ListDockerTagsSortByEnum {
	values := make([]ListDockerTagsSortByEnum, 0)
	for _, v := range mappingListDockerTagsSortBy {
		values = append(values, v)
	}
	return values
}

// ListDockerTagsSortOrderEnum Enum with underlying type: string
type ListDockerTagsSortOrderEnum string

// Set of constants representing the allowable values for ListDockerTagsSortOrderEnum
const (
	ListDockerTagsSortOrderAsc  ListDockerTagsSortOrderEnum = "ASC"
	ListDockerTagsSortOrderDesc ListDockerTagsSortOrderEnum = "DESC"
)

var mappingListDockerTagsSortOrder = map[string]ListDockerTagsSortOrderEnum{
	"ASC":  ListDockerTagsSortOrderAsc,
	"DESC": ListDockerTagsSortOrderDesc,
}

// GetListDockerTagsSortOrderEnumValues Enumerates the set of values for ListDockerTagsSortOrderEnum
func GetListDockerTagsSortOrderEnumValues() []ListDockerTagsSortOrderEnum {
	values := make([]ListDockerTagsSortOrderEnum, 0)
	for _, v := range mappingListDockerTagsSortOrder {
		values = append(values, v)
	}
	return values
}
