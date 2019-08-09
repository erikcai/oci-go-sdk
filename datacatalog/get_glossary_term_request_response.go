// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetGlossaryTermRequest wrapper for the GetGlossaryTerm operation
type GetGlossaryTermRequest struct {

	// unique Catalog identifier
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique Glossary key.
	GlossaryKey *string `mandatory:"true" contributesTo:"path" name:"glossaryKey"`

	// Unique Glossary Term key.
	TermKey *string `mandatory:"true" contributesTo:"path" name:"termKey"`

	// Used to control which fields are returned in a Term response.
	Fields []GetGlossaryTermFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetGlossaryTermRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetGlossaryTermRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetGlossaryTermRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetGlossaryTermResponse wrapper for the GetGlossaryTerm operation
type GetGlossaryTermResponse struct {

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

func (response GetGlossaryTermResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetGlossaryTermResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetGlossaryTermFieldsEnum Enum with underlying type: string
type GetGlossaryTermFieldsEnum string

// Set of constants representing the allowable values for GetGlossaryTermFieldsEnum
const (
	GetGlossaryTermFieldsKey                       GetGlossaryTermFieldsEnum = "key"
	GetGlossaryTermFieldsDisplayname               GetGlossaryTermFieldsEnum = "displayName"
	GetGlossaryTermFieldsDescription               GetGlossaryTermFieldsEnum = "description"
	GetGlossaryTermFieldsGlossarykey               GetGlossaryTermFieldsEnum = "glossaryKey"
	GetGlossaryTermFieldsParenttermkey             GetGlossaryTermFieldsEnum = "parentTermKey"
	GetGlossaryTermFieldsIsallowedtohavechildterms GetGlossaryTermFieldsEnum = "isAllowedToHaveChildTerms"
	GetGlossaryTermFieldsPath                      GetGlossaryTermFieldsEnum = "path"
	GetGlossaryTermFieldsLifecyclestate            GetGlossaryTermFieldsEnum = "lifecycleState"
	GetGlossaryTermFieldsTimecreated               GetGlossaryTermFieldsEnum = "timeCreated"
	GetGlossaryTermFieldsTimeupdated               GetGlossaryTermFieldsEnum = "timeUpdated"
	GetGlossaryTermFieldsCreatedbyid               GetGlossaryTermFieldsEnum = "createdById"
	GetGlossaryTermFieldsUpdatedbyid               GetGlossaryTermFieldsEnum = "updatedById"
	GetGlossaryTermFieldsUri                       GetGlossaryTermFieldsEnum = "uri"
	GetGlossaryTermFieldsRelatedterms              GetGlossaryTermFieldsEnum = "relatedTerms"
	GetGlossaryTermFieldsAssociatedobjects         GetGlossaryTermFieldsEnum = "associatedObjects"
)

var mappingGetGlossaryTermFields = map[string]GetGlossaryTermFieldsEnum{
	"key":                       GetGlossaryTermFieldsKey,
	"displayName":               GetGlossaryTermFieldsDisplayname,
	"description":               GetGlossaryTermFieldsDescription,
	"glossaryKey":               GetGlossaryTermFieldsGlossarykey,
	"parentTermKey":             GetGlossaryTermFieldsParenttermkey,
	"isAllowedToHaveChildTerms": GetGlossaryTermFieldsIsallowedtohavechildterms,
	"path":              GetGlossaryTermFieldsPath,
	"lifecycleState":    GetGlossaryTermFieldsLifecyclestate,
	"timeCreated":       GetGlossaryTermFieldsTimecreated,
	"timeUpdated":       GetGlossaryTermFieldsTimeupdated,
	"createdById":       GetGlossaryTermFieldsCreatedbyid,
	"updatedById":       GetGlossaryTermFieldsUpdatedbyid,
	"uri":               GetGlossaryTermFieldsUri,
	"relatedTerms":      GetGlossaryTermFieldsRelatedterms,
	"associatedObjects": GetGlossaryTermFieldsAssociatedobjects,
}

// GetGetGlossaryTermFieldsEnumValues Enumerates the set of values for GetGlossaryTermFieldsEnum
func GetGetGlossaryTermFieldsEnumValues() []GetGlossaryTermFieldsEnum {
	values := make([]GetGlossaryTermFieldsEnum, 0)
	for _, v := range mappingGetGlossaryTermFields {
		values = append(values, v)
	}
	return values
}
