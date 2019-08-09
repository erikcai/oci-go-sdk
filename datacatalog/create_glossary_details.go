// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateGlossaryDetails Properties used in Glossary create operations.
type CreateGlossaryDetails struct {

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Detailed description of the Glossary.
	Description *string `mandatory:"false" json:"description"`
}

func (m CreateGlossaryDetails) String() string {
	return common.PointerString(m)
}
