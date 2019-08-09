// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListFoldersRequest wrapper for the ListFolders operation
type ListFoldersRequest struct {

	// unique Catalog identifier
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique Data Asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState LifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique folder key.
	ParentFolderKey *string `mandatory:"false" contributesTo:"query" name:"parentFolderKey"`

	// Full path of the resource for resources that support paths.
	Path *string `mandatory:"false" contributesTo:"query" name:"path"`

	// Unique external identifier of this resource in the external source system.
	ExternalKey *string `mandatory:"false" contributesTo:"query" name:"externalKey"`

	// Time that the Resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the Resource was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// Id (OCID) of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// Id of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExternal"`

	// Harvest Status of the harvestable resource as updated by the harvest process.
	HarvestStatus HarvestStatusEnum `mandatory:"false" contributesTo:"query" name:"harvestStatus" omitEmpty:"true"`

	// Key of the last harvest process to update this resource.
	LastJobKey *string `mandatory:"false" contributesTo:"query" name:"lastJobKey"`

	// Used to control which fields are returned in a Folder summary response.
	Fields []ListFoldersFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListFoldersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListFoldersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFoldersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFoldersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFoldersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFoldersResponse wrapper for the ListFolders operation
type ListFoldersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FolderSummary instances
	Items []FolderSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFoldersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFoldersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFoldersFieldsEnum Enum with underlying type: string
type ListFoldersFieldsEnum string

// Set of constants representing the allowable values for ListFoldersFieldsEnum
const (
	ListFoldersFieldsKey             ListFoldersFieldsEnum = "key"
	ListFoldersFieldsDisplayname     ListFoldersFieldsEnum = "displayName"
	ListFoldersFieldsDescription     ListFoldersFieldsEnum = "description"
	ListFoldersFieldsParentfolderkey ListFoldersFieldsEnum = "parentFolderKey"
	ListFoldersFieldsPath            ListFoldersFieldsEnum = "path"
	ListFoldersFieldsDataassetkey    ListFoldersFieldsEnum = "dataAssetKey"
	ListFoldersFieldsExternalkey     ListFoldersFieldsEnum = "externalKey"
	ListFoldersFieldsTimecreated     ListFoldersFieldsEnum = "timeCreated"
	ListFoldersFieldsLifecyclestate  ListFoldersFieldsEnum = "lifecycleState"
	ListFoldersFieldsUri             ListFoldersFieldsEnum = "uri"
)

var mappingListFoldersFields = map[string]ListFoldersFieldsEnum{
	"key":             ListFoldersFieldsKey,
	"displayName":     ListFoldersFieldsDisplayname,
	"description":     ListFoldersFieldsDescription,
	"parentFolderKey": ListFoldersFieldsParentfolderkey,
	"path":            ListFoldersFieldsPath,
	"dataAssetKey":    ListFoldersFieldsDataassetkey,
	"externalKey":     ListFoldersFieldsExternalkey,
	"timeCreated":     ListFoldersFieldsTimecreated,
	"lifecycleState":  ListFoldersFieldsLifecyclestate,
	"uri":             ListFoldersFieldsUri,
}

// GetListFoldersFieldsEnumValues Enumerates the set of values for ListFoldersFieldsEnum
func GetListFoldersFieldsEnumValues() []ListFoldersFieldsEnum {
	values := make([]ListFoldersFieldsEnum, 0)
	for _, v := range mappingListFoldersFields {
		values = append(values, v)
	}
	return values
}

// ListFoldersSortByEnum Enum with underlying type: string
type ListFoldersSortByEnum string

// Set of constants representing the allowable values for ListFoldersSortByEnum
const (
	ListFoldersSortByTimecreated ListFoldersSortByEnum = "TIMECREATED"
	ListFoldersSortByDisplayname ListFoldersSortByEnum = "DISPLAYNAME"
)

var mappingListFoldersSortBy = map[string]ListFoldersSortByEnum{
	"TIMECREATED": ListFoldersSortByTimecreated,
	"DISPLAYNAME": ListFoldersSortByDisplayname,
}

// GetListFoldersSortByEnumValues Enumerates the set of values for ListFoldersSortByEnum
func GetListFoldersSortByEnumValues() []ListFoldersSortByEnum {
	values := make([]ListFoldersSortByEnum, 0)
	for _, v := range mappingListFoldersSortBy {
		values = append(values, v)
	}
	return values
}

// ListFoldersSortOrderEnum Enum with underlying type: string
type ListFoldersSortOrderEnum string

// Set of constants representing the allowable values for ListFoldersSortOrderEnum
const (
	ListFoldersSortOrderAsc  ListFoldersSortOrderEnum = "ASC"
	ListFoldersSortOrderDesc ListFoldersSortOrderEnum = "DESC"
)

var mappingListFoldersSortOrder = map[string]ListFoldersSortOrderEnum{
	"ASC":  ListFoldersSortOrderAsc,
	"DESC": ListFoldersSortOrderDesc,
}

// GetListFoldersSortOrderEnumValues Enumerates the set of values for ListFoldersSortOrderEnum
func GetListFoldersSortOrderEnumValues() []ListFoldersSortOrderEnum {
	values := make([]ListFoldersSortOrderEnum, 0)
	for _, v := range mappingListFoldersSortOrder {
		values = append(values, v)
	}
	return values
}
