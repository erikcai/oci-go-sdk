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

// CreateScriptDetails The request body used to create new Script.
// Only Side or Java Script content types are supported. Content should be in side or java script formats only.
type CreateScriptDetails struct {

	// A user-friendly name. Must be unique, and it can be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Content type of script.
	ContentType ContentTypesEnum `mandatory:"true" json:"contentType"`

	// Script Content. It may contain custom defined tags which can be used for setting dynamic parameters.
	// Format is `<ORAP><ON>param name</ON><OV>param value</OV><OS>isParamValueSecret(true/false)</OS></ORAP>`.
	// Param value and isParamValueSecret are optional, default value for isParamValueSecret is false. Few examples -
	// with mandatory param name : `<ORAP><ON>param name</ON></ORAP>`
	// with param name and value : `<ORAP><ON>param name</ON><OV>param value</OV></ORAP>`
	// Content is valid if content is matching given content type. If content type is SIDE, then the content should be in side format. If content type is JS, then the content should be in java script format.
	Content *string `mandatory:"true" json:"content"`

	// Filename of uploaded script content.
	ContentFileName *string `mandatory:"false" json:"contentFileName"`

	// List of script parameters. Example: `[{"paramName": "userid", "paramValue":"testuser", "isSecret": false}]`
	Parameters []ScriptParameter `mandatory:"false" json:"parameters"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateScriptDetails) String() string {
	return common.PointerString(m)
}
