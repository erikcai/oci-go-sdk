// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Secrets Management APIs
//
// Secrets Management APIs
//

package vault

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SecretVersionSummary The secret version summary object, which doesn't include the contents of the secret.
type SecretVersionSummary struct {

	// The OCID of the secret.
	SecretId *string `mandatory:"true" json:"secretId"`

	// A optional property indicating when the secret version was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The version number of the secret.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	// The type of the content
	ContentType SecretVersionSummaryContentTypeEnum `mandatory:"false" json:"contentType,omitempty"`

	// The name of a secret version. A name is unique in a secret.
	Name *string `mandatory:"false" json:"name"`

	// A list of possible rotation states for the secret version.
	Stages []SecretVersionSummaryStagesEnum `mandatory:"false" json:"stages,omitempty"`

	// An optional property indicating when to delete the secret version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// An optional property indicating when the secret version will expire, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfExpiry *common.SDKTime `mandatory:"false" json:"timeOfExpiry"`
}

func (m SecretVersionSummary) String() string {
	return common.PointerString(m)
}

// SecretVersionSummaryContentTypeEnum Enum with underlying type: string
type SecretVersionSummaryContentTypeEnum string

// Set of constants representing the allowable values for SecretVersionSummaryContentTypeEnum
const (
	SecretVersionSummaryContentTypeBase64 SecretVersionSummaryContentTypeEnum = "BASE64"
)

var mappingSecretVersionSummaryContentType = map[string]SecretVersionSummaryContentTypeEnum{
	"BASE64": SecretVersionSummaryContentTypeBase64,
}

// GetSecretVersionSummaryContentTypeEnumValues Enumerates the set of values for SecretVersionSummaryContentTypeEnum
func GetSecretVersionSummaryContentTypeEnumValues() []SecretVersionSummaryContentTypeEnum {
	values := make([]SecretVersionSummaryContentTypeEnum, 0)
	for _, v := range mappingSecretVersionSummaryContentType {
		values = append(values, v)
	}
	return values
}

// SecretVersionSummaryStagesEnum Enum with underlying type: string
type SecretVersionSummaryStagesEnum string

// Set of constants representing the allowable values for SecretVersionSummaryStagesEnum
const (
	SecretVersionSummaryStagesCurrent    SecretVersionSummaryStagesEnum = "CURRENT"
	SecretVersionSummaryStagesPending    SecretVersionSummaryStagesEnum = "PENDING"
	SecretVersionSummaryStagesLatest     SecretVersionSummaryStagesEnum = "LATEST"
	SecretVersionSummaryStagesPrevious   SecretVersionSummaryStagesEnum = "PREVIOUS"
	SecretVersionSummaryStagesDeprecated SecretVersionSummaryStagesEnum = "DEPRECATED"
)

var mappingSecretVersionSummaryStages = map[string]SecretVersionSummaryStagesEnum{
	"CURRENT":    SecretVersionSummaryStagesCurrent,
	"PENDING":    SecretVersionSummaryStagesPending,
	"LATEST":     SecretVersionSummaryStagesLatest,
	"PREVIOUS":   SecretVersionSummaryStagesPrevious,
	"DEPRECATED": SecretVersionSummaryStagesDeprecated,
}

// GetSecretVersionSummaryStagesEnumValues Enumerates the set of values for SecretVersionSummaryStagesEnum
func GetSecretVersionSummaryStagesEnumValues() []SecretVersionSummaryStagesEnum {
	values := make([]SecretVersionSummaryStagesEnum, 0)
	for _, v := range mappingSecretVersionSummaryStages {
		values = append(values, v)
	}
	return values
}
