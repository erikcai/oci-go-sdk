// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package containerregistry

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDockerRepositoriesRequest wrapper for the ListDockerRepositories operation
type ListDockerRepositoriesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the resource
	// belongs to. Use tenancyId to search in the root compartment.
	// Example: `ocid1.compartment.oc1..exampleuniqueID`
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter for repository id.
	// Will match the repository whose ocid matches the provided value
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A repository name
	// Filters results to the provided repository
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The optional field to sort the results by.
	SortBy ListDockerRepositoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The optional order in which to sort the results.
	SortOrder ListDockerRepositoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDockerRepositoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDockerRepositoriesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDockerRepositoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDockerRepositoriesResponse wrapper for the ListDockerRepositories operation
type ListDockerRepositoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DockerRepositoryCollection instances
	DockerRepositoryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of
	// results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDockerRepositoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDockerRepositoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDockerRepositoriesSortByEnum Enum with underlying type: string
type ListDockerRepositoriesSortByEnum string

// Set of constants representing the allowable values for ListDockerRepositoriesSortByEnum
const (
	ListDockerRepositoriesSortByName        ListDockerRepositoriesSortByEnum = "name"
	ListDockerRepositoriesSortByTimecreated ListDockerRepositoriesSortByEnum = "timeCreated"
)

var mappingListDockerRepositoriesSortBy = map[string]ListDockerRepositoriesSortByEnum{
	"name":        ListDockerRepositoriesSortByName,
	"timeCreated": ListDockerRepositoriesSortByTimecreated,
}

// GetListDockerRepositoriesSortByEnumValues Enumerates the set of values for ListDockerRepositoriesSortByEnum
func GetListDockerRepositoriesSortByEnumValues() []ListDockerRepositoriesSortByEnum {
	values := make([]ListDockerRepositoriesSortByEnum, 0)
	for _, v := range mappingListDockerRepositoriesSortBy {
		values = append(values, v)
	}
	return values
}

// ListDockerRepositoriesSortOrderEnum Enum with underlying type: string
type ListDockerRepositoriesSortOrderEnum string

// Set of constants representing the allowable values for ListDockerRepositoriesSortOrderEnum
const (
	ListDockerRepositoriesSortOrderAsc  ListDockerRepositoriesSortOrderEnum = "ASC"
	ListDockerRepositoriesSortOrderDesc ListDockerRepositoriesSortOrderEnum = "DESC"
)

var mappingListDockerRepositoriesSortOrder = map[string]ListDockerRepositoriesSortOrderEnum{
	"ASC":  ListDockerRepositoriesSortOrderAsc,
	"DESC": ListDockerRepositoriesSortOrderDesc,
}

// GetListDockerRepositoriesSortOrderEnumValues Enumerates the set of values for ListDockerRepositoriesSortOrderEnum
func GetListDockerRepositoriesSortOrderEnumValues() []ListDockerRepositoriesSortOrderEnum {
	values := make([]ListDockerRepositoriesSortOrderEnum, 0)
	for _, v := range mappingListDockerRepositoriesSortOrder {
		values = append(values, v)
	}
	return values
}
