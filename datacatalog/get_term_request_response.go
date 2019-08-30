// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetTermRequest wrapper for the GetTerm operation
type GetTermRequest struct {

	// unique Catalog identifier
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique Glossary key.
	GlossaryKey *string `mandatory:"true" contributesTo:"path" name:"glossaryKey"`

	// Unique Glossary Term key.
	TermKey *string `mandatory:"true" contributesTo:"path" name:"termKey"`

	// Used to control which fields are returned in a Term response.
	Fields []GetTermFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetTermRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetTermRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetTermRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetTermResponse wrapper for the GetTerm operation
type GetTermResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Term instance
	Term `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTermResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetTermResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetTermFieldsEnum Enum with underlying type: string
type GetTermFieldsEnum string

// Set of constants representing the allowable values for GetTermFieldsEnum
const (
	GetTermFieldsKey                       GetTermFieldsEnum = "key"
	GetTermFieldsDisplayname               GetTermFieldsEnum = "displayName"
	GetTermFieldsDescription               GetTermFieldsEnum = "description"
	GetTermFieldsGlossarykey               GetTermFieldsEnum = "glossaryKey"
	GetTermFieldsParenttermkey             GetTermFieldsEnum = "parentTermKey"
	GetTermFieldsIsallowedtohavechildterms GetTermFieldsEnum = "isAllowedToHaveChildTerms"
	GetTermFieldsPath                      GetTermFieldsEnum = "path"
	GetTermFieldsLifecyclestate            GetTermFieldsEnum = "lifecycleState"
	GetTermFieldsTimecreated               GetTermFieldsEnum = "timeCreated"
	GetTermFieldsTimeupdated               GetTermFieldsEnum = "timeUpdated"
	GetTermFieldsCreatedbyid               GetTermFieldsEnum = "createdById"
	GetTermFieldsUpdatedbyid               GetTermFieldsEnum = "updatedById"
	GetTermFieldsUri                       GetTermFieldsEnum = "uri"
	GetTermFieldsRelatedterms              GetTermFieldsEnum = "relatedTerms"
	GetTermFieldsAssociatedobjectcount     GetTermFieldsEnum = "associatedObjectCount"
	GetTermFieldsAssociatedobjects         GetTermFieldsEnum = "associatedObjects"
)

var mappingGetTermFields = map[string]GetTermFieldsEnum{
	"key":                       GetTermFieldsKey,
	"displayName":               GetTermFieldsDisplayname,
	"description":               GetTermFieldsDescription,
	"glossaryKey":               GetTermFieldsGlossarykey,
	"parentTermKey":             GetTermFieldsParenttermkey,
	"isAllowedToHaveChildTerms": GetTermFieldsIsallowedtohavechildterms,
	"path":                  GetTermFieldsPath,
	"lifecycleState":        GetTermFieldsLifecyclestate,
	"timeCreated":           GetTermFieldsTimecreated,
	"timeUpdated":           GetTermFieldsTimeupdated,
	"createdById":           GetTermFieldsCreatedbyid,
	"updatedById":           GetTermFieldsUpdatedbyid,
	"uri":                   GetTermFieldsUri,
	"relatedTerms":          GetTermFieldsRelatedterms,
	"associatedObjectCount": GetTermFieldsAssociatedobjectcount,
	"associatedObjects":     GetTermFieldsAssociatedobjects,
}

// GetGetTermFieldsEnumValues Enumerates the set of values for GetTermFieldsEnum
func GetGetTermFieldsEnumValues() []GetTermFieldsEnum {
	values := make([]GetTermFieldsEnum, 0)
	for _, v := range mappingGetTermFields {
		values = append(values, v)
	}
	return values
}