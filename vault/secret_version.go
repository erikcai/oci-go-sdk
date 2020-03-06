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

// SecretVersion The details of the secret version, excluding the contents of the secret.
type SecretVersion struct {

	// The type of the content
	ContentType SecretVersionContentTypeEnum `mandatory:"false" json:"contentType,omitempty"`

	// The name of a secret version. A name is unique in a secret.
	Name *string `mandatory:"false" json:"name"`

	// The OCID of the secret
	SecretId *string `mandatory:"false" json:"secretId"`

	// A list of possible rotation states for the secret version. A secret version marked `CURRENT` is currently in use. A secret version
	// marked `PENDING` is staged and available for use, but has not been applied on the target system and, therefore, has not been rotated
	// into current, active use. The secret most recently uploaded to a vault is always marked `LATEST`. (The first version of a secret is
	// always marked as both `CURRENT` and `LATEST`.) A secret version marked `PREVIOUS` is the secret version that was most recently marked
	// `CURRENT`, before the last secret version rotation. A secret version marked `DEPRECATED` is neither current, pending, nor the previous
	// one in use. Only secret versions marked `DEPRECATED` can be scheduled for deletion.
	Stages []SecretVersionStagesEnum `mandatory:"false" json:"stages,omitempty"`

	// A optional property indicating when the secret version was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// An optional property indicating when to delete the secret version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// An optional property indicating when the current secret version will expire, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfCurrentVersionExpiry *common.SDKTime `mandatory:"false" json:"timeOfCurrentVersionExpiry"`

	// The version number of the secret.
	VersionNumber *int64 `mandatory:"false" json:"versionNumber"`
}

func (m SecretVersion) String() string {
	return common.PointerString(m)
}

// SecretVersionContentTypeEnum Enum with underlying type: string
type SecretVersionContentTypeEnum string

// Set of constants representing the allowable values for SecretVersionContentTypeEnum
const (
	SecretVersionContentTypeBase64 SecretVersionContentTypeEnum = "BASE64"
)

var mappingSecretVersionContentType = map[string]SecretVersionContentTypeEnum{
	"BASE64": SecretVersionContentTypeBase64,
}

// GetSecretVersionContentTypeEnumValues Enumerates the set of values for SecretVersionContentTypeEnum
func GetSecretVersionContentTypeEnumValues() []SecretVersionContentTypeEnum {
	values := make([]SecretVersionContentTypeEnum, 0)
	for _, v := range mappingSecretVersionContentType {
		values = append(values, v)
	}
	return values
}

// SecretVersionStagesEnum Enum with underlying type: string
type SecretVersionStagesEnum string

// Set of constants representing the allowable values for SecretVersionStagesEnum
const (
	SecretVersionStagesCurrent    SecretVersionStagesEnum = "CURRENT"
	SecretVersionStagesPending    SecretVersionStagesEnum = "PENDING"
	SecretVersionStagesLatest     SecretVersionStagesEnum = "LATEST"
	SecretVersionStagesPrevious   SecretVersionStagesEnum = "PREVIOUS"
	SecretVersionStagesDeprecated SecretVersionStagesEnum = "DEPRECATED"
)

var mappingSecretVersionStages = map[string]SecretVersionStagesEnum{
	"CURRENT":    SecretVersionStagesCurrent,
	"PENDING":    SecretVersionStagesPending,
	"LATEST":     SecretVersionStagesLatest,
	"PREVIOUS":   SecretVersionStagesPrevious,
	"DEPRECATED": SecretVersionStagesDeprecated,
}

// GetSecretVersionStagesEnumValues Enumerates the set of values for SecretVersionStagesEnum
func GetSecretVersionStagesEnumValues() []SecretVersionStagesEnum {
	values := make([]SecretVersionStagesEnum, 0)
	for _, v := range mappingSecretVersionStages {
		values = append(values, v)
	}
	return values
}