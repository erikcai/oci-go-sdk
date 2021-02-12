// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"net/http"
)

// RecommendationsRequest wrapper for the Recommendations operation
type RecommendationsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter used to return only recommendations of the specified type.
	RecommendationType []RecommendationsRecommendationTypeEnum `contributesTo:"query" name:"recommendationType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter used to provide the unique identifier of the source object, for which a list of recommendations will be returned for review.
	SourceObjectKey *string `mandatory:"true" contributesTo:"query" name:"sourceObjectKey"`

	// A filter used to provide the type of the source object, for which a list of recommendations will be returned for review.
	SourceObjectType RecommendationsSourceObjectTypeEnum `mandatory:"true" contributesTo:"query" name:"sourceObjectType" omitEmpty:"true"`

	// A filter used to return only recommendations having the requested status.
	RecommendationStatus RecommendationsRecommendationStatusEnum `mandatory:"false" contributesTo:"query" name:"recommendationStatus" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RecommendationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RecommendationsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RecommendationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// RecommendationsResponse wrapper for the Recommendations operation
type RecommendationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The RecommendationCollection instance
	RecommendationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response RecommendationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RecommendationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RecommendationsRecommendationTypeEnum Enum with underlying type: string
type RecommendationsRecommendationTypeEnum string

// Set of constants representing the allowable values for RecommendationsRecommendationTypeEnum
const (
	RecommendationsRecommendationTypeLinkGlossaryTerm RecommendationsRecommendationTypeEnum = "LINK_GLOSSARY_TERM"
)

var mappingRecommendationsRecommendationType = map[string]RecommendationsRecommendationTypeEnum{
	"LINK_GLOSSARY_TERM": RecommendationsRecommendationTypeLinkGlossaryTerm,
}

// GetRecommendationsRecommendationTypeEnumValues Enumerates the set of values for RecommendationsRecommendationTypeEnum
func GetRecommendationsRecommendationTypeEnumValues() []RecommendationsRecommendationTypeEnum {
	values := make([]RecommendationsRecommendationTypeEnum, 0)
	for _, v := range mappingRecommendationsRecommendationType {
		values = append(values, v)
	}
	return values
}

// RecommendationsSourceObjectTypeEnum Enum with underlying type: string
type RecommendationsSourceObjectTypeEnum string

// Set of constants representing the allowable values for RecommendationsSourceObjectTypeEnum
const (
	RecommendationsSourceObjectTypeDataEntity RecommendationsSourceObjectTypeEnum = "DATA_ENTITY"
	RecommendationsSourceObjectTypeAttribute  RecommendationsSourceObjectTypeEnum = "ATTRIBUTE"
	RecommendationsSourceObjectTypeTerm       RecommendationsSourceObjectTypeEnum = "TERM"
	RecommendationsSourceObjectTypeCategory   RecommendationsSourceObjectTypeEnum = "CATEGORY"
)

var mappingRecommendationsSourceObjectType = map[string]RecommendationsSourceObjectTypeEnum{
	"DATA_ENTITY": RecommendationsSourceObjectTypeDataEntity,
	"ATTRIBUTE":   RecommendationsSourceObjectTypeAttribute,
	"TERM":        RecommendationsSourceObjectTypeTerm,
	"CATEGORY":    RecommendationsSourceObjectTypeCategory,
}

// GetRecommendationsSourceObjectTypeEnumValues Enumerates the set of values for RecommendationsSourceObjectTypeEnum
func GetRecommendationsSourceObjectTypeEnumValues() []RecommendationsSourceObjectTypeEnum {
	values := make([]RecommendationsSourceObjectTypeEnum, 0)
	for _, v := range mappingRecommendationsSourceObjectType {
		values = append(values, v)
	}
	return values
}

// RecommendationsRecommendationStatusEnum Enum with underlying type: string
type RecommendationsRecommendationStatusEnum string

// Set of constants representing the allowable values for RecommendationsRecommendationStatusEnum
const (
	RecommendationsRecommendationStatusAccepted RecommendationsRecommendationStatusEnum = "ACCEPTED"
	RecommendationsRecommendationStatusRejected RecommendationsRecommendationStatusEnum = "REJECTED"
	RecommendationsRecommendationStatusInferred RecommendationsRecommendationStatusEnum = "INFERRED"
)

var mappingRecommendationsRecommendationStatus = map[string]RecommendationsRecommendationStatusEnum{
	"ACCEPTED": RecommendationsRecommendationStatusAccepted,
	"REJECTED": RecommendationsRecommendationStatusRejected,
	"INFERRED": RecommendationsRecommendationStatusInferred,
}

// GetRecommendationsRecommendationStatusEnumValues Enumerates the set of values for RecommendationsRecommendationStatusEnum
func GetRecommendationsRecommendationStatusEnumValues() []RecommendationsRecommendationStatusEnum {
	values := make([]RecommendationsRecommendationStatusEnum, 0)
	for _, v := range mappingRecommendationsRecommendationStatus {
		values = append(values, v)
	}
	return values
}
