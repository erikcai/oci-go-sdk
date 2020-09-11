// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ValidatePatternDetails Validate pattern using the expression and file list.
type ValidatePatternDetails struct {

	// The new yet to be updated expression used in the pattern that may include qualifiers. If this is not provided,existing expression will be used to validate
	Expression *string `mandatory:"false" json:"expression"`

	// The list of file paths on which the expression is applied as a filter.
	FilePathList []string `mandatory:"false" json:"filePathList"`
}

func (m ValidatePatternDetails) String() string {
	return common.PointerString(m)
}
