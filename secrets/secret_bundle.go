// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Secrets
//
// APIs to retrieve secrets
//

package secrets

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// SecretBundle A SecretBundle contains the actual secret content as well as meta-data of the secret bundle.
type SecretBundle struct {

	// The OCID of the secret.
	SecretId *string `mandatory:"true" json:"secretId"`

	// The version number of the secret.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	// The time when the secret bundle was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The name of the particular secret version.
	VersionName *string `mandatory:"false" json:"versionName"`

	SecretBundleContent SecretBundleContentDetails `mandatory:"false" json:"secretBundleContent"`

	// An optional property indicating when to delete the secret version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// An optional property indicating when the secret version will expire, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfExpiry *common.SDKTime `mandatory:"false" json:"timeOfExpiry"`

	// A list of stages of the secretVersion
	Stages []SecretBundleStagesEnum `mandatory:"false" json:"stages,omitempty"`

	// Metadata of the secret.
	Metadata map[string]interface{} `mandatory:"false" json:"metadata"`
}

func (m SecretBundle) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *SecretBundle) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated         *common.SDKTime            `json:"timeCreated"`
		VersionName         *string                    `json:"versionName"`
		SecretBundleContent secretbundlecontentdetails `json:"secretBundleContent"`
		TimeOfDeletion      *common.SDKTime            `json:"timeOfDeletion"`
		TimeOfExpiry        *common.SDKTime            `json:"timeOfExpiry"`
		Stages              []SecretBundleStagesEnum   `json:"stages"`
		Metadata            map[string]interface{}     `json:"metadata"`
		SecretId            *string                    `json:"secretId"`
		VersionNumber       *int64                     `json:"versionNumber"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.VersionName = model.VersionName

	nn, e = model.SecretBundleContent.UnmarshalPolymorphicJSON(model.SecretBundleContent.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecretBundleContent = nn.(SecretBundleContentDetails)
	} else {
		m.SecretBundleContent = nil
	}

	m.TimeOfDeletion = model.TimeOfDeletion

	m.TimeOfExpiry = model.TimeOfExpiry

	m.Stages = make([]SecretBundleStagesEnum, len(model.Stages))
	for i, n := range model.Stages {
		m.Stages[i] = n
	}

	m.Metadata = model.Metadata

	m.SecretId = model.SecretId

	m.VersionNumber = model.VersionNumber
	return
}

// SecretBundleStagesEnum Enum with underlying type: string
type SecretBundleStagesEnum string

// Set of constants representing the allowable values for SecretBundleStagesEnum
const (
	SecretBundleStagesCurrent    SecretBundleStagesEnum = "CURRENT"
	SecretBundleStagesPending    SecretBundleStagesEnum = "PENDING"
	SecretBundleStagesLatest     SecretBundleStagesEnum = "LATEST"
	SecretBundleStagesPrevious   SecretBundleStagesEnum = "PREVIOUS"
	SecretBundleStagesDeprecated SecretBundleStagesEnum = "DEPRECATED"
)

var mappingSecretBundleStages = map[string]SecretBundleStagesEnum{
	"CURRENT":    SecretBundleStagesCurrent,
	"PENDING":    SecretBundleStagesPending,
	"LATEST":     SecretBundleStagesLatest,
	"PREVIOUS":   SecretBundleStagesPrevious,
	"DEPRECATED": SecretBundleStagesDeprecated,
}

// GetSecretBundleStagesEnumValues Enumerates the set of values for SecretBundleStagesEnum
func GetSecretBundleStagesEnumValues() []SecretBundleStagesEnum {
	values := make([]SecretBundleStagesEnum, 0)
	for _, v := range mappingSecretBundleStages {
		values = append(values, v)
	}
	return values
}
