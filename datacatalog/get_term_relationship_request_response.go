// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetTermRelationshipRequest wrapper for the GetTermRelationship operation
type GetTermRelationshipRequest struct {

	// unique Catalog identifier
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique Glossary key.
	GlossaryKey *string `mandatory:"true" contributesTo:"path" name:"glossaryKey"`

	// Unique Glossary Term key.
	TermKey *string `mandatory:"true" contributesTo:"path" name:"termKey"`

	// Unique Glossary Term Relationship key.
	TermRelationshipKey *string `mandatory:"true" contributesTo:"path" name:"termRelationshipKey"`

	// Used to control which fields are returned in a Term Relationship response.
	Fields []GetTermRelationshipFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetTermRelationshipRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetTermRelationshipRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetTermRelationshipRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetTermRelationshipResponse wrapper for the GetTermRelationship operation
type GetTermRelationshipResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TermRelationship instance
	TermRelationship `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTermRelationshipResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetTermRelationshipResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetTermRelationshipFieldsEnum Enum with underlying type: string
type GetTermRelationshipFieldsEnum string

// Set of constants representing the allowable values for GetTermRelationshipFieldsEnum
const (
	GetTermRelationshipFieldsKey                    GetTermRelationshipFieldsEnum = "key"
	GetTermRelationshipFieldsDisplayname            GetTermRelationshipFieldsEnum = "displayName"
	GetTermRelationshipFieldsDescription            GetTermRelationshipFieldsEnum = "description"
	GetTermRelationshipFieldsRelatedtermkey         GetTermRelationshipFieldsEnum = "relatedTermKey"
	GetTermRelationshipFieldsRelatedtermdisplayname GetTermRelationshipFieldsEnum = "relatedTermDisplayName"
	GetTermRelationshipFieldsParenttermkey          GetTermRelationshipFieldsEnum = "parentTermKey"
	GetTermRelationshipFieldsLifecyclestate         GetTermRelationshipFieldsEnum = "lifecycleState"
	GetTermRelationshipFieldsTimecreated            GetTermRelationshipFieldsEnum = "timeCreated"
	GetTermRelationshipFieldsUri                    GetTermRelationshipFieldsEnum = "uri"
)

var mappingGetTermRelationshipFields = map[string]GetTermRelationshipFieldsEnum{
	"key":                    GetTermRelationshipFieldsKey,
	"displayName":            GetTermRelationshipFieldsDisplayname,
	"description":            GetTermRelationshipFieldsDescription,
	"relatedTermKey":         GetTermRelationshipFieldsRelatedtermkey,
	"relatedTermDisplayName": GetTermRelationshipFieldsRelatedtermdisplayname,
	"parentTermKey":          GetTermRelationshipFieldsParenttermkey,
	"lifecycleState":         GetTermRelationshipFieldsLifecyclestate,
	"timeCreated":            GetTermRelationshipFieldsTimecreated,
	"uri":                    GetTermRelationshipFieldsUri,
}

// GetGetTermRelationshipFieldsEnumValues Enumerates the set of values for GetTermRelationshipFieldsEnum
func GetGetTermRelationshipFieldsEnumValues() []GetTermRelationshipFieldsEnum {
	values := make([]GetTermRelationshipFieldsEnum, 0)
	for _, v := range mappingGetTermRelationshipFields {
		values = append(values, v)
	}
	return values
}
