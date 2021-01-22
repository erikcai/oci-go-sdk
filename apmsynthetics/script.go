// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Synthetic
//
// API for APM Synthetic service. Use this API to query Synthethetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// Script The information about the script.
type Script struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the script.
	// scriptId is mandatory for creation of SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For any other monitor types, it should be set to null.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Must be unique, and it can be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Content type of the script.
	ContentType ContentTypesEnum `mandatory:"true" json:"contentType"`

	MonitorStatusCountMap *MonitorStatusCountMap `mandatory:"true" json:"monitorStatusCountMap"`

	// Script Content. It may contain custom defined tags which can be used for setting dynamic parameters.
	// Format is `<ORAP><ON>param name</ON><OV>param value</OV><OS>isParamValueSecret(true/false)</OS></ORAP>`.
	// Param value and isParamValueSecret are optional, default value for isParamValueSecret is false. Few examples -
	// with mandatory param name : `<ORAP><ON>param name</ON></ORAP>`
	// with param name and value : `<ORAP><ON>param name</ON><OV>param value</OV></ORAP>`
	// Content is valid if content is matching given content type. If content type is SIDE, then the content should be in side format. If content type is JS, then the content should be in java script format.
	Content *string `mandatory:"false" json:"content"`

	// Time when script uploaded.
	TimeUploaded *common.SDKTime `mandatory:"false" json:"timeUploaded"`

	// Size of the script content.
	ContentSizeInBytes *int `mandatory:"false" json:"contentSizeInBytes"`

	// Filename of the uploaded script content.
	ContentFileName *string `mandatory:"false" json:"contentFileName"`

	// List of script parameter information. Example: `[{"scriptParameter": {"paramName": "userid", "paramValue":"testuser", "isSecret": false}, "isOverwritten": false}]`
	Parameters []ScriptParameterInfo `mandatory:"false" json:"parameters"`

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Script) String() string {
	return common.PointerString(m)
}
