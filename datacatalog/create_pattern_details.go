// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// CreatePatternDetails Properties used in data asset create operations.
type CreatePatternDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Detailed description of the Pattern.
	Description *string `mandatory:"false" json:"description"`

	// The expression used in the pattern that may include qualifiers. If a filePathList is sent and no logical entities
	// are derived or all files are in the UNMATCHED group, then the expression results in validation failure.
	Expression *string `mandatory:"false" json:"expression"`

	// The list of file paths on which the expression is applied as a filter. Used for validating the expression.
	FilePathList []string `mandatory:"false" json:"filePathList"`

	// Indicates whether this pattern creation is rejected resulting in error if the expression is invalid.
	IsRejectedOnValidationFailure *bool `mandatory:"false" json:"isRejectedOnValidationFailure"`

	// Indicates the max number of UNMATCHED files for the expression to be invalid.
	RejectionLimit *int `mandatory:"false" json:"rejectionLimit"`

	// A map of maps that contains the properties which are specific to the pattern type. Each pattern type
	// definition defines it's set of required and optional properties.
	// Example: `{"properties": { "default": { "tbd"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m CreatePatternDetails) String() string {
	return common.PointerString(m)
}
